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
	"fmt"
	"github.com/bytedance/arishem/tools"
	"math/big"
	"strings"
)

// CompareOperator will invoke func(cmp int) bool
type CompareOperator struct {
	opr func(cmp int) bool
}

func EqualOperate(cmp int) bool {
	return cmp == 0
}

func NotEqualOperate(cmp int) bool {
	return cmp != 0
}

func GreaterOperate(cmp int) bool {
	return cmp > 0
}

func LessOperate(cmp int) bool {
	return cmp < 0
}

func GreaterOrEqualOperate(cmp int) bool {
	return cmp >= 0
}

func LessOrEqualOperate(cmp int) bool {
	return cmp <= 0
}

// Operate the left and value whether valid by opr, will convert left to right type implicit.
func (c *CompareOperator) Operate(left, right interface{}) (bool, error) {
	if c.opr == nil {
		return false, errors.New("compare operator is nil")
	}
	// convert to left type to right type
	switch right.(type) {
	case int, int8, int16, int32, int64:
		lv, err := tools.ConvToInt64(left)
		if err != nil {
			return false, fmt.Errorf("left to right type(int) error, left(%T,%v)", left, left)
		}
		rv, _ := tools.ConvToInt64(right)
		if lv < rv {
			return c.opr(-1), nil
		} else if lv == rv {
			return c.opr(0), nil
		} else {
			return c.opr(+1), nil
		}
	case uint, uint8, uint16, uint32, uint64:
		lv, err := tools.ConvToUint64(left)
		if err != nil {
			return false, fmt.Errorf("left to right type(uint) error, left(%T,%v)", left, left)
		}
		rv, _ := tools.ConvToUint64(right)
		if lv < rv {
			return c.opr(-1), nil
		} else if lv == rv {
			return c.opr(0), nil
		} else {
			return c.opr(+1), nil
		}
	case float32, float64:
		lv, err := tools.ConvToFloat64(left)
		if err != nil {
			return false, fmt.Errorf("left to right type(float) error, left(%T,%v)", left, left)
		}
		rv, _ := tools.ConvToFloat64(right)
		return c.opr(big.NewFloat(lv).Cmp(big.NewFloat(rv))), nil
	case bool:
		lv, err := tools.ConvToBool(left)
		if err != nil {
			return false, fmt.Errorf("left to right type(bool) error, left(%T,%v)", left, left)
		}
		rv := right.(bool)
		if lv == rv {
			return c.opr(0), nil
		} else if lv {
			return c.opr(1), nil
		} else {
			return c.opr(-1), nil
		}
	case string:
		return c.opr(strings.Compare(fmt.Sprint(left), right.(string))), nil
	default:
		rj, err := tools.ConvToJSONValidStr(right)
		if err != nil {
			return false, fmt.Errorf("right(%T,%v) to json valid str err=>%s", right, right, err.Error())
		}
		lj, err := tools.ConvToJSONValidStr(left)
		if err != nil {
			return false, fmt.Errorf("left(%T,%v) to json valid str err=>%s", left, left, err.Error())
		}
		return c.opr(strings.Compare(lj, rj)), nil
	}
}
