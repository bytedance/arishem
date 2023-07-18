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
)

type NoParamFn func() interface{}
type MapParamFn func(params map[string]interface{}) interface{}
type ListParamFn func(params []interface{}) interface{}

var (
	MapParamFunc = map[string]MapParamFn{
		"ListLength":             ListLength,
		"ListJoin":               ListJoin,
		"DateToUnix":             DateToUnix,
		"UnixToDate":             UnixToDate,
		"RandomUUIDWithReplacer": RandomUUIDWithReplacer,
		"MarshalString":          MarshalString,
	}
	ListParamFunc = map[string]ListParamFn{
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
	NoParamFunc = map[string]NoParamFn{
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

func RegNoPramFunc(funcName string, fn NoParamFn) {
	if NoParamFunc == nil {
		NoParamFunc = map[string]NoParamFn{}
	}
	if tools.IsBlank(funcName) || fn == nil {
		return
	}
	NoParamFunc[funcName] = fn
}

func RegMapParamFunc(funcName string, fn MapParamFn) {
	if MapParamFunc == nil {
		MapParamFunc = map[string]MapParamFn{}
	}
	if tools.IsBlank(funcName) || fn == nil {
		return
	}
	MapParamFunc[funcName] = fn
}

func RegListParamFunc(funcName string, fn ListParamFn) {
	if ListParamFunc == nil {
		ListParamFunc = map[string]ListParamFn{}
	}
	if tools.IsBlank(funcName) || fn == nil {
		return
	}
	ListParamFunc[funcName] = fn
}
