/*
 *  Copyright Bytedance Ltd. and/or its affiliates.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package arishem

import (
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"sort"
	"sync"
)

// ExecuteSingleRule will execute a single rule, if condition pass, rr.Aim() will not be null.
func ExecuteSingleRule(rule RuleTarget, dc typedef.DataCtx, opts ...ExecuteOption) (rr RuleResult) {
	if rule == nil || dc == nil {
		return
	}
	if len(opts) > 0 {
		// consume data context option
		ApplyExecuteOptions(nil, dc, opts...)
	}
	cdtFeats := rule.CondFeatParams()
	if len(cdtFeats) > 0 {
		dc.PrefetchFeatures(cdtFeats)
	}
	r := &ruleResult{
		name:   rule.Identifier(),
		passed: false,
		aim:    nil,
	}
	rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
	if len(opts) > 0 {
		// consume visit context option
		ApplyExecuteOptions(rv, nil, opts...)
	}
	r.passed = rv.VisitCondition(rule.ConditionPTree(), dc, rule)
	if !r.passed {
		rr = r
		return
	}
	aimFeats := rule.AimFeatParams()
	if len(aimFeats) > 0 {
		dc.PrefetchFeatures(aimFeats)
	}
	r.aim = rv.VisitAim(rule.AimPTree(), dc, rule)
	rr = r
	return
}

// ExecuteRules will execute all rule target, and returns rule results which passed by condition.
func ExecuteRules(rules []RuleTarget, dc typedef.DataCtx, opts ...ExecuteOption) (rrs []RuleResult) {
	rulesSize := len(rules)
	if rulesSize <= 0 || dc == nil {
		return
	}
	// distribute different rules
	wg := sync.WaitGroup{}
	priRules, noPriRules := distributeRules(rules)
	var priRR RuleResult
	var noPriRR []RuleResult
	if len(priRules) > 0 {
		wg.Add(1)
		arishemConfiguration.RuleComputePool.Submit(func(param interface{}) {
			defer wg.Done()
			priRR = executePriorityRules(priRules, dc, opts...)
		}, nil)

	}
	if len(noPriRules) > 0 {
		wg.Add(1)
		arishemConfiguration.RuleComputePool.Submit(func(param interface{}) {
			defer wg.Done()
			noPriRR = executeNoPriorityRules(noPriRules, dc, opts...)
		}, nil)
	}
	wg.Wait()

	if priRR != nil {
		rrs = append(rrs, priRR)
	}
	rrs = append(rrs, noPriRR...)
	return
}

func executePriorityRules(priRules []RuleTarget, dc typedef.DataCtx, opts ...ExecuteOption) (rr RuleResult) {
	if len(priRules) <= 0 {
		return
	}
	sort.Slice(priRules, func(i, j int) bool {
		return priRules[i].Compare(priRules[j]) < 0
	})
	// batch rules
	batchedRules := batchRules(priRules, arishemConfiguration.Granularity(len(priRules), ExecuteModePriority))
	// create visitor
	rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
	ApplyExecuteOptions(rv, dc, opts...)
outer:
	for idx, rules := range batchedRules {
		featParams := groupedConditionPreFetchFeatures(idx, batchedRules)
		if len(featParams) > 0 {
			dc.PrefetchFeatures(featParams)
		}
		// execute rule computation orderly
		for _, rule := range rules {
			pass := rv.VisitCondition(rule.ConditionPTree(), dc, rule)
			if pass {
				if len(rule.AimFeatParams()) > 0 {
					dc.PrefetchFeatures(rule.AimFeatParams())
				}
				aim := rv.VisitAim(rule.AimPTree(), dc, rule)
				// rr no race here
				rr = &ruleResult{name: rule.Identifier(), passed: pass, aim: aim}
				// break if passed
				break outer
			}
		}
	}
	return
}

func executeNoPriorityRules(noPriRules []RuleTarget, dc typedef.DataCtx, opts ...ExecuteOption) (rrs []RuleResult) {
	if len(noPriRules) <= 0 {
		return
	}
	// batch no priority rules
	batchedRules := batchRules(noPriRules, arishemConfiguration.Granularity(len(noPriRules), ExecuteModeNoPriority))
	ApplyExecuteOptions(nil, dc, opts...)
	raceCheck := sync.Mutex{}
	for idx, rules := range batchedRules {
		wg := sync.WaitGroup{}
		featParams := groupedConditionPreFetchFeatures(idx, batchedRules)
		if len(featParams) > 0 {
			dc.PrefetchFeatures(featParams)
		}
		for _, rule := range rules {
			wg.Add(1)
			arishemConfiguration.RuleComputePool.Submit(func(rI interface{}) {
				defer wg.Done()
				r := rI.(RuleTarget)

				rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
				ApplyExecuteOptions(rv, nil, opts...)
				pass := rv.VisitCondition(r.ConditionPTree(), dc, r)
				if pass {
					if len(rule.AimFeatParams()) > 0 {
						dc.PrefetchFeatures(rule.AimFeatParams())
					}
					aim := rv.VisitAim(r.AimPTree(), dc, r)
					raceCheck.Lock()
					rrs = append(rrs, &ruleResult{name: r.Identifier(), passed: pass, aim: aim})
					raceCheck.Unlock()
				}
			}, rule)
		}
		wg.Wait()
	}
	return
}

func groupedConditionPreFetchFeatures(currIdx int, batchedRules [][]RuleTarget) (featParams []typedef.FeatureParam) {
	if len(batchedRules) <= 0 {
		return
	}
	set := tools.NewHashSet()
	for _, rule := range batchedRules[currIdx] {
		for _, param := range rule.CondFeatParams() {
			set.Add(param.(typedef.Hashable))
		}
	}
	// if prefetching is configured, the next batch of features will be requested in advance when the rule is executed
	if arishemConfiguration.Prefetch && currIdx < len(batchedRules)-1 {
		for _, rule := range batchedRules[currIdx+1] {
			for _, param := range rule.CondFeatParams() {
				set.Add(param.(typedef.Hashable))
			}
		}
	}
	for _, hashable := range set.AsList() {
		featParams = append(featParams, hashable.(typedef.FeatureParam))
	}
	return
}

func distributeRules(rules []RuleTarget) (priority []RuleTarget, noPriority []RuleTarget) {
	if len(rules) <= 0 {
		return
	}
	for _, rule := range rules {
		switch r := rule.(type) {
		case *PriorityRule:
			priority = append(priority, r)
		case *NoPriorityRule:
			noPriority = append(noPriority, r)
		default:
			// default to priority rule, use to priority customize
			priority = append(priority, r)
		}
	}
	return
}

func batchRules(rules []RuleTarget, batchSize int) [][]RuleTarget {
	if len(rules) <= 0 || batchSize <= 0 {
		return [][]RuleTarget{rules}
	}
	var batches [][]RuleTarget
	for batchSize < len(rules) {
		rules, batches = rules[batchSize:], append(batches, rules[0:batchSize:batchSize])
	}
	// elements left
	batches = append(batches, rules)
	return batches
}
