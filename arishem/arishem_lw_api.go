/*
 *  Copyright 2023 ByteDance Inc.
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
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/tools"
	"github.com/bytedance/arishem/internal/typedef"
	"strings"
)

var (
	emptyDataCtx typedef.DataCtx
)

type dummyVisitTarget struct{ name string }

func (d *dummyVisitTarget) Identifier() string { return d.name }

// DataContext will create a new DataContext by fact meta json.
func DataContext(json string) (dc typedef.DataCtx, err error) {
	dc, err = core.NewArishemDataCtx(json, arishemConfiguration.FeatFetcherFactory())
	return
}

// ParseCondition will parse a condition string and return the rule tree with feature parameter pre-parsed.
func ParseCondition(condition string) (tree *core.RuleTree, err error) {
	condition = strings.TrimSpace(condition)
	if condition == "" {
		err = errors.New("empty cond text")
		return
	}
	condition, err = tools.Compat(condition)
	if err != nil {
		return
	}
	var ok bool
	tree, ok = arishemConfiguration.TCache.Get(condition)
	if ok && tree != nil {
		return
	}
	tree, err = parseRuleTree(condition, exprTypeCondition)
	if err != nil {
		return
	}
	if tree.Tree == nil {
		err = errors.New("tree parsed is null")
		return
	}
	return
}

// JudgeCondition will judge the condition string, pass will be set true if condition passed.
func JudgeCondition(condition string, opts ...core.ExecuteOption) (pass bool, err error) {
	return JudgeConditionWithFactMeta(condition, "{}", opts...)
}

// JudgeConditionWithFactMeta will judge the condition string with fact meta, pass will be set true if condition passed.
func JudgeConditionWithFactMeta(condition, factMeta string, opts ...core.ExecuteOption) (pass bool, err error) {
	tree, err := ParseCondition(condition)
	if err != nil {
		return
	}
	var dc typedef.DataCtx
	factMeta = strings.TrimSpace(factMeta)
	if factMeta == "" {
		if emptyDataCtx == nil {
			emptyDataCtx, _ = core.NewArishemDataCtx("{}", arishemConfiguration.FeatFetcherFactory())
		}
		dc = emptyDataCtx
	} else {
		dc, err = core.NewArishemDataCtx(factMeta, arishemConfiguration.FeatFetcherFactory())
	}
	if err != nil {
		return
	}
	rv := core.NewArishemRuleVisitor()
	core.ApplyExecuteOptions(rv, dc, opts...)
	if len(tree.FeatParams) > 0 {
		dc.PrefetchFeatures(tree.FeatParams)
	}
	pass = rv.VisitCondition(tree.Tree, dc, &dummyVisitTarget{name: condition})
	return
}
