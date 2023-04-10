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
	"fmt"
	"github.com/bytedance/arishem/internal/operator"
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgeConditionWithFactMeta(t *testing.T) {
	testCases := []struct {
		condition string
		factMeat  string
		check     func(pass bool, err error)
	}{
		{

			`{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"&&","Conditions":[{"Operator":"STRING_START_WITH","Lhs":{"VarExpr":"username"},"Rhs":{"Const":{"StrConst":"And"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#1"},"Rhs":{"Const":{"StrConst":"e"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"ack"}}},{"Operator":"CONTAIN_REGULAR","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"^J.*"}}}]},{"OpLogic":"and","ConditionGroups":[{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"number1"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":100},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"numbers#0"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":999},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"LIST_CONTAINS","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":false},{"NumConst":-3.1415926}]}},{"Operator":"!LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}}]}]}]}`,
			`{}`,
			func(pass bool, err error) {
				assert.Nil(t, err)
				assert.False(t, pass)
			},
		},
		{
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			`{}`,
			func(pass bool, err error) {
				assert.Nil(t, err)
				assert.True(t, pass)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(JudgeConditionWithFactMeta(tc.condition, tc.factMeat))
	}
}

func case1() {
	condition := `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "NumConst": 1
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 1
                }
            }
        }
    ]
}
`
	pass, err := JudgeCondition(condition)
	if err != nil {
		// handle error here
		// ...
		println(err.Error())
	}
	if pass {
		// your business code here
		// ...
		println("condition passed!")
	}
}

func case2() {
	condition := `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "VarExpr": "user.age"
            },
            "Rhs": {
                "VarExpr": "user_ages#1"
            }
        }
    ]
}
`
	pass, err := JudgeConditionWithFactMeta(condition, `
{
	"user": {
		"name": "KJ",
		"age": 24
	},
	"user_ages": [
		15,
		20,
		32
	]
}
`)
	if err != nil {

	}
	if pass {
		println("KJ's age is greater than 20!")
	}
}

func case3() {
	condition := `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">=",
            "Lhs": {
                "VarExpr": "user.age"
            },
            "Rhs": {
                "VarExpr": "user_ages#1"
            }
        }
    ]
}
`
	// create an expression aim
	aim := `
{
	"Const": {
		"StrConst": "rule passed!"
	}
}
`
	rule, err := NewNoPriorityRule("rule1", condition, aim)
	if err != nil {
		// handle error here
		return
	}
	dc, err := DataContext(`
{
	"user": {
		"name": "kouji",
		"age": 24
	},
	"user_ages": [
		20,
		18,
		32
	]
}
`)
	if err != nil {
		// handle error here
		return
	}
	rr := ExecuteSingleRule(rule, dc)
	if rr.Passed() {
		fmt.Printf("%s pass, output=>%s", rr.Identifier(), rr.Aim().AsExpr())
	}
}

func case4() {
	condGroup := NewConditionsCondGroup(OpLogicAnd)
	cond1 := NewCondition(operator.Equal)
	cond1.Lhs = NewConstExpr(NewNumConst(1.0))
	cond1.Rhs = NewConstExpr(NewNumConst(1.0))
	condGroup.AddConditions(cond1)
	expr, _ := condGroup.Build()
	println(expr)
}

type WeightsPriorityRule struct {
	*NoPriorityRule
	weight   float64
	priority float64
}

func (w *WeightsPriorityRule) Compare(other typedef.Comparable) int {
	o := other.(*WeightsPriorityRule)
	myWeighted := w.weight * w.priority
	otherWeighted := o.weight * o.priority
	if myWeighted < otherWeighted {
		return -1
	} else if myWeighted == otherWeighted {
		return 0
	} else {
		return 1
	}
}

func TestLWAPI(t *testing.T) {
	case1()
	case2()
	case3()
	case4()
}
