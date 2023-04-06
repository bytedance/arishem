/*
 *  Copyright 2023 ByteDance Inc.
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
	"github.com/bytedance/arishem/internal/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecuteSingleRule(t *testing.T) {
	testCases := []struct {
		name     string
		rule     func() core.RuleTarget
		factMeta string
		check    func(rrs core.RuleResult)
	}{
		{
			"rule3",
			func() (rt core.RuleTarget) {
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
			func(rr core.RuleResult) {
				assert.NotNil(t, rr)
				assert.False(t, rr.Passed())
				assert.Nil(t, rr.Aim())
			},
		},
		{
			"rule5",
			func() (rt core.RuleTarget) {
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
			func(rr core.RuleResult) {
				assert.NotNil(t, rr)
				assert.True(t, rr.Passed())
				assert.NotNil(t, rr.Aim())
				assert.Equal(t, "p-rule5 passed!", rr.Aim().AsExpr())
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		dc, err := core.NewArishemDataCtx(tc.factMeta, FeatureFetcherFactory())
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
		rules    func() []core.RuleTarget
		factMeta string
		check    func(rrs []core.RuleResult)
	}{
		{
			"multi rules with priority",
			func() (rt []core.RuleTarget) {
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
			func(rrs []core.RuleResult) {
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
		dc, err := core.NewArishemDataCtx(tc.factMeta, FeatureFetcherFactory())
		if err != nil {
			t.Log(err)
		} else {
			tc.check(ExecuteRules(tc.rules(), dc))
		}
	}
}
