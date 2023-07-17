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

package typedef

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// VisitTarget can be identified.
type VisitTarget Identifiable

// JudgeNode is the node of every single condition.
type JudgeNode interface {
	Passed() bool
	Left() interface{}
	LeftExpr() string
	Right() interface{}
	RightExpr() string
	Operator() string
	Error() error
}

// ActionAim is the interface that defines the aim of action.
type ActionAim interface {
	ActionName() string
	Parameter() interface{}
	ParamAsMap() (paramMap map[string]interface{})
	ParamAsList() (paramList []interface{})
}

// Aim is the interface that supports type convert.
type Aim interface {
	AsExpr() interface{}
	AsActionList() []ActionAim
	AsAction() ActionAim
}

// RuleVisitor is the interface that defines rule visit behaviour.
type RuleVisitor interface {
	ObservableVisitor

	// VisitCondition will visit condition context and return condition whether passed.
	VisitCondition(cdtCtx antlr.ParseTree, dCtx DataCtx, vt VisitTarget) (pass bool)
	// VisitAim will visit aim context and return aim value.
	VisitAim(aimCtx antlr.ParseTree, dCtx DataCtx, vt VisitTarget) (aim Aim)
}
