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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcListLength(t *testing.T) {
	paramKeyList := "list"
	paramKeyFilter := "filter_null"
	params := map[string]interface{}{}
	var length interface{}
	// @case 1
	params[paramKeyList] = nil
	params[paramKeyFilter] = true
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Zero(t, length)
	// @case 2
	params[paramKeyList] = 3
	params[paramKeyFilter] = true
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Equal(t, 1, length)
	// @case 3
	params[paramKeyList] = "123"
	params[paramKeyFilter] = true
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Equal(t, 1, length)
	// @case 4
	params[paramKeyList] = []int64{}
	params[paramKeyFilter] = true
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Zero(t, length)
	// @case 5
	params[paramKeyList] = []interface{}{nil, nil}
	params[paramKeyFilter] = true
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Zero(t, length)
	// @case 6
	params[paramKeyList] = []interface{}{nil, nil}
	params[paramKeyFilter] = false
	length = ListLength(params)
	assert.NotNil(t, length)
	assert.Equal(t, 2, length)
}
