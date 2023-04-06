// Code generated from arishem.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // arishem

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type arishemParser struct {
	*antlr.BaseParser
}

var arishemParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func arishemParserInit() {
	staticData := &arishemParserStaticData
	staticData.literalNames = []string{
		"", "'Conditions'", "'ConditionGroups'", "'ListExpr'", "'MapExpr'",
		"'MathExpr'", "'FuncExpr'", "'VarExpr'", "'FeatureExpr'", "'FeaturePath'",
		"'Const'", "'ConstList'", "'ParamMap'", "'ParamList'", "'BuiltinParam'",
		"'ActionName'", "'OpLogic'", "'OpMath'", "'FuncName'", "'Operator'",
		"'NOT '", "'Lhs'", "'Rhs'", "'true'", "'false'", "'\\'", "' '", "'#'",
		"'$'", "'&'", "'''", "'('", "')'", "'|'", "';'", "'='", "'?'", "'@'",
		"'^'", "'_'", "'`'", "'~'", "'StrConst'", "'NumConst'", "'BoolConst'",
		"", "", "'null'", "'.'", "','", "':'", "'\"'", "'['", "']'", "'{'",
		"'}'", "", "", "", "", "'&&'", "'||'", "'!'", "'+'", "'-'", "'/'", "'*'",
		"'%'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "OPERATOR_LOGIC_STR", "OPERATOR",
		"NULL", "DOT", "COMMA", "COLON", "QUOTATION", "L_BRACKET", "R_BRACKET",
		"L_BRACE", "R_BRACE", "INT", "EXP", "NAME_KEY_TYPE", "SP_UNICODE", "AND",
		"OR", "NOT", "PLUS", "SUB", "DIV", "MUL", "MOD", "EQUALS", "NOT_EQUALS",
		"GT", "LT", "GTE", "LTE", "WS",
	}
	staticData.ruleNames = []string{
		"condEntity", "aimEntity", "conditionGroup", "condition", "aim", "actionAimType",
		"expr", "exprType", "variable", "math", "function", "exprs", "feature",
		"constant", "constantList", "conditionsKey", "conditionsVal", "conditionsPair",
		"conditionGroupsKey", "conditionGroupsVal", "conditionGroupsPair", "listExprKey",
		"mapExprKey", "mathExprKey", "funcExprKey", "varExprKey", "featureExprKey",
		"featurePathKey", "constKey", "constListKey", "varPath", "varPathVal",
		"featurePath", "featurePathVal", "paramMapKey", "paramListKey", "builtinParamKey",
		"actionAimKey", "actionAimVal", "actionAimPair", "nameKey", "keyValue",
		"keyValues", "opLogicKey", "opLogicVal", "opLogicPair", "opLogic", "opMathKey",
		"opMathVal", "opMathPair", "mathOperator", "funcNameKey", "funcNamePair",
		"operatorKey", "operatorVal", "operatorPair", "operatorType", "lhsKey",
		"lhsPair", "rhsKey", "rhsPair", "boolType", "numberType", "stringType",
		"specialChars", "strConstKey", "numConstKey", "boolConstKey",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 623, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 154, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 165, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		171, 8, 4, 10, 4, 12, 4, 174, 9, 4, 3, 4, 176, 8, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 182, 8, 4, 1, 4, 1, 4, 3, 4, 186, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 195, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 206, 8, 5, 1, 5, 1, 5, 3, 5, 210, 8, 5, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 217, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 252, 8, 7, 1, 8, 1, 8, 3, 8, 256, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 275, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 284, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 295, 8, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 300, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 306, 8, 11, 10, 11,
		12, 11, 309, 9, 11, 3, 11, 311, 8, 11, 1, 11, 1, 11, 3, 11, 315, 8, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 326,
		8, 12, 1, 12, 1, 12, 1, 12, 3, 12, 331, 8, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 352, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 358, 8, 14, 10, 14, 12, 14, 361, 9, 14, 3, 14, 363, 8,
		14, 1, 14, 1, 14, 3, 14, 367, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 5, 16, 377, 8, 16, 10, 16, 12, 16, 380, 9, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 5, 19, 396, 8, 19, 10, 19, 12, 19, 399, 9, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 5, 31, 450, 8, 31, 10, 31,
		12, 31, 453, 9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 4,
		33, 462, 8, 33, 11, 33, 12, 33, 463, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 5,
		42, 502, 8, 42, 10, 42, 12, 42, 505, 9, 42, 3, 42, 507, 8, 42, 1, 42, 1,
		42, 3, 42, 511, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56,
		3, 56, 562, 8, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1,
		58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 61, 1, 61, 1, 62, 3, 62, 585, 8, 62, 1, 62, 1, 62, 1, 62, 3, 62, 590,
		8, 62, 1, 62, 3, 62, 593, 8, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1,
		63, 1, 63, 5, 63, 602, 8, 63, 10, 63, 12, 63, 605, 9, 63, 1, 63, 1, 63,
		1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1,
		67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 603, 0, 68, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
		120, 122, 124, 126, 128, 130, 132, 134, 0, 6, 2, 0, 45, 45, 60, 62, 1,
		0, 63, 67, 2, 0, 20, 20, 62, 62, 1, 0, 23, 24, 2, 0, 25, 25, 51, 51, 4,
		0, 25, 41, 48, 55, 62, 67, 70, 71, 609, 0, 136, 1, 0, 0, 0, 2, 138, 1,
		0, 0, 0, 4, 153, 1, 0, 0, 0, 6, 164, 1, 0, 0, 0, 8, 185, 1, 0, 0, 0, 10,
		209, 1, 0, 0, 0, 12, 216, 1, 0, 0, 0, 14, 251, 1, 0, 0, 0, 16, 255, 1,
		0, 0, 0, 18, 274, 1, 0, 0, 0, 20, 299, 1, 0, 0, 0, 22, 314, 1, 0, 0, 0,
		24, 330, 1, 0, 0, 0, 26, 351, 1, 0, 0, 0, 28, 366, 1, 0, 0, 0, 30, 368,
		1, 0, 0, 0, 32, 372, 1, 0, 0, 0, 34, 383, 1, 0, 0, 0, 36, 387, 1, 0, 0,
		0, 38, 391, 1, 0, 0, 0, 40, 402, 1, 0, 0, 0, 42, 406, 1, 0, 0, 0, 44, 410,
		1, 0, 0, 0, 46, 414, 1, 0, 0, 0, 48, 418, 1, 0, 0, 0, 50, 422, 1, 0, 0,
		0, 52, 426, 1, 0, 0, 0, 54, 430, 1, 0, 0, 0, 56, 434, 1, 0, 0, 0, 58, 438,
		1, 0, 0, 0, 60, 442, 1, 0, 0, 0, 62, 446, 1, 0, 0, 0, 64, 454, 1, 0, 0,
		0, 66, 458, 1, 0, 0, 0, 68, 465, 1, 0, 0, 0, 70, 469, 1, 0, 0, 0, 72, 473,
		1, 0, 0, 0, 74, 477, 1, 0, 0, 0, 76, 481, 1, 0, 0, 0, 78, 485, 1, 0, 0,
		0, 80, 489, 1, 0, 0, 0, 82, 493, 1, 0, 0, 0, 84, 510, 1, 0, 0, 0, 86, 512,
		1, 0, 0, 0, 88, 516, 1, 0, 0, 0, 90, 520, 1, 0, 0, 0, 92, 524, 1, 0, 0,
		0, 94, 526, 1, 0, 0, 0, 96, 530, 1, 0, 0, 0, 98, 534, 1, 0, 0, 0, 100,
		538, 1, 0, 0, 0, 102, 540, 1, 0, 0, 0, 104, 544, 1, 0, 0, 0, 106, 548,
		1, 0, 0, 0, 108, 552, 1, 0, 0, 0, 110, 556, 1, 0, 0, 0, 112, 561, 1, 0,
		0, 0, 114, 565, 1, 0, 0, 0, 116, 569, 1, 0, 0, 0, 118, 573, 1, 0, 0, 0,
		120, 577, 1, 0, 0, 0, 122, 581, 1, 0, 0, 0, 124, 584, 1, 0, 0, 0, 126,
		594, 1, 0, 0, 0, 128, 608, 1, 0, 0, 0, 130, 610, 1, 0, 0, 0, 132, 614,
		1, 0, 0, 0, 134, 618, 1, 0, 0, 0, 136, 137, 3, 4, 2, 0, 137, 1, 1, 0, 0,
		0, 138, 139, 3, 8, 4, 0, 139, 3, 1, 0, 0, 0, 140, 141, 5, 54, 0, 0, 141,
		142, 3, 90, 45, 0, 142, 143, 5, 49, 0, 0, 143, 144, 3, 40, 20, 0, 144,
		145, 5, 55, 0, 0, 145, 154, 1, 0, 0, 0, 146, 147, 5, 54, 0, 0, 147, 148,
		3, 90, 45, 0, 148, 149, 5, 49, 0, 0, 149, 150, 3, 34, 17, 0, 150, 151,
		5, 55, 0, 0, 151, 154, 1, 0, 0, 0, 152, 154, 5, 47, 0, 0, 153, 140, 1,
		0, 0, 0, 153, 146, 1, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154, 5, 1, 0, 0, 0,
		155, 156, 5, 54, 0, 0, 156, 157, 3, 110, 55, 0, 157, 158, 5, 49, 0, 0,
		158, 159, 3, 116, 58, 0, 159, 160, 5, 49, 0, 0, 160, 161, 3, 120, 60, 0,
		161, 162, 5, 55, 0, 0, 162, 165, 1, 0, 0, 0, 163, 165, 5, 47, 0, 0, 164,
		155, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 7, 1, 0, 0, 0, 166, 175, 5,
		52, 0, 0, 167, 172, 3, 10, 5, 0, 168, 169, 5, 49, 0, 0, 169, 171, 3, 10,
		5, 0, 170, 168, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175,
		167, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 186,
		5, 53, 0, 0, 178, 186, 3, 10, 5, 0, 179, 181, 5, 54, 0, 0, 180, 182, 3,
		14, 7, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0,
		0, 183, 186, 5, 55, 0, 0, 184, 186, 5, 47, 0, 0, 185, 166, 1, 0, 0, 0,
		185, 178, 1, 0, 0, 0, 185, 179, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186,
		9, 1, 0, 0, 0, 187, 188, 5, 54, 0, 0, 188, 194, 3, 78, 39, 0, 189, 190,
		5, 49, 0, 0, 190, 191, 3, 68, 34, 0, 191, 192, 5, 50, 0, 0, 192, 193, 3,
		84, 42, 0, 193, 195, 1, 0, 0, 0, 194, 189, 1, 0, 0, 0, 194, 195, 1, 0,
		0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 5, 55, 0, 0, 197, 210, 1, 0, 0, 0,
		198, 199, 5, 54, 0, 0, 199, 205, 3, 78, 39, 0, 200, 201, 5, 49, 0, 0, 201,
		202, 3, 70, 35, 0, 202, 203, 5, 50, 0, 0, 203, 204, 3, 22, 11, 0, 204,
		206, 1, 0, 0, 0, 205, 200, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 207,
		1, 0, 0, 0, 207, 208, 5, 55, 0, 0, 208, 210, 1, 0, 0, 0, 209, 187, 1, 0,
		0, 0, 209, 198, 1, 0, 0, 0, 210, 11, 1, 0, 0, 0, 211, 212, 5, 54, 0, 0,
		212, 213, 3, 14, 7, 0, 213, 214, 5, 55, 0, 0, 214, 217, 1, 0, 0, 0, 215,
		217, 5, 47, 0, 0, 216, 211, 1, 0, 0, 0, 216, 215, 1, 0, 0, 0, 217, 13,
		1, 0, 0, 0, 218, 219, 3, 42, 21, 0, 219, 220, 5, 50, 0, 0, 220, 221, 3,
		22, 11, 0, 221, 252, 1, 0, 0, 0, 222, 223, 3, 44, 22, 0, 223, 224, 5, 50,
		0, 0, 224, 225, 3, 84, 42, 0, 225, 252, 1, 0, 0, 0, 226, 227, 3, 56, 28,
		0, 227, 228, 5, 50, 0, 0, 228, 229, 3, 26, 13, 0, 229, 252, 1, 0, 0, 0,
		230, 231, 3, 58, 29, 0, 231, 232, 5, 50, 0, 0, 232, 233, 3, 28, 14, 0,
		233, 252, 1, 0, 0, 0, 234, 235, 3, 46, 23, 0, 235, 236, 5, 50, 0, 0, 236,
		237, 3, 18, 9, 0, 237, 252, 1, 0, 0, 0, 238, 239, 3, 48, 24, 0, 239, 240,
		5, 50, 0, 0, 240, 241, 3, 20, 10, 0, 241, 252, 1, 0, 0, 0, 242, 243, 3,
		50, 25, 0, 243, 244, 5, 50, 0, 0, 244, 245, 3, 16, 8, 0, 245, 252, 1, 0,
		0, 0, 246, 247, 3, 52, 26, 0, 247, 248, 5, 50, 0, 0, 248, 249, 3, 24, 12,
		0, 249, 252, 1, 0, 0, 0, 250, 252, 5, 47, 0, 0, 251, 218, 1, 0, 0, 0, 251,
		222, 1, 0, 0, 0, 251, 226, 1, 0, 0, 0, 251, 230, 1, 0, 0, 0, 251, 234,
		1, 0, 0, 0, 251, 238, 1, 0, 0, 0, 251, 242, 1, 0, 0, 0, 251, 246, 1, 0,
		0, 0, 251, 250, 1, 0, 0, 0, 252, 15, 1, 0, 0, 0, 253, 256, 3, 60, 30, 0,
		254, 256, 5, 47, 0, 0, 255, 253, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0, 256,
		17, 1, 0, 0, 0, 257, 258, 5, 54, 0, 0, 258, 259, 3, 98, 49, 0, 259, 260,
		5, 49, 0, 0, 260, 261, 3, 116, 58, 0, 261, 262, 5, 49, 0, 0, 262, 263,
		3, 120, 60, 0, 263, 264, 5, 55, 0, 0, 264, 275, 1, 0, 0, 0, 265, 266, 5,
		54, 0, 0, 266, 267, 3, 98, 49, 0, 267, 268, 5, 49, 0, 0, 268, 269, 3, 70,
		35, 0, 269, 270, 5, 50, 0, 0, 270, 271, 3, 22, 11, 0, 271, 272, 5, 55,
		0, 0, 272, 275, 1, 0, 0, 0, 273, 275, 5, 47, 0, 0, 274, 257, 1, 0, 0, 0,
		274, 265, 1, 0, 0, 0, 274, 273, 1, 0, 0, 0, 275, 19, 1, 0, 0, 0, 276, 277,
		5, 54, 0, 0, 277, 283, 3, 104, 52, 0, 278, 279, 5, 49, 0, 0, 279, 280,
		3, 68, 34, 0, 280, 281, 5, 50, 0, 0, 281, 282, 3, 84, 42, 0, 282, 284,
		1, 0, 0, 0, 283, 278, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 1, 0,
		0, 0, 285, 286, 5, 55, 0, 0, 286, 300, 1, 0, 0, 0, 287, 288, 5, 54, 0,
		0, 288, 294, 3, 104, 52, 0, 289, 290, 5, 49, 0, 0, 290, 291, 3, 70, 35,
		0, 291, 292, 5, 50, 0, 0, 292, 293, 3, 22, 11, 0, 293, 295, 1, 0, 0, 0,
		294, 289, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296,
		297, 5, 55, 0, 0, 297, 300, 1, 0, 0, 0, 298, 300, 5, 47, 0, 0, 299, 276,
		1, 0, 0, 0, 299, 287, 1, 0, 0, 0, 299, 298, 1, 0, 0, 0, 300, 21, 1, 0,
		0, 0, 301, 310, 5, 52, 0, 0, 302, 307, 3, 12, 6, 0, 303, 304, 5, 49, 0,
		0, 304, 306, 3, 12, 6, 0, 305, 303, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307,
		305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307,
		1, 0, 0, 0, 310, 302, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 1, 0,
		0, 0, 312, 315, 5, 53, 0, 0, 313, 315, 5, 47, 0, 0, 314, 301, 1, 0, 0,
		0, 314, 313, 1, 0, 0, 0, 315, 23, 1, 0, 0, 0, 316, 317, 5, 54, 0, 0, 317,
		318, 3, 54, 27, 0, 318, 319, 5, 50, 0, 0, 319, 325, 3, 64, 32, 0, 320,
		321, 5, 49, 0, 0, 321, 322, 3, 72, 36, 0, 322, 323, 5, 50, 0, 0, 323, 324,
		3, 84, 42, 0, 324, 326, 1, 0, 0, 0, 325, 320, 1, 0, 0, 0, 325, 326, 1,
		0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 5, 55, 0, 0, 328, 331, 1, 0, 0,
		0, 329, 331, 5, 47, 0, 0, 330, 316, 1, 0, 0, 0, 330, 329, 1, 0, 0, 0, 331,
		25, 1, 0, 0, 0, 332, 333, 5, 54, 0, 0, 333, 334, 3, 132, 66, 0, 334, 335,
		5, 50, 0, 0, 335, 336, 3, 124, 62, 0, 336, 337, 5, 55, 0, 0, 337, 352,
		1, 0, 0, 0, 338, 339, 5, 54, 0, 0, 339, 340, 3, 130, 65, 0, 340, 341, 5,
		50, 0, 0, 341, 342, 3, 126, 63, 0, 342, 343, 5, 55, 0, 0, 343, 352, 1,
		0, 0, 0, 344, 345, 5, 54, 0, 0, 345, 346, 3, 134, 67, 0, 346, 347, 5, 50,
		0, 0, 347, 348, 3, 122, 61, 0, 348, 349, 5, 55, 0, 0, 349, 352, 1, 0, 0,
		0, 350, 352, 5, 47, 0, 0, 351, 332, 1, 0, 0, 0, 351, 338, 1, 0, 0, 0, 351,
		344, 1, 0, 0, 0, 351, 350, 1, 0, 0, 0, 352, 27, 1, 0, 0, 0, 353, 362, 5,
		52, 0, 0, 354, 359, 3, 26, 13, 0, 355, 356, 5, 49, 0, 0, 356, 358, 3, 26,
		13, 0, 357, 355, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0,
		359, 360, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362,
		354, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 367,
		5, 53, 0, 0, 365, 367, 5, 47, 0, 0, 366, 353, 1, 0, 0, 0, 366, 365, 1,
		0, 0, 0, 367, 29, 1, 0, 0, 0, 368, 369, 5, 51, 0, 0, 369, 370, 5, 1, 0,
		0, 370, 371, 5, 51, 0, 0, 371, 31, 1, 0, 0, 0, 372, 373, 5, 52, 0, 0, 373,
		378, 3, 6, 3, 0, 374, 375, 5, 49, 0, 0, 375, 377, 3, 6, 3, 0, 376, 374,
		1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0,
		0, 0, 379, 381, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 381, 382, 5, 53, 0, 0,
		382, 33, 1, 0, 0, 0, 383, 384, 3, 30, 15, 0, 384, 385, 5, 50, 0, 0, 385,
		386, 3, 32, 16, 0, 386, 35, 1, 0, 0, 0, 387, 388, 5, 51, 0, 0, 388, 389,
		5, 2, 0, 0, 389, 390, 5, 51, 0, 0, 390, 37, 1, 0, 0, 0, 391, 392, 5, 52,
		0, 0, 392, 397, 3, 4, 2, 0, 393, 394, 5, 49, 0, 0, 394, 396, 3, 4, 2, 0,
		395, 393, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397,
		398, 1, 0, 0, 0, 398, 400, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 400, 401,
		5, 53, 0, 0, 401, 39, 1, 0, 0, 0, 402, 403, 3, 36, 18, 0, 403, 404, 5,
		50, 0, 0, 404, 405, 3, 38, 19, 0, 405, 41, 1, 0, 0, 0, 406, 407, 5, 51,
		0, 0, 407, 408, 5, 3, 0, 0, 408, 409, 5, 51, 0, 0, 409, 43, 1, 0, 0, 0,
		410, 411, 5, 51, 0, 0, 411, 412, 5, 4, 0, 0, 412, 413, 5, 51, 0, 0, 413,
		45, 1, 0, 0, 0, 414, 415, 5, 51, 0, 0, 415, 416, 5, 5, 0, 0, 416, 417,
		5, 51, 0, 0, 417, 47, 1, 0, 0, 0, 418, 419, 5, 51, 0, 0, 419, 420, 5, 6,
		0, 0, 420, 421, 5, 51, 0, 0, 421, 49, 1, 0, 0, 0, 422, 423, 5, 51, 0, 0,
		423, 424, 5, 7, 0, 0, 424, 425, 5, 51, 0, 0, 425, 51, 1, 0, 0, 0, 426,
		427, 5, 51, 0, 0, 427, 428, 5, 8, 0, 0, 428, 429, 5, 51, 0, 0, 429, 53,
		1, 0, 0, 0, 430, 431, 5, 51, 0, 0, 431, 432, 5, 9, 0, 0, 432, 433, 5, 51,
		0, 0, 433, 55, 1, 0, 0, 0, 434, 435, 5, 51, 0, 0, 435, 436, 5, 10, 0, 0,
		436, 437, 5, 51, 0, 0, 437, 57, 1, 0, 0, 0, 438, 439, 5, 51, 0, 0, 439,
		440, 5, 11, 0, 0, 440, 441, 5, 51, 0, 0, 441, 59, 1, 0, 0, 0, 442, 443,
		5, 51, 0, 0, 443, 444, 3, 62, 31, 0, 444, 445, 5, 51, 0, 0, 445, 61, 1,
		0, 0, 0, 446, 451, 5, 58, 0, 0, 447, 448, 5, 48, 0, 0, 448, 450, 5, 58,
		0, 0, 449, 447, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0,
		451, 452, 1, 0, 0, 0, 452, 63, 1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455,
		5, 51, 0, 0, 455, 456, 3, 66, 33, 0, 456, 457, 5, 51, 0, 0, 457, 65, 1,
		0, 0, 0, 458, 461, 5, 58, 0, 0, 459, 460, 5, 48, 0, 0, 460, 462, 5, 58,
		0, 0, 461, 459, 1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 461, 1, 0, 0, 0,
		463, 464, 1, 0, 0, 0, 464, 67, 1, 0, 0, 0, 465, 466, 5, 51, 0, 0, 466,
		467, 5, 12, 0, 0, 467, 468, 5, 51, 0, 0, 468, 69, 1, 0, 0, 0, 469, 470,
		5, 51, 0, 0, 470, 471, 5, 13, 0, 0, 471, 472, 5, 51, 0, 0, 472, 71, 1,
		0, 0, 0, 473, 474, 5, 51, 0, 0, 474, 475, 5, 14, 0, 0, 475, 476, 5, 51,
		0, 0, 476, 73, 1, 0, 0, 0, 477, 478, 5, 51, 0, 0, 478, 479, 5, 15, 0, 0,
		479, 480, 5, 51, 0, 0, 480, 75, 1, 0, 0, 0, 481, 482, 5, 51, 0, 0, 482,
		483, 5, 58, 0, 0, 483, 484, 5, 51, 0, 0, 484, 77, 1, 0, 0, 0, 485, 486,
		3, 74, 37, 0, 486, 487, 5, 50, 0, 0, 487, 488, 3, 76, 38, 0, 488, 79, 1,
		0, 0, 0, 489, 490, 5, 51, 0, 0, 490, 491, 5, 58, 0, 0, 491, 492, 5, 51,
		0, 0, 492, 81, 1, 0, 0, 0, 493, 494, 3, 80, 40, 0, 494, 495, 5, 50, 0,
		0, 495, 496, 3, 12, 6, 0, 496, 83, 1, 0, 0, 0, 497, 506, 5, 54, 0, 0, 498,
		503, 3, 82, 41, 0, 499, 500, 5, 49, 0, 0, 500, 502, 3, 82, 41, 0, 501,
		499, 1, 0, 0, 0, 502, 505, 1, 0, 0, 0, 503, 501, 1, 0, 0, 0, 503, 504,
		1, 0, 0, 0, 504, 507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 506, 498, 1, 0,
		0, 0, 506, 507, 1, 0, 0, 0, 507, 508, 1, 0, 0, 0, 508, 511, 5, 55, 0, 0,
		509, 511, 5, 47, 0, 0, 510, 497, 1, 0, 0, 0, 510, 509, 1, 0, 0, 0, 511,
		85, 1, 0, 0, 0, 512, 513, 5, 51, 0, 0, 513, 514, 5, 16, 0, 0, 514, 515,
		5, 51, 0, 0, 515, 87, 1, 0, 0, 0, 516, 517, 5, 51, 0, 0, 517, 518, 3, 92,
		46, 0, 518, 519, 5, 51, 0, 0, 519, 89, 1, 0, 0, 0, 520, 521, 3, 86, 43,
		0, 521, 522, 5, 50, 0, 0, 522, 523, 3, 88, 44, 0, 523, 91, 1, 0, 0, 0,
		524, 525, 7, 0, 0, 0, 525, 93, 1, 0, 0, 0, 526, 527, 5, 51, 0, 0, 527,
		528, 5, 17, 0, 0, 528, 529, 5, 51, 0, 0, 529, 95, 1, 0, 0, 0, 530, 531,
		5, 51, 0, 0, 531, 532, 3, 100, 50, 0, 532, 533, 5, 51, 0, 0, 533, 97, 1,
		0, 0, 0, 534, 535, 3, 94, 47, 0, 535, 536, 5, 50, 0, 0, 536, 537, 3, 96,
		48, 0, 537, 99, 1, 0, 0, 0, 538, 539, 7, 1, 0, 0, 539, 101, 1, 0, 0, 0,
		540, 541, 5, 51, 0, 0, 541, 542, 5, 18, 0, 0, 542, 543, 5, 51, 0, 0, 543,
		103, 1, 0, 0, 0, 544, 545, 3, 102, 51, 0, 545, 546, 5, 50, 0, 0, 546, 547,
		3, 80, 40, 0, 547, 105, 1, 0, 0, 0, 548, 549, 5, 51, 0, 0, 549, 550, 5,
		19, 0, 0, 550, 551, 5, 51, 0, 0, 551, 107, 1, 0, 0, 0, 552, 553, 5, 51,
		0, 0, 553, 554, 3, 112, 56, 0, 554, 555, 5, 51, 0, 0, 555, 109, 1, 0, 0,
		0, 556, 557, 3, 106, 53, 0, 557, 558, 5, 50, 0, 0, 558, 559, 3, 108, 54,
		0, 559, 111, 1, 0, 0, 0, 560, 562, 7, 2, 0, 0, 561, 560, 1, 0, 0, 0, 561,
		562, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 564, 5, 46, 0, 0, 564, 113,
		1, 0, 0, 0, 565, 566, 5, 51, 0, 0, 566, 567, 5, 21, 0, 0, 567, 568, 5,
		51, 0, 0, 568, 115, 1, 0, 0, 0, 569, 570, 3, 114, 57, 0, 570, 571, 5, 50,
		0, 0, 571, 572, 3, 12, 6, 0, 572, 117, 1, 0, 0, 0, 573, 574, 5, 51, 0,
		0, 574, 575, 5, 22, 0, 0, 575, 576, 5, 51, 0, 0, 576, 119, 1, 0, 0, 0,
		577, 578, 3, 118, 59, 0, 578, 579, 5, 50, 0, 0, 579, 580, 3, 12, 6, 0,
		580, 121, 1, 0, 0, 0, 581, 582, 7, 3, 0, 0, 582, 123, 1, 0, 0, 0, 583,
		585, 5, 64, 0, 0, 584, 583, 1, 0, 0, 0, 584, 585, 1, 0, 0, 0, 585, 586,
		1, 0, 0, 0, 586, 589, 5, 56, 0, 0, 587, 588, 5, 48, 0, 0, 588, 590, 5,
		56, 0, 0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 592, 1, 0, 0,
		0, 591, 593, 5, 57, 0, 0, 592, 591, 1, 0, 0, 0, 592, 593, 1, 0, 0, 0, 593,
		125, 1, 0, 0, 0, 594, 603, 5, 51, 0, 0, 595, 602, 9, 0, 0, 0, 596, 597,
		5, 25, 0, 0, 597, 602, 9, 0, 0, 0, 598, 602, 8, 4, 0, 0, 599, 602, 3, 128,
		64, 0, 600, 602, 5, 59, 0, 0, 601, 595, 1, 0, 0, 0, 601, 596, 1, 0, 0,
		0, 601, 598, 1, 0, 0, 0, 601, 599, 1, 0, 0, 0, 601, 600, 1, 0, 0, 0, 602,
		605, 1, 0, 0, 0, 603, 604, 1, 0, 0, 0, 603, 601, 1, 0, 0, 0, 604, 606,
		1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 606, 607, 5, 51, 0, 0, 607, 127, 1, 0,
		0, 0, 608, 609, 7, 5, 0, 0, 609, 129, 1, 0, 0, 0, 610, 611, 5, 51, 0, 0,
		611, 612, 5, 42, 0, 0, 612, 613, 5, 51, 0, 0, 613, 131, 1, 0, 0, 0, 614,
		615, 5, 51, 0, 0, 615, 616, 5, 43, 0, 0, 616, 617, 5, 51, 0, 0, 617, 133,
		1, 0, 0, 0, 618, 619, 5, 51, 0, 0, 619, 620, 5, 44, 0, 0, 620, 621, 5,
		51, 0, 0, 621, 135, 1, 0, 0, 0, 38, 153, 164, 172, 175, 181, 185, 194,
		205, 209, 216, 251, 255, 274, 283, 294, 299, 307, 310, 314, 325, 330, 351,
		359, 362, 366, 378, 397, 451, 463, 503, 506, 510, 561, 584, 589, 592, 601,
		603,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// arishemParserInit initializes any static state used to implement arishemParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewarishemParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArishemParserInit() {
	staticData := &arishemParserStaticData
	staticData.once.Do(arishemParserInit)
}

// NewarishemParser produces a new parser instance for the optional input antlr.TokenStream.
func NewarishemParser(input antlr.TokenStream) *arishemParser {
	ArishemParserInit()
	this := new(arishemParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &arishemParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "arishem.g4"

	return this
}

// arishemParser tokens.
const (
	arishemParserEOF                = antlr.TokenEOF
	arishemParserT__0               = 1
	arishemParserT__1               = 2
	arishemParserT__2               = 3
	arishemParserT__3               = 4
	arishemParserT__4               = 5
	arishemParserT__5               = 6
	arishemParserT__6               = 7
	arishemParserT__7               = 8
	arishemParserT__8               = 9
	arishemParserT__9               = 10
	arishemParserT__10              = 11
	arishemParserT__11              = 12
	arishemParserT__12              = 13
	arishemParserT__13              = 14
	arishemParserT__14              = 15
	arishemParserT__15              = 16
	arishemParserT__16              = 17
	arishemParserT__17              = 18
	arishemParserT__18              = 19
	arishemParserT__19              = 20
	arishemParserT__20              = 21
	arishemParserT__21              = 22
	arishemParserT__22              = 23
	arishemParserT__23              = 24
	arishemParserT__24              = 25
	arishemParserT__25              = 26
	arishemParserT__26              = 27
	arishemParserT__27              = 28
	arishemParserT__28              = 29
	arishemParserT__29              = 30
	arishemParserT__30              = 31
	arishemParserT__31              = 32
	arishemParserT__32              = 33
	arishemParserT__33              = 34
	arishemParserT__34              = 35
	arishemParserT__35              = 36
	arishemParserT__36              = 37
	arishemParserT__37              = 38
	arishemParserT__38              = 39
	arishemParserT__39              = 40
	arishemParserT__40              = 41
	arishemParserT__41              = 42
	arishemParserT__42              = 43
	arishemParserT__43              = 44
	arishemParserOPERATOR_LOGIC_STR = 45
	arishemParserOPERATOR           = 46
	arishemParserNULL               = 47
	arishemParserDOT                = 48
	arishemParserCOMMA              = 49
	arishemParserCOLON              = 50
	arishemParserQUOTATION          = 51
	arishemParserL_BRACKET          = 52
	arishemParserR_BRACKET          = 53
	arishemParserL_BRACE            = 54
	arishemParserR_BRACE            = 55
	arishemParserINT                = 56
	arishemParserEXP                = 57
	arishemParserNAME_KEY_TYPE      = 58
	arishemParserSP_UNICODE         = 59
	arishemParserAND                = 60
	arishemParserOR                 = 61
	arishemParserNOT                = 62
	arishemParserPLUS               = 63
	arishemParserSUB                = 64
	arishemParserDIV                = 65
	arishemParserMUL                = 66
	arishemParserMOD                = 67
	arishemParserEQUALS             = 68
	arishemParserNOT_EQUALS         = 69
	arishemParserGT                 = 70
	arishemParserLT                 = 71
	arishemParserGTE                = 72
	arishemParserLTE                = 73
	arishemParserWS                 = 74
)

// arishemParser rules.
const (
	arishemParserRULE_condEntity          = 0
	arishemParserRULE_aimEntity           = 1
	arishemParserRULE_conditionGroup      = 2
	arishemParserRULE_condition           = 3
	arishemParserRULE_aim                 = 4
	arishemParserRULE_actionAimType       = 5
	arishemParserRULE_expr                = 6
	arishemParserRULE_exprType            = 7
	arishemParserRULE_variable            = 8
	arishemParserRULE_math                = 9
	arishemParserRULE_function            = 10
	arishemParserRULE_exprs               = 11
	arishemParserRULE_feature             = 12
	arishemParserRULE_constant            = 13
	arishemParserRULE_constantList        = 14
	arishemParserRULE_conditionsKey       = 15
	arishemParserRULE_conditionsVal       = 16
	arishemParserRULE_conditionsPair      = 17
	arishemParserRULE_conditionGroupsKey  = 18
	arishemParserRULE_conditionGroupsVal  = 19
	arishemParserRULE_conditionGroupsPair = 20
	arishemParserRULE_listExprKey         = 21
	arishemParserRULE_mapExprKey          = 22
	arishemParserRULE_mathExprKey         = 23
	arishemParserRULE_funcExprKey         = 24
	arishemParserRULE_varExprKey          = 25
	arishemParserRULE_featureExprKey      = 26
	arishemParserRULE_featurePathKey      = 27
	arishemParserRULE_constKey            = 28
	arishemParserRULE_constListKey        = 29
	arishemParserRULE_varPath             = 30
	arishemParserRULE_varPathVal          = 31
	arishemParserRULE_featurePath         = 32
	arishemParserRULE_featurePathVal      = 33
	arishemParserRULE_paramMapKey         = 34
	arishemParserRULE_paramListKey        = 35
	arishemParserRULE_builtinParamKey     = 36
	arishemParserRULE_actionAimKey        = 37
	arishemParserRULE_actionAimVal        = 38
	arishemParserRULE_actionAimPair       = 39
	arishemParserRULE_nameKey             = 40
	arishemParserRULE_keyValue            = 41
	arishemParserRULE_keyValues           = 42
	arishemParserRULE_opLogicKey          = 43
	arishemParserRULE_opLogicVal          = 44
	arishemParserRULE_opLogicPair         = 45
	arishemParserRULE_opLogic             = 46
	arishemParserRULE_opMathKey           = 47
	arishemParserRULE_opMathVal           = 48
	arishemParserRULE_opMathPair          = 49
	arishemParserRULE_mathOperator        = 50
	arishemParserRULE_funcNameKey         = 51
	arishemParserRULE_funcNamePair        = 52
	arishemParserRULE_operatorKey         = 53
	arishemParserRULE_operatorVal         = 54
	arishemParserRULE_operatorPair        = 55
	arishemParserRULE_operatorType        = 56
	arishemParserRULE_lhsKey              = 57
	arishemParserRULE_lhsPair             = 58
	arishemParserRULE_rhsKey              = 59
	arishemParserRULE_rhsPair             = 60
	arishemParserRULE_boolType            = 61
	arishemParserRULE_numberType          = 62
	arishemParserRULE_stringType          = 63
	arishemParserRULE_specialChars        = 64
	arishemParserRULE_strConstKey         = 65
	arishemParserRULE_numConstKey         = 66
	arishemParserRULE_boolConstKey        = 67
)

// ICondEntityContext is an interface to support dynamic dispatch.
type ICondEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCondEntityContext differentiates from other interfaces.
	IsCondEntityContext()
}

type CondEntityContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondEntityContext() *CondEntityContext {
	var p = new(CondEntityContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_condEntity
	return p
}

func (*CondEntityContext) IsCondEntityContext() {}

func NewCondEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondEntityContext {
	var p = new(CondEntityContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_condEntity

	return p
}

func (s *CondEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *CondEntityContext) CopyFrom(ctx *CondEntityContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CondEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArishemCondContext struct {
	*CondEntityContext
}

func NewArishemCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArishemCondContext {
	var p = new(ArishemCondContext)

	p.CondEntityContext = NewEmptyCondEntityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CondEntityContext))

	return p
}

func (s *ArishemCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArishemCondContext) ConditionGroup() IConditionGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionGroupContext)
}

func (s *ArishemCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitArishemCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) CondEntity() (localctx ICondEntityContext) {
	this := p
	_ = this

	localctx = NewCondEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, arishemParserRULE_condEntity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewArishemCondContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.ConditionGroup()
	}

	return localctx
}

// IAimEntityContext is an interface to support dynamic dispatch.
type IAimEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAimEntityContext differentiates from other interfaces.
	IsAimEntityContext()
}

type AimEntityContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyAimEntityContext() *AimEntityContext {
	var p = new(AimEntityContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_aimEntity
	return p
}

func (*AimEntityContext) IsAimEntityContext() {}

func NewAimEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AimEntityContext {
	var p = new(AimEntityContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_aimEntity

	return p
}

func (s *AimEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *AimEntityContext) CopyFrom(ctx *AimEntityContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AimEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AimEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArishemAimContext struct {
	*AimEntityContext
}

func NewArishemAimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArishemAimContext {
	var p = new(ArishemAimContext)

	p.AimEntityContext = NewEmptyAimEntityContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AimEntityContext))

	return p
}

func (s *ArishemAimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArishemAimContext) Aim() IAimContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAimContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAimContext)
}

func (s *ArishemAimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitArishemAim(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) AimEntity() (localctx IAimEntityContext) {
	this := p
	_ = this

	localctx = NewAimEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, arishemParserRULE_aimEntity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewArishemAimContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Aim()
	}

	return localctx
}

// IConditionGroupContext is an interface to support dynamic dispatch.
type IConditionGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionGroupContext differentiates from other interfaces.
	IsConditionGroupContext()
}

type ConditionGroupContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionGroupContext() *ConditionGroupContext {
	var p = new(ConditionGroupContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionGroup
	return p
}

func (*ConditionGroupContext) IsConditionGroupContext() {}

func NewConditionGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionGroupContext {
	var p = new(ConditionGroupContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionGroup

	return p
}

func (s *ConditionGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionGroupContext) CopyFrom(ctx *ConditionGroupContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConditionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LeafConditionGroupContext struct {
	*ConditionGroupContext
}

func NewLeafConditionGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeafConditionGroupContext {
	var p = new(LeafConditionGroupContext)

	p.ConditionGroupContext = NewEmptyConditionGroupContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionGroupContext))

	return p
}

func (s *LeafConditionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeafConditionGroupContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *LeafConditionGroupContext) OpLogicPair() IOpLogicPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpLogicPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpLogicPairContext)
}

func (s *LeafConditionGroupContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *LeafConditionGroupContext) ConditionsPair() IConditionsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsPairContext)
}

func (s *LeafConditionGroupContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *LeafConditionGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLeafConditionGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type NestConditionGroupContext struct {
	*ConditionGroupContext
}

func NewNestConditionGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestConditionGroupContext {
	var p = new(NestConditionGroupContext)

	p.ConditionGroupContext = NewEmptyConditionGroupContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionGroupContext))

	return p
}

func (s *NestConditionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestConditionGroupContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *NestConditionGroupContext) OpLogicPair() IOpLogicPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpLogicPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpLogicPairContext)
}

func (s *NestConditionGroupContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *NestConditionGroupContext) ConditionGroupsPair() IConditionGroupsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionGroupsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionGroupsPairContext)
}

func (s *NestConditionGroupContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *NestConditionGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNestConditionGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullConditionGroupContext struct {
	*ConditionGroupContext
}

func NewNullConditionGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullConditionGroupContext {
	var p = new(NullConditionGroupContext)

	p.ConditionGroupContext = NewEmptyConditionGroupContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionGroupContext))

	return p
}

func (s *NullConditionGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullConditionGroupContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullConditionGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullConditionGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionGroup() (localctx IConditionGroupContext) {
	this := p
	_ = this

	localctx = NewConditionGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, arishemParserRULE_conditionGroup)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNestConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(141)
			p.OpLogicPair()
		}
		{
			p.SetState(142)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(143)
			p.ConditionGroupsPair()
		}
		{
			p.SetState(144)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewLeafConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(147)
			p.OpLogicPair()
		}
		{
			p.SetState(148)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(149)
			p.ConditionsPair()
		}
		{
			p.SetState(150)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyFrom(ctx *ConditionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullConditionContext struct {
	*ConditionContext
}

func NewNullConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullConditionContext {
	var p = new(NullConditionContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *NullConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullConditionContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullConditionContext struct {
	*ConditionContext
}

func NewFullConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullConditionContext {
	var p = new(FullConditionContext)

	p.ConditionContext = NewEmptyConditionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConditionContext))

	return p
}

func (s *FullConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullConditionContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullConditionContext) OperatorPair() IOperatorPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorPairContext)
}

func (s *FullConditionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *FullConditionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *FullConditionContext) LhsPair() ILhsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsPairContext)
}

func (s *FullConditionContext) RhsPair() IRhsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsPairContext)
}

func (s *FullConditionContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, arishemParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(156)
			p.OperatorPair()
		}
		{
			p.SetState(157)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(158)
			p.LhsPair()
		}
		{
			p.SetState(159)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(160)
			p.RhsPair()
		}
		{
			p.SetState(161)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAimContext is an interface to support dynamic dispatch.
type IAimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAimContext differentiates from other interfaces.
	IsAimContext()
}

type AimContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyAimContext() *AimContext {
	var p = new(AimContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_aim
	return p
}

func (*AimContext) IsAimContext() {}

func NewAimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AimContext {
	var p = new(AimContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_aim

	return p
}

func (s *AimContext) GetParser() antlr.Parser { return s.parser }

func (s *AimContext) CopyFrom(ctx *AimContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ActionAimListContext struct {
	*AimContext
}

func NewActionAimListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ActionAimListContext {
	var p = new(ActionAimListContext)

	p.AimContext = NewEmptyAimContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AimContext))

	return p
}

func (s *ActionAimListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimListContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *ActionAimListContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *ActionAimListContext) AllActionAimType() []IActionAimTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionAimTypeContext); ok {
			len++
		}
	}

	tst := make([]IActionAimTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionAimTypeContext); ok {
			tst[i] = t.(IActionAimTypeContext)
			i++
		}
	}

	return tst
}

func (s *ActionAimListContext) ActionAimType(i int) IActionAimTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimTypeContext)
}

func (s *ActionAimListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *ActionAimListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *ActionAimListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitActionAimList(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullAimContext struct {
	*AimContext
}

func NewNullAimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullAimContext {
	var p = new(NullAimContext)

	p.AimContext = NewEmptyAimContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AimContext))

	return p
}

func (s *NullAimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullAimContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullAimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullAim(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAimContext struct {
	*AimContext
}

func NewExprAimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAimContext {
	var p = new(ExprAimContext)

	p.AimContext = NewEmptyAimContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AimContext))

	return p
}

func (s *ExprAimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAimContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ExprAimContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ExprAimContext) ExprType() IExprTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprTypeContext)
}

func (s *ExprAimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitExprAim(s)

	default:
		return t.VisitChildren(s)
	}
}

type ActionAimContext struct {
	*AimContext
}

func NewActionAimContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ActionAimContext {
	var p = new(ActionAimContext)

	p.AimContext = NewEmptyAimContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AimContext))

	return p
}

func (s *ActionAimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimContext) ActionAimType() IActionAimTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimTypeContext)
}

func (s *ActionAimContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitActionAim(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Aim() (localctx IAimContext) {
	this := p
	_ = this

	localctx = NewAimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, arishemParserRULE_aim)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewActionAimListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserL_BRACE {
			{
				p.SetState(167)
				p.ActionAimType()
			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(168)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(169)
					p.ActionAimType()
				}

				p.SetState(174)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(177)
			p.Match(arishemParserR_BRACKET)
		}

	case 2:
		localctx = NewActionAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.ActionAimType()
		}

	case 3:
		localctx = NewExprAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(179)
			p.Match(arishemParserL_BRACE)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserQUOTATION {
			{
				p.SetState(180)
				p.ExprType()
			}

		}
		{
			p.SetState(183)
			p.Match(arishemParserR_BRACE)
		}

	case 4:
		localctx = NewNullAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(184)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IActionAimTypeContext is an interface to support dynamic dispatch.
type IActionAimTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsActionAimTypeContext differentiates from other interfaces.
	IsActionAimTypeContext()
}

type ActionAimTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAimTypeContext() *ActionAimTypeContext {
	var p = new(ActionAimTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_actionAimType
	return p
}

func (*ActionAimTypeContext) IsActionAimTypeContext() {}

func NewActionAimTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAimTypeContext {
	var p = new(ActionAimTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_actionAimType

	return p
}

func (s *ActionAimTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAimTypeContext) CopyFrom(ctx *ActionAimTypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ActionAimTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamMapActionContext struct {
	*ActionAimTypeContext
}

func NewParamMapActionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamMapActionContext {
	var p = new(ParamMapActionContext)

	p.ActionAimTypeContext = NewEmptyActionAimTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ActionAimTypeContext))

	return p
}

func (s *ParamMapActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamMapActionContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ParamMapActionContext) ActionAimPair() IActionAimPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimPairContext)
}

func (s *ParamMapActionContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ParamMapActionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *ParamMapActionContext) ParamMapKey() IParamMapKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamMapKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamMapKeyContext)
}

func (s *ParamMapActionContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ParamMapActionContext) KeyValues() IKeyValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValuesContext)
}

func (s *ParamMapActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamMapAction(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamListActionContext struct {
	*ActionAimTypeContext
}

func NewParamListActionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListActionContext {
	var p = new(ParamListActionContext)

	p.ActionAimTypeContext = NewEmptyActionAimTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ActionAimTypeContext))

	return p
}

func (s *ParamListActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListActionContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ParamListActionContext) ActionAimPair() IActionAimPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimPairContext)
}

func (s *ParamListActionContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ParamListActionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *ParamListActionContext) ParamListKey() IParamListKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListKeyContext)
}

func (s *ParamListActionContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ParamListActionContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *ParamListActionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamListAction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ActionAimType() (localctx IActionAimTypeContext) {
	this := p
	_ = this

	localctx = NewActionAimTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, arishemParserRULE_actionAimType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamMapActionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(188)
			p.ActionAimPair()
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(189)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(190)
				p.ParamMapKey()
			}
			{
				p.SetState(191)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(192)
				p.KeyValues()
			}

		}
		{
			p.SetState(196)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewParamListActionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(199)
			p.ActionAimPair()
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(200)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(201)
				p.ParamListKey()
			}
			{
				p.SetState(202)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(203)
				p.Exprs()
			}

		}
		{
			p.SetState(207)
			p.Match(arishemParserR_BRACE)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullExprContext struct {
	*ExprContext
}

func NewNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprContext {
	var p = new(NullExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullExprContext struct {
	*ExprContext
}

func NewFullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullExprContext {
	var p = new(FullExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullExprContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullExprContext) ExprType() IExprTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprTypeContext)
}

func (s *FullExprContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, arishemParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(212)
			p.ExprType()
		}
		{
			p.SetState(213)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprTypeContext is an interface to support dynamic dispatch.
type IExprTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprTypeContext differentiates from other interfaces.
	IsExprTypeContext()
}

type ExprTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprTypeContext() *ExprTypeContext {
	var p = new(ExprTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_exprType
	return p
}

func (*ExprTypeContext) IsExprTypeContext() {}

func NewExprTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprTypeContext {
	var p = new(ExprTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_exprType

	return p
}

func (s *ExprTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprTypeContext) CopyFrom(ctx *ExprTypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FeatureExprTypeContext struct {
	*ExprTypeContext
}

func NewFeatureExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FeatureExprTypeContext {
	var p = new(FeatureExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *FeatureExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureExprTypeContext) FeatureExprKey() IFeatureExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeatureExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeatureExprKeyContext)
}

func (s *FeatureExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *FeatureExprTypeContext) Feature() IFeatureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeatureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeatureContext)
}

func (s *FeatureExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFeatureExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstExprTypeContext struct {
	*ExprTypeContext
}

func NewConstExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstExprTypeContext {
	var p = new(ConstExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *ConstExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExprTypeContext) ConstKey() IConstKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstKeyContext)
}

func (s *ConstExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ConstExprTypeContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConstExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListExprTypeContext struct {
	*ExprTypeContext
}

func NewListExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExprTypeContext {
	var p = new(ListExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *ListExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExprTypeContext) ListExprKey() IListExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListExprKeyContext)
}

func (s *ListExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ListExprTypeContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *ListExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapExprTypeContext struct {
	*ExprTypeContext
}

func NewMapExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapExprTypeContext {
	var p = new(MapExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *MapExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprTypeContext) MapExprKey() IMapExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapExprKeyContext)
}

func (s *MapExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *MapExprTypeContext) KeyValues() IKeyValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValuesContext)
}

func (s *MapExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitMapExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarExprTypeContext struct {
	*ExprTypeContext
}

func NewVarExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarExprTypeContext {
	var p = new(VarExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *VarExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprTypeContext) VarExprKey() IVarExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarExprKeyContext)
}

func (s *VarExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *VarExprTypeContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VarExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitVarExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExprTypeContext struct {
	*ExprTypeContext
}

func NewNullExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprTypeContext {
	var p = new(NullExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *NullExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprTypeContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstListExprTypeContext struct {
	*ExprTypeContext
}

func NewConstListExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstListExprTypeContext {
	var p = new(ConstListExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *ConstListExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstListExprTypeContext) ConstListKey() IConstListKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstListKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstListKeyContext)
}

func (s *ConstListExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ConstListExprTypeContext) ConstantList() IConstantListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantListContext)
}

func (s *ConstListExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConstListExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type MathExprTypeContext struct {
	*ExprTypeContext
}

func NewMathExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathExprTypeContext {
	var p = new(MathExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *MathExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExprTypeContext) MathExprKey() IMathExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathExprKeyContext)
}

func (s *MathExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *MathExprTypeContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MathExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitMathExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncExprTypeContext struct {
	*ExprTypeContext
}

func NewFuncExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExprTypeContext {
	var p = new(FuncExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *FuncExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprTypeContext) FuncExprKey() IFuncExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncExprKeyContext)
}

func (s *FuncExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *FuncExprTypeContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FuncExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFuncExprType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ExprType() (localctx IExprTypeContext) {
	this := p
	_ = this

	localctx = NewExprTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, arishemParserRULE_exprType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.ListExprKey()
		}
		{
			p.SetState(219)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(220)
			p.Exprs()
		}

	case 2:
		localctx = NewMapExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.MapExprKey()
		}
		{
			p.SetState(223)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(224)
			p.KeyValues()
		}

	case 3:
		localctx = NewConstExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.ConstKey()
		}
		{
			p.SetState(227)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(228)
			p.Constant()
		}

	case 4:
		localctx = NewConstListExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(230)
			p.ConstListKey()
		}
		{
			p.SetState(231)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(232)
			p.ConstantList()
		}

	case 5:
		localctx = NewMathExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(234)
			p.MathExprKey()
		}
		{
			p.SetState(235)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(236)
			p.Math()
		}

	case 6:
		localctx = NewFuncExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(238)
			p.FuncExprKey()
		}
		{
			p.SetState(239)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(240)
			p.Function()
		}

	case 7:
		localctx = NewVarExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(242)
			p.VarExprKey()
		}
		{
			p.SetState(243)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(244)
			p.Variable()
		}

	case 8:
		localctx = NewFeatureExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(246)
			p.FeatureExprKey()
		}
		{
			p.SetState(247)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(248)
			p.Feature()
		}

	case 9:
		localctx = NewNullExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(250)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) CopyFrom(ctx *VariableContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullVarContext struct {
	*VariableContext
}

func NewNullVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullVarContext {
	var p = new(NullVarContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *NullVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullVarContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullVarContext struct {
	*VariableContext
}

func NewFullVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullVarContext {
	var p = new(FullVarContext)

	p.VariableContext = NewEmptyVariableContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VariableContext))

	return p
}

func (s *FullVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullVarContext) VarPath() IVarPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarPathContext)
}

func (s *FullVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, arishemParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserQUOTATION:
		localctx = NewFullVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.VarPath()
		}

	case arishemParserNULL:
		localctx = NewNullVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMathContext is an interface to support dynamic dispatch.
type IMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMathContext differentiates from other interfaces.
	IsMathContext()
}

type MathContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathContext() *MathContext {
	var p = new(MathContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_math
	return p
}

func (*MathContext) IsMathContext() {}

func NewMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathContext {
	var p = new(MathContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_math

	return p
}

func (s *MathContext) GetParser() antlr.Parser { return s.parser }

func (s *MathContext) CopyFrom(ctx *MathContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullMathContext struct {
	*MathContext
}

func NewNullMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullMathContext {
	var p = new(NullMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *NullMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullMathContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullMath(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullMathContext struct {
	*MathContext
}

func NewFullMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullMathContext {
	var p = new(FullMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *FullMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullMathContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullMathContext) OpMathPair() IOpMathPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMathPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMathPairContext)
}

func (s *FullMathContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *FullMathContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *FullMathContext) LhsPair() ILhsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsPairContext)
}

func (s *FullMathContext) RhsPair() IRhsPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsPairContext)
}

func (s *FullMathContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullMath(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListMathContext struct {
	*MathContext
}

func NewListMathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListMathContext {
	var p = new(ListMathContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *ListMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListMathContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ListMathContext) OpMathPair() IOpMathPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMathPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMathPairContext)
}

func (s *ListMathContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *ListMathContext) ParamListKey() IParamListKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListKeyContext)
}

func (s *ListMathContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ListMathContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *ListMathContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ListMathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListMath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Math() (localctx IMathContext) {
	this := p
	_ = this

	localctx = NewMathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, arishemParserRULE_math)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFullMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(258)
			p.OpMathPair()
		}
		{
			p.SetState(259)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(260)
			p.LhsPair()
		}
		{
			p.SetState(261)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(262)
			p.RhsPair()
		}
		{
			p.SetState(263)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewListMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(266)
			p.OpMathPair()
		}
		{
			p.SetState(267)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(268)
			p.ParamListKey()
		}
		{
			p.SetState(269)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(270)
			p.Exprs()
		}
		{
			p.SetState(271)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CopyFrom(ctx *FunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullFuncContext struct {
	*FunctionContext
}

func NewNullFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullFuncContext {
	var p = new(NullFuncContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *NullFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullFuncContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamListFuncContext struct {
	*FunctionContext
}

func NewParamListFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListFuncContext {
	var p = new(ParamListFuncContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *ParamListFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListFuncContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ParamListFuncContext) FuncNamePair() IFuncNamePairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncNamePairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncNamePairContext)
}

func (s *ParamListFuncContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ParamListFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *ParamListFuncContext) ParamListKey() IParamListKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListKeyContext)
}

func (s *ParamListFuncContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ParamListFuncContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *ParamListFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamListFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamMapFuncContext struct {
	*FunctionContext
}

func NewParamMapFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamMapFuncContext {
	var p = new(ParamMapFuncContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *ParamMapFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamMapFuncContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *ParamMapFuncContext) FuncNamePair() IFuncNamePairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncNamePairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncNamePairContext)
}

func (s *ParamMapFuncContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *ParamMapFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *ParamMapFuncContext) ParamMapKey() IParamMapKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamMapKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamMapKeyContext)
}

func (s *ParamMapFuncContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ParamMapFuncContext) KeyValues() IKeyValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValuesContext)
}

func (s *ParamMapFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamMapFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, arishemParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamMapFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(277)
			p.FuncNamePair()
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(278)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(279)
				p.ParamMapKey()
			}
			{
				p.SetState(280)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(281)
				p.KeyValues()
			}

		}
		{
			p.SetState(285)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewParamListFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(288)
			p.FuncNamePair()
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(289)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(290)
				p.ParamListKey()
			}
			{
				p.SetState(291)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(292)
				p.Exprs()
			}

		}
		{
			p.SetState(296)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(298)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IExprsContext is an interface to support dynamic dispatch.
type IExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprsContext differentiates from other interfaces.
	IsExprsContext()
}

type ExprsContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprsContext() *ExprsContext {
	var p = new(ExprsContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_exprs
	return p
}

func (*ExprsContext) IsExprsContext() {}

func NewExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprsContext {
	var p = new(ExprsContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_exprs

	return p
}

func (s *ExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprsContext) CopyFrom(ctx *ExprsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullExprsContext struct {
	*ExprsContext
}

func NewFullExprsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullExprsContext {
	var p = new(FullExprsContext)

	p.ExprsContext = NewEmptyExprsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprsContext))

	return p
}

func (s *FullExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullExprsContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *FullExprsContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *FullExprsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FullExprsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FullExprsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *FullExprsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *FullExprsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullExprs(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullExprsContext struct {
	*ExprsContext
}

func NewNullExprsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullExprsContext {
	var p = new(NullExprsContext)

	p.ExprsContext = NewEmptyExprsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprsContext))

	return p
}

func (s *NullExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullExprsContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullExprsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullExprs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Exprs() (localctx IExprsContext) {
	this := p
	_ = this

	localctx = NewExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, arishemParserRULE_exprs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACKET:
		localctx = NewFullExprsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserL_BRACE {
			{
				p.SetState(302)
				p.Expr()
			}
			p.SetState(307)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(303)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(304)
					p.Expr()
				}

				p.SetState(309)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(312)
			p.Match(arishemParserR_BRACKET)
		}

	case arishemParserNULL:
		localctx = NewNullExprsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFeatureContext is an interface to support dynamic dispatch.
type IFeatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFeatureContext differentiates from other interfaces.
	IsFeatureContext()
}

type FeatureContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureContext() *FeatureContext {
	var p = new(FeatureContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_feature
	return p
}

func (*FeatureContext) IsFeatureContext() {}

func NewFeatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureContext {
	var p = new(FeatureContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_feature

	return p
}

func (s *FeatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureContext) CopyFrom(ctx *FeatureContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullFeatureContext struct {
	*FeatureContext
}

func NewNullFeatureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullFeatureContext {
	var p = new(NullFeatureContext)

	p.FeatureContext = NewEmptyFeatureContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FeatureContext))

	return p
}

func (s *NullFeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullFeatureContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullFeatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullFeature(s)

	default:
		return t.VisitChildren(s)
	}
}

type FullFeatureContext struct {
	*FeatureContext
}

func NewFullFeatureContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullFeatureContext {
	var p = new(FullFeatureContext)

	p.FeatureContext = NewEmptyFeatureContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FeatureContext))

	return p
}

func (s *FullFeatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullFeatureContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullFeatureContext) FeaturePathKey() IFeaturePathKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeaturePathKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeaturePathKeyContext)
}

func (s *FullFeatureContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOLON)
}

func (s *FullFeatureContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, i)
}

func (s *FullFeatureContext) FeaturePath() IFeaturePathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeaturePathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeaturePathContext)
}

func (s *FullFeatureContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullFeatureContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *FullFeatureContext) BuiltinParamKey() IBuiltinParamKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinParamKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltinParamKeyContext)
}

func (s *FullFeatureContext) KeyValues() IKeyValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValuesContext)
}

func (s *FullFeatureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullFeature(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Feature() (localctx IFeatureContext) {
	this := p
	_ = this

	localctx = NewFeatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, arishemParserRULE_feature)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(330)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullFeatureContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(317)
			p.FeaturePathKey()
		}
		{
			p.SetState(318)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(319)
			p.FeaturePath()
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(320)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(321)
				p.BuiltinParamKey()
			}
			{
				p.SetState(322)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(323)
				p.KeyValues()
			}

		}
		{
			p.SetState(327)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullFeatureContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CopyFrom(ctx *ConstantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolConstContext struct {
	*ConstantContext
}

func NewBoolConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolConstContext {
	var p = new(BoolConstContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *BoolConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolConstContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *BoolConstContext) BoolConstKey() IBoolConstKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolConstKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolConstKeyContext)
}

func (s *BoolConstContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *BoolConstContext) BoolType() IBoolTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolTypeContext)
}

func (s *BoolConstContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *BoolConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitBoolConst(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumConstContext struct {
	*ConstantContext
}

func NewNumConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumConstContext {
	var p = new(NumConstContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *NumConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumConstContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *NumConstContext) NumConstKey() INumConstKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumConstKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumConstKeyContext)
}

func (s *NumConstContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *NumConstContext) NumberType() INumberTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberTypeContext)
}

func (s *NumConstContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *NumConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNumConst(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrConstContext struct {
	*ConstantContext
}

func NewStrConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrConstContext {
	var p = new(StrConstContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *StrConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrConstContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *StrConstContext) StrConstKey() IStrConstKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrConstKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrConstKeyContext)
}

func (s *StrConstContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *StrConstContext) StringType() IStringTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringTypeContext)
}

func (s *StrConstContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *StrConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitStrConst(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullConstContext struct {
	*ConstantContext
}

func NewNullConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullConstContext {
	var p = new(NullConstContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *NullConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullConstContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullConstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullConst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, arishemParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(333)
			p.NumConstKey()
		}
		{
			p.SetState(334)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(335)
			p.NumberType()
		}
		{
			p.SetState(336)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewStrConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(338)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(339)
			p.StrConstKey()
		}
		{
			p.SetState(340)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(341)
			p.StringType()
		}
		{
			p.SetState(342)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewBoolConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(344)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(345)
			p.BoolConstKey()
		}
		{
			p.SetState(346)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(347)
			p.BoolType()
		}
		{
			p.SetState(348)
			p.Match(arishemParserR_BRACE)
		}

	case 4:
		localctx = NewNullConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(350)
			p.Match(arishemParserNULL)
		}

	}

	return localctx
}

// IConstantListContext is an interface to support dynamic dispatch.
type IConstantListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConstantListContext differentiates from other interfaces.
	IsConstantListContext()
}

type ConstantListContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantListContext() *ConstantListContext {
	var p = new(ConstantListContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_constantList
	return p
}

func (*ConstantListContext) IsConstantListContext() {}

func NewConstantListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantListContext {
	var p = new(ConstantListContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_constantList

	return p
}

func (s *ConstantListContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantListContext) CopyFrom(ctx *ConstantListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullConstListContext struct {
	*ConstantListContext
}

func NewFullConstListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullConstListContext {
	var p = new(FullConstListContext)

	p.ConstantListContext = NewEmptyConstantListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantListContext))

	return p
}

func (s *FullConstListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullConstListContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *FullConstListContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *FullConstListContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *FullConstListContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FullConstListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *FullConstListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *FullConstListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullConstList(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullConstListContext struct {
	*ConstantListContext
}

func NewNullConstListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullConstListContext {
	var p = new(NullConstListContext)

	p.ConstantListContext = NewEmptyConstantListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantListContext))

	return p
}

func (s *NullConstListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullConstListContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullConstListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullConstList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConstantList() (localctx IConstantListContext) {
	this := p
	_ = this

	localctx = NewConstantListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, arishemParserRULE_constantList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACKET:
		localctx = NewFullConstListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(353)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserL_BRACE {
			{
				p.SetState(354)
				p.Constant()
			}
			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(355)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(356)
					p.Constant()
				}

				p.SetState(361)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(364)
			p.Match(arishemParserR_BRACKET)
		}

	case arishemParserNULL:
		localctx = NewNullConstListContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionsKeyContext is an interface to support dynamic dispatch.
type IConditionsKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsConditionsKeyContext differentiates from other interfaces.
	IsConditionsKeyContext()
}

type ConditionsKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsKeyContext() *ConditionsKeyContext {
	var p = new(ConditionsKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionsKey
	return p
}

func (*ConditionsKeyContext) IsConditionsKeyContext() {}

func NewConditionsKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsKeyContext {
	var p = new(ConditionsKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionsKey

	return p
}

func (s *ConditionsKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ConditionsKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ConditionsKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionsKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionsKey() (localctx IConditionsKeyContext) {
	this := p
	_ = this

	localctx = NewConditionsKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, arishemParserRULE_conditionsKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(369)
		p.Match(arishemParserT__0)
	}
	{
		p.SetState(370)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IConditionsValContext is an interface to support dynamic dispatch.
type IConditionsValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext
	R_BRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsConditionsValContext differentiates from other interfaces.
	IsConditionsValContext()
}

type ConditionsValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsValContext() *ConditionsValContext {
	var p = new(ConditionsValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionsVal
	return p
}

func (*ConditionsValContext) IsConditionsValContext() {}

func NewConditionsValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsValContext {
	var p = new(ConditionsValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionsVal

	return p
}

func (s *ConditionsValContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsValContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *ConditionsValContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionsValContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsValContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *ConditionsValContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *ConditionsValContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *ConditionsValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionsVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionsVal() (localctx IConditionsValContext) {
	this := p
	_ = this

	localctx = NewConditionsValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, arishemParserRULE_conditionsVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(arishemParserL_BRACKET)
	}
	{
		p.SetState(373)
		p.Condition()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserCOMMA {
		{
			p.SetState(374)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(375)
			p.Condition()
		}

		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(381)
		p.Match(arishemParserR_BRACKET)
	}

	return localctx
}

// IConditionsPairContext is an interface to support dynamic dispatch.
type IConditionsPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionsKey() IConditionsKeyContext
	COLON() antlr.TerminalNode
	ConditionsVal() IConditionsValContext

	// IsConditionsPairContext differentiates from other interfaces.
	IsConditionsPairContext()
}

type ConditionsPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsPairContext() *ConditionsPairContext {
	var p = new(ConditionsPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionsPair
	return p
}

func (*ConditionsPairContext) IsConditionsPairContext() {}

func NewConditionsPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsPairContext {
	var p = new(ConditionsPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionsPair

	return p
}

func (s *ConditionsPairContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsPairContext) ConditionsKey() IConditionsKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsKeyContext)
}

func (s *ConditionsPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ConditionsPairContext) ConditionsVal() IConditionsValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionsValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionsValContext)
}

func (s *ConditionsPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionsPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionsPair() (localctx IConditionsPairContext) {
	this := p
	_ = this

	localctx = NewConditionsPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, arishemParserRULE_conditionsPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.ConditionsKey()
	}
	{
		p.SetState(384)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(385)
		p.ConditionsVal()
	}

	return localctx
}

// IConditionGroupsKeyContext is an interface to support dynamic dispatch.
type IConditionGroupsKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsConditionGroupsKeyContext differentiates from other interfaces.
	IsConditionGroupsKeyContext()
}

type ConditionGroupsKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionGroupsKeyContext() *ConditionGroupsKeyContext {
	var p = new(ConditionGroupsKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionGroupsKey
	return p
}

func (*ConditionGroupsKeyContext) IsConditionGroupsKeyContext() {}

func NewConditionGroupsKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionGroupsKeyContext {
	var p = new(ConditionGroupsKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionGroupsKey

	return p
}

func (s *ConditionGroupsKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionGroupsKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ConditionGroupsKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ConditionGroupsKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupsKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionGroupsKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionGroupsKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionGroupsKey() (localctx IConditionGroupsKeyContext) {
	this := p
	_ = this

	localctx = NewConditionGroupsKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, arishemParserRULE_conditionGroupsKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(388)
		p.Match(arishemParserT__1)
	}
	{
		p.SetState(389)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IConditionGroupsValContext is an interface to support dynamic dispatch.
type IConditionGroupsValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	AllConditionGroup() []IConditionGroupContext
	ConditionGroup(i int) IConditionGroupContext
	R_BRACKET() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsConditionGroupsValContext differentiates from other interfaces.
	IsConditionGroupsValContext()
}

type ConditionGroupsValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionGroupsValContext() *ConditionGroupsValContext {
	var p = new(ConditionGroupsValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionGroupsVal
	return p
}

func (*ConditionGroupsValContext) IsConditionGroupsValContext() {}

func NewConditionGroupsValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionGroupsValContext {
	var p = new(ConditionGroupsValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionGroupsVal

	return p
}

func (s *ConditionGroupsValContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionGroupsValContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *ConditionGroupsValContext) AllConditionGroup() []IConditionGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionGroupContext); ok {
			len++
		}
	}

	tst := make([]IConditionGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionGroupContext); ok {
			tst[i] = t.(IConditionGroupContext)
			i++
		}
	}

	return tst
}

func (s *ConditionGroupsValContext) ConditionGroup(i int) IConditionGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionGroupContext)
}

func (s *ConditionGroupsValContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *ConditionGroupsValContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *ConditionGroupsValContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *ConditionGroupsValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupsValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionGroupsValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionGroupsVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionGroupsVal() (localctx IConditionGroupsValContext) {
	this := p
	_ = this

	localctx = NewConditionGroupsValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, arishemParserRULE_conditionGroupsVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(arishemParserL_BRACKET)
	}
	{
		p.SetState(392)
		p.ConditionGroup()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserCOMMA {
		{
			p.SetState(393)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(394)
			p.ConditionGroup()
		}

		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(400)
		p.Match(arishemParserR_BRACKET)
	}

	return localctx
}

// IConditionGroupsPairContext is an interface to support dynamic dispatch.
type IConditionGroupsPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionGroupsKey() IConditionGroupsKeyContext
	COLON() antlr.TerminalNode
	ConditionGroupsVal() IConditionGroupsValContext

	// IsConditionGroupsPairContext differentiates from other interfaces.
	IsConditionGroupsPairContext()
}

type ConditionGroupsPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionGroupsPairContext() *ConditionGroupsPairContext {
	var p = new(ConditionGroupsPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_conditionGroupsPair
	return p
}

func (*ConditionGroupsPairContext) IsConditionGroupsPairContext() {}

func NewConditionGroupsPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionGroupsPairContext {
	var p = new(ConditionGroupsPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_conditionGroupsPair

	return p
}

func (s *ConditionGroupsPairContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionGroupsPairContext) ConditionGroupsKey() IConditionGroupsKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionGroupsKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionGroupsKeyContext)
}

func (s *ConditionGroupsPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ConditionGroupsPairContext) ConditionGroupsVal() IConditionGroupsValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionGroupsValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionGroupsValContext)
}

func (s *ConditionGroupsPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupsPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionGroupsPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConditionGroupsPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConditionGroupsPair() (localctx IConditionGroupsPairContext) {
	this := p
	_ = this

	localctx = NewConditionGroupsPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, arishemParserRULE_conditionGroupsPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.ConditionGroupsKey()
	}
	{
		p.SetState(403)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(404)
		p.ConditionGroupsVal()
	}

	return localctx
}

// IListExprKeyContext is an interface to support dynamic dispatch.
type IListExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsListExprKeyContext differentiates from other interfaces.
	IsListExprKeyContext()
}

type ListExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyListExprKeyContext() *ListExprKeyContext {
	var p = new(ListExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_listExprKey
	return p
}

func (*ListExprKeyContext) IsListExprKeyContext() {}

func NewListExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListExprKeyContext {
	var p = new(ListExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_listExprKey

	return p
}

func (s *ListExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ListExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ListExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ListExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ListExprKey() (localctx IListExprKeyContext) {
	this := p
	_ = this

	localctx = NewListExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, arishemParserRULE_listExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(407)
		p.Match(arishemParserT__2)
	}
	{
		p.SetState(408)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IMapExprKeyContext is an interface to support dynamic dispatch.
type IMapExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsMapExprKeyContext differentiates from other interfaces.
	IsMapExprKeyContext()
}

type MapExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapExprKeyContext() *MapExprKeyContext {
	var p = new(MapExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_mapExprKey
	return p
}

func (*MapExprKeyContext) IsMapExprKeyContext() {}

func NewMapExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapExprKeyContext {
	var p = new(MapExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_mapExprKey

	return p
}

func (s *MapExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MapExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *MapExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *MapExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitMapExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) MapExprKey() (localctx IMapExprKeyContext) {
	this := p
	_ = this

	localctx = NewMapExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, arishemParserRULE_mapExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(411)
		p.Match(arishemParserT__3)
	}
	{
		p.SetState(412)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IMathExprKeyContext is an interface to support dynamic dispatch.
type IMathExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsMathExprKeyContext differentiates from other interfaces.
	IsMathExprKeyContext()
}

type MathExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExprKeyContext() *MathExprKeyContext {
	var p = new(MathExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_mathExprKey
	return p
}

func (*MathExprKeyContext) IsMathExprKeyContext() {}

func NewMathExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExprKeyContext {
	var p = new(MathExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_mathExprKey

	return p
}

func (s *MathExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *MathExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *MathExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitMathExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) MathExprKey() (localctx IMathExprKeyContext) {
	this := p
	_ = this

	localctx = NewMathExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, arishemParserRULE_mathExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(415)
		p.Match(arishemParserT__4)
	}
	{
		p.SetState(416)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IFuncExprKeyContext is an interface to support dynamic dispatch.
type IFuncExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsFuncExprKeyContext differentiates from other interfaces.
	IsFuncExprKeyContext()
}

type FuncExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExprKeyContext() *FuncExprKeyContext {
	var p = new(FuncExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_funcExprKey
	return p
}

func (*FuncExprKeyContext) IsFuncExprKeyContext() {}

func NewFuncExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExprKeyContext {
	var p = new(FuncExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_funcExprKey

	return p
}

func (s *FuncExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *FuncExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *FuncExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFuncExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FuncExprKey() (localctx IFuncExprKeyContext) {
	this := p
	_ = this

	localctx = NewFuncExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, arishemParserRULE_funcExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(419)
		p.Match(arishemParserT__5)
	}
	{
		p.SetState(420)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IVarExprKeyContext is an interface to support dynamic dispatch.
type IVarExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsVarExprKeyContext differentiates from other interfaces.
	IsVarExprKeyContext()
}

type VarExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarExprKeyContext() *VarExprKeyContext {
	var p = new(VarExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_varExprKey
	return p
}

func (*VarExprKeyContext) IsVarExprKeyContext() {}

func NewVarExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarExprKeyContext {
	var p = new(VarExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_varExprKey

	return p
}

func (s *VarExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *VarExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *VarExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *VarExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitVarExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) VarExprKey() (localctx IVarExprKeyContext) {
	this := p
	_ = this

	localctx = NewVarExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, arishemParserRULE_varExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(423)
		p.Match(arishemParserT__6)
	}
	{
		p.SetState(424)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IFeatureExprKeyContext is an interface to support dynamic dispatch.
type IFeatureExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsFeatureExprKeyContext differentiates from other interfaces.
	IsFeatureExprKeyContext()
}

type FeatureExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeatureExprKeyContext() *FeatureExprKeyContext {
	var p = new(FeatureExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_featureExprKey
	return p
}

func (*FeatureExprKeyContext) IsFeatureExprKeyContext() {}

func NewFeatureExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeatureExprKeyContext {
	var p = new(FeatureExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_featureExprKey

	return p
}

func (s *FeatureExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *FeatureExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *FeatureExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *FeatureExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeatureExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeatureExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFeatureExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FeatureExprKey() (localctx IFeatureExprKeyContext) {
	this := p
	_ = this

	localctx = NewFeatureExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, arishemParserRULE_featureExprKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(426)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(427)
		p.Match(arishemParserT__7)
	}
	{
		p.SetState(428)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IFeaturePathKeyContext is an interface to support dynamic dispatch.
type IFeaturePathKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsFeaturePathKeyContext differentiates from other interfaces.
	IsFeaturePathKeyContext()
}

type FeaturePathKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeaturePathKeyContext() *FeaturePathKeyContext {
	var p = new(FeaturePathKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_featurePathKey
	return p
}

func (*FeaturePathKeyContext) IsFeaturePathKeyContext() {}

func NewFeaturePathKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeaturePathKeyContext {
	var p = new(FeaturePathKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_featurePathKey

	return p
}

func (s *FeaturePathKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *FeaturePathKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *FeaturePathKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *FeaturePathKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeaturePathKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeaturePathKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFeaturePathKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FeaturePathKey() (localctx IFeaturePathKeyContext) {
	this := p
	_ = this

	localctx = NewFeaturePathKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, arishemParserRULE_featurePathKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(431)
		p.Match(arishemParserT__8)
	}
	{
		p.SetState(432)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IConstKeyContext is an interface to support dynamic dispatch.
type IConstKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsConstKeyContext differentiates from other interfaces.
	IsConstKeyContext()
}

type ConstKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstKeyContext() *ConstKeyContext {
	var p = new(ConstKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_constKey
	return p
}

func (*ConstKeyContext) IsConstKeyContext() {}

func NewConstKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstKeyContext {
	var p = new(ConstKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_constKey

	return p
}

func (s *ConstKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ConstKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ConstKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConstKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConstKey() (localctx IConstKeyContext) {
	this := p
	_ = this

	localctx = NewConstKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, arishemParserRULE_constKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(435)
		p.Match(arishemParserT__9)
	}
	{
		p.SetState(436)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IConstListKeyContext is an interface to support dynamic dispatch.
type IConstListKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsConstListKeyContext differentiates from other interfaces.
	IsConstListKeyContext()
}

type ConstListKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstListKeyContext() *ConstListKeyContext {
	var p = new(ConstListKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_constListKey
	return p
}

func (*ConstListKeyContext) IsConstListKeyContext() {}

func NewConstListKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstListKeyContext {
	var p = new(ConstListKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_constListKey

	return p
}

func (s *ConstListKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstListKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ConstListKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ConstListKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstListKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstListKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitConstListKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ConstListKey() (localctx IConstListKeyContext) {
	this := p
	_ = this

	localctx = NewConstListKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, arishemParserRULE_constListKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(439)
		p.Match(arishemParserT__10)
	}
	{
		p.SetState(440)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IVarPathContext is an interface to support dynamic dispatch.
type IVarPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	VarPathVal() IVarPathValContext

	// IsVarPathContext differentiates from other interfaces.
	IsVarPathContext()
}

type VarPathContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarPathContext() *VarPathContext {
	var p = new(VarPathContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_varPath
	return p
}

func (*VarPathContext) IsVarPathContext() {}

func NewVarPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarPathContext {
	var p = new(VarPathContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_varPath

	return p
}

func (s *VarPathContext) GetParser() antlr.Parser { return s.parser }

func (s *VarPathContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *VarPathContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *VarPathContext) VarPathVal() IVarPathValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarPathValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarPathValContext)
}

func (s *VarPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitVarPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) VarPath() (localctx IVarPathContext) {
	this := p
	_ = this

	localctx = NewVarPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, arishemParserRULE_varPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(443)
		p.VarPathVal()
	}
	{
		p.SetState(444)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IVarPathValContext is an interface to support dynamic dispatch.
type IVarPathValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME_KEY_TYPE() []antlr.TerminalNode
	NAME_KEY_TYPE(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsVarPathValContext differentiates from other interfaces.
	IsVarPathValContext()
}

type VarPathValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarPathValContext() *VarPathValContext {
	var p = new(VarPathValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_varPathVal
	return p
}

func (*VarPathValContext) IsVarPathValContext() {}

func NewVarPathValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarPathValContext {
	var p = new(VarPathValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_varPathVal

	return p
}

func (s *VarPathValContext) GetParser() antlr.Parser { return s.parser }

func (s *VarPathValContext) AllNAME_KEY_TYPE() []antlr.TerminalNode {
	return s.GetTokens(arishemParserNAME_KEY_TYPE)
}

func (s *VarPathValContext) NAME_KEY_TYPE(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserNAME_KEY_TYPE, i)
}

func (s *VarPathValContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(arishemParserDOT)
}

func (s *VarPathValContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserDOT, i)
}

func (s *VarPathValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarPathValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarPathValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitVarPathVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) VarPathVal() (localctx IVarPathValContext) {
	this := p
	_ = this

	localctx = NewVarPathValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, arishemParserRULE_varPathVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserDOT {
		{
			p.SetState(447)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(448)
			p.Match(arishemParserNAME_KEY_TYPE)
		}

		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFeaturePathContext is an interface to support dynamic dispatch.
type IFeaturePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	FeaturePathVal() IFeaturePathValContext

	// IsFeaturePathContext differentiates from other interfaces.
	IsFeaturePathContext()
}

type FeaturePathContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeaturePathContext() *FeaturePathContext {
	var p = new(FeaturePathContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_featurePath
	return p
}

func (*FeaturePathContext) IsFeaturePathContext() {}

func NewFeaturePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeaturePathContext {
	var p = new(FeaturePathContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_featurePath

	return p
}

func (s *FeaturePathContext) GetParser() antlr.Parser { return s.parser }

func (s *FeaturePathContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *FeaturePathContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *FeaturePathContext) FeaturePathVal() IFeaturePathValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFeaturePathValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFeaturePathValContext)
}

func (s *FeaturePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeaturePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeaturePathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFeaturePath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FeaturePath() (localctx IFeaturePathContext) {
	this := p
	_ = this

	localctx = NewFeaturePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, arishemParserRULE_featurePath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(455)
		p.FeaturePathVal()
	}
	{
		p.SetState(456)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IFeaturePathValContext is an interface to support dynamic dispatch.
type IFeaturePathValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNAME_KEY_TYPE() []antlr.TerminalNode
	NAME_KEY_TYPE(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFeaturePathValContext differentiates from other interfaces.
	IsFeaturePathValContext()
}

type FeaturePathValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFeaturePathValContext() *FeaturePathValContext {
	var p = new(FeaturePathValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_featurePathVal
	return p
}

func (*FeaturePathValContext) IsFeaturePathValContext() {}

func NewFeaturePathValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FeaturePathValContext {
	var p = new(FeaturePathValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_featurePathVal

	return p
}

func (s *FeaturePathValContext) GetParser() antlr.Parser { return s.parser }

func (s *FeaturePathValContext) AllNAME_KEY_TYPE() []antlr.TerminalNode {
	return s.GetTokens(arishemParserNAME_KEY_TYPE)
}

func (s *FeaturePathValContext) NAME_KEY_TYPE(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserNAME_KEY_TYPE, i)
}

func (s *FeaturePathValContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(arishemParserDOT)
}

func (s *FeaturePathValContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserDOT, i)
}

func (s *FeaturePathValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FeaturePathValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FeaturePathValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFeaturePathVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FeaturePathVal() (localctx IFeaturePathValContext) {
	this := p
	_ = this

	localctx = NewFeaturePathValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, arishemParserRULE_featurePathVal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == arishemParserDOT {
		{
			p.SetState(459)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(460)
			p.Match(arishemParserNAME_KEY_TYPE)
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamMapKeyContext is an interface to support dynamic dispatch.
type IParamMapKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsParamMapKeyContext differentiates from other interfaces.
	IsParamMapKeyContext()
}

type ParamMapKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamMapKeyContext() *ParamMapKeyContext {
	var p = new(ParamMapKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_paramMapKey
	return p
}

func (*ParamMapKeyContext) IsParamMapKeyContext() {}

func NewParamMapKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamMapKeyContext {
	var p = new(ParamMapKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_paramMapKey

	return p
}

func (s *ParamMapKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamMapKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ParamMapKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ParamMapKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamMapKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamMapKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamMapKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ParamMapKey() (localctx IParamMapKeyContext) {
	this := p
	_ = this

	localctx = NewParamMapKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, arishemParserRULE_paramMapKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(466)
		p.Match(arishemParserT__11)
	}
	{
		p.SetState(467)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IParamListKeyContext is an interface to support dynamic dispatch.
type IParamListKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsParamListKeyContext differentiates from other interfaces.
	IsParamListKeyContext()
}

type ParamListKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListKeyContext() *ParamListKeyContext {
	var p = new(ParamListKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_paramListKey
	return p
}

func (*ParamListKeyContext) IsParamListKeyContext() {}

func NewParamListKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListKeyContext {
	var p = new(ParamListKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_paramListKey

	return p
}

func (s *ParamListKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ParamListKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ParamListKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitParamListKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ParamListKey() (localctx IParamListKeyContext) {
	this := p
	_ = this

	localctx = NewParamListKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, arishemParserRULE_paramListKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(470)
		p.Match(arishemParserT__12)
	}
	{
		p.SetState(471)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IBuiltinParamKeyContext is an interface to support dynamic dispatch.
type IBuiltinParamKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsBuiltinParamKeyContext differentiates from other interfaces.
	IsBuiltinParamKeyContext()
}

type BuiltinParamKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinParamKeyContext() *BuiltinParamKeyContext {
	var p = new(BuiltinParamKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_builtinParamKey
	return p
}

func (*BuiltinParamKeyContext) IsBuiltinParamKeyContext() {}

func NewBuiltinParamKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinParamKeyContext {
	var p = new(BuiltinParamKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_builtinParamKey

	return p
}

func (s *BuiltinParamKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinParamKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *BuiltinParamKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *BuiltinParamKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinParamKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinParamKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitBuiltinParamKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) BuiltinParamKey() (localctx IBuiltinParamKeyContext) {
	this := p
	_ = this

	localctx = NewBuiltinParamKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, arishemParserRULE_builtinParamKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(474)
		p.Match(arishemParserT__13)
	}
	{
		p.SetState(475)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IActionAimKeyContext is an interface to support dynamic dispatch.
type IActionAimKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsActionAimKeyContext differentiates from other interfaces.
	IsActionAimKeyContext()
}

type ActionAimKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAimKeyContext() *ActionAimKeyContext {
	var p = new(ActionAimKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_actionAimKey
	return p
}

func (*ActionAimKeyContext) IsActionAimKeyContext() {}

func NewActionAimKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAimKeyContext {
	var p = new(ActionAimKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_actionAimKey

	return p
}

func (s *ActionAimKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAimKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ActionAimKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ActionAimKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionAimKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitActionAimKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ActionAimKey() (localctx IActionAimKeyContext) {
	this := p
	_ = this

	localctx = NewActionAimKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, arishemParserRULE_actionAimKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(478)
		p.Match(arishemParserT__14)
	}
	{
		p.SetState(479)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IActionAimValContext is an interface to support dynamic dispatch.
type IActionAimValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	NAME_KEY_TYPE() antlr.TerminalNode

	// IsActionAimValContext differentiates from other interfaces.
	IsActionAimValContext()
}

type ActionAimValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAimValContext() *ActionAimValContext {
	var p = new(ActionAimValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_actionAimVal
	return p
}

func (*ActionAimValContext) IsActionAimValContext() {}

func NewActionAimValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAimValContext {
	var p = new(ActionAimValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_actionAimVal

	return p
}

func (s *ActionAimValContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAimValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ActionAimValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ActionAimValContext) NAME_KEY_TYPE() antlr.TerminalNode {
	return s.GetToken(arishemParserNAME_KEY_TYPE, 0)
}

func (s *ActionAimValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionAimValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitActionAimVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ActionAimVal() (localctx IActionAimValContext) {
	this := p
	_ = this

	localctx = NewActionAimValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, arishemParserRULE_actionAimVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(482)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	{
		p.SetState(483)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IActionAimPairContext is an interface to support dynamic dispatch.
type IActionAimPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ActionAimKey() IActionAimKeyContext
	COLON() antlr.TerminalNode
	ActionAimVal() IActionAimValContext

	// IsActionAimPairContext differentiates from other interfaces.
	IsActionAimPairContext()
}

type ActionAimPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionAimPairContext() *ActionAimPairContext {
	var p = new(ActionAimPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_actionAimPair
	return p
}

func (*ActionAimPairContext) IsActionAimPairContext() {}

func NewActionAimPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionAimPairContext {
	var p = new(ActionAimPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_actionAimPair

	return p
}

func (s *ActionAimPairContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionAimPairContext) ActionAimKey() IActionAimKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimKeyContext)
}

func (s *ActionAimPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ActionAimPairContext) ActionAimVal() IActionAimValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionAimValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionAimValContext)
}

func (s *ActionAimPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionAimPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionAimPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitActionAimPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ActionAimPair() (localctx IActionAimPairContext) {
	this := p
	_ = this

	localctx = NewActionAimPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, arishemParserRULE_actionAimPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.ActionAimKey()
	}
	{
		p.SetState(486)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(487)
		p.ActionAimVal()
	}

	return localctx
}

// INameKeyContext is an interface to support dynamic dispatch.
type INameKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	NAME_KEY_TYPE() antlr.TerminalNode

	// IsNameKeyContext differentiates from other interfaces.
	IsNameKeyContext()
}

type NameKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameKeyContext() *NameKeyContext {
	var p = new(NameKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_nameKey
	return p
}

func (*NameKeyContext) IsNameKeyContext() {}

func NewNameKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameKeyContext {
	var p = new(NameKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_nameKey

	return p
}

func (s *NameKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *NameKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *NameKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *NameKeyContext) NAME_KEY_TYPE() antlr.TerminalNode {
	return s.GetToken(arishemParserNAME_KEY_TYPE, 0)
}

func (s *NameKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNameKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) NameKey() (localctx INameKeyContext) {
	this := p
	_ = this

	localctx = NewNameKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, arishemParserRULE_nameKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(490)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	{
		p.SetState(491)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameKey() INameKeyContext
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) NameKey() INameKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameKeyContext)
}

func (s *KeyValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *KeyValueContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) KeyValue() (localctx IKeyValueContext) {
	this := p
	_ = this

	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, arishemParserRULE_keyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		p.NameKey()
	}
	{
		p.SetState(494)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(495)
		p.Expr()
	}

	return localctx
}

// IKeyValuesContext is an interface to support dynamic dispatch.
type IKeyValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeyValuesContext differentiates from other interfaces.
	IsKeyValuesContext()
}

type KeyValuesContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValuesContext() *KeyValuesContext {
	var p = new(KeyValuesContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_keyValues
	return p
}

func (*KeyValuesContext) IsKeyValuesContext() {}

func NewKeyValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValuesContext {
	var p = new(KeyValuesContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_keyValues

	return p
}

func (s *KeyValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValuesContext) CopyFrom(ctx *KeyValuesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *KeyValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullKeyValuesContext struct {
	*KeyValuesContext
}

func NewFullKeyValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullKeyValuesContext {
	var p = new(FullKeyValuesContext)

	p.KeyValuesContext = NewEmptyKeyValuesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*KeyValuesContext))

	return p
}

func (s *FullKeyValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullKeyValuesContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullKeyValuesContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullKeyValuesContext) AllKeyValue() []IKeyValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyValueContext); ok {
			len++
		}
	}

	tst := make([]IKeyValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyValueContext); ok {
			tst[i] = t.(IKeyValueContext)
			i++
		}
	}

	return tst
}

func (s *FullKeyValuesContext) KeyValue(i int) IKeyValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *FullKeyValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(arishemParserCOMMA)
}

func (s *FullKeyValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, i)
}

func (s *FullKeyValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullKeyValues(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullKeyValuesContext struct {
	*KeyValuesContext
}

func NewNullKeyValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullKeyValuesContext {
	var p = new(NullKeyValuesContext)

	p.KeyValuesContext = NewEmptyKeyValuesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*KeyValuesContext))

	return p
}

func (s *NullKeyValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullKeyValuesContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullKeyValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullKeyValues(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) KeyValues() (localctx IKeyValuesContext) {
	this := p
	_ = this

	localctx = NewKeyValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, arishemParserRULE_keyValues)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(510)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullKeyValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)
			p.Match(arishemParserL_BRACE)
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserQUOTATION {
			{
				p.SetState(498)
				p.KeyValue()
			}
			p.SetState(503)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(499)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(500)
					p.KeyValue()
				}

				p.SetState(505)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(508)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullKeyValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(509)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOpLogicKeyContext is an interface to support dynamic dispatch.
type IOpLogicKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsOpLogicKeyContext differentiates from other interfaces.
	IsOpLogicKeyContext()
}

type OpLogicKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpLogicKeyContext() *OpLogicKeyContext {
	var p = new(OpLogicKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opLogicKey
	return p
}

func (*OpLogicKeyContext) IsOpLogicKeyContext() {}

func NewOpLogicKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpLogicKeyContext {
	var p = new(OpLogicKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opLogicKey

	return p
}

func (s *OpLogicKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OpLogicKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OpLogicKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OpLogicKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpLogicKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpLogicKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpLogicKey() (localctx IOpLogicKeyContext) {
	this := p
	_ = this

	localctx = NewOpLogicKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, arishemParserRULE_opLogicKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(513)
		p.Match(arishemParserT__15)
	}
	{
		p.SetState(514)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOpLogicValContext is an interface to support dynamic dispatch.
type IOpLogicValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	OpLogic() IOpLogicContext

	// IsOpLogicValContext differentiates from other interfaces.
	IsOpLogicValContext()
}

type OpLogicValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpLogicValContext() *OpLogicValContext {
	var p = new(OpLogicValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opLogicVal
	return p
}

func (*OpLogicValContext) IsOpLogicValContext() {}

func NewOpLogicValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpLogicValContext {
	var p = new(OpLogicValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opLogicVal

	return p
}

func (s *OpLogicValContext) GetParser() antlr.Parser { return s.parser }

func (s *OpLogicValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OpLogicValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OpLogicValContext) OpLogic() IOpLogicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpLogicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpLogicContext)
}

func (s *OpLogicValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpLogicValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpLogicVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpLogicVal() (localctx IOpLogicValContext) {
	this := p
	_ = this

	localctx = NewOpLogicValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, arishemParserRULE_opLogicVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(516)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(517)
		p.OpLogic()
	}
	{
		p.SetState(518)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOpLogicPairContext is an interface to support dynamic dispatch.
type IOpLogicPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpLogicKey() IOpLogicKeyContext
	COLON() antlr.TerminalNode
	OpLogicVal() IOpLogicValContext

	// IsOpLogicPairContext differentiates from other interfaces.
	IsOpLogicPairContext()
}

type OpLogicPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpLogicPairContext() *OpLogicPairContext {
	var p = new(OpLogicPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opLogicPair
	return p
}

func (*OpLogicPairContext) IsOpLogicPairContext() {}

func NewOpLogicPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpLogicPairContext {
	var p = new(OpLogicPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opLogicPair

	return p
}

func (s *OpLogicPairContext) GetParser() antlr.Parser { return s.parser }

func (s *OpLogicPairContext) OpLogicKey() IOpLogicKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpLogicKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpLogicKeyContext)
}

func (s *OpLogicPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *OpLogicPairContext) OpLogicVal() IOpLogicValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpLogicValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpLogicValContext)
}

func (s *OpLogicPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpLogicPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpLogicPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpLogicPair() (localctx IOpLogicPairContext) {
	this := p
	_ = this

	localctx = NewOpLogicPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, arishemParserRULE_opLogicPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(520)
		p.OpLogicKey()
	}
	{
		p.SetState(521)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(522)
		p.OpLogicVal()
	}

	return localctx
}

// IOpLogicContext is an interface to support dynamic dispatch.
type IOpLogicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATOR_LOGIC_STR() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsOpLogicContext differentiates from other interfaces.
	IsOpLogicContext()
}

type OpLogicContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpLogicContext() *OpLogicContext {
	var p = new(OpLogicContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opLogic
	return p
}

func (*OpLogicContext) IsOpLogicContext() {}

func NewOpLogicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpLogicContext {
	var p = new(OpLogicContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opLogic

	return p
}

func (s *OpLogicContext) GetParser() antlr.Parser { return s.parser }

func (s *OpLogicContext) OPERATOR_LOGIC_STR() antlr.TerminalNode {
	return s.GetToken(arishemParserOPERATOR_LOGIC_STR, 0)
}

func (s *OpLogicContext) AND() antlr.TerminalNode {
	return s.GetToken(arishemParserAND, 0)
}

func (s *OpLogicContext) OR() antlr.TerminalNode {
	return s.GetToken(arishemParserOR, 0)
}

func (s *OpLogicContext) NOT() antlr.TerminalNode {
	return s.GetToken(arishemParserNOT, 0)
}

func (s *OpLogicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpLogicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpLogicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpLogic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpLogic() (localctx IOpLogicContext) {
	this := p
	_ = this

	localctx = NewOpLogicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, arishemParserRULE_opLogic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070485716620017664) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOpMathKeyContext is an interface to support dynamic dispatch.
type IOpMathKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsOpMathKeyContext differentiates from other interfaces.
	IsOpMathKeyContext()
}

type OpMathKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpMathKeyContext() *OpMathKeyContext {
	var p = new(OpMathKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opMathKey
	return p
}

func (*OpMathKeyContext) IsOpMathKeyContext() {}

func NewOpMathKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpMathKeyContext {
	var p = new(OpMathKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opMathKey

	return p
}

func (s *OpMathKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OpMathKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OpMathKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OpMathKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpMathKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpMathKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpMathKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpMathKey() (localctx IOpMathKeyContext) {
	this := p
	_ = this

	localctx = NewOpMathKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, arishemParserRULE_opMathKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(527)
		p.Match(arishemParserT__16)
	}
	{
		p.SetState(528)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOpMathValContext is an interface to support dynamic dispatch.
type IOpMathValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	MathOperator() IMathOperatorContext

	// IsOpMathValContext differentiates from other interfaces.
	IsOpMathValContext()
}

type OpMathValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpMathValContext() *OpMathValContext {
	var p = new(OpMathValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opMathVal
	return p
}

func (*OpMathValContext) IsOpMathValContext() {}

func NewOpMathValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpMathValContext {
	var p = new(OpMathValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opMathVal

	return p
}

func (s *OpMathValContext) GetParser() antlr.Parser { return s.parser }

func (s *OpMathValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OpMathValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OpMathValContext) MathOperator() IMathOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathOperatorContext)
}

func (s *OpMathValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpMathValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpMathValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpMathVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpMathVal() (localctx IOpMathValContext) {
	this := p
	_ = this

	localctx = NewOpMathValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, arishemParserRULE_opMathVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(531)
		p.MathOperator()
	}
	{
		p.SetState(532)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOpMathPairContext is an interface to support dynamic dispatch.
type IOpMathPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpMathKey() IOpMathKeyContext
	COLON() antlr.TerminalNode
	OpMathVal() IOpMathValContext

	// IsOpMathPairContext differentiates from other interfaces.
	IsOpMathPairContext()
}

type OpMathPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpMathPairContext() *OpMathPairContext {
	var p = new(OpMathPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_opMathPair
	return p
}

func (*OpMathPairContext) IsOpMathPairContext() {}

func NewOpMathPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpMathPairContext {
	var p = new(OpMathPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_opMathPair

	return p
}

func (s *OpMathPairContext) GetParser() antlr.Parser { return s.parser }

func (s *OpMathPairContext) OpMathKey() IOpMathKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMathKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMathKeyContext)
}

func (s *OpMathPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *OpMathPairContext) OpMathVal() IOpMathValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMathValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMathValContext)
}

func (s *OpMathPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpMathPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpMathPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOpMathPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OpMathPair() (localctx IOpMathPairContext) {
	this := p
	_ = this

	localctx = NewOpMathPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, arishemParserRULE_opMathPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.OpMathKey()
	}
	{
		p.SetState(535)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(536)
		p.OpMathVal()
	}

	return localctx
}

// IMathOperatorContext is an interface to support dynamic dispatch.
type IMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsMathOperatorContext differentiates from other interfaces.
	IsMathOperatorContext()
}

type MathOperatorContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathOperatorContext() *MathOperatorContext {
	var p = new(MathOperatorContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_mathOperator
	return p
}

func (*MathOperatorContext) IsMathOperatorContext() {}

func NewMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathOperatorContext {
	var p = new(MathOperatorContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_mathOperator

	return p
}

func (s *MathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(arishemParserPLUS, 0)
}

func (s *MathOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(arishemParserSUB, 0)
}

func (s *MathOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(arishemParserMUL, 0)
}

func (s *MathOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(arishemParserDIV, 0)
}

func (s *MathOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(arishemParserMOD, 0)
}

func (s *MathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) MathOperator() (localctx IMathOperatorContext) {
	this := p
	_ = this

	localctx = NewMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, arishemParserRULE_mathOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFuncNameKeyContext is an interface to support dynamic dispatch.
type IFuncNameKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsFuncNameKeyContext differentiates from other interfaces.
	IsFuncNameKeyContext()
}

type FuncNameKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNameKeyContext() *FuncNameKeyContext {
	var p = new(FuncNameKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_funcNameKey
	return p
}

func (*FuncNameKeyContext) IsFuncNameKeyContext() {}

func NewFuncNameKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNameKeyContext {
	var p = new(FuncNameKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_funcNameKey

	return p
}

func (s *FuncNameKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNameKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *FuncNameKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *FuncNameKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNameKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNameKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFuncNameKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FuncNameKey() (localctx IFuncNameKeyContext) {
	this := p
	_ = this

	localctx = NewFuncNameKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, arishemParserRULE_funcNameKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(541)
		p.Match(arishemParserT__17)
	}
	{
		p.SetState(542)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IFuncNamePairContext is an interface to support dynamic dispatch.
type IFuncNamePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncNameKey() IFuncNameKeyContext
	COLON() antlr.TerminalNode
	NameKey() INameKeyContext

	// IsFuncNamePairContext differentiates from other interfaces.
	IsFuncNamePairContext()
}

type FuncNamePairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncNamePairContext() *FuncNamePairContext {
	var p = new(FuncNamePairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_funcNamePair
	return p
}

func (*FuncNamePairContext) IsFuncNamePairContext() {}

func NewFuncNamePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncNamePairContext {
	var p = new(FuncNamePairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_funcNamePair

	return p
}

func (s *FuncNamePairContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncNamePairContext) FuncNameKey() IFuncNameKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncNameKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncNameKeyContext)
}

func (s *FuncNamePairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *FuncNamePairContext) NameKey() INameKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameKeyContext)
}

func (s *FuncNamePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncNamePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncNamePairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFuncNamePair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) FuncNamePair() (localctx IFuncNamePairContext) {
	this := p
	_ = this

	localctx = NewFuncNamePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, arishemParserRULE_funcNamePair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
		p.FuncNameKey()
	}
	{
		p.SetState(545)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(546)
		p.NameKey()
	}

	return localctx
}

// IOperatorKeyContext is an interface to support dynamic dispatch.
type IOperatorKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsOperatorKeyContext differentiates from other interfaces.
	IsOperatorKeyContext()
}

type OperatorKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorKeyContext() *OperatorKeyContext {
	var p = new(OperatorKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_operatorKey
	return p
}

func (*OperatorKeyContext) IsOperatorKeyContext() {}

func NewOperatorKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorKeyContext {
	var p = new(OperatorKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_operatorKey

	return p
}

func (s *OperatorKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OperatorKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OperatorKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOperatorKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OperatorKey() (localctx IOperatorKeyContext) {
	this := p
	_ = this

	localctx = NewOperatorKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, arishemParserRULE_operatorKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(548)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(549)
		p.Match(arishemParserT__18)
	}
	{
		p.SetState(550)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOperatorValContext is an interface to support dynamic dispatch.
type IOperatorValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	OperatorType() IOperatorTypeContext

	// IsOperatorValContext differentiates from other interfaces.
	IsOperatorValContext()
}

type OperatorValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorValContext() *OperatorValContext {
	var p = new(OperatorValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_operatorVal
	return p
}

func (*OperatorValContext) IsOperatorValContext() {}

func NewOperatorValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorValContext {
	var p = new(OperatorValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_operatorVal

	return p
}

func (s *OperatorValContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *OperatorValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *OperatorValContext) OperatorType() IOperatorTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorTypeContext)
}

func (s *OperatorValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOperatorVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OperatorVal() (localctx IOperatorValContext) {
	this := p
	_ = this

	localctx = NewOperatorValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, arishemParserRULE_operatorVal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(552)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(553)
		p.OperatorType()
	}
	{
		p.SetState(554)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOperatorPairContext is an interface to support dynamic dispatch.
type IOperatorPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OperatorKey() IOperatorKeyContext
	COLON() antlr.TerminalNode
	OperatorVal() IOperatorValContext

	// IsOperatorPairContext differentiates from other interfaces.
	IsOperatorPairContext()
}

type OperatorPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorPairContext() *OperatorPairContext {
	var p = new(OperatorPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_operatorPair
	return p
}

func (*OperatorPairContext) IsOperatorPairContext() {}

func NewOperatorPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorPairContext {
	var p = new(OperatorPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_operatorPair

	return p
}

func (s *OperatorPairContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorPairContext) OperatorKey() IOperatorKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorKeyContext)
}

func (s *OperatorPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *OperatorPairContext) OperatorVal() IOperatorValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorValContext)
}

func (s *OperatorPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOperatorPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OperatorPair() (localctx IOperatorPairContext) {
	this := p
	_ = this

	localctx = NewOperatorPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, arishemParserRULE_operatorPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(556)
		p.OperatorKey()
	}
	{
		p.SetState(557)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(558)
		p.OperatorVal()
	}

	return localctx
}

// IOperatorTypeContext is an interface to support dynamic dispatch.
type IOperatorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPERATOR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsOperatorTypeContext differentiates from other interfaces.
	IsOperatorTypeContext()
}

type OperatorTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorTypeContext() *OperatorTypeContext {
	var p = new(OperatorTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_operatorType
	return p
}

func (*OperatorTypeContext) IsOperatorTypeContext() {}

func NewOperatorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorTypeContext {
	var p = new(OperatorTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_operatorType

	return p
}

func (s *OperatorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorTypeContext) OPERATOR() antlr.TerminalNode {
	return s.GetToken(arishemParserOPERATOR, 0)
}

func (s *OperatorTypeContext) NOT() antlr.TerminalNode {
	return s.GetToken(arishemParserNOT, 0)
}

func (s *OperatorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOperatorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) OperatorType() (localctx IOperatorTypeContext) {
	this := p
	_ = this

	localctx = NewOperatorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, arishemParserRULE_operatorType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserT__19 || _la == arishemParserNOT {
		{
			p.SetState(560)
			_la = p.GetTokenStream().LA(1)

			if !(_la == arishemParserT__19 || _la == arishemParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(563)
		p.Match(arishemParserOPERATOR)
	}

	return localctx
}

// ILhsKeyContext is an interface to support dynamic dispatch.
type ILhsKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsLhsKeyContext differentiates from other interfaces.
	IsLhsKeyContext()
}

type LhsKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsKeyContext() *LhsKeyContext {
	var p = new(LhsKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_lhsKey
	return p
}

func (*LhsKeyContext) IsLhsKeyContext() {}

func NewLhsKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsKeyContext {
	var p = new(LhsKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_lhsKey

	return p
}

func (s *LhsKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *LhsKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *LhsKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLhsKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LhsKey() (localctx ILhsKeyContext) {
	this := p
	_ = this

	localctx = NewLhsKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, arishemParserRULE_lhsKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(565)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(566)
		p.Match(arishemParserT__20)
	}
	{
		p.SetState(567)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ILhsPairContext is an interface to support dynamic dispatch.
type ILhsPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LhsKey() ILhsKeyContext
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsLhsPairContext differentiates from other interfaces.
	IsLhsPairContext()
}

type LhsPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsPairContext() *LhsPairContext {
	var p = new(LhsPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_lhsPair
	return p
}

func (*LhsPairContext) IsLhsPairContext() {}

func NewLhsPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsPairContext {
	var p = new(LhsPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_lhsPair

	return p
}

func (s *LhsPairContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsPairContext) LhsKey() ILhsKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsKeyContext)
}

func (s *LhsPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *LhsPairContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LhsPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LhsPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLhsPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LhsPair() (localctx ILhsPairContext) {
	this := p
	_ = this

	localctx = NewLhsPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, arishemParserRULE_lhsPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.LhsKey()
	}
	{
		p.SetState(570)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(571)
		p.Expr()
	}

	return localctx
}

// IRhsKeyContext is an interface to support dynamic dispatch.
type IRhsKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsRhsKeyContext differentiates from other interfaces.
	IsRhsKeyContext()
}

type RhsKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsKeyContext() *RhsKeyContext {
	var p = new(RhsKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_rhsKey
	return p
}

func (*RhsKeyContext) IsRhsKeyContext() {}

func NewRhsKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsKeyContext {
	var p = new(RhsKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_rhsKey

	return p
}

func (s *RhsKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *RhsKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *RhsKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitRhsKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) RhsKey() (localctx IRhsKeyContext) {
	this := p
	_ = this

	localctx = NewRhsKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, arishemParserRULE_rhsKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(574)
		p.Match(arishemParserT__21)
	}
	{
		p.SetState(575)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IRhsPairContext is an interface to support dynamic dispatch.
type IRhsPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RhsKey() IRhsKeyContext
	COLON() antlr.TerminalNode
	Expr() IExprContext

	// IsRhsPairContext differentiates from other interfaces.
	IsRhsPairContext()
}

type RhsPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsPairContext() *RhsPairContext {
	var p = new(RhsPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_rhsPair
	return p
}

func (*RhsPairContext) IsRhsPairContext() {}

func NewRhsPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsPairContext {
	var p = new(RhsPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_rhsPair

	return p
}

func (s *RhsPairContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsPairContext) RhsKey() IRhsKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRhsKeyContext)
}

func (s *RhsPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *RhsPairContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RhsPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitRhsPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) RhsPair() (localctx IRhsPairContext) {
	this := p
	_ = this

	localctx = NewRhsPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, arishemParserRULE_rhsPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(577)
		p.RhsKey()
	}
	{
		p.SetState(578)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(579)
		p.Expr()
	}

	return localctx
}

// IBoolTypeContext is an interface to support dynamic dispatch.
type IBoolTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBoolTypeContext differentiates from other interfaces.
	IsBoolTypeContext()
}

type BoolTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolTypeContext() *BoolTypeContext {
	var p = new(BoolTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_boolType
	return p
}

func (*BoolTypeContext) IsBoolTypeContext() {}

func NewBoolTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolTypeContext {
	var p = new(BoolTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_boolType

	return p
}

func (s *BoolTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) BoolType() (localctx IBoolTypeContext) {
	this := p
	_ = this

	localctx = NewBoolTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, arishemParserRULE_boolType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		_la = p.GetTokenStream().LA(1)

		if !(_la == arishemParserT__22 || _la == arishemParserT__23) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumberTypeContext is an interface to support dynamic dispatch.
type INumberTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllINT() []antlr.TerminalNode
	INT(i int) antlr.TerminalNode
	SUB() antlr.TerminalNode
	DOT() antlr.TerminalNode
	EXP() antlr.TerminalNode

	// IsNumberTypeContext differentiates from other interfaces.
	IsNumberTypeContext()
}

type NumberTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberTypeContext() *NumberTypeContext {
	var p = new(NumberTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_numberType
	return p
}

func (*NumberTypeContext) IsNumberTypeContext() {}

func NewNumberTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberTypeContext {
	var p = new(NumberTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_numberType

	return p
}

func (s *NumberTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberTypeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(arishemParserINT)
}

func (s *NumberTypeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserINT, i)
}

func (s *NumberTypeContext) SUB() antlr.TerminalNode {
	return s.GetToken(arishemParserSUB, 0)
}

func (s *NumberTypeContext) DOT() antlr.TerminalNode {
	return s.GetToken(arishemParserDOT, 0)
}

func (s *NumberTypeContext) EXP() antlr.TerminalNode {
	return s.GetToken(arishemParserEXP, 0)
}

func (s *NumberTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNumberType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) NumberType() (localctx INumberTypeContext) {
	this := p
	_ = this

	localctx = NewNumberTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, arishemParserRULE_numberType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserSUB {
		{
			p.SetState(583)
			p.Match(arishemParserSUB)
		}

	}
	{
		p.SetState(586)
		p.Match(arishemParserINT)
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserDOT {
		{
			p.SetState(587)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(588)
			p.Match(arishemParserINT)
		}

	}
	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserEXP {
		{
			p.SetState(591)
			p.Match(arishemParserEXP)
		}

	}

	return localctx
}

// IStringTypeContext is an interface to support dynamic dispatch.
type IStringTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	AllSpecialChars() []ISpecialCharsContext
	SpecialChars(i int) ISpecialCharsContext
	AllSP_UNICODE() []antlr.TerminalNode
	SP_UNICODE(i int) antlr.TerminalNode

	// IsStringTypeContext differentiates from other interfaces.
	IsStringTypeContext()
}

type StringTypeContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringTypeContext() *StringTypeContext {
	var p = new(StringTypeContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_stringType
	return p
}

func (*StringTypeContext) IsStringTypeContext() {}

func NewStringTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringTypeContext {
	var p = new(StringTypeContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_stringType

	return p
}

func (s *StringTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StringTypeContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *StringTypeContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *StringTypeContext) AllSpecialChars() []ISpecialCharsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecialCharsContext); ok {
			len++
		}
	}

	tst := make([]ISpecialCharsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecialCharsContext); ok {
			tst[i] = t.(ISpecialCharsContext)
			i++
		}
	}

	return tst
}

func (s *StringTypeContext) SpecialChars(i int) ISpecialCharsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecialCharsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecialCharsContext)
}

func (s *StringTypeContext) AllSP_UNICODE() []antlr.TerminalNode {
	return s.GetTokens(arishemParserSP_UNICODE)
}

func (s *StringTypeContext) SP_UNICODE(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserSP_UNICODE, i)
}

func (s *StringTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitStringType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) StringType() (localctx IStringTypeContext) {
	this := p
	_ = this

	localctx = NewStringTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, arishemParserRULE_stringType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(arishemParserQUOTATION)
	}
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(601)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				p.SetState(595)
				p.MatchWildcard()

			case 2:
				{
					p.SetState(596)
					p.Match(arishemParserT__24)
				}
				p.SetState(597)
				p.MatchWildcard()

			case 3:
				{
					p.SetState(598)
					_la = p.GetTokenStream().LA(1)

					if _la <= 0 || _la == arishemParserT__24 || _la == arishemParserQUOTATION {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 4:
				{
					p.SetState(599)
					p.SpecialChars()
				}

			case 5:
				{
					p.SetState(600)
					p.Match(arishemParserSP_UNICODE)
				}

			}

		}
		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}
	{
		p.SetState(606)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ISpecialCharsContext is an interface to support dynamic dispatch.
type ISpecialCharsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	QUOTATION() antlr.TerminalNode
	MOD() antlr.TerminalNode
	MUL() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SUB() antlr.TerminalNode
	DOT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	COLON() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	L_BRACKET() antlr.TerminalNode
	R_BRACKET() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	R_BRACE() antlr.TerminalNode

	// IsSpecialCharsContext differentiates from other interfaces.
	IsSpecialCharsContext()
}

type SpecialCharsContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecialCharsContext() *SpecialCharsContext {
	var p = new(SpecialCharsContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_specialChars
	return p
}

func (*SpecialCharsContext) IsSpecialCharsContext() {}

func NewSpecialCharsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecialCharsContext {
	var p = new(SpecialCharsContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_specialChars

	return p
}

func (s *SpecialCharsContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecialCharsContext) NOT() antlr.TerminalNode {
	return s.GetToken(arishemParserNOT, 0)
}

func (s *SpecialCharsContext) QUOTATION() antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, 0)
}

func (s *SpecialCharsContext) MOD() antlr.TerminalNode {
	return s.GetToken(arishemParserMOD, 0)
}

func (s *SpecialCharsContext) MUL() antlr.TerminalNode {
	return s.GetToken(arishemParserMUL, 0)
}

func (s *SpecialCharsContext) PLUS() antlr.TerminalNode {
	return s.GetToken(arishemParserPLUS, 0)
}

func (s *SpecialCharsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(arishemParserCOMMA, 0)
}

func (s *SpecialCharsContext) SUB() antlr.TerminalNode {
	return s.GetToken(arishemParserSUB, 0)
}

func (s *SpecialCharsContext) DOT() antlr.TerminalNode {
	return s.GetToken(arishemParserDOT, 0)
}

func (s *SpecialCharsContext) DIV() antlr.TerminalNode {
	return s.GetToken(arishemParserDIV, 0)
}

func (s *SpecialCharsContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *SpecialCharsContext) LT() antlr.TerminalNode {
	return s.GetToken(arishemParserLT, 0)
}

func (s *SpecialCharsContext) GT() antlr.TerminalNode {
	return s.GetToken(arishemParserGT, 0)
}

func (s *SpecialCharsContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACKET, 0)
}

func (s *SpecialCharsContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACKET, 0)
}

func (s *SpecialCharsContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *SpecialCharsContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *SpecialCharsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecialCharsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecialCharsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitSpecialChars(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) SpecialChars() (localctx ISpecialCharsContext) {
	this := p
	_ = this

	localctx = NewSpecialCharsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, arishemParserRULE_specialChars)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-25)) & ^0x3f) == 0 && ((int64(1)<<(_la-25))&114213909561343) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStrConstKeyContext is an interface to support dynamic dispatch.
type IStrConstKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsStrConstKeyContext differentiates from other interfaces.
	IsStrConstKeyContext()
}

type StrConstKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrConstKeyContext() *StrConstKeyContext {
	var p = new(StrConstKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_strConstKey
	return p
}

func (*StrConstKeyContext) IsStrConstKeyContext() {}

func NewStrConstKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrConstKeyContext {
	var p = new(StrConstKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_strConstKey

	return p
}

func (s *StrConstKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *StrConstKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *StrConstKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *StrConstKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrConstKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrConstKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitStrConstKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) StrConstKey() (localctx IStrConstKeyContext) {
	this := p
	_ = this

	localctx = NewStrConstKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, arishemParserRULE_strConstKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(610)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(611)
		p.Match(arishemParserT__41)
	}
	{
		p.SetState(612)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// INumConstKeyContext is an interface to support dynamic dispatch.
type INumConstKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsNumConstKeyContext differentiates from other interfaces.
	IsNumConstKeyContext()
}

type NumConstKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumConstKeyContext() *NumConstKeyContext {
	var p = new(NumConstKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_numConstKey
	return p
}

func (*NumConstKeyContext) IsNumConstKeyContext() {}

func NewNumConstKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumConstKeyContext {
	var p = new(NumConstKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_numConstKey

	return p
}

func (s *NumConstKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *NumConstKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *NumConstKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *NumConstKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumConstKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumConstKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNumConstKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) NumConstKey() (localctx INumConstKeyContext) {
	this := p
	_ = this

	localctx = NewNumConstKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, arishemParserRULE_numConstKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(614)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(615)
		p.Match(arishemParserT__42)
	}
	{
		p.SetState(616)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IBoolConstKeyContext is an interface to support dynamic dispatch.
type IBoolConstKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsBoolConstKeyContext differentiates from other interfaces.
	IsBoolConstKeyContext()
}

type BoolConstKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolConstKeyContext() *BoolConstKeyContext {
	var p = new(BoolConstKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_boolConstKey
	return p
}

func (*BoolConstKeyContext) IsBoolConstKeyContext() {}

func NewBoolConstKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolConstKeyContext {
	var p = new(BoolConstKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_boolConstKey

	return p
}

func (s *BoolConstKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolConstKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *BoolConstKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *BoolConstKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolConstKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolConstKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitBoolConstKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) BoolConstKey() (localctx IBoolConstKeyContext) {
	this := p
	_ = this

	localctx = NewBoolConstKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, arishemParserRULE_boolConstKey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(619)
		p.Match(arishemParserT__43)
	}
	{
		p.SetState(620)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}
