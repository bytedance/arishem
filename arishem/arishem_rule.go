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
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"strings"
	"sync"
)

type RuleTarget interface {
	typedef.Comparable
	typedef.VisitTarget

	ConditionPTree() antlr.ParseTree
	AimPTree() antlr.ParseTree

	CondFeatParams() []typedef.FeatureParam
	AimFeatParams() []typedef.FeatureParam
}

type RuleResult interface {
	typedef.VisitTarget

	Passed() bool
	Aim() typedef.Aim
}

type exprType int

const (
	exprTypeCondition = iota
	exprTypeAim
)

type ruleResult struct {
	name   string
	passed bool
	aim    typedef.Aim
}

func (r *ruleResult) Identifier() string {
	return r.name
}

func (r *ruleResult) Passed() bool {
	return r.passed
}

func (r *ruleResult) Aim() typedef.Aim {
	return r.aim
}

type NoPriorityRule struct {
	name    string
	cond    parser.ICondEntityContext
	aim     parser.IAimEntityContext
	condFps []typedef.FeatureParam
	aimFps  []typedef.FeatureParam
}

type PriorityRule struct {
	*NoPriorityRule
	priority int
}

// NewPriorityRule will create a new rule with priority.
func NewPriorityRule(name string, priority int, cdtExpr, aimExpr string) (pr RuleTarget, err error) {
	var npr RuleTarget
	npr, err = NewNoPriorityRule(name, cdtExpr, aimExpr)
	if err != nil {
		return
	}
	pr = &PriorityRule{NoPriorityRule: npr.(*NoPriorityRule), priority: priority}
	return
}

// NewNoPriorityRule will create a new rule with no priority.
func NewNoPriorityRule(name string, cdtExpr, aimExpr string) (npr RuleTarget, err error) {
	cdtExpr = strings.TrimSpace(cdtExpr)
	if cdtExpr == "" {
		err = errors.New("cond expression is empty")
		return
	}
	aimExpr = strings.TrimSpace(aimExpr)
	if aimExpr == "" {
		err = errors.New("aim expression is empty")
		return
	}
	cdtTree, aimTree, err := tryParse(cdtExpr, aimExpr)
	if err != nil {
		return
	}
	npr = &NoPriorityRule{
		name:    name,
		cond:    cdtTree.Tree.(parser.ICondEntityContext),
		condFps: cdtTree.FeatParams,
		aim:     aimTree.Tree.(parser.IAimEntityContext),
		aimFps:  aimTree.FeatParams,
	}
	return
}

func (n *NoPriorityRule) Compare(other typedef.Comparable) int {
	return 0
}

func (n *NoPriorityRule) Identifier() string {
	return n.name
}

func (n *NoPriorityRule) ConditionPTree() antlr.ParseTree {
	return n.cond
}

func (n *NoPriorityRule) AimPTree() antlr.ParseTree {
	return n.aim
}

func (n *NoPriorityRule) CondFeatParams() []typedef.FeatureParam {
	return n.condFps
}

func (n *NoPriorityRule) AimFeatParams() []typedef.FeatureParam {
	return n.aimFps
}

func (p *PriorityRule) Compare(other typedef.Comparable) int {
	cmp := 0
	switch t := other.(type) {
	case *NoPriorityRule:
		cmp = 1
	case *PriorityRule:
		if p.priority < t.priority {
			cmp = -1
		} else if p.priority == t.priority {
			cmp = 0
		} else {
			cmp = 1
		}
	}
	return cmp
}

func tryParse(cdtExpr, aimExpr string) (cdtTree, aimTree *RuleTree, err error) {
	// compat two expressions
	cdtExpr, err = tools.Compat(cdtExpr)
	if err != nil {
		return
	}
	aimExpr, err = tools.Compat(aimExpr)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}

	var cdtErr error
	treeCache, ok := arishemConfiguration.TCache.Get(cdtExpr)
	if ok {
		// remove debug info
		//logger.Info("cdt cache hit!")
		cdtTree = treeCache
	} else {
		wg.Add(1)
		arishemConfiguration.RuleComputePool.Submit(func(interface{}) {
			defer wg.Done()
			cdtTree, cdtErr = parseRuleTree(cdtExpr, exprTypeCondition)
		}, nil)
	}

	var aimErr error
	treeCache, ok = arishemConfiguration.TCache.Get(aimExpr)
	if ok {
		// remove debug info
		//logger.Info("aim cache hit!")
		aimTree = treeCache
	} else {
		wg.Add(1)
		arishemConfiguration.RuleComputePool.Submit(func(interface{}) {
			defer wg.Done()
			aimTree, aimErr = parseRuleTree(aimExpr, exprTypeAim)
		}, nil)
	}
	// wait for parsing
	wg.Wait()

	if cdtTree != nil {
		arishemConfiguration.TCache.Put(cdtExpr, cdtTree)
	}
	if aimTree != nil {
		arishemConfiguration.TCache.Put(aimExpr, aimTree)
	}

	var errMsg string
	if cdtErr != nil {
		errMsg = "cond parse error " + cdtErr.Error()
	}
	if aimErr != nil {
		errMsg += ", aim parse error " + aimErr.Error()
	}
	if errMsg != "" {
		err = errors.New(errMsg)
	}
	return
}

func parseRuleTree(exprStr string, expT exprType) (exprTree *RuleTree, err error) {
	fppl := core.NewFeaturePreParseListener()

	var treeCtx antlr.ParseTree
	if expT == exprTypeCondition {
		treeCtx, err = parser.ParseArishemCondition(exprStr, fppl)
	} else {
		treeCtx, err = parser.ParseArishemAim(exprStr, fppl)
	}
	featParams := fppl.Data()
	fperr := fppl.Error()

	if fperr != nil && err == nil {
		err = fperr
	}
	if err != nil {
		return
	}
	exprTree = &RuleTree{Tree: treeCtx, FeatParams: featParams}
	return
}
