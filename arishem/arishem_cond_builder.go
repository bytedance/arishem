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

import "github.com/bytedance/sonic"

// NewConditionsCondGroup create a new condition group with conditions.
func NewConditionsCondGroup(opLogic string) *ConditionGroupBuilder {
	return &ConditionGroupBuilder{OpLogic: opLogic, cgType: condGroupTypeConditions}
}

// NewCondGroupsCondGroup create a new condition group with condition groups.
func NewCondGroupsCondGroup(opLogic string) *ConditionGroupBuilder {
	return &ConditionGroupBuilder{OpLogic: opLogic, cgType: condGroupTypeCondGroups}
}

type ConditionGroupBuilder struct {
	cgType          condGroupType
	OpLogic         string                   `json:"OpLogic"`
	ConditionGroups []*ConditionGroupBuilder `json:"ConditionGroups,omitempty"`
	Conditions      []*Condition             `json:"Conditions,omitempty"`
}

func (c *ConditionGroupBuilder) Build() (expr string, err error) {
	return sonic.MarshalString(c)
}

type Condition struct {
	Operator string `json:"Operator"`
	Lhs      *Expr  `json:"Lhs,omitempty"`
	Rhs      *Expr  `json:"Rhs,omitempty"`
}

func NewCondition(operator string, not ...Not) *Condition {
	if len(not) > 0 {
		operator = "!" + operator
	}
	return &Condition{Operator: operator}
}

func (c *ConditionGroupBuilder) AddConditionGroups(condGroups ...*ConditionGroupBuilder) {
	if len(condGroups) <= 0 {
		return
	}
	if c.cgType != condGroupTypeCondGroups {
		return
	}
	c.ConditionGroups = append(c.ConditionGroups, condGroups...)
	return
}

func (c *ConditionGroupBuilder) AddConditions(cond ...*Condition) {
	if len(cond) <= 0 {
		return
	}
	if c.cgType != condGroupTypeConditions {
		return
	}
	c.Conditions = append(c.Conditions, cond...)
	return
}
