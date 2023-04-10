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
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/internal/tools"
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/mnogu/go-calculator"
	"strconv"
	"strings"
	"sync"
)

var (
	featVisitCache typedef.SharedVisitCache
	featVCOnce     = &sync.Once{}
)

func SetFeaturePreParseCache(sharedVC typedef.SharedVisitCache) {
	if sharedVC == nil {
		panic(errors.New("feature pre-parse shared visit cache is nil"))
	}
	featVCOnce.Do(func() {
		featVisitCache = sharedVC
	})
}

func NewFeaturePreParseListener() typedef.FeaturePreParser {
	return &featurePreParseListener{
		featurePreParserVisitor: &featurePreParserVisitor{
			BasearishemVisitor: &parser.BasearishemVisitor{
				BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
			},
			errMsgs: tools.NewStringSet(),
		},
		featsParam: tools.NewHashSet(),
	}
}

type featurePreParseListener struct {
	*featurePreParserVisitor
	featsParam typedef.Set
}

func (f *featurePreParseListener) Data() []typedef.FeatureParam {
	var data []typedef.FeatureParam
	for _, fp := range f.featsParam.AsList() {
		data = append(data, fp.(typedef.FeatureParam))
	}
	return data
}

func (f *featurePreParseListener) Error() error {
	errMsgs := f.errMsgs.AsList()
	if len(errMsgs) <= 0 {
		return nil
	}
	return errors.New(strings.Join(errMsgs, ","))
}

func (f *featurePreParseListener) VisitTerminal(node antlr.TerminalNode, tokenErrMsgs, parseErrMsgs []string) {
}

func (f *featurePreParseListener) VisitErrorNode(node antlr.ErrorNode, tokenErrMsgs, parseErrMsgs []string) {
}

func (f *featurePreParseListener) EnterEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string) {
}

func (f *featurePreParseListener) ExitEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string) {
	if len(tokenErrMsgs) > 0 || len(parseErrMsgs) > 0 {
		return
	}
	switch ctx.(type) {
	case parser.IFeatureContext:
		val := f.Visit(ctx)
		if val != nil {
			fp, ok := val.(typedef.FeatureParam)
			if ok {
				f.featsParam.Add(fp)
			}
		}
	}
}

type featurePreParserVisitor struct {
	*parser.BasearishemVisitor
	errMsgs *tools.StringSet
}

func (f *featurePreParserVisitor) Visit(tree antlr.ParseTree) interface{} {
	// try to get value from cache first
	var key, cache interface{}
	var ok bool
	switch t := tree.(type) {
	case antlr.RuleContext:
		key = t.GetAltNumber()
		cache, ok = featVisitCache.Get(key)
	default:
		key = t.GetText()
		cache, ok = featVisitCache.Get(key)
	}
	if !ok {
		cache = tree.Accept(f)
		featVisitCache.Set(key, cache)
	} else {
		// remove debug info
		//logger.Info("[feature cache hit]")
	}
	return cache
}

func (f *featurePreParserVisitor) VisitFullFeature(ctx *parser.FullFeatureContext) interface{} {
	const node = "VisitFullFeature"
	// get feature name and field path, feature path at least contains two name
	featPath := strings.Split(ctx.FeaturePath().FeaturePathVal().GetText(), strDot)
	// it should be prevented from token verified period
	if len(featPath) <= 1 {
		f.errMsgs.Add(node + ": feature path must contain at least one field path")
		return nil
	}
	// check if feature carries parameters
	var builtinParam map[string]interface{}
	kvCtx := ctx.KeyValues()
	if kvCtx != nil {
		paramVisit := f.Visit(kvCtx)
		if paramVisit != nil {
			builtinParam = paramVisit.(map[string]interface{})
		}
	}
	// return feature param
	return newArishemFeatureParam(featPath[0], builtinParam)
}

func (f *featurePreParserVisitor) VisitNullFeature(ctx *parser.NullFeatureContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitFullExpr(ctx *parser.FullExprContext) interface{} {
	return f.Visit(ctx.ExprType())
}

func (f *featurePreParserVisitor) VisitNullExpr(ctx *parser.NullExprContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitListExprType(ctx *parser.ListExprTypeContext) interface{} {
	return f.Visit(ctx.Exprs())
}

func (f *featurePreParserVisitor) VisitMapExprType(ctx *parser.MapExprTypeContext) interface{} {
	return f.Visit(ctx.KeyValues())
}

func (f *featurePreParserVisitor) VisitConstExprType(ctx *parser.ConstExprTypeContext) interface{} {
	return f.Visit(ctx.Constant())
}

func (f *featurePreParserVisitor) VisitConstListExprType(ctx *parser.ConstListExprTypeContext) interface{} {
	return f.Visit(ctx.ConstantList())
}

func (f *featurePreParserVisitor) VisitMathExprType(ctx *parser.MathExprTypeContext) interface{} {
	return f.Visit(ctx.Math())
}

func (f *featurePreParserVisitor) VisitFuncExprType(ctx *parser.FuncExprTypeContext) interface{} {
	return f.Visit(ctx.Function())
}

func (f *featurePreParserVisitor) VisitVarExprType(ctx *parser.VarExprTypeContext) interface{} {
	return f.Visit(ctx.Variable())
}

func (f *featurePreParserVisitor) VisitFeatureExprType(ctx *parser.FeatureExprTypeContext) interface{} {
	return f.Visit(ctx.Feature())
}

func (f *featurePreParserVisitor) VisitNullExprType(ctx *parser.NullExprTypeContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitNullVar(ctx *parser.NullVarContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitFullMath(ctx *parser.FullMathContext) interface{} {
	const node = "VisitFullMath"
	// get math operator
	mathOp := ctx.OpMathPair().OpMathVal().MathOperator().GetText()
	// get left value
	lhsExpr := ctx.LhsPair().Expr()
	rhsExpr := ctx.RhsPair().Expr()
	left := f.Visit(lhsExpr)
	if left == nil {
		f.errMsgs.Add(node + ":math operator: left is nil, left expr=>" + lhsExpr.GetText())
		return nil
	}
	right := f.Visit(rhsExpr)
	if right == nil {
		return left
	}
	if (mathOp == mathDivide || mathOp == mathMod) && fmt.Sprint(right) == "0" {
		f.errMsgs.Add(node + ":math operator illegal, divide or mod zero is not allowed")
		return nil
	}

	val, err := calculator.Calculate(fmt.Sprint(left) + mathOp + fmt.Sprint(right))
	if err != nil {
		f.errMsgs.Add(node + ":calculate math expression err=>" + err.Error())
		return nil
	}
	return val
}

func (f *featurePreParserVisitor) VisitListMath(ctx *parser.ListMathContext) interface{} {
	const node = "VisitListMath"
	// get math operator
	mathOp := ctx.OpMathPair().OpMathVal().MathOperator().GetText()
	// get left value
	vals := f.Visit(ctx.Exprs())
	if vals == nil {
		f.errMsgs.Add(node + ":visit list math error: expr list is empty")
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
			f.errMsgs.Add(node + ":math operator illegal, divide or mod zero is not allowed")
			continue
		}

		var err error
		left, err = calculator.Calculate(fmt.Sprint(left) + mathOp + fmt.Sprint(right))
		if err != nil {
			f.errMsgs.Add(node + ":calculate math expression err=>" + err.Error())
			return nil
		}
	}
	return left
}

func (f *featurePreParserVisitor) VisitNullMath(ctx *parser.NullMathContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitParamMapFunc(ctx *parser.ParamMapFuncContext) interface{} {
	const node = "VisitParamMapFunc"
	// get function name
	funcname := ctx.FuncNamePair().NameKey().NAME_KEY_TYPE().GetText()
	// param map
	var paramMap map[string]interface{}
	kvCtx := ctx.KeyValues()
	if kvCtx != nil {
		paramI := f.Visit(kvCtx)
		if paramI != nil {
			paramMap = paramI.(map[string]interface{})
		}
	}
	res, err := funcs.InvokeFunc(funcname, paramMap, nil)
	if err != nil {
		f.errMsgs.Add(node + ":" + err.Error())
		return nil
	}
	return res
}

func (f *featurePreParserVisitor) VisitParamListFunc(ctx *parser.ParamListFuncContext) interface{} {
	const node = "VisitParamListFunc"
	// get function name
	funcname := ctx.FuncNamePair().NameKey().NAME_KEY_TYPE().GetText()
	// param map
	var paramList []interface{}
	exprsCtx := ctx.Exprs()
	if exprsCtx != nil {
		paramI := f.Visit(exprsCtx)
		if paramI != nil {
			paramList = paramI.([]interface{})
		}
	}
	res, err := funcs.InvokeFunc(funcname, nil, paramList)
	if err != nil {
		f.errMsgs.Add(node + ":" + err.Error())
		return nil
	}
	return res
}

func (f *featurePreParserVisitor) VisitNullFunc(ctx *parser.NullFuncContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitFullExprs(ctx *parser.FullExprsContext) interface{} {
	var exprs []interface{}
	for _, expr := range ctx.AllExpr() {
		exprs = append(exprs, f.Visit(expr))
	}
	return exprs
}

func (f *featurePreParserVisitor) VisitNullExprs(ctx *parser.NullExprsContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitNumConst(ctx *parser.NumConstContext) interface{} {
	const node = "VisitNumConst"
	numStr := ctx.NumberType().GetText()
	if tools.IsBlank(numStr) {
		f.errMsgs.Add(node + ":number text is empty")
		return nil
	}
	num, err := tools.StringToNumber(numStr)
	if err != nil {
		f.errMsgs.Add(node + ":" + err.Error())
		return nil
	}
	return num
}

func (f *featurePreParserVisitor) VisitStrConst(ctx *parser.StrConstContext) interface{} {
	return StrippingQuotation(ctx.StringType().GetText())
}

func (f *featurePreParserVisitor) VisitBoolConst(ctx *parser.BoolConstContext) interface{} {
	boolStr := ctx.BoolType().GetText()
	bv, _ := strconv.ParseBool(boolStr)
	return bv
}

func (f *featurePreParserVisitor) VisitNullConst(ctx *parser.NullConstContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitFullConstList(ctx *parser.FullConstListContext) interface{} {
	// at least one constant
	var consts []interface{}
	for _, constCtx := range ctx.AllConstant() {
		consts = append(consts, f.Visit(constCtx))
	}
	return consts
}

func (f *featurePreParserVisitor) VisitNullConstList(ctx *parser.NullConstListContext) interface{} {
	return nil
}

func (f *featurePreParserVisitor) VisitFullKeyValues(ctx *parser.FullKeyValuesContext) interface{} {
	kv := make(map[string]interface{}, 2)
	for _, kvCtx := range ctx.AllKeyValue() {
		kv[kvCtx.NameKey().NAME_KEY_TYPE().GetText()] = f.Visit(kvCtx.Expr())
	}
	return kv
}

func (f *featurePreParserVisitor) VisitNullKeyValues(ctx *parser.NullKeyValuesContext) interface{} {
	return nil
}
