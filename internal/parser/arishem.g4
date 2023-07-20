/*
 * Copyright Bytedance Ltd. and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

grammar arishem;

options {contextSuperClass=ArishemParserRuleContext;}

// Definition of Arishem Condition and aim
condEntity  :   conditionGroup                                  #arishemCond;
aimEntity   :   aim                                             #arishemAim;

conditionGroup
    :   L_BRACE opLogicPair COMMA   conditionGroupsPair R_BRACE     #nestConditionGroup
    |   L_BRACE opLogicPair COMMA   conditionsPair      R_BRACE     #leafConditionGroup
    |   NULL                                                        #nullConditionGroup
    ;

condition
    :   L_BRACE operatorPair COMMA lhsPair COMMA rhsPair R_BRACE    #fullCondition
    |   NULL                                                        #nullCondition
    ;

aim
    :   L_BRACKET   (actionAimType  (COMMA actionAimType)*)?    R_BRACKET   #actionAimList
    |   actionAimType                                                       #actionAim
    |   L_BRACE     exprType?           R_BRACE                             #exprAim
    |   NULL                                                                #nullAim
    ;

actionAimType
    :   L_BRACE actionAimPair (COMMA paramMapKey  COLON keyValues)?   R_BRACE   #paramMapAction
    |   L_BRACE actionAimPair (COMMA paramListKey COLON exprs)?       R_BRACE   #paramListAction
    ;

expr
    :   L_BRACE exprType    R_BRACE             #fullExpr
    |   NULL                                    #nullExpr
    ;

exprType
    :   listExprKey     COLON   exprs           #listExprType
    |   mapExprKey      COLON   keyValues       #mapExprType
    |   constKey        COLON   constant        #constExprType
    |   constListKey    COLON   constantList    #constListExprType
    |   mathExprKey     COLON   math            #mathExprType
    |   funcExprKey     COLON   function        #funcExprType
    |   varExprKey      COLON   variable        #varExprType
    |   featureExprKey  COLON   feature         #featureExprType
    |   NULL                                    #nullExprType
    ;

variable
    :   varPath                                                     #fullVar
    |   NULL                                                        #nullVar
    ;

math
    :   L_BRACE lrOpMathPair    COMMA lhsPair COMMA rhsPair R_BRACE      #fullMath
    |   L_BRACE listOpMathPair  COMMA paramListKey COLON exprs R_BRACE   #listMath
    |   NULL                                                             #nullMath
    ;

function
    :   L_BRACE funcNamePair    (COMMA  paramMapKey     COLON   keyValues)? R_BRACE     #paramMapFunc
    |   L_BRACE funcNamePair    (COMMA  paramListKey    COLON   exprs)?     R_BRACE     #paramListFunc
    |   NULL                                                                            #nullFunc
    ;

exprs
    :   L_BRACKET (expr (COMMA expr)*)? R_BRACKET                                       #fullExprs
    |   NULL                                                                            #nullExprs
    ;

feature
    :   L_BRACE featurePathKey COLON featurePath (COMMA  builtinParamKey  COLON   keyValues)? R_BRACE   #fullFeature
    |   NULL                                                                                            #nullFeature
    ;

constant
    :   L_BRACE numConstKey     COLON   numberType  R_BRACE     #numConst
    |   L_BRACE strConstKey     COLON   stringType  R_BRACE     #strConst
    |   L_BRACE boolConstKey    COLON   boolType    R_BRACE     #boolConst
    |   NULL                                                    #nullConst
    ;

constantList
    :   L_BRACKET (constant (COMMA constant)*)? R_BRACKET       #fullConstList
    |   NULL                                                    #nullConstList
    ;

conditionsKey       :   QUOTATION 'Conditions' QUOTATION ;
conditionsVal       :   L_BRACKET condition (COMMA condition)* R_BRACKET ;
conditionsPair      :   conditionsKey COLON conditionsVal ;
conditionGroupsKey  :   QUOTATION 'ConditionGroups' QUOTATION ;
conditionGroupsVal  :   L_BRACKET conditionGroup (COMMA conditionGroup)* R_BRACKET ;
conditionGroupsPair :   conditionGroupsKey COLON conditionGroupsVal ;

listExprKey     :   QUOTATION   'ListExpr'      QUOTATION   ;
mapExprKey      :   QUOTATION   'MapExpr'       QUOTATION   ;
mathExprKey     :   QUOTATION   'MathExpr'      QUOTATION   ;
funcExprKey     :   QUOTATION   'FuncExpr'      QUOTATION   ;
varExprKey      :   QUOTATION   'VarExpr'       QUOTATION   ;
featureExprKey  :   QUOTATION   'FeatureExpr'   QUOTATION   ;
featurePathKey  :   QUOTATION   'FeaturePath'   QUOTATION   ;
constKey        :   QUOTATION   'Const'         QUOTATION   ;
constListKey    :   QUOTATION   'ConstList'     QUOTATION   ;

varPath         :   QUOTATION   varPathVal      QUOTATION   ;
varPathVal      :   NAME_KEY_TYPE   (DOT NAME_KEY_TYPE)*    ;
featurePath     :   QUOTATION   featurePathVal  QUOTATION   ;
featurePathVal  :   NAME_KEY_TYPE   (DOT NAME_KEY_TYPE)+    ;

paramMapKey     :   QUOTATION   'ParamMap'      QUOTATION   ;
paramListKey    :   QUOTATION   'ParamList'     QUOTATION   ;
builtinParamKey :   QUOTATION   'BuiltinParam'  QUOTATION   ;

actionAimKey    :   QUOTATION   'ActionName'    QUOTATION   ;
actionAimVal    :   QUOTATION   NAME_KEY_TYPE   QUOTATION   ;
actionAimPair   :   actionAimKey    COLON   actionAimVal    ;

nameKey     :   QUOTATION   NAME_KEY_TYPE   QUOTATION   ;
keyValue    :   nameKey     COLON   expr                ;
keyValues
    :   L_BRACE (keyValue (COMMA keyValue)*)? R_BRACE   #fullKeyValues
    |   NULL                                            #nullKeyValues
    ;

opLogicKey  : QUOTATION     'OpLogic'       QUOTATION   ;
opLogicVal  : QUOTATION     opLogic         QUOTATION   ;
opLogicPair : opLogicKey    COLON           opLogicVal  ;
opLogic
    :   OPERATOR_LOGIC_STR
    |   AND
    |   OR
    |   NOT
    ;

opMathKey       :   QUOTATION   'OpMath'    QUOTATION   ;
lrOpMathVal     :   QUOTATION   lrMathOperator      QUOTATION   ;
listOpMathVal   :   QUOTATION   listMathOperator    QUOTATION   ;
lrOpMathPair    :   opMathKey   COLON       lrOpMathVal     ;
listOpMathPair  :   opMathKey   COLON       listOpMathVal   ;

listMathOperator
    :   PLUS
    |   SUB
    |   MUL
    |   DIV
    |   MOD
    ;

lrMathOperator
    :   listMathOperator
    |   logicMathOperator
    ;

logicMathOperator
    :   EQUALS
    |   NOT_EQUALS
    |   GT
    |   GTE
    |   LT
    |   LTE
    ;

funcNameKey     :   QUOTATION   'FuncName'  QUOTATION   ;
funcNamePair    :   funcNameKey COLON       nameKey     ;

operatorKey     :   QUOTATION   'Operator'      QUOTATION   ;
operatorVal     :   QUOTATION   operatorType    QUOTATION   ;
operatorPair    :   operatorKey COLON           operatorVal ;
operatorType    :   ('NOT ' | NOT)?   operator          ;

lhsKey  :   QUOTATION   'Lhs'   QUOTATION   ;
lhsPair :   lhsKey      COLON   expr        ;
rhsKey  :   QUOTATION   'Rhs'   QUOTATION   ;
rhsPair :   rhsKey      COLON   expr        ;

boolType
    :   'true'
    |   'false'
    ;
numberType  :   SUB? INT (DOT INT)? EXP?    ;
stringType  :   QUOTATION ( . | '\\'. | ~('"'| '\\') | specialChars | SP_UNICODE )*?  QUOTATION  ;

specialChars
    :   ' ' | '!' | '"' | '#' | '$' | '%' | '&' | '\'' | '(' | ')' | '*' | '|' | '+' | ',' | '-' | '.' | '/'
    |   ':'	| ';' |	'<'	| '=' |	'>'	| '?'
    |   '@' | '[' | '\\' | '|' | ']' | '^' | '_'
    |   '`' | '{' | '|' | '}' | '~'
    ;

strConstKey     :   QUOTATION   'StrConst'  QUOTATION   ;
numConstKey     :   QUOTATION   'NumConst'  QUOTATION   ;
boolConstKey    :   QUOTATION   'BoolConst' QUOTATION   ;


fragment A  : [aA] ;
fragment B  : [bB] ;
fragment C  : [cC] ;
fragment D  : [dD] ;
fragment E  : [eE] ;
fragment F  : [fF] ;
fragment G  : [gG] ;
fragment H  : [hH] ;
fragment I  : [iI] ;
fragment J  : [jJ] ;
fragment K  : [kK] ;
fragment L  : [lL] ;
fragment M  : [mM] ;
fragment N  : [nN] ;
fragment O  : [oO] ;
fragment P  : [pP] ;
fragment Q  : [qQ] ;
fragment R  : [rR] ;
fragment S  : [sS] ;
fragment T  : [tT] ;
fragment U  : [uU] ;
fragment V  : [vV] ;
fragment W  : [wW] ;
fragment X  : [xX] ;
fragment Y  : [yY] ;
fragment Z  : [zZ] ;

OPERATOR_LOGIC_STR
    :   A N D
    |   O R
    |   N O T
    ;

operator
    :   EQUALS
    |   NOT_EQUALS
    |   GT
    |   GTE
    |   LT
    |   LTE
    |   'IS_NULL'
    |   'STRING_START_WITH'
    |   'STRING_END_WITH'
    |   'STRING_CONTAINS'
    |   'LIST_IN'
    |   'LIST_CONTAINS'
    |   'LIST_RETAIN'
    |   'CONTAIN_REGULAR'
    |   'SUB_LIST_IN'
    |   'SUB_LIST_CONTAINS'
    |   'CONTAINS_KEYWORDS'
    |   'BETWEEN_ALL_CLOSE'
    |   'BETWEEN_ALL_OPEN'
    |   'BETWEEN_LEFT_OPEN_RIGHT_CLOSE'
    |   'BETWEEN_LEFT_CLOSE_RIGHT_OPEN'
    ;

NULL : 'null' ;

DOT         : '.' ;
COMMA       : ',' ;
COLON       : ':' ;
QUOTATION   : '"' ;
L_BRACKET   : '[' ;
R_BRACKET   : ']' ;
L_BRACE     : '{' ;
R_BRACE     : '}' ;

INT : [0-9]+ ;
EXP : [Ee] [+\-]? INT;
NAME_KEY_TYPE : [a-zA-Z][a-zA-Z0-9_#]* ;
SP_UNICODE : '\u0080'..'\uFFFF' ;

AND : '&&'  ;
OR  : '||'  ;
NOT : '!'   ;

PLUS    : '+' ;
SUB     : '-' ;
DIV     : '/' ;
MUL     : '*' ;
MOD     : '%' ;

EQUALS      : '==' ;
NOT_EQUALS  : '!=' ;
GT          : '>'  ;
LT          : '<'  ;
GTE         : '>=' ;
LTE         : '<=' ;

// skip tab, newline, return
WS  :   [\t\n\r]+ -> skip ;