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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/typedef"
	"testing"
)

var (
	arishemBenchRules       []core.RuleTarget
	arishemBenchComplexRule core.RuleTarget
	arishemBenchComplexDc   typedef.DataCtx
	arishemBenchDc          typedef.DataCtx
)

func init() {
	Initialize(DefaultConfiguration())
	arishemBenchRules = buildArishemRules()
	arishemBenchComplexRule = buildArishemComplexRule()
	arishemBenchDc, _ = core.NewArishemDataCtx(`{}`, FeatureFetcherFactory())
	arishemBenchComplexDc, _ = core.NewArishemDataCtx(`{"username":"Andrew","usernames":["Jack","Mike","Andrew"],"news":"Jack hanged out with Mike last weekend.","number1":100,"numbers":["10",99.9,0]}`, FeatureFetcherFactory())
}

type nodeCounter struct{ count int }

func (n *nodeCounter) VisitTerminal(node antlr.TerminalNode, tokenErrMsgs, parseErrMsgs []string) {}

func (n *nodeCounter) VisitErrorNode(node antlr.ErrorNode, tokenErrMsgs, parseErrMsgs []string) {}

func (n *nodeCounter) EnterEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string) {
}

func (n *nodeCounter) ExitEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string) {
	n.count++
}

func buildArishemComplexRule() core.RuleTarget {
	condCounter := &nodeCounter{}
	aimCounter := &nodeCounter{}
	condition, err := parser.ParseArishemCondition(`{"OpLogic":"||","ConditionGroups":[{"OpLogic":"&&","Conditions":[{"Operator":"STRING_START_WITH","Lhs":{"VarExpr":"username"},"Rhs":{"Const":{"StrConst":"Banana"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#1"},"Rhs":{"Const":{"StrConst":"A"}}},{"Operator":"STRING_END_WITH","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"hahaha"}}},{"Operator":"CONTAIN_REGULAR","Lhs":{"VarExpr":"usernames#0"},"Rhs":{"Const":{"StrConst":"^M.*"}}}]},{"OpLogic":"and","ConditionGroups":[{"OpLogic":"&&","ConditionGroups":[{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"VarExpr":"number1"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":98},{"NumConst":101},{"NumConst":1.32e-3},{"NumConst":12.234},{"NumConst":-1}]}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"numbers#0"},"Rhs":{"ConstList":[{"NumConst":1},{"NumConst":99},{"NumConst":999},{"NumConst":1.32e-3},{"NumConst":10},{"NumConst":-1}]}},{"Operator":"LIST_CONTAINS","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":6}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":6}},{"Const":{"StrConst":"5"}}]}}},{"Operator":"LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}},{"Operator":"!LIST_RETAIN","Lhs":{"VarExpr":"numbers"},"Rhs":{"ConstList":[{"StrConst":"-1"},{"BoolConst":true},{"NumConst":-3.1415926}]}}]}]},{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}}]}]}`, condCounter)
	if err != nil {
		panic(err)
	}
	aim, err := parser.ParseArishemAim(`
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
`, aimCounter)
	if err != nil {
		panic(err)
	}
	println(condCounter.count + aimCounter.count)
	return &NoPriorityRule{name: "complex rule", cond: condition, aim: aim}
}

func buildArishemRules() []core.RuleTarget {
	var rt []core.RuleTarget
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
                    "NumConst": 9.999999
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
	return rt
}

func BenchmarkBuildArishemRules(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildArishemRules()
	}
}

func BenchmarkExecuteArishemRule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecuteRules(arishemBenchRules, arishemBenchDc)
	}
}

func BenchmarkSingleComplexRule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecuteRules([]core.RuleTarget{arishemBenchComplexRule}, arishemBenchComplexDc)
	}
}
