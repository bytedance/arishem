// Code generated from arishem.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // arishem

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by arishemParser.
type arishemVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by arishemParser#arishemCond.
	VisitArishemCond(ctx *ArishemCondContext) interface{}

	// Visit a parse tree produced by arishemParser#arishemAim.
	VisitArishemAim(ctx *ArishemAimContext) interface{}

	// Visit a parse tree produced by arishemParser#nestConditionGroup.
	VisitNestConditionGroup(ctx *NestConditionGroupContext) interface{}

	// Visit a parse tree produced by arishemParser#leafConditionGroup.
	VisitLeafConditionGroup(ctx *LeafConditionGroupContext) interface{}

	// Visit a parse tree produced by arishemParser#nullConditionGroup.
	VisitNullConditionGroup(ctx *NullConditionGroupContext) interface{}

	// Visit a parse tree produced by arishemParser#fullCondition.
	VisitFullCondition(ctx *FullConditionContext) interface{}

	// Visit a parse tree produced by arishemParser#nullCondition.
	VisitNullCondition(ctx *NullConditionContext) interface{}

	// Visit a parse tree produced by arishemParser#actionAimList.
	VisitActionAimList(ctx *ActionAimListContext) interface{}

	// Visit a parse tree produced by arishemParser#actionAim.
	VisitActionAim(ctx *ActionAimContext) interface{}

	// Visit a parse tree produced by arishemParser#exprAim.
	VisitExprAim(ctx *ExprAimContext) interface{}

	// Visit a parse tree produced by arishemParser#nullAim.
	VisitNullAim(ctx *NullAimContext) interface{}

	// Visit a parse tree produced by arishemParser#paramMapAction.
	VisitParamMapAction(ctx *ParamMapActionContext) interface{}

	// Visit a parse tree produced by arishemParser#paramListAction.
	VisitParamListAction(ctx *ParamListActionContext) interface{}

	// Visit a parse tree produced by arishemParser#fullExpr.
	VisitFullExpr(ctx *FullExprContext) interface{}

	// Visit a parse tree produced by arishemParser#nullExpr.
	VisitNullExpr(ctx *NullExprContext) interface{}

	// Visit a parse tree produced by arishemParser#listExprType.
	VisitListExprType(ctx *ListExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#mapExprType.
	VisitMapExprType(ctx *MapExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#constExprType.
	VisitConstExprType(ctx *ConstExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#constListExprType.
	VisitConstListExprType(ctx *ConstListExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#mathExprType.
	VisitMathExprType(ctx *MathExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#funcExprType.
	VisitFuncExprType(ctx *FuncExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#varExprType.
	VisitVarExprType(ctx *VarExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#featureExprType.
	VisitFeatureExprType(ctx *FeatureExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#subCondExprType.
	VisitSubCondExprType(ctx *SubCondExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#nullExprType.
	VisitNullExprType(ctx *NullExprTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#fullVar.
	VisitFullVar(ctx *FullVarContext) interface{}

	// Visit a parse tree produced by arishemParser#nullVar.
	VisitNullVar(ctx *NullVarContext) interface{}

	// Visit a parse tree produced by arishemParser#fullMath.
	VisitFullMath(ctx *FullMathContext) interface{}

	// Visit a parse tree produced by arishemParser#listMath.
	VisitListMath(ctx *ListMathContext) interface{}

	// Visit a parse tree produced by arishemParser#nullMath.
	VisitNullMath(ctx *NullMathContext) interface{}

	// Visit a parse tree produced by arishemParser#paramMapFunc.
	VisitParamMapFunc(ctx *ParamMapFuncContext) interface{}

	// Visit a parse tree produced by arishemParser#paramListFunc.
	VisitParamListFunc(ctx *ParamListFuncContext) interface{}

	// Visit a parse tree produced by arishemParser#nullFunc.
	VisitNullFunc(ctx *NullFuncContext) interface{}

	// Visit a parse tree produced by arishemParser#fullExprs.
	VisitFullExprs(ctx *FullExprsContext) interface{}

	// Visit a parse tree produced by arishemParser#nullExprs.
	VisitNullExprs(ctx *NullExprsContext) interface{}

	// Visit a parse tree produced by arishemParser#fullFeature.
	VisitFullFeature(ctx *FullFeatureContext) interface{}

	// Visit a parse tree produced by arishemParser#nullFeature.
	VisitNullFeature(ctx *NullFeatureContext) interface{}

	// Visit a parse tree produced by arishemParser#fullSubCond.
	VisitFullSubCond(ctx *FullSubCondContext) interface{}

	// Visit a parse tree produced by arishemParser#nullSubCond.
	VisitNullSubCond(ctx *NullSubCondContext) interface{}

	// Visit a parse tree produced by arishemParser#numConst.
	VisitNumConst(ctx *NumConstContext) interface{}

	// Visit a parse tree produced by arishemParser#strConst.
	VisitStrConst(ctx *StrConstContext) interface{}

	// Visit a parse tree produced by arishemParser#boolConst.
	VisitBoolConst(ctx *BoolConstContext) interface{}

	// Visit a parse tree produced by arishemParser#nullConst.
	VisitNullConst(ctx *NullConstContext) interface{}

	// Visit a parse tree produced by arishemParser#fullConstList.
	VisitFullConstList(ctx *FullConstListContext) interface{}

	// Visit a parse tree produced by arishemParser#nullConstList.
	VisitNullConstList(ctx *NullConstListContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionsKey.
	VisitConditionsKey(ctx *ConditionsKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionsVal.
	VisitConditionsVal(ctx *ConditionsValContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionsPair.
	VisitConditionsPair(ctx *ConditionsPairContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionGroupsKey.
	VisitConditionGroupsKey(ctx *ConditionGroupsKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionGroupsVal.
	VisitConditionGroupsVal(ctx *ConditionGroupsValContext) interface{}

	// Visit a parse tree produced by arishemParser#conditionGroupsPair.
	VisitConditionGroupsPair(ctx *ConditionGroupsPairContext) interface{}

	// Visit a parse tree produced by arishemParser#listExprKey.
	VisitListExprKey(ctx *ListExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#mapExprKey.
	VisitMapExprKey(ctx *MapExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#mathExprKey.
	VisitMathExprKey(ctx *MathExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#funcExprKey.
	VisitFuncExprKey(ctx *FuncExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#varExprKey.
	VisitVarExprKey(ctx *VarExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#featureExprKey.
	VisitFeatureExprKey(ctx *FeatureExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#featurePathKey.
	VisitFeaturePathKey(ctx *FeaturePathKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#constKey.
	VisitConstKey(ctx *ConstKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#constListKey.
	VisitConstListKey(ctx *ConstListKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#subCondExprKey.
	VisitSubCondExprKey(ctx *SubCondExprKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#condNameKey.
	VisitCondNameKey(ctx *CondNameKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#varPath.
	VisitVarPath(ctx *VarPathContext) interface{}

	// Visit a parse tree produced by arishemParser#varPathVal.
	VisitVarPathVal(ctx *VarPathValContext) interface{}

	// Visit a parse tree produced by arishemParser#featurePath.
	VisitFeaturePath(ctx *FeaturePathContext) interface{}

	// Visit a parse tree produced by arishemParser#featurePathVal.
	VisitFeaturePathVal(ctx *FeaturePathValContext) interface{}

	// Visit a parse tree produced by arishemParser#paramMapKey.
	VisitParamMapKey(ctx *ParamMapKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#paramListKey.
	VisitParamListKey(ctx *ParamListKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#builtinParamKey.
	VisitBuiltinParamKey(ctx *BuiltinParamKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#actionAimKey.
	VisitActionAimKey(ctx *ActionAimKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#actionAimVal.
	VisitActionAimVal(ctx *ActionAimValContext) interface{}

	// Visit a parse tree produced by arishemParser#actionAimPair.
	VisitActionAimPair(ctx *ActionAimPairContext) interface{}

	// Visit a parse tree produced by arishemParser#nameKey.
	VisitNameKey(ctx *NameKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by arishemParser#fullKeyValues.
	VisitFullKeyValues(ctx *FullKeyValuesContext) interface{}

	// Visit a parse tree produced by arishemParser#nullKeyValues.
	VisitNullKeyValues(ctx *NullKeyValuesContext) interface{}

	// Visit a parse tree produced by arishemParser#opLogicKey.
	VisitOpLogicKey(ctx *OpLogicKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#opLogicVal.
	VisitOpLogicVal(ctx *OpLogicValContext) interface{}

	// Visit a parse tree produced by arishemParser#opLogicPair.
	VisitOpLogicPair(ctx *OpLogicPairContext) interface{}

	// Visit a parse tree produced by arishemParser#opLogic.
	VisitOpLogic(ctx *OpLogicContext) interface{}

	// Visit a parse tree produced by arishemParser#opMathKey.
	VisitOpMathKey(ctx *OpMathKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#lrOpMathVal.
	VisitLrOpMathVal(ctx *LrOpMathValContext) interface{}

	// Visit a parse tree produced by arishemParser#listOpMathVal.
	VisitListOpMathVal(ctx *ListOpMathValContext) interface{}

	// Visit a parse tree produced by arishemParser#lrOpMathPair.
	VisitLrOpMathPair(ctx *LrOpMathPairContext) interface{}

	// Visit a parse tree produced by arishemParser#listOpMathPair.
	VisitListOpMathPair(ctx *ListOpMathPairContext) interface{}

	// Visit a parse tree produced by arishemParser#listMathOperator.
	VisitListMathOperator(ctx *ListMathOperatorContext) interface{}

	// Visit a parse tree produced by arishemParser#lrMathOperator.
	VisitLrMathOperator(ctx *LrMathOperatorContext) interface{}

	// Visit a parse tree produced by arishemParser#logicMathOperator.
	VisitLogicMathOperator(ctx *LogicMathOperatorContext) interface{}

	// Visit a parse tree produced by arishemParser#funcNameKey.
	VisitFuncNameKey(ctx *FuncNameKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#funcNamePair.
	VisitFuncNamePair(ctx *FuncNamePairContext) interface{}

	// Visit a parse tree produced by arishemParser#operatorKey.
	VisitOperatorKey(ctx *OperatorKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#operatorVal.
	VisitOperatorVal(ctx *OperatorValContext) interface{}

	// Visit a parse tree produced by arishemParser#operatorPair.
	VisitOperatorPair(ctx *OperatorPairContext) interface{}

	// Visit a parse tree produced by arishemParser#operatorType.
	VisitOperatorType(ctx *OperatorTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#lhsKey.
	VisitLhsKey(ctx *LhsKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#lhsPair.
	VisitLhsPair(ctx *LhsPairContext) interface{}

	// Visit a parse tree produced by arishemParser#rhsKey.
	VisitRhsKey(ctx *RhsKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#rhsPair.
	VisitRhsPair(ctx *RhsPairContext) interface{}

	// Visit a parse tree produced by arishemParser#boolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#numberType.
	VisitNumberType(ctx *NumberTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by arishemParser#specialChars.
	VisitSpecialChars(ctx *SpecialCharsContext) interface{}

	// Visit a parse tree produced by arishemParser#strConstKey.
	VisitStrConstKey(ctx *StrConstKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#numConstKey.
	VisitNumConstKey(ctx *NumConstKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#boolConstKey.
	VisitBoolConstKey(ctx *BoolConstKeyContext) interface{}

	// Visit a parse tree produced by arishemParser#operator.
	VisitOperator(ctx *OperatorContext) interface{}
}
