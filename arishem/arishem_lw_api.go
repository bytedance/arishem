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
	"context"
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"strings"
)

var (
	emptyDataCtx typedef.DataCtx
)

type dummyVisitTarget struct{ name string }

func (d *dummyVisitTarget) Identifier() string { return d.name }

// DataContext will create a new typedef.DataCtx by fact meta json.
func DataContext(ctx context.Context, json string) (dc typedef.DataCtx, err error) {
	dc, err = core.NewArishemDataCtx(ctx, json, arishemConfiguration.FeatFetcherFactory())
	return
}

// DataContextFromMeta will create a new typedef.DataCtx by fact meta map
func DataContextFromMeta(ctx context.Context, meta typedef.MetaType) (dc typedef.DataCtx, err error) {
	dc, err = core.NewArishemDataCtxFromMeta(ctx, meta, arishemConfiguration.FeatFetcherFactory())
	return
}

// ParseCondition will parse a condition string and return the rule tree with feature parameter pre-parsed.
func ParseCondition(condition string) (tree *RuleTree, err error) {
	return ParseRuleTree(condition, ExprTypeCondition)
}

// AddSubCondition adds a sub condition to arishem, it will be used at the moment of judging a condition which right hand type is SubCondExpr,
// pairs allow users to add some conditions when enable arishem's sub-condition feature
func AddSubCondition(pairs ...string) error {
	size := len(pairs)
	if size%2 != 0 {
		return errors.New("pair number not valid")
	}
	if arishemConfiguration.SubCond == nil {
		return errors.New("sub condition not enabled")
	}
	for i := 0; i < size; i += 2 {
		condName := pairs[i]
		condition := pairs[i+1]
		tree, err := ParseRuleTree(condition, ExprTypeCondition)
		if err != nil {
			return err
		}
		if tree != nil && tree.Tree != nil {
			arishemConfiguration.SubCond.WhenConditionParsed(condName, condition, tree.Tree)
		}
	}
	return nil
}

// ParseAim will parse an aim string and return the rule tree with feature parameter pre-parsed.
func ParseAim(aim string) (tree *RuleTree, err error) {
	return ParseRuleTree(aim, ExprTypeAim)
}

// ParseRuleTree will try parsing expr to antlr parse tree, err will not be null when failed.
func ParseRuleTree(expr string, exprType ExprType) (exprTree *RuleTree, err error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		err = errors.New("empty expression")
		return
	}
	expr, err = tools.Compat(expr)
	if err != nil {
		return
	}
	var ok bool
	exprTree, ok = arishemConfiguration.TCache.Get(expr)
	if ok && exprTree != nil {
		return
	}
	exprTree, err = parseRuleTree(expr, exprType)
	if err != nil {
		return
	}
	if exprTree.Tree == nil {
		err = errors.New("tree parsed is null")
		return
	} else {
		arishemConfiguration.TCache.Put(expr, exprTree)
	}
	return
}

// ParseFeature input a feature expression string and return the typedef.FeatureParam
func ParseFeature(featureExpr string) (typedef.FeatureParam, error) {
	fppl := core.NewFeaturePreParseListener()
	err := parser.ParseSingleFeature(featureExpr, fppl)
	if err != nil {
		return nil, err
	}
	featureParams := fppl.Data()
	if len(featureParams) <= 0 {
		return nil, errors.New("parse failed: empty feature params")
	}
	return featureParams[0], nil
}

// WalkAim will try to parse the aimExpr into the antlr parse tree and visit it to get the result, err will not be null when parse aim expression failed.
func WalkAim(aimExpr string, dc typedef.DataCtx, opts ...ExecuteOption) (aim typedef.Aim, err error) {
	var tree *RuleTree
	tree, err = ParseAim(aimExpr)
	if err != nil {
		return
	}
	rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
	ApplyExecuteOptions(rv, dc, opts...)
	if len(tree.FeatParams) > 0 {
		dc.PrefetchFeatures(tree.FeatParams)
	}
	aim = rv.VisitAim(tree.Tree, dc, &dummyVisitTarget{name: aimExpr})
	return
}

// WalkAimTree will visit the aimTree to get the result
func WalkAimTree(aimTree antlr.ParseTree, dc typedef.DataCtx, opts ...ExecuteOption) (aim typedef.Aim) {
	rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
	ApplyExecuteOptions(rv, dc, opts...)
	aim = rv.VisitAim(aimTree, dc, &dummyVisitTarget{name: aimTree.GetText()})
	return
}

// JudgeCondition will judge the condition string, pass will be set true if condition passed.
func JudgeCondition(ctx context.Context, condition string, opts ...ExecuteOption) (pass bool, err error) {
	return JudgeConditionWithFactMeta(ctx, condition, "{}", opts...)
}

// JudgeConditionWithFactMeta will judge the condition string with fact meta, pass will be set true if condition passed.
func JudgeConditionWithFactMeta(ctx context.Context, condition, factMeta string, opts ...ExecuteOption) (pass bool, err error) {
	tree, err := ParseCondition(condition)
	if err != nil {
		return
	}
	var dc typedef.DataCtx
	factMeta = strings.TrimSpace(factMeta)
	if factMeta == "" {
		if emptyDataCtx == nil {
			emptyDataCtx, _ = core.NewArishemDataCtx(ctx, "{}", arishemConfiguration.FeatFetcherFactory())
		}
		dc = emptyDataCtx
	} else {
		dc, err = core.NewArishemDataCtx(ctx, factMeta, arishemConfiguration.FeatFetcherFactory())
	}
	if err != nil {
		return
	}
	rv := core.NewArishemRuleVisitor(arishemConfiguration.getConditionFinder())
	ApplyExecuteOptions(rv, dc, opts...)
	if len(tree.FeatParams) > 0 {
		dc.PrefetchFeatures(tree.FeatParams)
	}
	pass = rv.VisitCondition(tree.Tree, dc, &dummyVisitTarget{name: condition})
	return
}
