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
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/stretchr/testify/assert"
	"testing"
)

type myError struct {
}

func (m *myError) Error() string {
	return "myError"
}

func TestIsNil(t *testing.T) {
	var nilMyError *myError = nil
	var nilInt *int = nil
	var nilString *int64 = nil
	var nilInterface typedef.Hashable = nil
	testCases := []struct {
		args  interface{}
		check func(is1 bool, is2 bool)
	}{

		{
			nil,
			func(is1 bool, is2 bool) {
				assert.True(t, is1)
				assert.True(t, is2)
			},
		},
		{
			nilMyError,
			func(is1 bool, is2 bool) {
				assert.True(t, is1)
				assert.False(t, is2)
			},
		},
		{
			nilInt,
			func(is1 bool, is2 bool) {
				assert.True(t, is1)
				assert.False(t, is2)
			},
		},
		{
			nilString,
			func(is1 bool, is2 bool) {
				assert.True(t, is1)
				assert.False(t, is2)
			},
		},
		{
			nilInterface,
			func(is1 bool, is2 bool) {
				assert.True(t, is1)
				assert.True(t, is2)
			},
		},
		{
			1000,
			func(is1 bool, is2 bool) {
				assert.False(t, is1)
				assert.False(t, is2)
			},
		},
		{
			"1000",
			func(is1 bool, is2 bool) {
				assert.False(t, is1)
				assert.False(t, is2)
			},
		},
	}
	for _, testCase := range testCases {
		testCase.check(IsNil(testCase.args), isNil(testCase.args))
	}
}

func isNil(x interface{}) bool {
	return x == nil
}
