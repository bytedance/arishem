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

package core

import (
	"github.com/bytedance/arishem/typedef"
)

type arishemCondition struct {
	passed    bool
	left      interface{}
	leftExpr  string
	right     interface{}
	rightExpr string
	operator  string
	error     error
}

func newArishemCondition(
	passed bool,
	left interface{},
	leftExpr string,
	right interface{},
	rightExpr string,
	operator string,
	error error,
) typedef.JudgeNode {
	return &arishemCondition{
		passed:    passed,
		left:      left,
		leftExpr:  leftExpr,
		right:     right,
		rightExpr: rightExpr,
		operator:  operator,
		error:     error,
	}
}

func (a *arishemCondition) Passed() bool {
	return a.passed
}

func (a *arishemCondition) Left() interface{} {
	return a.left
}

func (a *arishemCondition) LeftExpr() string {
	return a.leftExpr
}

func (a *arishemCondition) Right() interface{} {
	return a.right
}

func (a *arishemCondition) RightExpr() string {
	return a.rightExpr
}

func (a *arishemCondition) Operator() string {
	return a.operator
}

func (a *arishemCondition) Error() error {
	return a.error
}

type arishemForeachItem struct {
	Index   int
	ItemVal interface{}
	Pass    bool
}

func (a *arishemForeachItem) Idx() int {
	return a.Index
}

func (a *arishemForeachItem) ItemValue() interface{} {
	return a.ItemVal
}

func (a *arishemForeachItem) Passed() bool {
	return a.Pass
}

func newArishemForeachItem(idx int, itemVal interface{}, pass bool) typedef.ForeachItem {
	return &arishemForeachItem{
		Index:   idx,
		ItemVal: itemVal,
		Pass:    pass,
	}
}

type arishemForeachCondition struct {
	*arishemCondition
	foreachOperator string
	foreachLogic    string
	leftItems       []typedef.ForeachItem
}

func newArishemForeachCondition(
	passed bool,
	left interface{},
	leftExpr string,
	right interface{},
	rightExpr string,
	operator string,
	error error,
	foreachOperator string,
	foreachLogic string,
	foreachItems []typedef.ForeachItem,
) typedef.JudgeNode {
	return &arishemForeachCondition{
		arishemCondition: &arishemCondition{
			passed:    passed,
			left:      left,
			leftExpr:  leftExpr,
			right:     right,
			rightExpr: rightExpr,
			operator:  operator,
			error:     error,
		},
		foreachOperator: foreachOperator,
		foreachLogic:    foreachLogic,
		leftItems:       foreachItems,
	}
}

func (a *arishemForeachCondition) ForeachLogic() string {
	return a.foreachLogic
}

func (a *arishemForeachCondition) ForeachOperator() string {
	return a.foreachOperator
}

func (a *arishemForeachCondition) ForeachItems() []typedef.ForeachItem {
	return a.leftItems
}
