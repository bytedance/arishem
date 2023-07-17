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

package funcs

import (
	"fmt"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/sonic"
	"math"
)

func MinFlatWithInt(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	var minSave int64 = math.MaxInt64
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			var i64 int64
			i64, err = tools.ConvToInt64(ele)
			if err != nil {
				continue
			}
			if i64 < minSave {
				minSave = i64
			}
		}
	}
	if minSave == math.MaxInt64 {
		minSave = 0
	}
	return minSave
}

func MinFlatWithFloat(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	var minSave float64 = math.MaxInt64
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			var f64 float64
			f64, err = tools.ConvToFloat64(ele)
			if err != nil {
				continue
			}
			if f64 < minSave {
				minSave = f64
			}
		}
	}
	if minSave == math.MaxInt64 {
		minSave = 0
	}
	return minSave
}

func MaxFlatWithInt(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	var maxSave int64 = math.MinInt64
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			var i64 int64
			i64, err = tools.ConvToInt64(ele)
			if err != nil {
				continue
			}
			if i64 > maxSave {
				maxSave = i64
			}
		}
	}
	if maxSave == math.MinInt64 {
		maxSave = 0
	}
	return maxSave
}

func MaxFlatWithFloat(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	var maxSave float64 = math.MinInt64
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			var f64 float64
			f64, err = tools.ConvToFloat64(ele)
			if err != nil {
				continue
			}
			if f64 > maxSave {
				maxSave = f64
			}
		}
	}
	if maxSave == math.MinInt64 {
		maxSave = 0
	}
	return maxSave
}

//ListLength calculate the length of targetList
func ListLength(params map[string]interface{}) interface{} {
	const (
		keyList       = "list"
		keyFilterNull = "filter_null"
	)

	listVar, ok := params[keyList]
	if !ok {
		return 0
	}
	list, err := tools.ConvToSliceUnifyType(listVar)
	if err != nil {
		return 0
	}
	if list == nil {
		return 0
	}
	filterNull := false
	filterNullStr, ok := params[keyFilterNull]
	if ok {
		filterNull, _ = tools.ConvToBool(filterNullStr)
	}
	length := 0
	for _, ele := range list {
		if filterNull && ele == nil {
			continue
		}
		length++
	}
	return length
}

//ListFlatJoin flat params into one list, params[0] as the separator
func ListFlatJoin(params []interface{}) interface{} {
	if len(params) <= 1 {
		return nil
	}
	// separator
	flatList := ListAdd(params[1:])
	return tools.JoinString(flatList, fmt.Sprint(params[0]))
}

func ListJoin(params map[string]interface{}) interface{} {
	const (
		keyList      = "list"
		keySeparator = "sep"
	)
	listI, ok := params[keyList]
	if !ok {
		return ""
	}
	separator, ok := params[keySeparator]
	if !ok {
		return ""
	}
	return tools.JoinString(listI, fmt.Sprint(separator))
}

func ListAdd(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	elements := make([]interface{}, 0, len(params)*4)
	for _, param := range params {
		list, err := tools.ConvToSliceUnifyType(param)
		if err != nil {
			continue
		}
		for _, ele := range list {
			elements = append(elements, ele)
		}
	}
	return elements
}

func ListRemove(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	firstList, err := tools.ConvToSliceUnifyType(params[0])
	if err != nil {
		return nil
	}
	recordMap := make(map[interface{}]bool)
	for _, e := range firstList {
		if e == nil {
			continue
		}
		pv := propertyVal(e)
		if pv == nil {
			continue
		}
		recordMap[pv] = true
	}
	for i := range params {
		if i == 0 {
			continue
		}
		var list []interface{}
		list, err = tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, e := range list {
			if e == nil {
				continue
			}
			pv := propertyVal(e)
			if pv == nil {
				continue
			}
			recordMap[pv] = false
		}
	}
	finalList := make([]interface{}, 0, len(recordMap))
	for k, v := range recordMap {
		if !v {
			continue
		}
		finalList = append(finalList, k)
	}
	return finalList
}

func ListUnion(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	recordMap := make(map[interface{}]bool, len(params)*4)
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			if ele == nil {
				continue
			}
			pv := propertyVal(ele)
			if pv == nil {
				continue
			}
			recordMap[pv] = true
		}
	}
	finalList := make([]interface{}, 0, len(recordMap))
	for k := range recordMap {
		if k == nil {
			continue
		}
		finalList = append(finalList, k)
	}
	return finalList
}

func ListIntersect(params []interface{}) interface{} {
	if len(params) <= 0 {
		return nil
	}
	if len(params) == 1 {
		list, err := tools.ConvToSliceUnifyType(params[0])
		if err != nil {
			return nil
		}
		return list
	}
	recordMap := make(map[interface{}]int, len(params)*4)
	for i := range params {
		list, err := tools.ConvToSliceUnifyType(params[i])
		if err != nil {
			continue
		}
		for _, ele := range list {
			if ele == nil {
				continue
			}
			pv := propertyVal(ele)
			if pv == nil {
				continue
			}
			if _, ok := recordMap[pv]; ok {
				recordMap[pv]++
			} else {
				recordMap[pv] = 1
			}
		}
	}
	finalList := make([]interface{}, 0, len(recordMap))
	for k, v := range recordMap {
		if v != len(params) {
			continue
		}
		finalList = append(finalList, k)
	}
	return finalList
}

func propertyVal(target interface{}) interface{} {
	if target == nil {
		return nil
	}
	if tools.IsNumber(target) {
		num, err := tools.StringToNumber(fmt.Sprint(target))
		if err != nil {
			return nil
		}
		return num
	} else if tools.IsString(target) {
		return target
	}
	m, err := sonic.MarshalString(target)
	if err != nil {
		return nil
	}
	return m
}
