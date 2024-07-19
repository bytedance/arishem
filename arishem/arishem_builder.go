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

import "github.com/bytedance/arishem/internal/operator"

type Not string

const NOT Not = "!"

const (
	IsNull                    = operator.IsNull
	Equal                     = operator.Equal
	NotEqual                  = operator.NotEqual
	Greater                   = operator.Greater
	Less                      = operator.Less
	GreaterOrEqual            = operator.GreaterOrEqual
	LessOrEqual               = operator.LessOrEqual
	StringStartWith           = operator.StringStartWith
	StringEndWith             = operator.StringEndWith
	StringContains            = operator.StringContains
	ContainRegexp             = operator.ContainRegexp
	ListIn                    = operator.ListIn
	ListContains              = operator.ListContains
	ListRetain                = operator.ListRetain
	SubListIn                 = operator.SubListIn
	SubListContains           = operator.SubListContains
	ListVagueContains         = operator.ListVagueContains
	BetweenAllClose           = operator.BetweenAllClose
	BetweenAllOpen            = operator.BetweenAllOpen
	BetweenLeftOpenRightClose = operator.BetweenLeftOpenRightClose
	BetweenLeftCloseRightOpen = operator.BetweenLeftCloseRightOpen
)

const (
	OpLogicAnd = "&&"
	OpLogicOr  = "||"
	OpLogicNot = "!"
)

type condGroupType int
type actionAimType int

const (
	condGroupTypeCondGroups condGroupType = iota
	condGroupTypeConditions
)

const (
	actionAimTypeParamMap actionAimType = iota
	actionAimTypeParamList
)

type Builder interface {
	Build() (string, error)
}
