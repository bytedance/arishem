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

package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseArishemConditionContext(t *testing.T) {
	testCases := []struct {
		name  string
		args  string
		check func(ctx ICondEntityContext, err error)
	}{
		{
			"INVALID: EMPTY string",
			``,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"INVALID: EMPTY JSON string",
			`{}`,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"VALID: null JSON string",
			`null`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE condition",
			`{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"BoolConst":true}}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE condition: NOT compressed",
			`{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "VarExpr": "upgrade_time"
            },
            "Rhs": {
                "FeatureExpr": {
                    "FeaturePath": "user.username"
                }
            }
        }
    ]
}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"INVALID: wrong value type in BuiltinParam",
			`{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]}}},"Rhs":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"VALID: SINGLE condition: compressed",
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"FeatureExpr":{"FeaturePath":"user.username"}}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE condition: LIST_IN operator",
			`{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ConstList":[{"NumConst":100},{"NumConst":100},{"NumConst":100}]}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE condition: NOT LIST_IN operator",
			`{"OpLogic":"&&","Conditions":[{"Operator":"NOT LIST_IN","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ConstList":[{"NumConst":100},{"NumConst":100},{"NumConst":100}]}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE condition: special characters",
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ConstList":[{"StrConst":"测试中文"},{"StrConst":"测试特殊字符          ～！@#¥%……&*（）-=——+【】、「」｜；’：”，。/《》？"},{"StrConst":"Test English"},{"StrConst":"Test Special Characters\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/"},{"StrConst":"Test Special Alphabet œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"}]}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: null condition group",
			`{"OpLogic":"AND","Conditions":[null,null]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: null condition",
			`{"OpLogic":"AND","ConditionGroups":[null,null]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: MULTI conditions: multi expr type",
			`{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"NumConst":123421342}}},{"Operator":">","Lhs":{"VarExpr":"current_time"},"Rhs":{"Const":{"BoolConst":false}}},{"Operator":">","Lhs":{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"UserID":{"VarExpr":"user_id"}}}},"Rhs":{"Const":{"StrConst":"hello!! Mike\n"}}},{"Operator":"LIST_IN","Lhs":{"VarExpr":"last_time"},"Rhs":{"ConstList":[{"NumConst":1314520},{"BoolConst":true},{"StrConst":"Programming!\n"},{"StrConst":"With special Char {}[]';<>?!~@#$%^&*()-+.,/\\ \t\n\r\"'"}]}},{"Operator":"LIST_IN","Lhs":{"VarExpr":"last_time"},"Rhs":{"MapExpr":{"UserName":{"Const":{"StrConst":"Micheal"}},"Age":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}},"Gener":{"FuncExpr":{"FuncName":"GetGender"}},"Personalities":{"ListExpr":[{"Const":{"BoolConst":true}},{"MapExpr":{"WhenMonday":{"Const":{"StrConst":"Happy!"}}}}]}}}}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: NEST condition groups",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":-1}},"Rhs":{"Const":{"NumConst":1}}},{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"StrConst":"1"}}},{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"BoolConst":true}}}]}]}]}]}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"INVALID: NEST condition groups, but no condition",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[]}]}]}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"VALID: NEST condition groups, with null condition",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[null]}]}]}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: NEST condition groups with conditions at same level@1",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"NumConst":100}}}]}]}]}]}]},{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"NumConst":100}}}]},{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"BoolConst":true}}}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: NEST condition groups with conditions at same level@2",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"NumConst":100}}}]},{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"BoolConst":true}}}]},{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"Const":{"NumConst":100}}}]}]}]}]}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"INVALID: NODES with WRONG ORDER",
			`{"ConditionGroups":[{"Conditions":[{"Rhs":{"Const":{"BoolConst":true}},"Lhs":{"VarExpr":"upgrade_time"},"Operator":"=="}],"OpLogic":"AND"}],"OpLogic":"AND"}`,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"INVALID: feature type only contains one name",
			`{"OpLogic":"AND","ConditionGroups":[{"OpLogic":"AND","Conditions":[{"Operator":"==","Lhs":{"FeatureExpr":{"FeaturePath":"upgrade_time"}},"Rhs":{"Const":{"BoolConst":true}}}]}]}`,
			func(ctx ICondEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.name)
		testCase.check(ParseArishemCondition(testCase.args))
	}
}

func TestParseArishemAimContext(t *testing.T) {
	testCases := []struct {
		name  string
		args  string
		check func(ctx IAimEntityContext, err error)
	}{
		{
			"INVALID: EMPTY string",
			``,
			func(ctx IAimEntityContext, err error) {
				assert.Nil(t, ctx)
				assert.NotNil(t, err)
				t.Log(err.Error())
			},
		},
		{
			"VALID: EMPTY JSON string",
			`{}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: null JSON string",
			`null`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: Const expression",
			`{"Const":{"StrConst":"{\"passed\":true}"}}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: MapExpr",
			`{"MapExpr":{"UserName":{"VarExpr":"user.username"},"UserAge":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}},"UserWeight":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":100}},{"Const":{"NumConst":20}}]}},"UserPersonality":{"FeatureExpr":{"FeaturePath":"user_center.user_personality","BuiltinParam":{"user_id":{"VarExpr":"user.id"}}}}}}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: ActionAim with a single name",
			`{"ActionName":"Greeting"}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: ActionAim with a single name",
			`{
    "ActionName": "Greeting"
}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: ActionAim with MapParam",
			`{"ActionName":"Greeting","ParamMap":{"UserName":{"FeatureExpr":{"FeaturePath":"user_profile.username"}}}}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: ActionAim with MapParam",
			`{
    "ActionName": "Greeting",
    "ParamMap": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		}, {
			"VALID: ActionAim with ListParam",
			`{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]}`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: EMPTY ActionAim",
			`[]`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
		{
			"VALID: SINGLE ActionAim",
			`[{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]}]`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		}, {
			"VALID: MULTI ActionAim",
			`[{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]},{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]}]`,
			func(ctx IAimEntityContext, err error) {
				assert.NotNil(t, ctx)
				assert.Nil(t, err)
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.name)
		testCase.check(ParseArishemAim(testCase.args))
	}
}
