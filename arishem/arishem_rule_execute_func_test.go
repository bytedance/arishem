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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuncExprNoParam(t *testing.T) {
	condition := `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "FuncExpr": {
    				"FuncName": "GetCurrentYear"
				}
            },
            "Rhs": {
                "Const": {
                    "NumConst": 2023
                }
            }
        }
    ]
}`
	pass, err := JudgeCondition(condition)
	assert.Nil(t, err)
	assert.True(t, pass)
}

func TestFuncExprParamList(t *testing.T) {
	condition := `{"OpLogic":"&&","Conditions":[{"Operator":"LIST_CONTAINS","Lhs":{"FuncExpr":{"FuncName":"ListAdd","ParamList":[{"ConstList":[{"NumConst":1},{"NumConst":1}]},{"Const":{"NumConst":1}}]}},"Rhs":{"Const":{"NumConst":1}}}]}`
	pass, err := JudgeCondition(condition)
	assert.Nil(t, err)
	assert.True(t, pass)
}

func TestFuncExprParamMap(t *testing.T) {
	condition := `{"OpLogic":"&&","Conditions":[{"Operator":"LIST_CONTAINS","Lhs":{"FuncExpr":{"FuncName":"ListLength","ParamMap":{"list":{"ConstList":[{"NumConst":1},{"NumConst":1}]},"filter_null":{"Const":{"BoolConst":true}}}}},"Rhs":{"Const":{"NumConst":2}}}]}`
	pass, err := JudgeCondition(condition)
	assert.Nil(t, err)
	assert.True(t, pass)
}

func TestCustomFunc(t *testing.T) {
	condition := `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "FuncExpr": {
    				"FuncName": "NPHello"
				}
            },
            "Rhs": {
                "Const": {
                    "StrConst": "Hello Arishem"
                }
            }
        }
    ]
}`
	pass, err := JudgeCondition(condition)
	assert.Nil(t, err)
	assert.True(t, pass)
}
