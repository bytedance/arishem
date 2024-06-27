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
	"context"
	"github.com/bytedance/arishem/internal/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecuteSingleRule(t *testing.T) {
	testCases := []struct {
		name     string
		rule     func() RuleTarget
		factMeta string
		check    func(rrs RuleResult)
	}{
		{
			"rule3",
			func() (rt RuleTarget) {
				rt, _ = NewPriorityRule("p-rule3", 3, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "NumConst": 99.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule3 passed!"
    }
}`)
				return
			},
			`{}`,
			func(rr RuleResult) {
				assert.NotNil(t, rr)
				assert.False(t, rr.Passed())
				assert.Nil(t, rr.Aim())
			},
		},
		{
			"rule5",
			func() (rt RuleTarget) {
				rt, _ = NewPriorityRule("p-rule5", 5, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "Const": {
                    "NumConst": 199.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule5 passed!"
    }
}`)
				return
			},
			`{}`,
			func(rr RuleResult) {
				assert.NotNil(t, rr)
				assert.True(t, rr.Passed())
				assert.NotNil(t, rr.Aim())
				assert.Equal(t, "p-rule5 passed!", rr.Aim().AsExpr())
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		dc, err := core.NewArishemDataCtx(context.Background(), tc.factMeta, FeatureFetcherFactory())
		if err != nil {
			t.Log(err)
		} else {
			tc.check(ExecuteSingleRule(tc.rule(), dc))
		}
	}
}

func TestExecuteRules(t *testing.T) {
	testCases := []struct {
		name     string
		rules    func() []RuleTarget
		factMeta string
		check    func(rrs []RuleResult)
	}{
		{
			"multi rules with priority",
			func() (rt []RuleTarget) {
				pr, err := NewPriorityRule("p-rule3", 3, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "NumConst": 99.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule3 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, pr)
				}
				pr, err = NewPriorityRule("p-rule5", 5, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "Const": {
                    "NumConst": 99.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule5 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, pr)
				}
				pr, err = NewPriorityRule("p-rule4", 4, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "Const": {
                    "NumConst": 99.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule4 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, pr)
				}
				pr, err = NewPriorityRule("p-rule2", 2, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "LIST_IN",
            "Lhs": {
                "Const": {
                    "NumConst": 99.999999
                }
            },
            "Rhs": {
                "ConstList": [
                    {
                        "NumConst": 10
                    },
                    {
                        "NumConst": 98
                    },
                    {
                        "NumConst": 101
                    },
                    {
                        "NumConst": -1.12e3
                    }
                ]
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule2 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, pr)
				}
				pr, err = NewPriorityRule("p-rule1", 1, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "Const": {
                    "NumConst": 199.999999
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 100
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule1 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, pr)
				}
				npr, err := NewNoPriorityRule("np-rule1", `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "BoolConst": true
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 1
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "np-rule1 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, npr)
				}
				npr, err = NewNoPriorityRule("np-rule2", `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "BoolConst": true
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 1
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "np-rule2 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, npr)
				}
				npr, err = NewNoPriorityRule("np-rule3", `
{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "Const": {
                    "BoolConst": true
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 1
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "np-rule3 passed!"
    }
}`)
				if err == nil {
					rt = append(rt, npr)
				}
				return
			},
			"{}",
			func(rrs []RuleResult) {
				assert.Len(t, rrs, 4)
				for _, rr := range rrs {
					assert.Contains(t, []string{"p-rule1", "np-rule1", "np-rule2", "np-rule3"}, rr.Identifier())
					assert.Contains(t, []string{"p-rule1 passed!", "np-rule1 passed!", "np-rule2 passed!", "np-rule3 passed!"}, rr.Aim().AsExpr().(string))
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		dc, err := core.NewArishemDataCtx(context.Background(), tc.factMeta, FeatureFetcherFactory())
		if err != nil {
			t.Log(err)
		} else {
			tc.check(ExecuteRules(tc.rules(), dc))
		}
	}
}

func TestForeachWithSubCondition(t *testing.T) {
	// case1-1 test OnlyPriceMatch, OnlyPriceMatch added when initialize the arishem
	pass, err := JudgeConditionWithFactMeta(
		context.Background(),
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND and","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"OnlyPriceMatch"}}}]}`,
		`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
	)
	assert.Nil(t, err)
	assert.True(t, pass)
	// case1-2 use rule execution while dependent on the sub condition
	rule, err := NewNoPriorityRule(
		"depend-OnlyPriceMatch",
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND and","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"OnlyPriceMatch"}}}]}`,
		`{
    "Const": {
        "StrConst": "depend-OnlyPriceMatch"
    }
}`)
	assert.Nil(t, err)
	dc, err := DataContext(context.Background(), `{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`)
	assert.Nil(t, err)
	//rr := ExecuteSingleRule(rule, dc, WithVisitObserver(NewMyObserver("obs")))
	rr := ExecuteSingleRule(rule, dc)
	assert.NotNil(t, rr)
	assert.True(t, rr.Passed())

	// case2 test OnlyNameMatch, OnlyNameMatch added by create a rule
	rule, err = NewNoPriorityRule("OnlyNameMatch", `{"OpLogic":"&&","Conditions":[{"Operator":"STRING_CONTAINS","Lhs":{"VarExpr":"name"},"Rhs":{"Const":{"StrConst":"name@"}}}]}`, `{
		"Const": {
			"StrConst": "OnlyNameMatch"
		}
	}`)
	assert.Nil(t, err)
	pass, err = JudgeConditionWithFactMeta(
		context.Background(),
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND and","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"OnlyNameMatch"}}}]}`,
		`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
	)
	assert.Nil(t, err)
	assert.True(t, pass)
	// case3 test NameAndPriceMatch, NameAndPriceMatch added by AddSubCondition
	err = AddSubCondition("NameAndPriceMatch", `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"STRING_CONTAINS","Lhs":{"VarExpr":"name"},"Rhs":{"Const":{"StrConst":"name@"}}}]}`)
	assert.Nil(t, err)
	pass, err = JudgeConditionWithFactMeta(
		context.Background(),
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH SUB_COND and","Lhs":{"VarExpr":"item_info.item_list"},"Rhs":{"SubCondExpr":{"CondName":"NameAndPriceMatch"}}}]}`,
		`{"item_info":{"item_list":[{"name":"name@1","price":100},{"name":"name@2","price":102.13},{"name":"name@3","price":200},{"name":"name@4","price":100},{"name":"name@5","price":101},{"name":"name@6","price":303.1234}]}}`,
	)
	assert.Nil(t, err)
	assert.True(t, pass)
	// case4 normal foreach test
	pass, err = JudgeConditionWithFactMeta(
		context.Background(),
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH < or","Lhs":{"VarExpr":"item_info.price_list"},"Rhs":{"Const":{"NumConst":-1}}}]}`,
		`{"item_info":{"price_list":[100,102.13,200,100,101,303.1234]}}`,
	)
	assert.Nil(t, err)
	assert.False(t, pass)

	pass, err = JudgeConditionWithFactMeta(
		context.Background(),
		`{"OpLogic":"&&","Conditions":[{"Operator":"FOREACH > or","Lhs":{"VarExpr":"item_info.price_list"},"Rhs":{"Const":{"NumConst":-1}}}]}`,
		`{"item_info":{"price_list":[100,102.13,200,100,101,303.1234]}}`,
	)
	assert.Nil(t, err)
	assert.True(t, pass)
}
