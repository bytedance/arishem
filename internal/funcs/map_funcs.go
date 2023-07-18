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
	"context"
	"errors"
)

func MapUnion(ctx context.Context, params []interface{}) (interface{}, error) {
	if len(params) <= 0 {
		return nil, emptyParam
	}
	kv := make(map[string]interface{}, len(params)*4)
	for _, param := range params {
		m, ok := param.(map[string]interface{})
		if !ok {
			continue
		}
		for k, v := range m {
			kv[k] = v
		}
	}
	return kv, nil
}

func MapIntersect(ctx context.Context, params []interface{}) (interface{}, error) {
	if len(params) <= 0 {
		return nil, emptyParam
	}
	if len(params) == 1 {
		m, ok := params[0].(map[string]interface{})
		if !ok {
			return nil, errors.New("[MapIntersect] param type not map")
		}
		return m, nil
	}
	type bv struct {
		val   interface{}
		times int
	}
	kv := make(map[string]*bv, len(params)*4)
	for _, param := range params {
		m, ok := param.(map[string]interface{})
		if !ok {
			continue
		}
		for k, v := range m {
			if val, o := kv[k]; o {
				val.times++
			} else {
				kv[k] = &bv{val: v, times: 1}
			}
		}
	}
	finalMap := make(map[string]interface{}, len(kv))
	for k, v := range kv {
		if v == nil || v.times != len(params) {
			continue
		}
		finalMap[k] = v.val
	}
	return finalMap, nil
}
