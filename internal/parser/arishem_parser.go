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
		"'Const'", "'ConstList'", "'SubCondExpr'", "'CondName'", "'ParamMap'",
		"'ParamList'", "'BuiltinParam'", "'ActionName'", "'OpLogic'", "'OpMath'",
		"'FuncName'", "'Operator'", "'NOT '", "'Lhs'", "'Rhs'", "'true'", "'false'",
		"'\\'", "' '", "'#'", "'$'", "'&'", "'''", "'('", "')'", "'|'", "';'",
		"'='", "'?'", "'@'", "'^'", "'_'", "'`'", "'~'", "'StrConst'", "'NumConst'",
		"'BoolConst'", "'IS_NULL'", "'STRING_START_WITH'", "'STRING_END_WITH'",
		"'STRING_CONTAINS'", "'LIST_IN'", "'LIST_CONTAINS'", "'LIST_RETAIN'",
		"'CONTAIN_REGULAR'", "'SUB_LIST_IN'", "'SUB_LIST_CONTAINS'", "'CONTAINS_KEYWORDS'",
		"'BETWEEN_ALL_CLOSE'", "'BETWEEN_ALL_OPEN'", "'BETWEEN_LEFT_OPEN_RIGHT_CLOSE'",
		"'BETWEEN_LEFT_CLOSE_RIGHT_OPEN'", "'SUB_COND'", "'FOREACH'", "", "'null'",
		"'.'", "','", "':'", "'\"'", "'['", "']'", "'{'", "'}'", "", "", "",
		"", "'&&'", "'||'", "'!'", "'+'", "'-'", "'/'", "'*'", "'%'", "'=='",
		"'!='", "'>'", "'<'", "'>='", "'<='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "OPERATOR_LOGIC_STR",
		"NULL", "DOT", "COMMA", "COLON", "QUOTATION", "L_BRACKET", "R_BRACKET",
		"L_BRACE", "R_BRACE", "INT", "EXP", "NAME_KEY_TYPE", "SP_UNICODE", "AND",
		"OR", "NOT", "PLUS", "SUB", "DIV", "MUL", "MOD", "EQUALS", "NOT_EQUALS",
		"GT", "LT", "GTE", "LTE", "WS",
	}
	staticData.ruleNames = []string{
		"condEntity", "aimEntity", "conditionGroup", "condition", "aim", "actionAimType",
		"expr", "exprType", "variable", "math", "function", "exprs", "feature",
		"subCond", "constant", "constantList", "conditionsKey", "conditionsVal",
		"conditionsPair", "conditionGroupsKey", "conditionGroupsVal", "conditionGroupsPair",
		"listExprKey", "mapExprKey", "mathExprKey", "funcExprKey", "varExprKey",
		"featureExprKey", "featurePathKey", "constKey", "constListKey", "subCondExprKey",
		"condNameKey", "varPath", "varPathVal", "featurePath", "featurePathVal",
		"paramMapKey", "paramListKey", "builtinParamKey", "actionAimKey", "actionAimVal",
		"actionAimPair", "nameKey", "keyValue", "keyValues", "opLogicKey", "opLogicVal",
		"opLogicPair", "opLogic", "opMathKey", "lrOpMathVal", "listOpMathVal",
		"lrOpMathPair", "listOpMathPair", "listMathOperator", "lrMathOperator",
		"logicMathOperator", "funcNameKey", "funcNamePair", "operatorKey", "operatorVal",
		"operatorPair", "operatorType", "lhsKey", "lhsPair", "rhsKey", "rhsPair",
		"boolType", "numberType", "stringType", "specialChars", "strConstKey",
		"numConstKey", "boolConstKey", "operator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 92, 704, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 170,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 181,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 187, 8, 4, 10, 4, 12, 4, 190, 9, 4,
		3, 4, 192, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 198, 8, 4, 1, 4, 1, 4, 3,
		4, 202, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 211, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 222, 8, 5,
		1, 5, 1, 5, 3, 5, 226, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 233, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 3, 7, 272, 8, 7, 1, 8, 1, 8, 3, 8, 276, 8, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 295, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 304, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 315, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 320, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 326, 8, 11, 10, 11, 12, 11, 329,
		9, 11, 3, 11, 331, 8, 11, 1, 11, 1, 11, 3, 11, 335, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 346, 8, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 351, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 360, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 381, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5,
		15, 387, 8, 15, 10, 15, 12, 15, 390, 9, 15, 3, 15, 392, 8, 15, 1, 15, 1,
		15, 3, 15, 396, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 5, 17, 406, 8, 17, 10, 17, 12, 17, 409, 9, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 425, 8, 20, 10, 20, 12, 20, 428, 9, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 34, 5, 34, 487, 8, 34, 10, 34, 12, 34, 490, 9, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 4, 36, 499, 8, 36, 11, 36, 12, 36,
		500, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 539, 8, 45, 10, 45, 12, 45,
		542, 9, 45, 3, 45, 544, 8, 45, 1, 45, 1, 45, 3, 45, 548, 8, 45, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 3, 56, 588, 8, 56, 1, 57, 1, 57,
		1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63,
		3, 63, 613, 8, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1,
		65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 67,
		1, 68, 1, 68, 1, 69, 3, 69, 636, 8, 69, 1, 69, 1, 69, 1, 69, 3, 69, 641,
		8, 69, 1, 69, 3, 69, 644, 8, 69, 1, 70, 1, 70, 1, 70, 1, 70, 1, 70, 1,
		70, 1, 70, 5, 70, 653, 8, 70, 10, 70, 12, 70, 656, 9, 70, 1, 70, 1, 70,
		1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 73, 1,
		74, 1, 74, 1, 74, 1, 74, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75,
		1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1,
		75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75,
		3, 75, 702, 8, 75, 1, 75, 1, 654, 0, 76, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120,
		122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150,
		0, 7, 2, 0, 64, 64, 78, 80, 1, 0, 81, 85, 1, 0, 86, 91, 2, 0, 22, 22, 80,
		80, 1, 0, 25, 26, 2, 0, 27, 27, 69, 69, 4, 0, 27, 43, 66, 73, 80, 85, 88,
		89, 707, 0, 152, 1, 0, 0, 0, 2, 154, 1, 0, 0, 0, 4, 169, 1, 0, 0, 0, 6,
		180, 1, 0, 0, 0, 8, 201, 1, 0, 0, 0, 10, 225, 1, 0, 0, 0, 12, 232, 1, 0,
		0, 0, 14, 271, 1, 0, 0, 0, 16, 275, 1, 0, 0, 0, 18, 294, 1, 0, 0, 0, 20,
		319, 1, 0, 0, 0, 22, 334, 1, 0, 0, 0, 24, 350, 1, 0, 0, 0, 26, 359, 1,
		0, 0, 0, 28, 380, 1, 0, 0, 0, 30, 395, 1, 0, 0, 0, 32, 397, 1, 0, 0, 0,
		34, 401, 1, 0, 0, 0, 36, 412, 1, 0, 0, 0, 38, 416, 1, 0, 0, 0, 40, 420,
		1, 0, 0, 0, 42, 431, 1, 0, 0, 0, 44, 435, 1, 0, 0, 0, 46, 439, 1, 0, 0,
		0, 48, 443, 1, 0, 0, 0, 50, 447, 1, 0, 0, 0, 52, 451, 1, 0, 0, 0, 54, 455,
		1, 0, 0, 0, 56, 459, 1, 0, 0, 0, 58, 463, 1, 0, 0, 0, 60, 467, 1, 0, 0,
		0, 62, 471, 1, 0, 0, 0, 64, 475, 1, 0, 0, 0, 66, 479, 1, 0, 0, 0, 68, 483,
		1, 0, 0, 0, 70, 491, 1, 0, 0, 0, 72, 495, 1, 0, 0, 0, 74, 502, 1, 0, 0,
		0, 76, 506, 1, 0, 0, 0, 78, 510, 1, 0, 0, 0, 80, 514, 1, 0, 0, 0, 82, 518,
		1, 0, 0, 0, 84, 522, 1, 0, 0, 0, 86, 526, 1, 0, 0, 0, 88, 530, 1, 0, 0,
		0, 90, 547, 1, 0, 0, 0, 92, 549, 1, 0, 0, 0, 94, 553, 1, 0, 0, 0, 96, 557,
		1, 0, 0, 0, 98, 561, 1, 0, 0, 0, 100, 563, 1, 0, 0, 0, 102, 567, 1, 0,
		0, 0, 104, 571, 1, 0, 0, 0, 106, 575, 1, 0, 0, 0, 108, 579, 1, 0, 0, 0,
		110, 583, 1, 0, 0, 0, 112, 587, 1, 0, 0, 0, 114, 589, 1, 0, 0, 0, 116,
		591, 1, 0, 0, 0, 118, 595, 1, 0, 0, 0, 120, 599, 1, 0, 0, 0, 122, 603,
		1, 0, 0, 0, 124, 607, 1, 0, 0, 0, 126, 612, 1, 0, 0, 0, 128, 616, 1, 0,
		0, 0, 130, 620, 1, 0, 0, 0, 132, 624, 1, 0, 0, 0, 134, 628, 1, 0, 0, 0,
		136, 632, 1, 0, 0, 0, 138, 635, 1, 0, 0, 0, 140, 645, 1, 0, 0, 0, 142,
		659, 1, 0, 0, 0, 144, 661, 1, 0, 0, 0, 146, 665, 1, 0, 0, 0, 148, 669,
		1, 0, 0, 0, 150, 701, 1, 0, 0, 0, 152, 153, 3, 4, 2, 0, 153, 1, 1, 0, 0,
		0, 154, 155, 3, 8, 4, 0, 155, 3, 1, 0, 0, 0, 156, 157, 5, 72, 0, 0, 157,
		158, 3, 96, 48, 0, 158, 159, 5, 67, 0, 0, 159, 160, 3, 42, 21, 0, 160,
		161, 5, 73, 0, 0, 161, 170, 1, 0, 0, 0, 162, 163, 5, 72, 0, 0, 163, 164,
		3, 96, 48, 0, 164, 165, 5, 67, 0, 0, 165, 166, 3, 36, 18, 0, 166, 167,
		5, 73, 0, 0, 167, 170, 1, 0, 0, 0, 168, 170, 5, 65, 0, 0, 169, 156, 1,
		0, 0, 0, 169, 162, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 5, 1, 0, 0, 0,
		171, 172, 5, 72, 0, 0, 172, 173, 3, 124, 62, 0, 173, 174, 5, 67, 0, 0,
		174, 175, 3, 130, 65, 0, 175, 176, 5, 67, 0, 0, 176, 177, 3, 134, 67, 0,
		177, 178, 5, 73, 0, 0, 178, 181, 1, 0, 0, 0, 179, 181, 5, 65, 0, 0, 180,
		171, 1, 0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 7, 1, 0, 0, 0, 182, 191, 5,
		70, 0, 0, 183, 188, 3, 10, 5, 0, 184, 185, 5, 67, 0, 0, 185, 187, 3, 10,
		5, 0, 186, 184, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0,
		188, 189, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191,
		183, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 202,
		5, 71, 0, 0, 194, 202, 3, 10, 5, 0, 195, 197, 5, 72, 0, 0, 196, 198, 3,
		14, 7, 0, 197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0,
		0, 199, 202, 5, 73, 0, 0, 200, 202, 5, 65, 0, 0, 201, 182, 1, 0, 0, 0,
		201, 194, 1, 0, 0, 0, 201, 195, 1, 0, 0, 0, 201, 200, 1, 0, 0, 0, 202,
		9, 1, 0, 0, 0, 203, 204, 5, 72, 0, 0, 204, 210, 3, 84, 42, 0, 205, 206,
		5, 67, 0, 0, 206, 207, 3, 74, 37, 0, 207, 208, 5, 68, 0, 0, 208, 209, 3,
		90, 45, 0, 209, 211, 1, 0, 0, 0, 210, 205, 1, 0, 0, 0, 210, 211, 1, 0,
		0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 5, 73, 0, 0, 213, 226, 1, 0, 0, 0,
		214, 215, 5, 72, 0, 0, 215, 221, 3, 84, 42, 0, 216, 217, 5, 67, 0, 0, 217,
		218, 3, 76, 38, 0, 218, 219, 5, 68, 0, 0, 219, 220, 3, 22, 11, 0, 220,
		222, 1, 0, 0, 0, 221, 216, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 224, 5, 73, 0, 0, 224, 226, 1, 0, 0, 0, 225, 203, 1, 0,
		0, 0, 225, 214, 1, 0, 0, 0, 226, 11, 1, 0, 0, 0, 227, 228, 5, 72, 0, 0,
		228, 229, 3, 14, 7, 0, 229, 230, 5, 73, 0, 0, 230, 233, 1, 0, 0, 0, 231,
		233, 5, 65, 0, 0, 232, 227, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 13,
		1, 0, 0, 0, 234, 235, 3, 44, 22, 0, 235, 236, 5, 68, 0, 0, 236, 237, 3,
		22, 11, 0, 237, 272, 1, 0, 0, 0, 238, 239, 3, 46, 23, 0, 239, 240, 5, 68,
		0, 0, 240, 241, 3, 90, 45, 0, 241, 272, 1, 0, 0, 0, 242, 243, 3, 58, 29,
		0, 243, 244, 5, 68, 0, 0, 244, 245, 3, 28, 14, 0, 245, 272, 1, 0, 0, 0,
		246, 247, 3, 60, 30, 0, 247, 248, 5, 68, 0, 0, 248, 249, 3, 30, 15, 0,
		249, 272, 1, 0, 0, 0, 250, 251, 3, 48, 24, 0, 251, 252, 5, 68, 0, 0, 252,
		253, 3, 18, 9, 0, 253, 272, 1, 0, 0, 0, 254, 255, 3, 50, 25, 0, 255, 256,
		5, 68, 0, 0, 256, 257, 3, 20, 10, 0, 257, 272, 1, 0, 0, 0, 258, 259, 3,
		52, 26, 0, 259, 260, 5, 68, 0, 0, 260, 261, 3, 16, 8, 0, 261, 272, 1, 0,
		0, 0, 262, 263, 3, 54, 27, 0, 263, 264, 5, 68, 0, 0, 264, 265, 3, 24, 12,
		0, 265, 272, 1, 0, 0, 0, 266, 267, 3, 62, 31, 0, 267, 268, 5, 68, 0, 0,
		268, 269, 3, 26, 13, 0, 269, 272, 1, 0, 0, 0, 270, 272, 5, 65, 0, 0, 271,
		234, 1, 0, 0, 0, 271, 238, 1, 0, 0, 0, 271, 242, 1, 0, 0, 0, 271, 246,
		1, 0, 0, 0, 271, 250, 1, 0, 0, 0, 271, 254, 1, 0, 0, 0, 271, 258, 1, 0,
		0, 0, 271, 262, 1, 0, 0, 0, 271, 266, 1, 0, 0, 0, 271, 270, 1, 0, 0, 0,
		272, 15, 1, 0, 0, 0, 273, 276, 3, 66, 33, 0, 274, 276, 5, 65, 0, 0, 275,
		273, 1, 0, 0, 0, 275, 274, 1, 0, 0, 0, 276, 17, 1, 0, 0, 0, 277, 278, 5,
		72, 0, 0, 278, 279, 3, 106, 53, 0, 279, 280, 5, 67, 0, 0, 280, 281, 3,
		130, 65, 0, 281, 282, 5, 67, 0, 0, 282, 283, 3, 134, 67, 0, 283, 284, 5,
		73, 0, 0, 284, 295, 1, 0, 0, 0, 285, 286, 5, 72, 0, 0, 286, 287, 3, 108,
		54, 0, 287, 288, 5, 67, 0, 0, 288, 289, 3, 76, 38, 0, 289, 290, 5, 68,
		0, 0, 290, 291, 3, 22, 11, 0, 291, 292, 5, 73, 0, 0, 292, 295, 1, 0, 0,
		0, 293, 295, 5, 65, 0, 0, 294, 277, 1, 0, 0, 0, 294, 285, 1, 0, 0, 0, 294,
		293, 1, 0, 0, 0, 295, 19, 1, 0, 0, 0, 296, 297, 5, 72, 0, 0, 297, 303,
		3, 118, 59, 0, 298, 299, 5, 67, 0, 0, 299, 300, 3, 74, 37, 0, 300, 301,
		5, 68, 0, 0, 301, 302, 3, 90, 45, 0, 302, 304, 1, 0, 0, 0, 303, 298, 1,
		0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 5, 73, 0,
		0, 306, 320, 1, 0, 0, 0, 307, 308, 5, 72, 0, 0, 308, 314, 3, 118, 59, 0,
		309, 310, 5, 67, 0, 0, 310, 311, 3, 76, 38, 0, 311, 312, 5, 68, 0, 0, 312,
		313, 3, 22, 11, 0, 313, 315, 1, 0, 0, 0, 314, 309, 1, 0, 0, 0, 314, 315,
		1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317, 5, 73, 0, 0, 317, 320, 1, 0,
		0, 0, 318, 320, 5, 65, 0, 0, 319, 296, 1, 0, 0, 0, 319, 307, 1, 0, 0, 0,
		319, 318, 1, 0, 0, 0, 320, 21, 1, 0, 0, 0, 321, 330, 5, 70, 0, 0, 322,
		327, 3, 12, 6, 0, 323, 324, 5, 67, 0, 0, 324, 326, 3, 12, 6, 0, 325, 323,
		1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0,
		0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 322, 1, 0, 0, 0,
		330, 331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 335, 5, 71, 0, 0, 333,
		335, 5, 65, 0, 0, 334, 321, 1, 0, 0, 0, 334, 333, 1, 0, 0, 0, 335, 23,
		1, 0, 0, 0, 336, 337, 5, 72, 0, 0, 337, 338, 3, 56, 28, 0, 338, 339, 5,
		68, 0, 0, 339, 345, 3, 70, 35, 0, 340, 341, 5, 67, 0, 0, 341, 342, 3, 78,
		39, 0, 342, 343, 5, 68, 0, 0, 343, 344, 3, 90, 45, 0, 344, 346, 1, 0, 0,
		0, 345, 340, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347,
		348, 5, 73, 0, 0, 348, 351, 1, 0, 0, 0, 349, 351, 5, 65, 0, 0, 350, 336,
		1, 0, 0, 0, 350, 349, 1, 0, 0, 0, 351, 25, 1, 0, 0, 0, 352, 353, 5, 72,
		0, 0, 353, 354, 3, 64, 32, 0, 354, 355, 5, 68, 0, 0, 355, 356, 3, 86, 43,
		0, 356, 357, 5, 73, 0, 0, 357, 360, 1, 0, 0, 0, 358, 360, 5, 65, 0, 0,
		359, 352, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360, 27, 1, 0, 0, 0, 361, 362,
		5, 72, 0, 0, 362, 363, 3, 146, 73, 0, 363, 364, 5, 68, 0, 0, 364, 365,
		3, 138, 69, 0, 365, 366, 5, 73, 0, 0, 366, 381, 1, 0, 0, 0, 367, 368, 5,
		72, 0, 0, 368, 369, 3, 144, 72, 0, 369, 370, 5, 68, 0, 0, 370, 371, 3,
		140, 70, 0, 371, 372, 5, 73, 0, 0, 372, 381, 1, 0, 0, 0, 373, 374, 5, 72,
		0, 0, 374, 375, 3, 148, 74, 0, 375, 376, 5, 68, 0, 0, 376, 377, 3, 136,
		68, 0, 377, 378, 5, 73, 0, 0, 378, 381, 1, 0, 0, 0, 379, 381, 5, 65, 0,
		0, 380, 361, 1, 0, 0, 0, 380, 367, 1, 0, 0, 0, 380, 373, 1, 0, 0, 0, 380,
		379, 1, 0, 0, 0, 381, 29, 1, 0, 0, 0, 382, 391, 5, 70, 0, 0, 383, 388,
		3, 28, 14, 0, 384, 385, 5, 67, 0, 0, 385, 387, 3, 28, 14, 0, 386, 384,
		1, 0, 0, 0, 387, 390, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0,
		0, 0, 389, 392, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 391, 383, 1, 0, 0, 0,
		391, 392, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 396, 5, 71, 0, 0, 394,
		396, 5, 65, 0, 0, 395, 382, 1, 0, 0, 0, 395, 394, 1, 0, 0, 0, 396, 31,
		1, 0, 0, 0, 397, 398, 5, 69, 0, 0, 398, 399, 5, 1, 0, 0, 399, 400, 5, 69,
		0, 0, 400, 33, 1, 0, 0, 0, 401, 402, 5, 70, 0, 0, 402, 407, 3, 6, 3, 0,
		403, 404, 5, 67, 0, 0, 404, 406, 3, 6, 3, 0, 405, 403, 1, 0, 0, 0, 406,
		409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 410,
		1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 411, 5, 71, 0, 0, 411, 35, 1, 0,
		0, 0, 412, 413, 3, 32, 16, 0, 413, 414, 5, 68, 0, 0, 414, 415, 3, 34, 17,
		0, 415, 37, 1, 0, 0, 0, 416, 417, 5, 69, 0, 0, 417, 418, 5, 2, 0, 0, 418,
		419, 5, 69, 0, 0, 419, 39, 1, 0, 0, 0, 420, 421, 5, 70, 0, 0, 421, 426,
		3, 4, 2, 0, 422, 423, 5, 67, 0, 0, 423, 425, 3, 4, 2, 0, 424, 422, 1, 0,
		0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0,
		427, 429, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 429, 430, 5, 71, 0, 0, 430,
		41, 1, 0, 0, 0, 431, 432, 3, 38, 19, 0, 432, 433, 5, 68, 0, 0, 433, 434,
		3, 40, 20, 0, 434, 43, 1, 0, 0, 0, 435, 436, 5, 69, 0, 0, 436, 437, 5,
		3, 0, 0, 437, 438, 5, 69, 0, 0, 438, 45, 1, 0, 0, 0, 439, 440, 5, 69, 0,
		0, 440, 441, 5, 4, 0, 0, 441, 442, 5, 69, 0, 0, 442, 47, 1, 0, 0, 0, 443,
		444, 5, 69, 0, 0, 444, 445, 5, 5, 0, 0, 445, 446, 5, 69, 0, 0, 446, 49,
		1, 0, 0, 0, 447, 448, 5, 69, 0, 0, 448, 449, 5, 6, 0, 0, 449, 450, 5, 69,
		0, 0, 450, 51, 1, 0, 0, 0, 451, 452, 5, 69, 0, 0, 452, 453, 5, 7, 0, 0,
		453, 454, 5, 69, 0, 0, 454, 53, 1, 0, 0, 0, 455, 456, 5, 69, 0, 0, 456,
		457, 5, 8, 0, 0, 457, 458, 5, 69, 0, 0, 458, 55, 1, 0, 0, 0, 459, 460,
		5, 69, 0, 0, 460, 461, 5, 9, 0, 0, 461, 462, 5, 69, 0, 0, 462, 57, 1, 0,
		0, 0, 463, 464, 5, 69, 0, 0, 464, 465, 5, 10, 0, 0, 465, 466, 5, 69, 0,
		0, 466, 59, 1, 0, 0, 0, 467, 468, 5, 69, 0, 0, 468, 469, 5, 11, 0, 0, 469,
		470, 5, 69, 0, 0, 470, 61, 1, 0, 0, 0, 471, 472, 5, 69, 0, 0, 472, 473,
		5, 12, 0, 0, 473, 474, 5, 69, 0, 0, 474, 63, 1, 0, 0, 0, 475, 476, 5, 69,
		0, 0, 476, 477, 5, 13, 0, 0, 477, 478, 5, 69, 0, 0, 478, 65, 1, 0, 0, 0,
		479, 480, 5, 69, 0, 0, 480, 481, 3, 68, 34, 0, 481, 482, 5, 69, 0, 0, 482,
		67, 1, 0, 0, 0, 483, 488, 5, 76, 0, 0, 484, 485, 5, 66, 0, 0, 485, 487,
		5, 76, 0, 0, 486, 484, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486, 1, 0,
		0, 0, 488, 489, 1, 0, 0, 0, 489, 69, 1, 0, 0, 0, 490, 488, 1, 0, 0, 0,
		491, 492, 5, 69, 0, 0, 492, 493, 3, 72, 36, 0, 493, 494, 5, 69, 0, 0, 494,
		71, 1, 0, 0, 0, 495, 498, 5, 76, 0, 0, 496, 497, 5, 66, 0, 0, 497, 499,
		5, 76, 0, 0, 498, 496, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 498, 1, 0,
		0, 0, 500, 501, 1, 0, 0, 0, 501, 73, 1, 0, 0, 0, 502, 503, 5, 69, 0, 0,
		503, 504, 5, 14, 0, 0, 504, 505, 5, 69, 0, 0, 505, 75, 1, 0, 0, 0, 506,
		507, 5, 69, 0, 0, 507, 508, 5, 15, 0, 0, 508, 509, 5, 69, 0, 0, 509, 77,
		1, 0, 0, 0, 510, 511, 5, 69, 0, 0, 511, 512, 5, 16, 0, 0, 512, 513, 5,
		69, 0, 0, 513, 79, 1, 0, 0, 0, 514, 515, 5, 69, 0, 0, 515, 516, 5, 17,
		0, 0, 516, 517, 5, 69, 0, 0, 517, 81, 1, 0, 0, 0, 518, 519, 5, 69, 0, 0,
		519, 520, 5, 76, 0, 0, 520, 521, 5, 69, 0, 0, 521, 83, 1, 0, 0, 0, 522,
		523, 3, 80, 40, 0, 523, 524, 5, 68, 0, 0, 524, 525, 3, 82, 41, 0, 525,
		85, 1, 0, 0, 0, 526, 527, 5, 69, 0, 0, 527, 528, 5, 76, 0, 0, 528, 529,
		5, 69, 0, 0, 529, 87, 1, 0, 0, 0, 530, 531, 3, 86, 43, 0, 531, 532, 5,
		68, 0, 0, 532, 533, 3, 12, 6, 0, 533, 89, 1, 0, 0, 0, 534, 543, 5, 72,
		0, 0, 535, 540, 3, 88, 44, 0, 536, 537, 5, 67, 0, 0, 537, 539, 3, 88, 44,
		0, 538, 536, 1, 0, 0, 0, 539, 542, 1, 0, 0, 0, 540, 538, 1, 0, 0, 0, 540,
		541, 1, 0, 0, 0, 541, 544, 1, 0, 0, 0, 542, 540, 1, 0, 0, 0, 543, 535,
		1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 548, 5, 73,
		0, 0, 546, 548, 5, 65, 0, 0, 547, 534, 1, 0, 0, 0, 547, 546, 1, 0, 0, 0,
		548, 91, 1, 0, 0, 0, 549, 550, 5, 69, 0, 0, 550, 551, 5, 18, 0, 0, 551,
		552, 5, 69, 0, 0, 552, 93, 1, 0, 0, 0, 553, 554, 5, 69, 0, 0, 554, 555,
		3, 98, 49, 0, 555, 556, 5, 69, 0, 0, 556, 95, 1, 0, 0, 0, 557, 558, 3,
		92, 46, 0, 558, 559, 5, 68, 0, 0, 559, 560, 3, 94, 47, 0, 560, 97, 1, 0,
		0, 0, 561, 562, 7, 0, 0, 0, 562, 99, 1, 0, 0, 0, 563, 564, 5, 69, 0, 0,
		564, 565, 5, 19, 0, 0, 565, 566, 5, 69, 0, 0, 566, 101, 1, 0, 0, 0, 567,
		568, 5, 69, 0, 0, 568, 569, 3, 112, 56, 0, 569, 570, 5, 69, 0, 0, 570,
		103, 1, 0, 0, 0, 571, 572, 5, 69, 0, 0, 572, 573, 3, 110, 55, 0, 573, 574,
		5, 69, 0, 0, 574, 105, 1, 0, 0, 0, 575, 576, 3, 100, 50, 0, 576, 577, 5,
		68, 0, 0, 577, 578, 3, 102, 51, 0, 578, 107, 1, 0, 0, 0, 579, 580, 3, 100,
		50, 0, 580, 581, 5, 68, 0, 0, 581, 582, 3, 104, 52, 0, 582, 109, 1, 0,
		0, 0, 583, 584, 7, 1, 0, 0, 584, 111, 1, 0, 0, 0, 585, 588, 3, 110, 55,
		0, 586, 588, 3, 114, 57, 0, 587, 585, 1, 0, 0, 0, 587, 586, 1, 0, 0, 0,
		588, 113, 1, 0, 0, 0, 589, 590, 7, 2, 0, 0, 590, 115, 1, 0, 0, 0, 591,
		592, 5, 69, 0, 0, 592, 593, 5, 20, 0, 0, 593, 594, 5, 69, 0, 0, 594, 117,
		1, 0, 0, 0, 595, 596, 3, 116, 58, 0, 596, 597, 5, 68, 0, 0, 597, 598, 3,
		86, 43, 0, 598, 119, 1, 0, 0, 0, 599, 600, 5, 69, 0, 0, 600, 601, 5, 21,
		0, 0, 601, 602, 5, 69, 0, 0, 602, 121, 1, 0, 0, 0, 603, 604, 5, 69, 0,
		0, 604, 605, 3, 126, 63, 0, 605, 606, 5, 69, 0, 0, 606, 123, 1, 0, 0, 0,
		607, 608, 3, 120, 60, 0, 608, 609, 5, 68, 0, 0, 609, 610, 3, 122, 61, 0,
		610, 125, 1, 0, 0, 0, 611, 613, 7, 3, 0, 0, 612, 611, 1, 0, 0, 0, 612,
		613, 1, 0, 0, 0, 613, 614, 1, 0, 0, 0, 614, 615, 3, 150, 75, 0, 615, 127,
		1, 0, 0, 0, 616, 617, 5, 69, 0, 0, 617, 618, 5, 23, 0, 0, 618, 619, 5,
		69, 0, 0, 619, 129, 1, 0, 0, 0, 620, 621, 3, 128, 64, 0, 621, 622, 5, 68,
		0, 0, 622, 623, 3, 12, 6, 0, 623, 131, 1, 0, 0, 0, 624, 625, 5, 69, 0,
		0, 625, 626, 5, 24, 0, 0, 626, 627, 5, 69, 0, 0, 627, 133, 1, 0, 0, 0,
		628, 629, 3, 132, 66, 0, 629, 630, 5, 68, 0, 0, 630, 631, 3, 12, 6, 0,
		631, 135, 1, 0, 0, 0, 632, 633, 7, 4, 0, 0, 633, 137, 1, 0, 0, 0, 634,
		636, 5, 82, 0, 0, 635, 634, 1, 0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 637,
		1, 0, 0, 0, 637, 640, 5, 74, 0, 0, 638, 639, 5, 66, 0, 0, 639, 641, 5,
		74, 0, 0, 640, 638, 1, 0, 0, 0, 640, 641, 1, 0, 0, 0, 641, 643, 1, 0, 0,
		0, 642, 644, 5, 75, 0, 0, 643, 642, 1, 0, 0, 0, 643, 644, 1, 0, 0, 0, 644,
		139, 1, 0, 0, 0, 645, 654, 5, 69, 0, 0, 646, 653, 9, 0, 0, 0, 647, 648,
		5, 27, 0, 0, 648, 653, 9, 0, 0, 0, 649, 653, 8, 5, 0, 0, 650, 653, 3, 142,
		71, 0, 651, 653, 5, 77, 0, 0, 652, 646, 1, 0, 0, 0, 652, 647, 1, 0, 0,
		0, 652, 649, 1, 0, 0, 0, 652, 650, 1, 0, 0, 0, 652, 651, 1, 0, 0, 0, 653,
		656, 1, 0, 0, 0, 654, 655, 1, 0, 0, 0, 654, 652, 1, 0, 0, 0, 655, 657,
		1, 0, 0, 0, 656, 654, 1, 0, 0, 0, 657, 658, 5, 69, 0, 0, 658, 141, 1, 0,
		0, 0, 659, 660, 7, 6, 0, 0, 660, 143, 1, 0, 0, 0, 661, 662, 5, 69, 0, 0,
		662, 663, 5, 44, 0, 0, 663, 664, 5, 69, 0, 0, 664, 145, 1, 0, 0, 0, 665,
		666, 5, 69, 0, 0, 666, 667, 5, 45, 0, 0, 667, 668, 5, 69, 0, 0, 668, 147,
		1, 0, 0, 0, 669, 670, 5, 69, 0, 0, 670, 671, 5, 46, 0, 0, 671, 672, 5,
		69, 0, 0, 672, 149, 1, 0, 0, 0, 673, 702, 5, 86, 0, 0, 674, 702, 5, 87,
		0, 0, 675, 702, 5, 88, 0, 0, 676, 702, 5, 90, 0, 0, 677, 702, 5, 89, 0,
		0, 678, 702, 5, 91, 0, 0, 679, 702, 5, 47, 0, 0, 680, 702, 5, 48, 0, 0,
		681, 702, 5, 49, 0, 0, 682, 702, 5, 50, 0, 0, 683, 702, 5, 51, 0, 0, 684,
		702, 5, 52, 0, 0, 685, 702, 5, 53, 0, 0, 686, 702, 5, 54, 0, 0, 687, 702,
		5, 55, 0, 0, 688, 702, 5, 56, 0, 0, 689, 702, 5, 57, 0, 0, 690, 702, 5,
		58, 0, 0, 691, 702, 5, 59, 0, 0, 692, 702, 5, 60, 0, 0, 693, 702, 5, 61,
		0, 0, 694, 702, 5, 62, 0, 0, 695, 696, 5, 63, 0, 0, 696, 697, 5, 28, 0,
		0, 697, 698, 3, 126, 63, 0, 698, 699, 5, 28, 0, 0, 699, 700, 3, 98, 49,
		0, 700, 702, 1, 0, 0, 0, 701, 673, 1, 0, 0, 0, 701, 674, 1, 0, 0, 0, 701,
		675, 1, 0, 0, 0, 701, 676, 1, 0, 0, 0, 701, 677, 1, 0, 0, 0, 701, 678,
		1, 0, 0, 0, 701, 679, 1, 0, 0, 0, 701, 680, 1, 0, 0, 0, 701, 681, 1, 0,
		0, 0, 701, 682, 1, 0, 0, 0, 701, 683, 1, 0, 0, 0, 701, 684, 1, 0, 0, 0,
		701, 685, 1, 0, 0, 0, 701, 686, 1, 0, 0, 0, 701, 687, 1, 0, 0, 0, 701,
		688, 1, 0, 0, 0, 701, 689, 1, 0, 0, 0, 701, 690, 1, 0, 0, 0, 701, 691,
		1, 0, 0, 0, 701, 692, 1, 0, 0, 0, 701, 693, 1, 0, 0, 0, 701, 694, 1, 0,
		0, 0, 701, 695, 1, 0, 0, 0, 702, 151, 1, 0, 0, 0, 41, 169, 180, 188, 191,
		197, 201, 210, 221, 225, 232, 271, 275, 294, 303, 314, 319, 327, 330, 334,
		345, 350, 359, 380, 388, 391, 395, 407, 426, 488, 500, 540, 543, 547, 587,
		612, 635, 640, 643, 652, 654, 701,
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
	arishemParserT__44              = 45
	arishemParserT__45              = 46
	arishemParserT__46              = 47
	arishemParserT__47              = 48
	arishemParserT__48              = 49
	arishemParserT__49              = 50
	arishemParserT__50              = 51
	arishemParserT__51              = 52
	arishemParserT__52              = 53
	arishemParserT__53              = 54
	arishemParserT__54              = 55
	arishemParserT__55              = 56
	arishemParserT__56              = 57
	arishemParserT__57              = 58
	arishemParserT__58              = 59
	arishemParserT__59              = 60
	arishemParserT__60              = 61
	arishemParserT__61              = 62
	arishemParserT__62              = 63
	arishemParserOPERATOR_LOGIC_STR = 64
	arishemParserNULL               = 65
	arishemParserDOT                = 66
	arishemParserCOMMA              = 67
	arishemParserCOLON              = 68
	arishemParserQUOTATION          = 69
	arishemParserL_BRACKET          = 70
	arishemParserR_BRACKET          = 71
	arishemParserL_BRACE            = 72
	arishemParserR_BRACE            = 73
	arishemParserINT                = 74
	arishemParserEXP                = 75
	arishemParserNAME_KEY_TYPE      = 76
	arishemParserSP_UNICODE         = 77
	arishemParserAND                = 78
	arishemParserOR                 = 79
	arishemParserNOT                = 80
	arishemParserPLUS               = 81
	arishemParserSUB                = 82
	arishemParserDIV                = 83
	arishemParserMUL                = 84
	arishemParserMOD                = 85
	arishemParserEQUALS             = 86
	arishemParserNOT_EQUALS         = 87
	arishemParserGT                 = 88
	arishemParserLT                 = 89
	arishemParserGTE                = 90
	arishemParserLTE                = 91
	arishemParserWS                 = 92
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
	arishemParserRULE_subCond             = 13
	arishemParserRULE_constant            = 14
	arishemParserRULE_constantList        = 15
	arishemParserRULE_conditionsKey       = 16
	arishemParserRULE_conditionsVal       = 17
	arishemParserRULE_conditionsPair      = 18
	arishemParserRULE_conditionGroupsKey  = 19
	arishemParserRULE_conditionGroupsVal  = 20
	arishemParserRULE_conditionGroupsPair = 21
	arishemParserRULE_listExprKey         = 22
	arishemParserRULE_mapExprKey          = 23
	arishemParserRULE_mathExprKey         = 24
	arishemParserRULE_funcExprKey         = 25
	arishemParserRULE_varExprKey          = 26
	arishemParserRULE_featureExprKey      = 27
	arishemParserRULE_featurePathKey      = 28
	arishemParserRULE_constKey            = 29
	arishemParserRULE_constListKey        = 30
	arishemParserRULE_subCondExprKey      = 31
	arishemParserRULE_condNameKey         = 32
	arishemParserRULE_varPath             = 33
	arishemParserRULE_varPathVal          = 34
	arishemParserRULE_featurePath         = 35
	arishemParserRULE_featurePathVal      = 36
	arishemParserRULE_paramMapKey         = 37
	arishemParserRULE_paramListKey        = 38
	arishemParserRULE_builtinParamKey     = 39
	arishemParserRULE_actionAimKey        = 40
	arishemParserRULE_actionAimVal        = 41
	arishemParserRULE_actionAimPair       = 42
	arishemParserRULE_nameKey             = 43
	arishemParserRULE_keyValue            = 44
	arishemParserRULE_keyValues           = 45
	arishemParserRULE_opLogicKey          = 46
	arishemParserRULE_opLogicVal          = 47
	arishemParserRULE_opLogicPair         = 48
	arishemParserRULE_opLogic             = 49
	arishemParserRULE_opMathKey           = 50
	arishemParserRULE_lrOpMathVal         = 51
	arishemParserRULE_listOpMathVal       = 52
	arishemParserRULE_lrOpMathPair        = 53
	arishemParserRULE_listOpMathPair      = 54
	arishemParserRULE_listMathOperator    = 55
	arishemParserRULE_lrMathOperator      = 56
	arishemParserRULE_logicMathOperator   = 57
	arishemParserRULE_funcNameKey         = 58
	arishemParserRULE_funcNamePair        = 59
	arishemParserRULE_operatorKey         = 60
	arishemParserRULE_operatorVal         = 61
	arishemParserRULE_operatorPair        = 62
	arishemParserRULE_operatorType        = 63
	arishemParserRULE_lhsKey              = 64
	arishemParserRULE_lhsPair             = 65
	arishemParserRULE_rhsKey              = 66
	arishemParserRULE_rhsPair             = 67
	arishemParserRULE_boolType            = 68
	arishemParserRULE_numberType          = 69
	arishemParserRULE_stringType          = 70
	arishemParserRULE_specialChars        = 71
	arishemParserRULE_strConstKey         = 72
	arishemParserRULE_numConstKey         = 73
	arishemParserRULE_boolConstKey        = 74
	arishemParserRULE_operator            = 75
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
		p.SetState(152)
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
		p.SetState(154)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNestConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(157)
			p.OpLogicPair()
		}
		{
			p.SetState(158)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(159)
			p.ConditionGroupsPair()
		}
		{
			p.SetState(160)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewLeafConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(163)
			p.OpLogicPair()
		}
		{
			p.SetState(164)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(165)
			p.ConditionsPair()
		}
		{
			p.SetState(166)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullConditionGroupContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
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

	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(172)
			p.OperatorPair()
		}
		{
			p.SetState(173)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(174)
			p.LhsPair()
		}
		{
			p.SetState(175)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(176)
			p.RhsPair()
		}
		{
			p.SetState(177)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewActionAimListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserL_BRACE {
			{
				p.SetState(183)
				p.ActionAimType()
			}
			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(184)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(185)
					p.ActionAimType()
				}

				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(193)
			p.Match(arishemParserR_BRACKET)
		}

	case 2:
		localctx = NewActionAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.ActionAimType()
		}

	case 3:
		localctx = NewExprAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Match(arishemParserL_BRACE)
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserQUOTATION {
			{
				p.SetState(196)
				p.ExprType()
			}

		}
		{
			p.SetState(199)
			p.Match(arishemParserR_BRACE)
		}

	case 4:
		localctx = NewNullAimContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(200)
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

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamMapActionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(204)
			p.ActionAimPair()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(205)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(206)
				p.ParamMapKey()
			}
			{
				p.SetState(207)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(208)
				p.KeyValues()
			}

		}
		{
			p.SetState(212)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewParamListActionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(215)
			p.ActionAimPair()
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(216)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(217)
				p.ParamListKey()
			}
			{
				p.SetState(218)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(219)
				p.Exprs()
			}

		}
		{
			p.SetState(223)
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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(228)
			p.ExprType()
		}
		{
			p.SetState(229)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
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

type SubCondExprTypeContext struct {
	*ExprTypeContext
}

func NewSubCondExprTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubCondExprTypeContext {
	var p = new(SubCondExprTypeContext)

	p.ExprTypeContext = NewEmptyExprTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprTypeContext))

	return p
}

func (s *SubCondExprTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubCondExprTypeContext) SubCondExprKey() ISubCondExprKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubCondExprKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubCondExprKeyContext)
}

func (s *SubCondExprTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *SubCondExprTypeContext) SubCond() ISubCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubCondContext)
}

func (s *SubCondExprTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitSubCondExprType(s)

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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewListExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.ListExprKey()
		}
		{
			p.SetState(235)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(236)
			p.Exprs()
		}

	case 2:
		localctx = NewMapExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.MapExprKey()
		}
		{
			p.SetState(239)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(240)
			p.KeyValues()
		}

	case 3:
		localctx = NewConstExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.ConstKey()
		}
		{
			p.SetState(243)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(244)
			p.Constant()
		}

	case 4:
		localctx = NewConstListExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.ConstListKey()
		}
		{
			p.SetState(247)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(248)
			p.ConstantList()
		}

	case 5:
		localctx = NewMathExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(250)
			p.MathExprKey()
		}
		{
			p.SetState(251)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(252)
			p.Math()
		}

	case 6:
		localctx = NewFuncExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
			p.FuncExprKey()
		}
		{
			p.SetState(255)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(256)
			p.Function()
		}

	case 7:
		localctx = NewVarExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(258)
			p.VarExprKey()
		}
		{
			p.SetState(259)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(260)
			p.Variable()
		}

	case 8:
		localctx = NewFeatureExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(262)
			p.FeatureExprKey()
		}
		{
			p.SetState(263)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(264)
			p.Feature()
		}

	case 9:
		localctx = NewSubCondExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(266)
			p.SubCondExprKey()
		}
		{
			p.SetState(267)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(268)
			p.SubCond()
		}

	case 10:
		localctx = NewNullExprTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(270)
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

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserQUOTATION:
		localctx = NewFullVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(273)
			p.VarPath()
		}

	case arishemParserNULL:
		localctx = NewNullVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
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

func (s *FullMathContext) LrOpMathPair() ILrOpMathPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILrOpMathPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILrOpMathPairContext)
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

func (s *ListMathContext) ListOpMathPair() IListOpMathPairContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListOpMathPairContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListOpMathPairContext)
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

	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFullMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(278)
			p.LrOpMathPair()
		}
		{
			p.SetState(279)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(280)
			p.LhsPair()
		}
		{
			p.SetState(281)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(282)
			p.RhsPair()
		}
		{
			p.SetState(283)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewListMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(286)
			p.ListOpMathPair()
		}
		{
			p.SetState(287)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(288)
			p.ParamListKey()
		}
		{
			p.SetState(289)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(290)
			p.Exprs()
		}
		{
			p.SetState(291)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullMathContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(293)
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

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamMapFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(297)
			p.FuncNamePair()
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(298)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(299)
				p.ParamMapKey()
			}
			{
				p.SetState(300)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(301)
				p.KeyValues()
			}

		}
		{
			p.SetState(305)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewParamListFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(308)
			p.FuncNamePair()
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(309)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(310)
				p.ParamListKey()
			}
			{
				p.SetState(311)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(312)
				p.Exprs()
			}

		}
		{
			p.SetState(316)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewNullFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(318)
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

	p.SetState(334)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACKET:
		localctx = NewFullExprsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserL_BRACE {
			{
				p.SetState(322)
				p.Expr()
			}
			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(323)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(324)
					p.Expr()
				}

				p.SetState(329)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(332)
			p.Match(arishemParserR_BRACKET)
		}

	case arishemParserNULL:
		localctx = NewNullExprsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
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

	p.SetState(350)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullFeatureContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(337)
			p.FeaturePathKey()
		}
		{
			p.SetState(338)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(339)
			p.FeaturePath()
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserCOMMA {
			{
				p.SetState(340)
				p.Match(arishemParserCOMMA)
			}
			{
				p.SetState(341)
				p.BuiltinParamKey()
			}
			{
				p.SetState(342)
				p.Match(arishemParserCOLON)
			}
			{
				p.SetState(343)
				p.KeyValues()
			}

		}
		{
			p.SetState(347)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullFeatureContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Match(arishemParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubCondContext is an interface to support dynamic dispatch.
type ISubCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSubCondContext differentiates from other interfaces.
	IsSubCondContext()
}

type SubCondContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptySubCondContext() *SubCondContext {
	var p = new(SubCondContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_subCond
	return p
}

func (*SubCondContext) IsSubCondContext() {}

func NewSubCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubCondContext {
	var p = new(SubCondContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_subCond

	return p
}

func (s *SubCondContext) GetParser() antlr.Parser { return s.parser }

func (s *SubCondContext) CopyFrom(ctx *SubCondContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SubCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullSubCondContext struct {
	*SubCondContext
}

func NewFullSubCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullSubCondContext {
	var p = new(FullSubCondContext)

	p.SubCondContext = NewEmptySubCondContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubCondContext))

	return p
}

func (s *FullSubCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullSubCondContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserL_BRACE, 0)
}

func (s *FullSubCondContext) CondNameKey() ICondNameKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondNameKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondNameKeyContext)
}

func (s *FullSubCondContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *FullSubCondContext) NameKey() INameKeyContext {
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

func (s *FullSubCondContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(arishemParserR_BRACE, 0)
}

func (s *FullSubCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitFullSubCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullSubCondContext struct {
	*SubCondContext
}

func NewNullSubCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullSubCondContext {
	var p = new(NullSubCondContext)

	p.SubCondContext = NewEmptySubCondContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubCondContext))

	return p
}

func (s *NullSubCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullSubCondContext) NULL() antlr.TerminalNode {
	return s.GetToken(arishemParserNULL, 0)
}

func (s *NullSubCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitNullSubCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) SubCond() (localctx ISubCondContext) {
	this := p
	_ = this

	localctx = NewSubCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, arishemParserRULE_subCond)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullSubCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(352)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(353)
			p.CondNameKey()
		}
		{
			p.SetState(354)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(355)
			p.NameKey()
		}
		{
			p.SetState(356)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullSubCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
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
	p.EnterRule(localctx, 28, arishemParserRULE_constant)

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

	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(362)
			p.NumConstKey()
		}
		{
			p.SetState(363)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(364)
			p.NumberType()
		}
		{
			p.SetState(365)
			p.Match(arishemParserR_BRACE)
		}

	case 2:
		localctx = NewStrConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(368)
			p.StrConstKey()
		}
		{
			p.SetState(369)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(370)
			p.StringType()
		}
		{
			p.SetState(371)
			p.Match(arishemParserR_BRACE)
		}

	case 3:
		localctx = NewBoolConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Match(arishemParserL_BRACE)
		}
		{
			p.SetState(374)
			p.BoolConstKey()
		}
		{
			p.SetState(375)
			p.Match(arishemParserCOLON)
		}
		{
			p.SetState(376)
			p.BoolType()
		}
		{
			p.SetState(377)
			p.Match(arishemParserR_BRACE)
		}

	case 4:
		localctx = NewNullConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(379)
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
	p.EnterRule(localctx, 30, arishemParserRULE_constantList)
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

	p.SetState(395)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACKET:
		localctx = NewFullConstListContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.Match(arishemParserL_BRACKET)
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserNULL || _la == arishemParserL_BRACE {
			{
				p.SetState(383)
				p.Constant()
			}
			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(384)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(385)
					p.Constant()
				}

				p.SetState(390)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(393)
			p.Match(arishemParserR_BRACKET)
		}

	case arishemParserNULL:
		localctx = NewNullConstListContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
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
	p.EnterRule(localctx, 32, arishemParserRULE_conditionsKey)

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
		p.SetState(397)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(398)
		p.Match(arishemParserT__0)
	}
	{
		p.SetState(399)
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
	p.EnterRule(localctx, 34, arishemParserRULE_conditionsVal)
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
		p.SetState(401)
		p.Match(arishemParserL_BRACKET)
	}
	{
		p.SetState(402)
		p.Condition()
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserCOMMA {
		{
			p.SetState(403)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(404)
			p.Condition()
		}

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(410)
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
	p.EnterRule(localctx, 36, arishemParserRULE_conditionsPair)

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
		p.SetState(412)
		p.ConditionsKey()
	}
	{
		p.SetState(413)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(414)
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
	p.EnterRule(localctx, 38, arishemParserRULE_conditionGroupsKey)

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
		p.SetState(416)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(417)
		p.Match(arishemParserT__1)
	}
	{
		p.SetState(418)
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
	p.EnterRule(localctx, 40, arishemParserRULE_conditionGroupsVal)
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
		p.SetState(420)
		p.Match(arishemParserL_BRACKET)
	}
	{
		p.SetState(421)
		p.ConditionGroup()
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserCOMMA {
		{
			p.SetState(422)
			p.Match(arishemParserCOMMA)
		}
		{
			p.SetState(423)
			p.ConditionGroup()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(429)
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
	p.EnterRule(localctx, 42, arishemParserRULE_conditionGroupsPair)

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
		p.SetState(431)
		p.ConditionGroupsKey()
	}
	{
		p.SetState(432)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(433)
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
	p.EnterRule(localctx, 44, arishemParserRULE_listExprKey)

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
		p.SetState(435)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(436)
		p.Match(arishemParserT__2)
	}
	{
		p.SetState(437)
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
	p.EnterRule(localctx, 46, arishemParserRULE_mapExprKey)

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
		p.SetState(439)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(440)
		p.Match(arishemParserT__3)
	}
	{
		p.SetState(441)
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
	p.EnterRule(localctx, 48, arishemParserRULE_mathExprKey)

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
		p.SetState(443)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(444)
		p.Match(arishemParserT__4)
	}
	{
		p.SetState(445)
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
	p.EnterRule(localctx, 50, arishemParserRULE_funcExprKey)

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
		p.SetState(447)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(448)
		p.Match(arishemParserT__5)
	}
	{
		p.SetState(449)
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
	p.EnterRule(localctx, 52, arishemParserRULE_varExprKey)

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
		p.SetState(451)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(452)
		p.Match(arishemParserT__6)
	}
	{
		p.SetState(453)
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
	p.EnterRule(localctx, 54, arishemParserRULE_featureExprKey)

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
		p.SetState(455)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(456)
		p.Match(arishemParserT__7)
	}
	{
		p.SetState(457)
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
	p.EnterRule(localctx, 56, arishemParserRULE_featurePathKey)

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
		p.SetState(459)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(460)
		p.Match(arishemParserT__8)
	}
	{
		p.SetState(461)
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
	p.EnterRule(localctx, 58, arishemParserRULE_constKey)

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
		p.SetState(463)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(464)
		p.Match(arishemParserT__9)
	}
	{
		p.SetState(465)
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
	p.EnterRule(localctx, 60, arishemParserRULE_constListKey)

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
		p.SetState(467)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(468)
		p.Match(arishemParserT__10)
	}
	{
		p.SetState(469)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ISubCondExprKeyContext is an interface to support dynamic dispatch.
type ISubCondExprKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsSubCondExprKeyContext differentiates from other interfaces.
	IsSubCondExprKeyContext()
}

type SubCondExprKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptySubCondExprKeyContext() *SubCondExprKeyContext {
	var p = new(SubCondExprKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_subCondExprKey
	return p
}

func (*SubCondExprKeyContext) IsSubCondExprKeyContext() {}

func NewSubCondExprKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubCondExprKeyContext {
	var p = new(SubCondExprKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_subCondExprKey

	return p
}

func (s *SubCondExprKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *SubCondExprKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *SubCondExprKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *SubCondExprKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubCondExprKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubCondExprKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitSubCondExprKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) SubCondExprKey() (localctx ISubCondExprKeyContext) {
	this := p
	_ = this

	localctx = NewSubCondExprKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, arishemParserRULE_subCondExprKey)

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
		p.SetState(471)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(472)
		p.Match(arishemParserT__11)
	}
	{
		p.SetState(473)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ICondNameKeyContext is an interface to support dynamic dispatch.
type ICondNameKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode

	// IsCondNameKeyContext differentiates from other interfaces.
	IsCondNameKeyContext()
}

type CondNameKeyContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondNameKeyContext() *CondNameKeyContext {
	var p = new(CondNameKeyContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_condNameKey
	return p
}

func (*CondNameKeyContext) IsCondNameKeyContext() {}

func NewCondNameKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondNameKeyContext {
	var p = new(CondNameKeyContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_condNameKey

	return p
}

func (s *CondNameKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *CondNameKeyContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *CondNameKeyContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *CondNameKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondNameKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondNameKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitCondNameKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) CondNameKey() (localctx ICondNameKeyContext) {
	this := p
	_ = this

	localctx = NewCondNameKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, arishemParserRULE_condNameKey)

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
		p.SetState(475)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(476)
		p.Match(arishemParserT__12)
	}
	{
		p.SetState(477)
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
	p.EnterRule(localctx, 66, arishemParserRULE_varPath)

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
		p.SetState(479)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(480)
		p.VarPathVal()
	}
	{
		p.SetState(481)
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
	p.EnterRule(localctx, 68, arishemParserRULE_varPathVal)
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
		p.SetState(483)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == arishemParserDOT {
		{
			p.SetState(484)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(485)
			p.Match(arishemParserNAME_KEY_TYPE)
		}

		p.SetState(490)
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
	p.EnterRule(localctx, 70, arishemParserRULE_featurePath)

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
		p.SetState(491)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(492)
		p.FeaturePathVal()
	}
	{
		p.SetState(493)
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
	p.EnterRule(localctx, 72, arishemParserRULE_featurePathVal)
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
		p.SetState(495)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == arishemParserDOT {
		{
			p.SetState(496)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(497)
			p.Match(arishemParserNAME_KEY_TYPE)
		}

		p.SetState(500)
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
	p.EnterRule(localctx, 74, arishemParserRULE_paramMapKey)

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
		p.SetState(502)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(503)
		p.Match(arishemParserT__13)
	}
	{
		p.SetState(504)
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
	p.EnterRule(localctx, 76, arishemParserRULE_paramListKey)

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
		p.SetState(506)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(507)
		p.Match(arishemParserT__14)
	}
	{
		p.SetState(508)
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
	p.EnterRule(localctx, 78, arishemParserRULE_builtinParamKey)

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
		p.SetState(510)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(511)
		p.Match(arishemParserT__15)
	}
	{
		p.SetState(512)
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
	p.EnterRule(localctx, 80, arishemParserRULE_actionAimKey)

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
		p.SetState(514)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(515)
		p.Match(arishemParserT__16)
	}
	{
		p.SetState(516)
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
	p.EnterRule(localctx, 82, arishemParserRULE_actionAimVal)

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
		p.SetState(518)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(519)
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	{
		p.SetState(520)
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
	p.EnterRule(localctx, 84, arishemParserRULE_actionAimPair)

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
		p.SetState(522)
		p.ActionAimKey()
	}
	{
		p.SetState(523)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(524)
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
	p.EnterRule(localctx, 86, arishemParserRULE_nameKey)

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
		p.Match(arishemParserNAME_KEY_TYPE)
	}
	{
		p.SetState(528)
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
	p.EnterRule(localctx, 88, arishemParserRULE_keyValue)

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
		p.NameKey()
	}
	{
		p.SetState(531)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(532)
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
	p.EnterRule(localctx, 90, arishemParserRULE_keyValues)
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

	p.SetState(547)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserL_BRACE:
		localctx = NewFullKeyValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Match(arishemParserL_BRACE)
		}
		p.SetState(543)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == arishemParserQUOTATION {
			{
				p.SetState(535)
				p.KeyValue()
			}
			p.SetState(540)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == arishemParserCOMMA {
				{
					p.SetState(536)
					p.Match(arishemParserCOMMA)
				}
				{
					p.SetState(537)
					p.KeyValue()
				}

				p.SetState(542)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(545)
			p.Match(arishemParserR_BRACE)
		}

	case arishemParserNULL:
		localctx = NewNullKeyValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(546)
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
	p.EnterRule(localctx, 92, arishemParserRULE_opLogicKey)

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
		p.SetState(549)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(550)
		p.Match(arishemParserT__17)
	}
	{
		p.SetState(551)
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
	p.EnterRule(localctx, 94, arishemParserRULE_opLogicVal)

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
		p.SetState(553)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(554)
		p.OpLogic()
	}
	{
		p.SetState(555)
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
	p.EnterRule(localctx, 96, arishemParserRULE_opLogicPair)

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
		p.SetState(557)
		p.OpLogicKey()
	}
	{
		p.SetState(558)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(559)
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
	p.EnterRule(localctx, 98, arishemParserRULE_opLogic)
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
		p.SetState(561)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&114689) != 0) {
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
	p.EnterRule(localctx, 100, arishemParserRULE_opMathKey)

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
		p.SetState(563)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(564)
		p.Match(arishemParserT__18)
	}
	{
		p.SetState(565)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ILrOpMathValContext is an interface to support dynamic dispatch.
type ILrOpMathValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	LrMathOperator() ILrMathOperatorContext

	// IsLrOpMathValContext differentiates from other interfaces.
	IsLrOpMathValContext()
}

type LrOpMathValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLrOpMathValContext() *LrOpMathValContext {
	var p = new(LrOpMathValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_lrOpMathVal
	return p
}

func (*LrOpMathValContext) IsLrOpMathValContext() {}

func NewLrOpMathValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LrOpMathValContext {
	var p = new(LrOpMathValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_lrOpMathVal

	return p
}

func (s *LrOpMathValContext) GetParser() antlr.Parser { return s.parser }

func (s *LrOpMathValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *LrOpMathValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *LrOpMathValContext) LrMathOperator() ILrMathOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILrMathOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILrMathOperatorContext)
}

func (s *LrOpMathValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrOpMathValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LrOpMathValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLrOpMathVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LrOpMathVal() (localctx ILrOpMathValContext) {
	this := p
	_ = this

	localctx = NewLrOpMathValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, arishemParserRULE_lrOpMathVal)

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
		p.SetState(567)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(568)
		p.LrMathOperator()
	}
	{
		p.SetState(569)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IListOpMathValContext is an interface to support dynamic dispatch.
type IListOpMathValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQUOTATION() []antlr.TerminalNode
	QUOTATION(i int) antlr.TerminalNode
	ListMathOperator() IListMathOperatorContext

	// IsListOpMathValContext differentiates from other interfaces.
	IsListOpMathValContext()
}

type ListOpMathValContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyListOpMathValContext() *ListOpMathValContext {
	var p = new(ListOpMathValContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_listOpMathVal
	return p
}

func (*ListOpMathValContext) IsListOpMathValContext() {}

func NewListOpMathValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListOpMathValContext {
	var p = new(ListOpMathValContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_listOpMathVal

	return p
}

func (s *ListOpMathValContext) GetParser() antlr.Parser { return s.parser }

func (s *ListOpMathValContext) AllQUOTATION() []antlr.TerminalNode {
	return s.GetTokens(arishemParserQUOTATION)
}

func (s *ListOpMathValContext) QUOTATION(i int) antlr.TerminalNode {
	return s.GetToken(arishemParserQUOTATION, i)
}

func (s *ListOpMathValContext) ListMathOperator() IListMathOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListMathOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListMathOperatorContext)
}

func (s *ListOpMathValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOpMathValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListOpMathValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListOpMathVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ListOpMathVal() (localctx IListOpMathValContext) {
	this := p
	_ = this

	localctx = NewListOpMathValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, arishemParserRULE_listOpMathVal)

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
		p.SetState(571)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(572)
		p.ListMathOperator()
	}
	{
		p.SetState(573)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// ILrOpMathPairContext is an interface to support dynamic dispatch.
type ILrOpMathPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpMathKey() IOpMathKeyContext
	COLON() antlr.TerminalNode
	LrOpMathVal() ILrOpMathValContext

	// IsLrOpMathPairContext differentiates from other interfaces.
	IsLrOpMathPairContext()
}

type LrOpMathPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLrOpMathPairContext() *LrOpMathPairContext {
	var p = new(LrOpMathPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_lrOpMathPair
	return p
}

func (*LrOpMathPairContext) IsLrOpMathPairContext() {}

func NewLrOpMathPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LrOpMathPairContext {
	var p = new(LrOpMathPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_lrOpMathPair

	return p
}

func (s *LrOpMathPairContext) GetParser() antlr.Parser { return s.parser }

func (s *LrOpMathPairContext) OpMathKey() IOpMathKeyContext {
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

func (s *LrOpMathPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *LrOpMathPairContext) LrOpMathVal() ILrOpMathValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILrOpMathValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILrOpMathValContext)
}

func (s *LrOpMathPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrOpMathPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LrOpMathPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLrOpMathPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LrOpMathPair() (localctx ILrOpMathPairContext) {
	this := p
	_ = this

	localctx = NewLrOpMathPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, arishemParserRULE_lrOpMathPair)

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
		p.SetState(575)
		p.OpMathKey()
	}
	{
		p.SetState(576)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(577)
		p.LrOpMathVal()
	}

	return localctx
}

// IListOpMathPairContext is an interface to support dynamic dispatch.
type IListOpMathPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpMathKey() IOpMathKeyContext
	COLON() antlr.TerminalNode
	ListOpMathVal() IListOpMathValContext

	// IsListOpMathPairContext differentiates from other interfaces.
	IsListOpMathPairContext()
}

type ListOpMathPairContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyListOpMathPairContext() *ListOpMathPairContext {
	var p = new(ListOpMathPairContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_listOpMathPair
	return p
}

func (*ListOpMathPairContext) IsListOpMathPairContext() {}

func NewListOpMathPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListOpMathPairContext {
	var p = new(ListOpMathPairContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_listOpMathPair

	return p
}

func (s *ListOpMathPairContext) GetParser() antlr.Parser { return s.parser }

func (s *ListOpMathPairContext) OpMathKey() IOpMathKeyContext {
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

func (s *ListOpMathPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(arishemParserCOLON, 0)
}

func (s *ListOpMathPairContext) ListOpMathVal() IListOpMathValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListOpMathValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListOpMathValContext)
}

func (s *ListOpMathPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOpMathPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListOpMathPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListOpMathPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ListOpMathPair() (localctx IListOpMathPairContext) {
	this := p
	_ = this

	localctx = NewListOpMathPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, arishemParserRULE_listOpMathPair)

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
		p.SetState(579)
		p.OpMathKey()
	}
	{
		p.SetState(580)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(581)
		p.ListOpMathVal()
	}

	return localctx
}

// IListMathOperatorContext is an interface to support dynamic dispatch.
type IListMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsListMathOperatorContext differentiates from other interfaces.
	IsListMathOperatorContext()
}

type ListMathOperatorContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyListMathOperatorContext() *ListMathOperatorContext {
	var p = new(ListMathOperatorContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_listMathOperator
	return p
}

func (*ListMathOperatorContext) IsListMathOperatorContext() {}

func NewListMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListMathOperatorContext {
	var p = new(ListMathOperatorContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_listMathOperator

	return p
}

func (s *ListMathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ListMathOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(arishemParserPLUS, 0)
}

func (s *ListMathOperatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(arishemParserSUB, 0)
}

func (s *ListMathOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(arishemParserMUL, 0)
}

func (s *ListMathOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(arishemParserDIV, 0)
}

func (s *ListMathOperatorContext) MOD() antlr.TerminalNode {
	return s.GetToken(arishemParserMOD, 0)
}

func (s *ListMathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListMathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListMathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitListMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) ListMathOperator() (localctx IListMathOperatorContext) {
	this := p
	_ = this

	localctx = NewListMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, arishemParserRULE_listMathOperator)
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
		p.SetState(583)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-81)) & ^0x3f) == 0 && ((int64(1)<<(_la-81))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILrMathOperatorContext is an interface to support dynamic dispatch.
type ILrMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ListMathOperator() IListMathOperatorContext
	LogicMathOperator() ILogicMathOperatorContext

	// IsLrMathOperatorContext differentiates from other interfaces.
	IsLrMathOperatorContext()
}

type LrMathOperatorContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLrMathOperatorContext() *LrMathOperatorContext {
	var p = new(LrMathOperatorContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_lrMathOperator
	return p
}

func (*LrMathOperatorContext) IsLrMathOperatorContext() {}

func NewLrMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LrMathOperatorContext {
	var p = new(LrMathOperatorContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_lrMathOperator

	return p
}

func (s *LrMathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LrMathOperatorContext) ListMathOperator() IListMathOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListMathOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListMathOperatorContext)
}

func (s *LrMathOperatorContext) LogicMathOperator() ILogicMathOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicMathOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicMathOperatorContext)
}

func (s *LrMathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrMathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LrMathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLrMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LrMathOperator() (localctx ILrMathOperatorContext) {
	this := p
	_ = this

	localctx = NewLrMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, arishemParserRULE_lrMathOperator)

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

	p.SetState(587)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserPLUS, arishemParserSUB, arishemParserDIV, arishemParserMUL, arishemParserMOD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(585)
			p.ListMathOperator()
		}

	case arishemParserEQUALS, arishemParserNOT_EQUALS, arishemParserGT, arishemParserLT, arishemParserGTE, arishemParserLTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(586)
			p.LogicMathOperator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILogicMathOperatorContext is an interface to support dynamic dispatch.
type ILogicMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUALS() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode

	// IsLogicMathOperatorContext differentiates from other interfaces.
	IsLogicMathOperatorContext()
}

type LogicMathOperatorContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicMathOperatorContext() *LogicMathOperatorContext {
	var p = new(LogicMathOperatorContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_logicMathOperator
	return p
}

func (*LogicMathOperatorContext) IsLogicMathOperatorContext() {}

func NewLogicMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicMathOperatorContext {
	var p = new(LogicMathOperatorContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_logicMathOperator

	return p
}

func (s *LogicMathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicMathOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(arishemParserEQUALS, 0)
}

func (s *LogicMathOperatorContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(arishemParserNOT_EQUALS, 0)
}

func (s *LogicMathOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(arishemParserGT, 0)
}

func (s *LogicMathOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(arishemParserGTE, 0)
}

func (s *LogicMathOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(arishemParserLT, 0)
}

func (s *LogicMathOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(arishemParserLTE, 0)
}

func (s *LogicMathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicMathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicMathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitLogicMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) LogicMathOperator() (localctx ILogicMathOperatorContext) {
	this := p
	_ = this

	localctx = NewLogicMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, arishemParserRULE_logicMathOperator)
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
		p.SetState(589)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-86)) & ^0x3f) == 0 && ((int64(1)<<(_la-86))&63) != 0) {
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
	p.EnterRule(localctx, 116, arishemParserRULE_funcNameKey)

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
		p.SetState(591)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(592)
		p.Match(arishemParserT__19)
	}
	{
		p.SetState(593)
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
	p.EnterRule(localctx, 118, arishemParserRULE_funcNamePair)

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
		p.SetState(595)
		p.FuncNameKey()
	}
	{
		p.SetState(596)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(597)
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
	p.EnterRule(localctx, 120, arishemParserRULE_operatorKey)

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
		p.SetState(599)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(600)
		p.Match(arishemParserT__20)
	}
	{
		p.SetState(601)
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
	p.EnterRule(localctx, 122, arishemParserRULE_operatorVal)

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
		p.SetState(603)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(604)
		p.OperatorType()
	}
	{
		p.SetState(605)
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
	p.EnterRule(localctx, 124, arishemParserRULE_operatorPair)

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
		p.SetState(607)
		p.OperatorKey()
	}
	{
		p.SetState(608)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(609)
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
	Operator() IOperatorContext
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

func (s *OperatorTypeContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
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
	p.EnterRule(localctx, 126, arishemParserRULE_operatorType)
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
	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserT__21 || _la == arishemParserNOT {
		{
			p.SetState(611)
			_la = p.GetTokenStream().LA(1)

			if !(_la == arishemParserT__21 || _la == arishemParserNOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(614)
		p.Operator()
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
	p.EnterRule(localctx, 128, arishemParserRULE_lhsKey)

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
		p.SetState(616)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(617)
		p.Match(arishemParserT__22)
	}
	{
		p.SetState(618)
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
	p.EnterRule(localctx, 130, arishemParserRULE_lhsPair)

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
		p.SetState(620)
		p.LhsKey()
	}
	{
		p.SetState(621)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(622)
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
	p.EnterRule(localctx, 132, arishemParserRULE_rhsKey)

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
		p.SetState(624)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(625)
		p.Match(arishemParserT__23)
	}
	{
		p.SetState(626)
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
	p.EnterRule(localctx, 134, arishemParserRULE_rhsPair)

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
		p.SetState(628)
		p.RhsKey()
	}
	{
		p.SetState(629)
		p.Match(arishemParserCOLON)
	}
	{
		p.SetState(630)
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
	p.EnterRule(localctx, 136, arishemParserRULE_boolType)
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
		p.SetState(632)
		_la = p.GetTokenStream().LA(1)

		if !(_la == arishemParserT__24 || _la == arishemParserT__25) {
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
	p.EnterRule(localctx, 138, arishemParserRULE_numberType)
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
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserSUB {
		{
			p.SetState(634)
			p.Match(arishemParserSUB)
		}

	}
	{
		p.SetState(637)
		p.Match(arishemParserINT)
	}
	p.SetState(640)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserDOT {
		{
			p.SetState(638)
			p.Match(arishemParserDOT)
		}
		{
			p.SetState(639)
			p.Match(arishemParserINT)
		}

	}
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == arishemParserEXP {
		{
			p.SetState(642)
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
	p.EnterRule(localctx, 140, arishemParserRULE_stringType)
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
		p.SetState(645)
		p.Match(arishemParserQUOTATION)
	}
	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(652)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				p.SetState(646)
				p.MatchWildcard()

			case 2:
				{
					p.SetState(647)
					p.Match(arishemParserT__26)
				}
				p.SetState(648)
				p.MatchWildcard()

			case 3:
				{
					p.SetState(649)
					_la = p.GetTokenStream().LA(1)

					if _la <= 0 || _la == arishemParserT__26 || _la == arishemParserQUOTATION {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 4:
				{
					p.SetState(650)
					p.SpecialChars()
				}

			case 5:
				{
					p.SetState(651)
					p.Match(arishemParserSP_UNICODE)
				}

			}

		}
		p.SetState(656)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}
	{
		p.SetState(657)
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
	p.EnterRule(localctx, 142, arishemParserRULE_specialChars)
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
		p.SetState(659)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-27)) & ^0x3f) == 0 && ((int64(1)<<(_la-27))&7485122768422436863) != 0) {
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
	p.EnterRule(localctx, 144, arishemParserRULE_strConstKey)

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
		p.SetState(661)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(662)
		p.Match(arishemParserT__43)
	}
	{
		p.SetState(663)
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
	p.EnterRule(localctx, 146, arishemParserRULE_numConstKey)

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
		p.SetState(665)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(666)
		p.Match(arishemParserT__44)
	}
	{
		p.SetState(667)
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
	p.EnterRule(localctx, 148, arishemParserRULE_boolConstKey)

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
		p.SetState(669)
		p.Match(arishemParserQUOTATION)
	}
	{
		p.SetState(670)
		p.Match(arishemParserT__45)
	}
	{
		p.SetState(671)
		p.Match(arishemParserQUOTATION)
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUALS() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode
	GT() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	OperatorType() IOperatorTypeContext
	OpLogic() IOpLogicContext

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*ArishemParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.ArishemParserRuleContext = NewArishemParserRuleContext(nil, -1)
	p.RuleIndex = arishemParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.ArishemParserRuleContext = NewArishemParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = arishemParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(arishemParserEQUALS, 0)
}

func (s *OperatorContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(arishemParserNOT_EQUALS, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(arishemParserGT, 0)
}

func (s *OperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(arishemParserGTE, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(arishemParserLT, 0)
}

func (s *OperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(arishemParserLTE, 0)
}

func (s *OperatorContext) OperatorType() IOperatorTypeContext {
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

func (s *OperatorContext) OpLogic() IOpLogicContext {
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

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case arishemVisitor:
		return t.VisitOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *arishemParser) Operator() (localctx IOperatorContext) {
	this := p
	_ = this

	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, arishemParserRULE_operator)

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

	p.SetState(701)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case arishemParserEQUALS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(673)
			p.Match(arishemParserEQUALS)
		}

	case arishemParserNOT_EQUALS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(674)
			p.Match(arishemParserNOT_EQUALS)
		}

	case arishemParserGT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(675)
			p.Match(arishemParserGT)
		}

	case arishemParserGTE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(676)
			p.Match(arishemParserGTE)
		}

	case arishemParserLT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(677)
			p.Match(arishemParserLT)
		}

	case arishemParserLTE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(678)
			p.Match(arishemParserLTE)
		}

	case arishemParserT__46:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(679)
			p.Match(arishemParserT__46)
		}

	case arishemParserT__47:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(680)
			p.Match(arishemParserT__47)
		}

	case arishemParserT__48:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(681)
			p.Match(arishemParserT__48)
		}

	case arishemParserT__49:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(682)
			p.Match(arishemParserT__49)
		}

	case arishemParserT__50:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(683)
			p.Match(arishemParserT__50)
		}

	case arishemParserT__51:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(684)
			p.Match(arishemParserT__51)
		}

	case arishemParserT__52:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(685)
			p.Match(arishemParserT__52)
		}

	case arishemParserT__53:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(686)
			p.Match(arishemParserT__53)
		}

	case arishemParserT__54:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(687)
			p.Match(arishemParserT__54)
		}

	case arishemParserT__55:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(688)
			p.Match(arishemParserT__55)
		}

	case arishemParserT__56:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(689)
			p.Match(arishemParserT__56)
		}

	case arishemParserT__57:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(690)
			p.Match(arishemParserT__57)
		}

	case arishemParserT__58:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(691)
			p.Match(arishemParserT__58)
		}

	case arishemParserT__59:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(692)
			p.Match(arishemParserT__59)
		}

	case arishemParserT__60:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(693)
			p.Match(arishemParserT__60)
		}

	case arishemParserT__61:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(694)
			p.Match(arishemParserT__61)
		}

	case arishemParserT__62:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(695)
			p.Match(arishemParserT__62)
		}
		{
			p.SetState(696)
			p.Match(arishemParserT__27)
		}
		{
			p.SetState(697)
			p.OperatorType()
		}
		{
			p.SetState(698)
			p.Match(arishemParserT__27)
		}
		{
			p.SetState(699)
			p.OpLogic()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
