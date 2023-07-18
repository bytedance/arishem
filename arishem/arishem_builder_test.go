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
	"github.com/bytedance/arishem/internal/operator"
	"github.com/bytedance/arishem/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildCondition(t *testing.T) {
	condGroup := NewConditionsCondGroup(OpLogicAnd)
	cond1 := NewCondition(operator.Equal)
	cond1.Lhs = NewConstExpr(NewNumConst(1.0))
	cond1.Rhs = NewConstExpr(NewNumConst(1.0))
	condGroup.AddConditions(cond1)
	expr, err := condGroup.Build()
	assert.Nil(t, err)
	assert.NotEmpty(t, expr)
	assert.Equal(t, `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}}]}`, expr)

	cond2 := NewCondition(operator.ListIn, NOT)
	var varPath VarExpr = "user.user_list#2.name"
	cond2.Lhs = NewVarExpr(&varPath)
	cond2.Rhs = NewConstListExpr([]*Const{
		NewStrConst("Jack"),
		NewStrConst("Jane"),
		NewStrConst("John"),
		NewStrConst("Ezreal"),
	})
	condGroup.AddConditions(cond2)
	expr, err = condGroup.Build()
	assert.Nil(t, err)
	assert.NotEmpty(t, expr)
	assert.Equal(t, `{"OpLogic":"&&","Conditions":[{"Operator":"==","Lhs":{"Const":{"NumConst":1}},"Rhs":{"Const":{"NumConst":1}}},{"Operator":"!LIST_IN","Lhs":{"VarExpr":"user.user_list#2.name"},"Rhs":{"ConstList":[{"StrConst":"Jack"},{"StrConst":"Jane"},{"StrConst":"John"},{"StrConst":"Ezreal"}]}}]}`, expr)

	pass, err := JudgeConditionWithFactMeta(context.Background(), expr, `{"user":{"user_list":[{"name":"Aatrox"},{"name":"Ahri"},{"name":"Ezreal"},{"name":"MalPhite"}]}}`)
	assert.Nil(t, err)
	assert.False(t, pass)
}

func TestBuildAction(t *testing.T) {
	aimStr, err := NewParamListAction("MyAction", []*Expr{NewConstExpr(NewNumConst(1.23))}).Build()
	assert.Nil(t, err)
	t.Log(aimStr)
	aim, err := parser.ParseArishemAim(aimStr)
	assert.Nil(t, err)
	assert.NotNil(t, aim)
	aimStr, err = NewParamMapAction("MyAction", map[string]*Expr{"rate": NewConstExpr(NewNumConst(1.23))}).Build()
	assert.Nil(t, err)
	t.Log(aimStr)
	aim, err = parser.ParseArishemAim(aimStr)
	assert.Nil(t, err)
	assert.NotNil(t, aim)
}
