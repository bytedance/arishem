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
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/typedef"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testFeatureFetcher struct {
}

func (t *testFeatureFetcher) AddFetchObserver(v ...typedef.FeatureFetchObserver) {}

func (t *testFeatureFetcher) GetFetchObservers() []typedef.FeatureFetchObserver { return nil }

func (t *testFeatureFetcher) ClearFetchObservers() {}

func (t *testFeatureFetcher) FetchFeature(ctx context.Context, feat typedef.FeatureParam, dc typedef.DataCtx) (typedef.MetaType, error) {
	//dc.GetFeatureValue(&testFeatureParam{})
	return nil, nil
}

type testVisitTarget struct {
	name string
}

func newArisVisitTarget(name string) *testVisitTarget {
	return &testVisitTarget{name: name}
}

func (a *testVisitTarget) Identifier() string {
	return a.name
}

type testArisVisitObserver struct {
	hash    string
	errMsgs []string
}

func (t *testArisVisitObserver) HashCode() string {
	return t.hash
}

func (t *testArisVisitObserver) OnJudgeNodeVisitEnd(ctx context.Context, info typedef.JudgeNode, vt typedef.VisitTarget) {
	if info != nil {
		foreachNode, ok := info.(typedef.ForeachJudgeNode)
		if !ok {
			logger.Infof("left: %v, leftExpr: %v, right: %v, rightExpr: %v, operator: %v, target: %v, err: %v",
				info.Left(), info.LeftExpr(), info.Right(), info.RightExpr(), info.Operator(), vt.Identifier(), info.Error(),
			)
		} else {
			itemsStr, _ := sonic.MarshalString(foreachNode.ForeachItems())
			logger.Infof("[ForeachJudgeNode] left: %v, leftExpr: %v, right: %v, rightExpr: %v, operator: %v, target: %v, err: %v, foreachOperator: %v, foreachLogic: %v, foreachItems: %v",
				foreachNode.Left(), info.LeftExpr(), foreachNode.Right(), info.RightExpr(), foreachNode.Operator(), vt.Identifier(), foreachNode.Error(),
				foreachNode.ForeachOperator(), foreachNode.ForeachLogic(), itemsStr,
			)
		}
	}
}

func (t *testArisVisitObserver) OnVisitError(ctx context.Context, node, errMsg string, vt typedef.VisitTarget) {
	t.errMsgs = append(t.errMsgs, errMsg)
}

func TestConditionVisit(t *testing.T) {
	testCases := []struct {
		name  string
		fact  string
		expr  string
		check func(pass bool, errMsg []string)
	}{
		{
			"CONDITION: single condition, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: single condition, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"<","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: single condition, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: SINGLE CONDITION, JUDGE WITH CONSTANT [OPERATOR]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"!=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: single condition, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":">=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: single condition, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi conditions, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi conditions, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi conditions, judge with constant [compare operator]",
			`{}`,
			`{"OpLogic":"oR","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},

		{
			"CONDITION: single condition, judge with variable in. [compare operator]",
			`{"number1":100,"numbers":["10",1,99.9]}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"number1"},"Rhs":{"VarExpr":"numbers#0"}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: single condition, judge with variable in. [compare operator]",
			`{"number1":100,"numbers":["10",1,99.9]}`,
			`{"OpLogic":"and","Conditions":[{"Operator":">","Lhs":{"VarExpr":"number1"},"Rhs":{"VarExpr":"numbers#0"}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi condition, judge with variable in. [compare operator]",
			`{"number1":100,"numbers":["10",1,99.9]}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"number1"},"Rhs":{"VarExpr":"numbers#0"}},{"Operator":"<=","Lhs":{"VarExpr":"numbers#1"},"Rhs":{"VarExpr":"numbers#0"}},{"Operator":">","Lhs":{"VarExpr":"numbers#2"},"Rhs":{"VarExpr":"numbers#0"}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi condition, judge with variable in. [list operator]",
			`{"number1":100,"numbers":["10",99.9,0]}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"number1"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":100},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"NOT LIST_IN","Lhs":{"VarExpr":"numbers#0"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":999},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"LIST_CONTAINS","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":false},{"NumConst":-3.1415926}]}},{"Operator":"!LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: multi condition, judge with variable in. [string operator]",
			`{"username":"Andrew","usernames":["Jack","Mike","Andrew"],"news":"Jack hanged out with Mike last weekend."}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"STRING_START_WITH","Lhs":{"VarExpr":"username"},"Rhs":{"Const":{"StrConst":"And"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#1"},"Rhs":{"Const":{"StrConst":"e"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"ack"}}},{"Operator":"CONTAIN_REGULAR","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"^J.*"}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"CONDITION: nest condition groups, judge with variable in. [all judgements above tests]",
			`{"username":"Andrew","usernames":["Jack","Mike","Andrew"],"news":"Jack hanged out with Mike last weekend.","number1":100,"numbers":["10",99.9,0]}`,
			`{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"&&","Conditions":[{"Operator":"STRING_START_WITH","Lhs":{"VarExpr":"username"},"Rhs":{"Const":{"StrConst":"And"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#1"},"Rhs":{"Const":{"StrConst":"e"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"ack"}}},{"Operator":"CONTAIN_REGULAR","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"^J.*"}}}]},{"OpLogic":"and","ConditionGroups":[{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"number1"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":100},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"numbers#0"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":999},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"LIST_CONTAINS","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":5}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":false},{"NumConst":-3.1415926}]}},{"Operator":"!LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}}]}]}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		}, {
			"CONDITION: [Pass] logic math operator by LrMath",
			`{"is_latest": true}`,
			`{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"==","Lhs":{"Const":{"NumConst":1.1232}},"Rhs":{"Const":{"NumConst":1.1232}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"!=","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":">","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"<","Lhs":{"Const":{"NumConst":-1}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":">=","Lhs":{"Const":{"NumConst":8}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"<=","Lhs":{"Const":{"NumConst":8}},"Rhs":{"Const":{"NumConst":8}}}}}]}`,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		}, {
			"CONDITION: [not Pass] logic math operator by LrMath",
			`{"is_latest": true}`,
			`{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"!=","Lhs":{"Const":{"NumConst":8}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"!=","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":">","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"<","Lhs":{"Const":{"NumConst":-1}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":">=","Lhs":{"Const":{"NumConst":8}},"Rhs":{"Const":{"NumConst":8}}}}},{"Operator":"==","Lhs":{"VarExpr":"is_latest"},"Rhs":{"MathExpr":{"OpMath":"<=","Lhs":{"Const":{"NumConst":8}},"Rhs":{"Const":{"NumConst":8}}}}}]}`,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		cdtCtx, err := parser.ParseArishemCondition(tc.expr)
		assert.Nil(t, err)
		dc, err := NewArishemDataCtx(context.Background(), tc.fact, &testFeatureFetcher{})
		assert.Nil(t, err)
		arisVst := NewArishemRuleVisitor(nil)
		tavo := &testArisVisitObserver{hash: tc.name}
		arisVst.AddVisitObserver(tavo)

		pass := arisVst.VisitCondition(cdtCtx, dc, newArisVisitTarget(tc.name))
		tc.check(pass, tavo.errMsgs)
	}
}

func TestAimVisit(t *testing.T) {
	testCases := []struct {
		name  string
		fact  string
		expr  string
		check func(aim typedef.Aim, errMsg []string)
	}{
		{
			"AIM: action aim param map with single no param feature",
			`{"user":{"name":"tadokoro kouji","age":24}}`,
			`{"ActionName":"Greeting","ParamMap":{"UserName":{"FeatureExpr":{"FeaturePath":"user_profile.username"}},"UserAge":{"VarExpr":"user.age"},"TempUsername":{"VarExpr":"user.name"},"UserBirthDay":{"FuncExpr":{"FuncName":"ListJoin","ParamMap":{"sep":{"Const":{"StrConst":""}},"list":{"ConstList":[{"NumConst":19},{"NumConst":19},{"NumConst":8},{"NumConst":10},{"NumConst":1},{"NumConst":14},{"NumConst":5},{"NumConst":14}]}}}}}}`,
			func(aim typedef.Aim, errMsg []string) {
				assert.Empty(t, errMsg)
				assert.NotNil(t, aim)
				action := aim.AsAction()
				assert.NotNil(t, action)
				assert.Equal(t, "Greeting", action.ActionName())
				paramMap := action.ParamAsMap()
				assert.NotNil(t, paramMap)
				// due to feature not implemented
				assert.Nil(t, paramMap["UserName"])
				assert.NotNil(t, paramMap["UserAge"])
				assert.NotNil(t, paramMap["TempUsername"])
				assert.NotNil(t, paramMap["UserBirthDay"])
				assert.Equal(t, int64(24), paramMap["UserAge"])
				assert.Equal(t, "tadokoro kouji", paramMap["TempUsername"])
				assert.Equal(t, "1919810114514", paramMap["UserBirthDay"])
			},
		},
		{
			"AIM: multi action aim, one with param map feature",
			`{}`,
			`[
    {
        "ActionName": "Greeting",
        "ParamMap": {
            "UserAge": {
                "VarExpr": "user.age"
            },
            "TempUsername": {
                "VarExpr": "user.name"
            }
        }
    },
    {
        "ActionName": "Greeting1",
        "ParamMap": {
            "UserAge": {
                "VarExpr": "user.age"
            },
            "TempUsername": {
                "VarExpr": "user.name"
            }
        }
    },
    {
        "ActionName": "Greeting2",
        "ParamMap": {
            "UserAge": {
                "VarExpr": "user.age"
            },
            "TempUsername": {
                "VarExpr": "user.name"
            }
        }
    }
]`,
			func(aim typedef.Aim, errMsg []string) {
				assert.Empty(t, errMsg)
				assert.NotNil(t, aim)
				actionList := aim.AsActionList()
				assert.NotNil(t, actionList)
				assert.Len(t, actionList, 3)

				for _, actionAim := range actionList {
					assert.NotNil(t, actionAim)
					assert.Contains(t, []string{"Greeting", "Greeting1", "Greeting2"}, actionAim.ActionName())
					paramMap := actionAim.ParamAsMap()
					assert.NotNil(t, paramMap)
					assert.Empty(t, paramMap["UserAge"])
					assert.Empty(t, paramMap["TempUsername"])
				}
			},
		},
		{
			"AIM: expr aim with MapExpr",
			`{}`,
			`{"MapExpr":{"UserName":{"VarExpr":"user.username"},"MessageText":{"Const":{"StrConst":"Hello World!\t\n"}},"UserAge":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}},"UserWeight":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":100}},{"Const":{"NumConst":20}}]}},"UserPersonality":{"FeatureExpr":{"FeaturePath":"user_center.user_personality","BuiltinParam":{"user_id":{"Const":{"NumConst":1234567}}}}}}}`,
			func(aim typedef.Aim, errMsg []string) {
				assert.Empty(t, errMsg)
				assert.NotNil(t, aim)
				assert.Nil(t, aim.AsAction())
				expr := aim.AsExpr()
				assert.NotNil(t, expr)
				mapExpr, ok := expr.(map[string]interface{})
				assert.True(t, ok)
				assert.NotEmpty(t, mapExpr)
				assert.Equal(t, float64(18), mapExpr["UserAge"])
				assert.Equal(t, float64(120), mapExpr["UserWeight"])
				assert.Equal(t, "Hello World!\\t\\n", mapExpr["MessageText"])
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		cdtCtx, err := parser.ParseArishemAim(tc.expr)
		assert.Nil(t, err)
		dc, err := NewArishemDataCtx(context.Background(), tc.fact, &testFeatureFetcher{})
		assert.Nil(t, err)
		arisVst := NewArishemRuleVisitor(nil)
		tavo := &testArisVisitObserver{hash: tc.name}
		arisVst.AddVisitObserver(tavo)

		aim := arisVst.VisitAim(cdtCtx, dc, newArisVisitTarget(tc.name))
		tc.check(aim, tavo.errMsgs)
	}
}

func TestSubCondition(t *testing.T) {
	testCases := []struct {
		name       string
		fact       string
		expr       string
		condFinder func(string) (antlr.ParseTree, error)
		check      func(pass bool, errMsg []string)
	}{
		{
			"SUB CONDITION: judge with INVALID left type",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"Const":{"NumConst":100}},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.NotEmpty(t, errMsg)
				t.Log(errMsg)
			},
		},
		{
			"SUB CONDITION: judge with INVALID right type",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"MapExpr":{"price":{"Const":{"NumConst":100}}}},"Rhs":{"Const":{"NumConst":100}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.NotEmpty(t, errMsg)
				t.Log(errMsg)
			},
		},
		{
			"SUB CONDITION: judge with INVALID condition finder",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"MapExpr":{"price":{"Const":{"NumConst":100}}}},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			nil,
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.NotEmpty(t, errMsg)
				t.Log(errMsg)
			},
		},
		{
			"SUB CONDITION: judge with VALID type@1",
			`{}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"MapExpr":{"price":{"Const":{"NumConst":100}}}},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"SUB CONDITION: judge with VALID type@2",
			`{"item_info":{"name":"xxx","price":100}}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"VarExpr":"item_info"},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		},
		{
			"SUB CONDITION: judge with VALID type@2",
			`{"item_info":{"name":"xxx","price":100}}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"SUB_COND","Lhs":{"VarExpr":"item_info"},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":"<=","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"==","Lhs":{"VarExpr":"name"},"Rhs":{"Const":{"StrConst":"xxx"}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		cdtCtx, err := parser.ParseArishemCondition(tc.expr)
		assert.Nil(t, err)
		dc, err := NewArishemDataCtx(context.Background(), tc.fact, &testFeatureFetcher{})
		assert.Nil(t, err)
		arisVst := NewArishemRuleVisitor(tc.condFinder)
		tavo := &testArisVisitObserver{hash: tc.name}
		arisVst.AddVisitObserver(tavo)

		pass := arisVst.VisitCondition(cdtCtx, dc, newArisVisitTarget(tc.name))
		tc.check(pass, tavo.errMsgs)
		t.Log("\n")
	}
}

func TestForeach(t *testing.T) {
	testCases := []struct {
		name       string
		fact       string
		expr       string
		condFinder func(string) (antlr.ParseTree, error)
		check      func(pass bool, errMsg []string)
	}{
		{
			"SUB CONDITION: judge with FOREACH VALID pre-defined operator: PASS",
			`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH >= and","Lhs":{"VarExpr":"item_info.item_list##price"},"Rhs":{"Const":{"NumConst":100}}}]}`,
			nil,
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		}, {
			"SUB CONDITION: judge with FOREACH VALID sub condition operator: PASS",
			`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND ||","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"STRING_CONTAINS","Lhs":{"VarExpr":"name"},"Rhs":{"Const":{"StrConst":"name@"}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.True(t, pass)
				assert.Empty(t, errMsg)
			},
		}, {
			"SUB CONDITION: judge with FOREACH VALID sub condition operator: NOT PASS",
			`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
			`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND &&","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"MyCond"}}}]}`,
			func(_ string) (antlr.ParseTree, error) {
				return parser.ParseArishemCondition(`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"STRING_CONTAINS","Lhs":{"VarExpr":"name"},"Rhs":{"Const":{"StrConst":"name@"}}}]}`)
			},
			func(pass bool, errMsg []string) {
				assert.False(t, pass)
				assert.Empty(t, errMsg)
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		cdtCtx, err := parser.ParseArishemCondition(tc.expr)
		assert.Nil(t, err)
		dc, err := NewArishemDataCtx(context.Background(), tc.fact, &testFeatureFetcher{})
		assert.Nil(t, err)
		arisVst := NewArishemRuleVisitor(tc.condFinder)
		tavo := &testArisVisitObserver{hash: tc.name}
		arisVst.AddVisitObserver(tavo)

		pass := arisVst.VisitCondition(cdtCtx, dc, newArisVisitTarget(tc.name))
		tc.check(pass, tavo.errMsgs)
		t.Log("\n")
	}
}
