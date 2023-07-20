// Code generated from arishem.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // arishem

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BasearishemVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasearishemVisitor) VisitArishemCond(ctx *ArishemCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitArishemAim(ctx *ArishemAimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNestConditionGroup(ctx *NestConditionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLeafConditionGroup(ctx *LeafConditionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullConditionGroup(ctx *NullConditionGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullCondition(ctx *FullConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullCondition(ctx *NullConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitActionAimList(ctx *ActionAimListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitActionAim(ctx *ActionAimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitExprAim(ctx *ExprAimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullAim(ctx *NullAimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamMapAction(ctx *ParamMapActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamListAction(ctx *ParamListActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullExpr(ctx *FullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullExpr(ctx *NullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListExprType(ctx *ListExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitMapExprType(ctx *MapExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConstExprType(ctx *ConstExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConstListExprType(ctx *ConstListExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitMathExprType(ctx *MathExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFuncExprType(ctx *FuncExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitVarExprType(ctx *VarExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFeatureExprType(ctx *FeatureExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullExprType(ctx *NullExprTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullVar(ctx *FullVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullVar(ctx *NullVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullMath(ctx *FullMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListMath(ctx *ListMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullMath(ctx *NullMathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamMapFunc(ctx *ParamMapFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamListFunc(ctx *ParamListFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullFunc(ctx *NullFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullExprs(ctx *FullExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullExprs(ctx *NullExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullFeature(ctx *FullFeatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullFeature(ctx *NullFeatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNumConst(ctx *NumConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitStrConst(ctx *StrConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitBoolConst(ctx *BoolConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullConst(ctx *NullConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullConstList(ctx *FullConstListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullConstList(ctx *NullConstListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionsKey(ctx *ConditionsKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionsVal(ctx *ConditionsValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionsPair(ctx *ConditionsPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionGroupsKey(ctx *ConditionGroupsKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionGroupsVal(ctx *ConditionGroupsValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConditionGroupsPair(ctx *ConditionGroupsPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListExprKey(ctx *ListExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitMapExprKey(ctx *MapExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitMathExprKey(ctx *MathExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFuncExprKey(ctx *FuncExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitVarExprKey(ctx *VarExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFeatureExprKey(ctx *FeatureExprKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFeaturePathKey(ctx *FeaturePathKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConstKey(ctx *ConstKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitConstListKey(ctx *ConstListKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitVarPath(ctx *VarPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitVarPathVal(ctx *VarPathValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFeaturePath(ctx *FeaturePathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFeaturePathVal(ctx *FeaturePathValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamMapKey(ctx *ParamMapKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitParamListKey(ctx *ParamListKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitBuiltinParamKey(ctx *BuiltinParamKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitActionAimKey(ctx *ActionAimKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitActionAimVal(ctx *ActionAimValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitActionAimPair(ctx *ActionAimPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNameKey(ctx *NameKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFullKeyValues(ctx *FullKeyValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNullKeyValues(ctx *NullKeyValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOpLogicKey(ctx *OpLogicKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOpLogicVal(ctx *OpLogicValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOpLogicPair(ctx *OpLogicPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOpLogic(ctx *OpLogicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOpMathKey(ctx *OpMathKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLrOpMathVal(ctx *LrOpMathValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListOpMathVal(ctx *ListOpMathValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLrOpMathPair(ctx *LrOpMathPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListOpMathPair(ctx *ListOpMathPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitListMathOperator(ctx *ListMathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLrMathOperator(ctx *LrMathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLogicMathOperator(ctx *LogicMathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFuncNameKey(ctx *FuncNameKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitFuncNamePair(ctx *FuncNamePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOperatorKey(ctx *OperatorKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOperatorVal(ctx *OperatorValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOperatorPair(ctx *OperatorPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOperatorType(ctx *OperatorTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLhsKey(ctx *LhsKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitLhsPair(ctx *LhsPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitRhsKey(ctx *RhsKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitRhsPair(ctx *RhsPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNumberType(ctx *NumberTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitSpecialChars(ctx *SpecialCharsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitStrConstKey(ctx *StrConstKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitNumConstKey(ctx *NumConstKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitBoolConstKey(ctx *BoolConstKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasearishemVisitor) VisitOperator(ctx *OperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
