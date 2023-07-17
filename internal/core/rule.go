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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/typedef"
)

type RuleTarget interface {
	typedef.Comparable
	typedef.VisitTarget

	ConditionPTree() antlr.ParseTree
	AimPTree() antlr.ParseTree

	CondFeatParams() []typedef.FeatureParam
	AimFeatParams() []typedef.FeatureParam
}

type RuleResult interface {
	typedef.VisitTarget

	Passed() bool
	Aim() typedef.Aim
}
