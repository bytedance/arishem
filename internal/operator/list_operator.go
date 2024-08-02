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

package operator

import (
	"errors"
	"github.com/bytedance/arishem/tools"
	"strings"
)

type ListOperator struct {
	opr func(left, right interface{}) (bool, error)
}

var equalCmp = CompareOperator{opr: EqualOperate}
var greaterCmp = CompareOperator{opr: GreaterOperate}
var greaterEqCmp = CompareOperator{opr: GreaterOrEqualOperate}
var lessCmp = CompareOperator{opr: LessOperate}
var lessEqCmp = CompareOperator{opr: LessOrEqualOperate}

func ListInOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, errors.New("left or right must not be null")
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	for _, r := range rSlice {
		if eq, e := equalCmp.Operate(left, r); eq && e == nil {
			return true, nil
		} else if e != nil {
			err = e
		}
	}
	return false, err
}

func ListNotInOperate(left, right interface{}) (bool, error) {
	res, err := ListInOperate(left, right)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func ListContainsOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, errors.New("left or right must not be null")
	}
	lSlice, err := tools.ConvToSliceUnifyType(left)
	if err != nil {
		return false, err
	}
	for _, l := range lSlice {
		// once hit, return true
		if eq, e := equalCmp.Operate(l, right); eq && e == nil {
			return true, nil
		} else if e != nil {
			err = e
		}
	}
	return false, err
}

func ListNotContainsOperate(left, right interface{}) (bool, error) {
	res, err := ListContainsOperate(left, right)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func ListRetainOperate(left, right interface{}) (bool, error) {
	if left == nil && right == nil {
		return true, nil
	}
	if left == nil || right == nil {
		return false, nil
	}
	lSlice, err := tools.ConvToSliceUnifyType(left)
	if err != nil {
		return false, err
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	leftLen, rightLen := len(lSlice), len(rSlice)
	if leftLen == 0 && rightLen == 0 {
		return true, nil
	}
	if leftLen == 0 && rightLen != 0 {
		return false, nil
	}
	if leftLen != 0 && rightLen == 0 {
		return false, nil
	}
	record := map[string]bool{}
	// O(m + n)
	// record right hit
	for i := 0; i < rightLen; i++ {
		record[tools.ConvToUnifiedStringType(rSlice[i])] = true
	}
	// then record left hit
	for i := 0; i < leftLen; i++ {
		if record[tools.ConvToUnifiedStringType(lSlice[i])] {
			return true, nil
		}
	}
	// finally, return the error to indicate
	return false, err
}

func ListNotRetainOperate(left, right interface{}) (bool, error) {
	res, err := ListRetainOperate(left, right)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func SubListInOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	lSlice, err := tools.ConvToSliceUnifyType(left)
	if err != nil {
		return false, err
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	leftLen, rightLen := len(lSlice), len(rSlice)
	if leftLen == 0 || rightLen == 0 {
		return false, nil
	}
	// record right element hit
	record := make(map[string]bool, rightLen)
	for i := 0; i < rightLen; i++ {
		record[tools.ConvToUnifiedStringType(rSlice[i])] = true
	}
	for i := 0; i < leftLen; i++ {
		if !record[tools.ConvToUnifiedStringType(lSlice[i])] {
			return false, nil
		}
	}
	return true, nil
}

func SubListNotInOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	res, err := SubListInOperate(left, right)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func SubListContainsOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	res, err := SubListInOperate(right, left)
	if err != nil {
		return false, err
	}
	return res, nil
}

func SubListNotContainsOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	res, err := SubListInOperate(right, left)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func ListVagueContainsOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, errors.New("left or right must not be null")
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	for _, r := range rSlice {
		lj, e := tools.ConvToJSONValidStr(left)
		if e != nil {
			return false, e
		}
		rj, e := tools.ConvToJSONValidStr(r)
		if e != nil {
			return false, e
		}
		if strings.Contains(lj, rj) {
			return true, nil
		}
	}
	return false, nil
}

func ListNotVagueContainsOperate(left, right interface{}) (bool, error) {
	res, err := ListVagueContainsOperate(left, right)
	if err != nil {
		return false, err
	}
	return !res, nil
}

func BetweenAllCloseOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	if len(rSlice) < 2 {
		return false, errors.New("between operate: right slice require at least two elements")
	}
	// a <= x <= b
	// a <= x
	greaterEq, err := greaterEqCmp.Operate(left, rSlice[0])
	if err != nil {
		return false, err
	}
	// x <= b
	lessEq, err := lessEqCmp.Operate(left, rSlice[1])
	if err != nil {
		return false, err
	}
	return greaterEq && lessEq, nil
}

func NotBetweenAllCloseOperate(left, right interface{}) (bool, error) {
	betweenAllClose, err := BetweenAllCloseOperate(left, right)
	if err != nil {
		return false, err
	}
	return !betweenAllClose, nil
}

func BetweenAllOpenOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	if len(rSlice) < 2 {
		return false, errors.New("between operate: right slice require at least two elements")
	}
	// a < x < b
	// a < x
	greater, err := greaterCmp.Operate(left, rSlice[0])
	if err != nil {
		return false, err
	}
	// x < b
	less, err := lessCmp.Operate(left, rSlice[1])
	if err != nil {
		return false, err
	}
	return greater && less, nil
}

func NotBetweenAllOpenOperate(left, right interface{}) (bool, error) {
	betweenAllOpen, err := BetweenAllOpenOperate(left, right)
	if err != nil {
		return false, err
	}
	return !betweenAllOpen, nil
}

func BetweenLeftOpenRightCloseOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil || rSlice == nil {
		return false, err
	}
	if len(rSlice) < 2 {
		return false, errors.New("between operate: right slice require at least two elements")
	}
	// a < x <= b
	// a < x
	greater, err := greaterCmp.Operate(left, rSlice[0])
	if err != nil {
		return false, err
	}
	// x <= b
	lessEq, err := lessEqCmp.Operate(left, rSlice[1])
	if err != nil {
		return false, err
	}
	return greater && lessEq, nil
}

func NotBetweenLeftOpenRightCloseOperate(left, right interface{}) (bool, error) {
	betweenLeftOpenRightClose, err := BetweenLeftOpenRightCloseOperate(left, right)
	if err != nil {
		return false, err
	}
	return !betweenLeftOpenRightClose, nil
}

func BetweenLeftCloseRightOpenOperate(left, right interface{}) (bool, error) {
	if left == nil || right == nil {
		return false, nil
	}
	rSlice, err := tools.ConvToSliceUnifyType(right)
	if err != nil {
		return false, err
	}
	if len(rSlice) < 2 {
		return false, errors.New("between operate: right slice require at least two elements")
	}
	// a <= x < b
	// a <= x
	greaterEq, err := greaterEqCmp.Operate(left, rSlice[0])
	if err != nil {
		return false, err
	}
	// x <= b
	less, err := lessCmp.Operate(left, rSlice[1])
	if err != nil {
		return false, err
	}
	return greaterEq && less, nil
}

func NotBetweenLeftCloseRightOpenOperate(left, right interface{}) (bool, error) {
	betweenLeftCloseRightOpen, err := BetweenLeftCloseRightOpenOperate(left, right)
	if err != nil {
		return false, err
	}
	return !betweenLeftCloseRightOpen, nil
}

func (b *ListOperator) Operate(left, right interface{}) (bool, error) {
	return b.opr(left, right)
}
