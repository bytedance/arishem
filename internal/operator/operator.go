/*
 * Copyright 2023 ByteDance Inc.
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

package operator

import (
	"fmt"
)

const (
	// IsNull object operators
	IsNull = "IS_NULL"
	// Equal math operators
	Equal          = "=="
	NotEqual       = "!="
	Greater        = ">"
	Less           = "<"
	GreaterOrEqual = ">="
	LessOrEqual    = "<="
	// StringStartWith string operators
	StringStartWith = "STRING_START_WITH"
	StringEndWith   = "STRING_END_WITH"
	StringContains  = "STRING_CONTAINS"
	ContainRegexp   = "CONTAIN_REGULAR"
	// ListIn list operators
	ListIn                    = "LIST_IN"
	ListContains              = "LIST_CONTAINS"
	ListRetain                = "LIST_RETAIN"
	SubListIn                 = "SUB_LIST_IN"
	SubListContains           = "SUB_LIST_CONTAINS"
	ListVagueContains         = "CONTAINS_KEYWORDS"
	BetweenAllClose           = "BETWEEN_ALL_CLOSE"
	BetweenAllOpen            = "BETWEEN_ALL_OPEN"
	BetweenLeftOpenRightClose = "BETWEEN_LEFT_OPEN_RIGHT_CLOSE"
	BetweenLeftCloseRightOpen = "BETWEEN_LEFT_CLOSE_RIGHT_OPEN"
)

type Operator interface {
	Operate(left, right interface{}) (bool, error)
}

type operatorPair struct {
	// oprKey is operator key
	oprKey string
	// posOpr is the positive operator
	posOpr Operator
	// negOpr is the negative operator
	negOpr Operator
}

var (
	operators map[string]Operator
	oprPairs  = []*operatorPair{
		{
			oprKey: IsNull,
			posOpr: &NullOperator{},
			negOpr: &NotNullOperator{},
		},
		{
			oprKey: StringStartWith,
			posOpr: &StringOperator{opr: StringStartWithOperate},
			negOpr: &StringOperator{opr: StringNotStartWithOperate},
		},
		{
			oprKey: StringEndWith,
			posOpr: &StringOperator{opr: StringEndWithOperate},
			negOpr: &StringOperator{opr: StringNotEndWithOperate},
		},
		{
			oprKey: StringContains,
			posOpr: &StringOperator{opr: StringContainsOperate},
			negOpr: &StringOperator{opr: StringNotContainsOperate},
		},
		{
			oprKey: ContainRegexp,
			posOpr: &StringOperator{opr: ContainRegexpOperate},
			negOpr: &StringOperator{opr: NotContainRegexpOperate},
		},
		{
			oprKey: ListIn,
			posOpr: &ListOperator{opr: ListInOperate},
			negOpr: &ListOperator{opr: ListNotInOperate},
		},
		{
			oprKey: ListContains,
			posOpr: &ListOperator{opr: ListContainsOperate},
			negOpr: &ListOperator{opr: ListNotContainsOperate},
		},
		{
			oprKey: ListRetain,
			posOpr: &ListOperator{opr: ListRetainOperate},
			negOpr: &ListOperator{opr: ListNotRetainOperate},
		},
		{
			oprKey: SubListIn,
			posOpr: &ListOperator{opr: SubListInOperate},
			negOpr: &ListOperator{opr: SubListNotInOperate},
		},
		{
			oprKey: SubListContains,
			posOpr: &ListOperator{opr: SubListContainsOperate},
			negOpr: &ListOperator{opr: SubListNotContainsOperate},
		},
		{
			oprKey: ListVagueContains,
			posOpr: &ListOperator{opr: ListVagueContainsOperate},
			negOpr: &ListOperator{opr: ListNotVagueContainsOperate},
		},
		{
			oprKey: BetweenAllClose,
			posOpr: &ListOperator{opr: BetweenAllCloseOperate},
			negOpr: &ListOperator{opr: NotBetweenAllCloseOperate},
		},
		{
			oprKey: BetweenAllOpen,
			posOpr: &ListOperator{opr: BetweenAllOpenOperate},
			negOpr: &ListOperator{opr: NotBetweenAllOpenOperate},
		},
		{
			oprKey: BetweenLeftOpenRightClose,
			posOpr: &ListOperator{opr: BetweenLeftOpenRightCloseOperate},
			negOpr: &ListOperator{opr: NotBetweenLeftOpenRightCloseOperate},
		},
		{
			oprKey: BetweenLeftCloseRightOpen,
			posOpr: &ListOperator{opr: BetweenLeftCloseRightOpenOperate},
			negOpr: &ListOperator{opr: NotBetweenLeftCloseRightOpenOperate},
		},
	}
	oprMath = map[string]Operator{
		Equal:          &CompareOperator{opr: EqualOperate},
		NotEqual:       &CompareOperator{opr: NotEqualOperate},
		Greater:        &CompareOperator{opr: GreaterOperate},
		Less:           &CompareOperator{opr: LessOperate},
		GreaterOrEqual: &CompareOperator{opr: GreaterOrEqualOperate},
		LessOrEqual:    &CompareOperator{opr: LessOrEqualOperate},
	}
)

// Judge left and right with operator find by operator
func Judge(left, right interface{}, opr string) (bool, error) {
	if _, ok := operators[opr]; !ok {
		// oprKey exist check
		return false, fmt.Errorf("oprKey=>%s not found", opr)
	}
	return operators[opr].Operate(left, right)
}

func init() {
	operators = make(map[string]Operator, 3*len(oprPairs)+len(oprMath))
	for oprKey, opr := range oprMath {
		operators[oprKey] = opr
	}
	for _, pair := range oprPairs {
		operators[pair.oprKey] = pair.posOpr
		// generator and register negative operator
		operators["NOT "+pair.oprKey] = pair.negOpr
		operators["!"+pair.oprKey] = pair.negOpr
	}
}
