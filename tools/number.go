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

package tools

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func IsNumber(target interface{}) bool {
	if target == nil {
		return false
	}
	switch target.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return true
	default:
		return false
	}
}

func IsFloatNumber(target interface{}) bool {
	if target == nil {
		return false
	}
	switch t := target.(type) {
	case float32, float64:
		return true
	case string:
		// a string first has '.' or 'eE' before it can be a float number
		if strings.Contains(t, ".") || strings.ContainsAny(t, "eE") {
			_, err := strconv.ParseFloat(t, 64)
			return err == nil
		}
		return false
	default:
		return false
	}
}

// StringToNumber convert string to number type, float format will be converted to float64 which support scientific format,
// and a signed number will be converted to int64, unsigned number will be converted to uint64.
func StringToNumber(numStr string) (num interface{}, err error) {
	numStr = strings.TrimSpace(numStr)
	if numStr == "" {
		err = errors.New("empty number string")
		return
	}
	// float determinate
	if strings.Contains(numStr, ".") || strings.ContainsAny(numStr, "eE") {
		num, err = StringToFloat(numStr)
	} else {
		num, err = strconv.ParseInt(numStr, 10, 64)
	}
	return
}

// StringToFloat support scientific format.
func StringToFloat(numStr string) (num float64, err error) {
	pos := strings.IndexAny(numStr, "eE")
	if pos < 0 {
		num, err = strconv.ParseFloat(numStr, 64)
	} else {
		baseStr := numStr[0:pos]
		var baseVal float64
		baseVal, err = strconv.ParseFloat(baseStr, 64)
		if err == nil {
			expStr := numStr[pos+1:]
			var expVal int
			expVal, err = strconv.Atoi(expStr)
			if err == nil {
				num = baseVal * math.Pow10(expVal)
			}
		}
	}
	return
}
