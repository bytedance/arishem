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

package core

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/funcs"
	"github.com/bytedance/arishem/internal/operator"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"github.com/mnogu/go-calculator"
	"strconv"
	"strings"
)

const foreachKey = "FOREACH"

type simpleVisitTarget struct {
	name string
}

func (s *simpleVisitTarget) Identifier() string {
	return s.name
}

type arishemRuleVisitor struct {
	*parser.BasearishemVisitor
	visitTarget typedef.VisitTarget
	dataCtx     typedef.DataCtx
	observers   map[string]typedef.VisitObserver
	condFinder  func(string) (antlr.ParseTree, error)
}

func NewArishemRuleVisitor(condFinder func(string) (antlr.ParseTree, error)) typedef.RuleVisitor {
	rv := &arishemRuleVisitor{
		BasearishemVisitor: &parser.BasearishemVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		observers: make(map[string]typedef.VisitObserver, 1),
	}
	if condFinder != nil {
		rv.condFinder = condFinder
	}
	return rv
}

func (a *arishemRuleVisitor) errorCallback(node, errMsg string) {
	for _, obsr := range a.observers {
		obsr.OnVisitError(a.dataCtx.Context(), node, errMsg, a.visitTarget)
	}
}

func (a *arishemRuleVisitor) AddVisitObserver(v ...typedef.VisitObserver) {
	for _, observer := range v {
		a.observers[observer.HashCode()] = observer
	}
}

func (a *arishemRuleVisitor) GetVisitObservers() []typedef.VisitObserver {
	var obs []typedef.VisitObserver
	for _, observer := range a.observers {
		obs = append(obs, observer)
	}
	return obs
}

func (a *arishemRuleVisitor) ClearVisitObservers() {
	a.observers = make(map[string]typedef.VisitObserver, 1)
}

func (a *arishemRuleVisitor) SetConditionFinder(cf func(string) (antlr.ParseTree, error)) {
	a.condFinder = cf
}

// VisitCondition visit condition entity to determinate condition is passed
func (a *arishemRuleVisitor) VisitCondition(cdtCtx antlr.ParseTree, dCtx typedef.DataCtx, vt typedef.VisitTarget) (pass bool) {
	const node = "VisitCondition"
	flag := false
	switch cdtCtx.(type) {
	case parser.ICondEntityContext, parser.IConditionGroupContext:
		flag = true
	default:
	}
	if !flag {
		a.errorCallback(node, "invalid condition context type")
		return
	}
	// assignment data context and visit target
	a.dataCtx = dCtx
	a.visitTarget = vt

	// visit condition
	return a.Visit(cdtCtx).(bool)
}

func (a *arishemRuleVisitor) VisitAim(aimCtx antlr.ParseTree, dCtx typedef.DataCtx, vt typedef.VisitTarget) (aim typedef.Aim) {
	const node = "VisitAim"
	flag := false
	switch aimCtx.(type) {
	case parser.IAimEntityContext, parser.IAimContext:
		flag = true
	default:
	}
	if !flag {
		a.errorCallback(node, "invalid aim context type")
		return
	}
	// assignment data context and visit target
	a.dataCtx = dCtx
	a.visitTarget = vt

	// visit aim
	v := a.Visit(aimCtx)
	if v != nil {
		aim = v.(typedef.Aim)
	}
	return
}

func (a *arishemRuleVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	// cache expression
	case parser.IExprContext:
		// try get from cache
		exprVal, ok := a.dataCtx.Get(t.GetAltNumber())
		if !ok {
			exprVal = tree.Accept(a)
			a.dataCtx.Set(t.GetAltNumber(), exprVal)
		} /*else {
			// remove debug log
			//logger.Infof("cache hit in visit!!! text=>%s", tree.GetText())
		}*/
		return exprVal
	}
	return tree.Accept(a)
}

func (a *arishemRuleVisitor) VisitArishemCond(ctx *parser.ArishemCondContext) interface{} {
	return a.Visit(ctx.ConditionGroup())
}

func (a *arishemRuleVisitor) VisitArishemAim(ctx *parser.ArishemAimContext) interface{} {
	return a.Visit(ctx.Aim())
}

func (a *arishemRuleVisitor) VisitNestConditionGroup(ctx *parser.NestConditionGroupContext) interface{} {
	// cdtGroups contains at least one condition
	cdtGroups := ctx.ConditionGroupsPair().ConditionGroupsVal().AllConditionGroup()
	opLogic := toStdLogic(ctx.OpLogicPair().OpLogicVal().OpLogic().GetText())
	var pass bool
	switch opLogic {
	case logicAnd:
		pass = true
		// execute conditions orderly
		for i := range cdtGroups {
			if pass = a.Visit(cdtGroups[i]).(bool); !pass {
				break
			}
		}
	case logicOr:
		pass = false
		// execute conditions orderly
		for i := range cdtGroups {
			if pass = a.Visit(cdtGroups[i]).(bool); pass {
				break
			}
		}
	case logicNot:
		pass = !(a.Visit(cdtGroups[0]).(bool))
	}
	return pass
}

func (a *arishemRuleVisitor) VisitLeafConditionGroup(ctx *parser.LeafConditionGroupContext) interface{} {
	// condition contains at least one condition
	cdts := ctx.ConditionsPair().ConditionsVal().AllCondition()
	opLogic := toStdLogic(ctx.OpLogicPair().OpLogicVal().OpLogic().GetText())
	var pass bool
	switch opLogic {
	case logicAnd:
		pass = true
		// execute condition orderly
		for i := range cdts {
			if pass = a.Visit(cdts[i]).(bool); !pass {
				break
			}
		}
	case logicOr:
		pass = false
		// execute condition orderly
		for i := range cdts {
			if pass = a.Visit(cdts[i]).(bool); pass {
				break
			}
		}
	case logicNot:
		pass = !(a.Visit(cdts[0]).(bool))
	}
	return pass
}

func (a *arishemRuleVisitor) VisitNullConditionGroup(_ *parser.NullConditionGroupContext) interface{} {
	return false
}

func (a *arishemRuleVisitor) VisitFullCondition(ctx *parser.FullConditionContext) interface{} {
	const node = "VisitFullCondition"
	// try finding cache in DataCtx
	if cdtCacheI, ok := a.dataCtx.Get(ctx.GetAltNumber()); ok {
		var cdtCc typedef.JudgeNode
		if cdtC, convO := cdtCacheI.(typedef.JudgeNode); convO {
			cdtCc = cdtC
		}
		for _, obsr := range a.observers {
			obsr.OnJudgeNodeVisitEnd(a.dataCtx.Context(), cdtCc, a.visitTarget)
		}
		// remove debug log
		//logger.Infof("cache hit in full condition!!! text=>%s", ctx.GetText())
		return cdtCc.Passed()
	}
	// get operator
	opr := ctx.OperatorPair().OperatorVal().OperatorType().GetText()
	// get left value expr
	lhsExpr := ctx.LhsPair().Expr()
	rhsExpr := ctx.RhsPair().Expr()
	// visit left and right
	left := a.Visit(lhsExpr)
	right := a.Visit(rhsExpr)

	var pass = false
	var err error
	var cdtInfo typedef.JudgeNode

	if !strings.HasPrefix(opr, foreachKey) {
		pass, err = evaluateLR(a, left, right, opr)
		cdtInfo = newArishemCondition(pass, left, lhsExpr.GetText(), right, rhsExpr.GetText(), opr, err)
	} else {
		// parse operator and logic
		const space = " "
		firstSpace := strings.Index(opr, space)
		lastSpace := strings.LastIndex(opr, space)
		foreachOpr := opr[firstSpace+1 : lastSpace]
		logicOpr := toStdLogic(opr[lastSpace+1:])
		// execute the foreach condition
		var foreachItems []typedef.ForeachItem
		foreachItems, pass, err = a.visitForeachConditionInner(left, right, foreachOpr, logicOpr)
		cdtInfo = newArishemForeachCondition(pass, left, lhsExpr.GetText(), right, rhsExpr.GetText(), opr, err, foreachOpr, logicOpr, foreachItems)
	}
	if err != nil {
		a.errorCallback(node, err.Error())
	}
	// put it into cache
	a.dataCtx.Set(ctx.GetAltNumber(), cdtInfo)
	// observers call back
	for _, obsr := range a.observers {
		obsr.OnJudgeNodeVisitEnd(a.dataCtx.Context(), cdtInfo, a.visitTarget)
	}
	return pass
}

// execute the foreach operator judgement
func (a *arishemRuleVisitor) visitForeachConditionInner(left, right interface{}, foreachOpr, logicOpr string) ([]typedef.ForeachItem, bool, error) {
	if strings.HasPrefix(foreachOpr, foreachKey) {
		return nil, false, errors.New("nested foreach operation is not supported")
	}
	// check the left type whether array or not
	lArray, err := tools.ConvToSliceUnifyType(left)
	if err != nil {
		return nil, false, errors.New("foreach operation is supported only when left is slice/array type")
	}
	// parse the operator, arishem grammar guarantee that here are at least two space
	items := make([]typedef.ForeachItem, 0, len(lArray)/2+1)
	pass := false
	for i, lVal := range lArray {
		pass, err = evaluateLR(a, lVal, right, foreachOpr)

		items = append(items, newArishemForeachItem(i, lVal, pass))

		if err != nil {
			break
		}
		if logicOpr == logicAnd && pass == false {
			break
		}
		if logicOpr == logicOr && pass == true {
			break
		}
	}
	return items, pass, err
}

func (a *arishemRuleVisitor) VisitNullCondition(_ *parser.NullConditionContext) interface{} {
	return false
}

func (a *arishemRuleVisitor) VisitActionAimList(ctx *parser.ActionAimListContext) interface{} {
	aimTypes := ctx.AllActionAimType()
	if len(aimTypes) <= 0 {
		return newArishemAim(make([]typedef.ActionAim, 0), aimTypeActionList)
	}
	var actionList []typedef.ActionAim
	for _, actCtx := range aimTypes {
		actAim := a.Visit(actCtx)
		if actAim != nil {
			actionList = append(actionList, actAim.(typedef.ActionAim))
		}
	}
	return newArishemAim(actionList, aimTypeActionList)
}

func (a *arishemRuleVisitor) VisitActionAim(ctx *parser.ActionAimContext) interface{} {
	return newArishemAim(a.Visit(ctx.ActionAimType()), aimTypeAction)
}

func (a *arishemRuleVisitor) VisitExprAim(ctx *parser.ExprAimContext) interface{} {
	return newArishemAim(a.Visit(ctx.ExprType()), aimTypeExpr)
}

func (a *arishemRuleVisitor) VisitNullAim(_ *parser.NullAimContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitParamMapAction(ctx *parser.ParamMapActionContext) interface{} {
	// get action name
	actionName := ctx.ActionAimPair().ActionAimVal().NAME_KEY_TYPE().GetText()
	// check if param map exist
	var paramMap interface{}
	if ctx.KeyValues() != nil {
		paramMap = a.Visit(ctx.KeyValues())
	}
	return newArishemActionAim(actionName, paramMap, paramTypeMap)
}

func (a *arishemRuleVisitor) VisitParamListAction(ctx *parser.ParamListActionContext) interface{} {
	// get action name
	actionName := ctx.ActionAimPair().ActionAimVal().NAME_KEY_TYPE().GetText()
	// check if param list exist
	var paramList interface{}
	if ctx.Exprs() != nil {
		paramList = a.Visit(ctx.Exprs())
	}
	return newArishemActionAim(actionName, paramList, paramTypeList)
}

func (a *arishemRuleVisitor) VisitFullExpr(ctx *parser.FullExprContext) interface{} {
	return a.Visit(ctx.ExprType())
}

func (a *arishemRuleVisitor) VisitNullExpr(_ *parser.NullExprContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitListExprType(ctx *parser.ListExprTypeContext) interface{} {
	return a.Visit(ctx.Exprs())
}

func (a *arishemRuleVisitor) VisitMapExprType(ctx *parser.MapExprTypeContext) interface{} {
	return a.Visit(ctx.KeyValues())
}

func (a *arishemRuleVisitor) VisitConstExprType(ctx *parser.ConstExprTypeContext) interface{} {
	return a.Visit(ctx.Constant())
}

func (a *arishemRuleVisitor) VisitConstListExprType(ctx *parser.ConstListExprTypeContext) interface{} {
	return a.Visit(ctx.ConstantList())
}

func (a *arishemRuleVisitor) VisitMathExprType(ctx *parser.MathExprTypeContext) interface{} {
	return a.Visit(ctx.Math())
}

func (a *arishemRuleVisitor) VisitFuncExprType(ctx *parser.FuncExprTypeContext) interface{} {
	return a.Visit(ctx.Function())
}

func (a *arishemRuleVisitor) VisitVarExprType(ctx *parser.VarExprTypeContext) interface{} {
	return a.Visit(ctx.Variable())
}

func (a *arishemRuleVisitor) VisitFeatureExprType(ctx *parser.FeatureExprTypeContext) interface{} {
	return a.Visit(ctx.Feature())
}

func (a *arishemRuleVisitor) VisitNullExprType(_ *parser.NullExprTypeContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullVar(ctx *parser.FullVarContext) interface{} {
	const node = "VisitFullVar"
	// 1. first check path whether contains double sharp '##'
	varPathStr := ctx.VarPath().VarPathVal().GetText()
	twinSharpCheck := strings.Split(varPathStr, "##")
	var val interface{}
	var err error
	if len(twinSharpCheck) > 1 {
		var elements []interface{}

		slicePathStr := twinSharpCheck[0]
		elementPathStr := twinSharpCheck[1]

		slicePath := strings.Split(slicePathStr, strDot)
		var slice interface{}
		slice, err = a.dataCtx.GetVarValue(slicePath)
		if err == nil {
			if arr, ok := slice.([]interface{}); ok {
				for _, eleObjI := range arr {
					var ele map[string]interface{}
					ele, ok = eleObjI.(map[string]interface{})
					if !ok {
						err = errors.New("var path use double sharp inner index of a wrong type")
						break
					}
					var element interface{}
					element, err = tools.MapIndex(ele, elementPathStr)
					if err != nil {
						break
					}
					elements = append(elements, element)
				}
			} else {
				err = errors.New("var path use double sharp index of a wrong type")
			}
		}
		// assign elements into val
		val = elements
	} else {
		varPath := strings.Split(twinSharpCheck[0], strDot)
		val, err = a.dataCtx.GetVarValue(varPath)
	}
	if err != nil {
		a.errorCallback(node, err.Error())
	}
	return val
}

func (a *arishemRuleVisitor) VisitNullVar(_ *parser.NullVarContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullMath(ctx *parser.FullMathContext) interface{} {
	const node = "VisitFullMath"
	// get math operator
	mathOp := ctx.LrOpMathPair().LrOpMathVal().LrMathOperator().GetText()
	// get left value
	lhsExpr := ctx.LhsPair().Expr()
	rhsExpr := ctx.RhsPair().Expr()
	left := a.Visit(lhsExpr)
	if left == nil {
		a.errorCallback(node, "math operator: left is nil, left expr=>"+lhsExpr.GetText())
		return nil
	}
	right := a.Visit(rhsExpr)
	if right == nil {
		return left
	}
	if (mathOp == mathDivide || mathOp == mathMod) && fmt.Sprint(right) == "0" {
		a.errorCallback(node, "math operator illegal, divide or mod zero is not allowed")
		return nil
	}
	if mathOp == operator.Equal || mathOp == operator.NotEqual || mathOp == operator.Greater ||
		mathOp == operator.Less || mathOp == operator.GreaterOrEqual || mathOp == operator.LessOrEqual {
		b, err := operator.Judge(left, right, mathOp)
		if err != nil {
			a.errorCallback(node, "math logic calculate err=>"+err.Error())
			return nil
		}
		return b
	}
	val, err := calculator.Calculate(fmt.Sprint(left) + mathOp + fmt.Sprint(right))
	if err != nil {
		a.errorCallback(node, "calculate math expression err=>"+err.Error())
		return nil
	}
	return val
}

func (a *arishemRuleVisitor) VisitListMath(ctx *parser.ListMathContext) interface{} {
	const node = "VisitListMath"
	// get math operator
	mathOp := ctx.ListOpMathPair().ListOpMathVal().ListMathOperator().GetText()
	// get left value
	vals := a.Visit(ctx.Exprs())
	if vals == nil {
		a.errorCallback(node, "visit list math error: expr list is empty")
		return nil
	}
	exprList := vals.([]interface{})
	if len(exprList) <= 1 {
		return exprList[0]
	}
	// first left value
	left := exprList[0]
	for i := 1; i < len(exprList); i++ {
		right := exprList[i]
		if right == nil {
			continue
		}
		if (mathOp == mathDivide || mathOp == mathMod) && fmt.Sprint(right) == "0" {
			a.errorCallback(node, "math operator illegal, divide or mod zero is not allowed")
			continue
		}
		var err error
		left, err = calculator.Calculate(fmt.Sprint(left) + mathOp + fmt.Sprint(right))
		if err != nil {
			a.errorCallback(node, "calculate math expression err=>"+err.Error())
			return nil
		}
	}
	return left
}

func (a *arishemRuleVisitor) VisitNullMath(_ *parser.NullMathContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitParamMapFunc(ctx *parser.ParamMapFuncContext) interface{} {
	const node = "VisitParamMapFunc"
	// get function name
	funcname := ctx.FuncNamePair().NameKey().NAME_KEY_TYPE().GetText()
	// param map
	var paramMap map[string]interface{}
	kvCtx := ctx.KeyValues()
	if kvCtx != nil {
		paramI := a.Visit(kvCtx)
		if paramI != nil {
			paramMap = paramI.(map[string]interface{})
		}
	}
	res, err := funcs.InvokeFunc(a.dataCtx.Context(), funcname, paramMap, nil)
	if err != nil {
		a.errorCallback(node, err.Error())
	}
	return res
}

func (a *arishemRuleVisitor) VisitParamListFunc(ctx *parser.ParamListFuncContext) interface{} {
	const node = "VisitParamListFunc"
	// get function name
	funcname := ctx.FuncNamePair().NameKey().NAME_KEY_TYPE().GetText()
	// param list
	var paramList []interface{}
	exprsCtx := ctx.Exprs()
	if exprsCtx != nil {
		paramI := a.Visit(exprsCtx)
		if paramI != nil {
			paramList = paramI.([]interface{})
		}
	}
	res, err := funcs.InvokeFunc(a.dataCtx.Context(), funcname, nil, paramList)
	if err != nil {
		a.errorCallback(node, err.Error())
	}
	return res
}

func (a *arishemRuleVisitor) VisitNullFunc(_ *parser.NullFuncContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullExprs(ctx *parser.FullExprsContext) interface{} {
	var exprs []interface{}
	for _, expr := range ctx.AllExpr() {
		exprs = append(exprs, a.Visit(expr))
	}
	return exprs
}

func (a *arishemRuleVisitor) VisitNullExprs(_ *parser.NullExprsContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullFeature(ctx *parser.FullFeatureContext) interface{} {
	const node = "VisitFullFeature"
	// get feature name and field path, feature path at least contains two name
	featPath := strings.Split(ctx.FeaturePath().FeaturePathVal().GetText(), strDot)
	// it should be prevented from token verified period
	if len(featPath) <= 1 {
		a.errorCallback(node, "feature path must contain at least one field path")
		return nil
	}
	// check if feature carries parameters
	var builtinParam map[string]interface{}
	kvCtx := ctx.KeyValues()
	if kvCtx != nil {
		paramVisit := a.Visit(kvCtx)
		if paramVisit != nil {
			builtinParam = paramVisit.(map[string]interface{})
		}
	}
	featVal, err := a.dataCtx.GetFeatureValue(newArishemFeatureParam(featPath[0], builtinParam), featPath[1:]...)
	if err != nil {
		a.errorCallback(node, err.Error())
	}
	return featVal
}

func (a *arishemRuleVisitor) VisitNullFeature(_ *parser.NullFeatureContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitNumConst(ctx *parser.NumConstContext) interface{} {
	const node = "VisitNumConst"
	numStr := ctx.NumberType().GetText()
	if tools.IsBlank(numStr) {
		a.errorCallback(node, "number text is empty")
		return nil
	}
	num, err := tools.StringToNumber(numStr)
	if err != nil {
		a.errorCallback(node, err.Error())
		return nil
	}
	return num
}

func (a *arishemRuleVisitor) VisitStrConst(ctx *parser.StrConstContext) interface{} {
	return tools.UnescapeString(StrippingQuotation(ctx.StringType().GetText()))
}

func (a *arishemRuleVisitor) VisitBoolConst(ctx *parser.BoolConstContext) interface{} {
	boolStr := ctx.BoolType().GetText()
	bv, _ := strconv.ParseBool(boolStr)
	return bv
}

func (a *arishemRuleVisitor) VisitNullConst(ctx *parser.NullConstContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullConstList(ctx *parser.FullConstListContext) interface{} {
	// at least one constant
	var consts []interface{}
	for _, constCtx := range ctx.AllConstant() {
		consts = append(consts, a.Visit(constCtx))
	}
	return consts
}

func (a *arishemRuleVisitor) VisitNullConstList(_ *parser.NullConstListContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitFullKeyValues(ctx *parser.FullKeyValuesContext) interface{} {
	kv := make(map[string]interface{}, 2)
	for _, kvCtx := range ctx.AllKeyValue() {
		kv[kvCtx.NameKey().NAME_KEY_TYPE().GetText()] = a.Visit(kvCtx.Expr())
	}
	return kv
}

func (a *arishemRuleVisitor) VisitNullKeyValues(_ *parser.NullKeyValuesContext) interface{} {
	return nil
}

func (a *arishemRuleVisitor) VisitSubCondExprType(ctx *parser.SubCondExprTypeContext) interface{} {
	return a.Visit(ctx.SubCond())
}

type FullSubCond struct {
	name string
	tree antlr.ParseTree
}

func (a *arishemRuleVisitor) VisitFullSubCond(ctx *parser.FullSubCondContext) interface{} {
	const node = "VisitFullSubCond"
	// 1. get the condition name
	condName := ctx.NameKey().NAME_KEY_TYPE().GetText()
	if tools.IsBlank(condName) {
		a.errorCallback(node, "sub condition name is empty")
		return nil
	}
	// 2. get then condition parse tree
	if a.condFinder == nil {
		a.errorCallback(node, "condition finder not set while using sub condition feature")
		return nil
	}
	condTree, err := a.condFinder(condName)
	if err != nil {
		a.errorCallback(node, "get the condition tree failed: "+err.Error())
		return nil
	}
	return &FullSubCond{
		name: condName,
		tree: condTree,
	}
}

func (a *arishemRuleVisitor) VisitNullSubCond(_ *parser.NullSubCondContext) interface{} {
	return nil
}

func evaluateLR(ori *arishemRuleVisitor, left, right interface{}, opr string) (bool, error) {
	var pass bool
	var err error
	if strings.Contains(opr, "SUB_COND") {
		lData, ok := left.(map[string]interface{})
		if !ok {
			return false, errors.New("sub condition left type is not object")
		}
		rSubCond, ok := right.(*FullSubCond)
		if !ok {
			return false, errors.New("sub condition right type is not sub condition expression")
		}
		pass, err = evaluateSubCond(ori, rSubCond.name, lData, rSubCond.tree)
		// if err is not nil, Pass will have to be false
		if err == nil && (strings.HasPrefix(opr, "!") || strings.HasPrefix(opr, "NOT")) {
			// revert the judge result
			pass = !pass
		}
	} else {
		// judge the condition by opr with left and right
		pass, err = operator.Judge(left, right, opr)
	}
	return pass, err
}

func evaluateSubCond(ori *arishemRuleVisitor, subCondName string, left map[string]interface{}, right antlr.ParseTree) (bool, error) {
	newDataCtx, err := NewArishemDataCtxFromMeta(ori.dataCtx.Context(), left, ori.dataCtx.FetcherFetcher())
	if err != nil {
		return false, err
	}
	newVisitor := &arishemRuleVisitor{
		BasearishemVisitor: &parser.BasearishemVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		},
		observers:  ori.observers,
		condFinder: ori.condFinder,
	}
	pass := newVisitor.VisitCondition(right, newDataCtx, &simpleVisitTarget{name: subCondName})
	return pass, nil
}
