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

package funcs

import (
	"fmt"
)

var (
	MapParamFunc = map[string]func(params map[string]interface{}) interface{}{
		"ListLength":             ListLength,
		"ListJoin":               ListJoin,
		"DateToUnix":             DateToUnix,
		"UnixToDate":             UnixToDate,
		"RandomUUIDWithReplacer": RandomUUIDWithReplacer,
		"MarshalString":          MarshalString,
	}
	ListParamFunc = map[string]func(params []interface{}) interface{}{
		"MinFlatWithInt":   MinFlatWithInt,
		"MinFlatWithFloat": MinFlatWithFloat,
		"MaxFlatWithInt":   MaxFlatWithInt,
		"MaxFlatWithFloat": MaxFlatWithFloat,
		"ListFlatJoin":     ListFlatJoin,
		"ListAdd":          ListAdd,
		"ListRemove":       ListRemove,
		"ListUnion":        ListUnion,
		"ListIntersect":    ListIntersect,
		"MapUnion":         MapUnion,
		"MapIntersect":     MapIntersect,
	}
	NoParamFunc = map[string]func() interface{}{
		"GetUnixSecond":      GetUnixSecond,
		"RandomUUID":         RandomUUID,
		"GetCurrentYear":     GetCurrentYear,
		"GetCurrentYearDay":  GetCurrentYearDay,
		"GetCurrentMonth":    GetCurrentMonth,
		"GetCurrentMonthDay": GetCurrentMonthDay,
		"GetCurrentWeekDay":  GetCurrentWeekDay,
		"GetCurrentHour":     GetCurrentHour,
		"GetCurrentMinute":   GetCurrentMinute,
		"GetCurrentSecond":   GetCurrentSecond,
		"GetCurrentLocation": GetCurrentLocation,
	}
)

func InvokeFunc(funcname string, paramMap map[string]interface{}, paramList []interface{}) (interface{}, error) {
	if fn, ok := NoParamFunc[funcname]; ok {
		return fn(), nil
	}
	if fn, ok := MapParamFunc[funcname]; ok {
		return fn(paramMap), nil
	}
	if fn, ok := ListParamFunc[funcname]; ok {
		return fn(paramList), nil
	}
	return nil, fmt.Errorf("no function named=>%s found", funcname)
}
