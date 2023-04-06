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

package core

import (
	"github.com/bytedance/arishem/internal/parser"
	"testing"
)

func BenchmarkConditionVisitNoFact(b *testing.B) {
	cdtExpr := `{"OpLogic":"||","Conditions":[{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":"<=","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}},{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`
	cdtCtx, _ := parser.ParseArishemCondition(cdtExpr, nil)
	aruleV := NewArishemRuleVisitor()
	dc, _ := NewArishemDataCtx("{}", &testFeatureFetcher{})
	for i := 0; i < b.N; i++ {
		aruleV.VisitCondition(cdtCtx, dc, nil)
	}
}

func BenchmarkConditionVisitFactVar(b *testing.B) {
	cdtExpr := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"number1"},"Rhs":{"VarExpr":"numbers#0"}},{"Operator":"<=","Lhs":{"VarExpr":"numbers#1"},"Rhs":{"VarExpr":"numbers#0"}},{"Operator":">","Lhs":{"VarExpr":"numbers#2"},"Rhs":{"VarExpr":"numbers#0"}}]}`
	cdtCtx, _ := parser.ParseArishemCondition(cdtExpr, nil)
	aruleV := NewArishemRuleVisitor()
	dc, _ := NewArishemDataCtx(`{"number1":100,"numbers":["10",1,99.9]}`, &testFeatureFetcher{})
	for i := 0; i < b.N; i++ {
		aruleV.VisitCondition(cdtCtx, dc, nil)
	}
}
