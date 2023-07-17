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
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestMapIndex(t *testing.T) {
	testCases := []struct {
		name  string
		args  string
		check func(target interface{}, err error)
	}{
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"name.first",
			func(target interface{}, err error) {
				assert.NotNil(t, target)
				assert.Nil(t, err)
				assert.Equal(t, "Tom", target)
			},
		},
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"age",
			func(target interface{}, err error) {
				assert.NotNil(t, target)
				assert.Nil(t, err)
				assert.Equal(t, int64(37), target)
			},
		},
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"children#2.name",
			func(target interface{}, err error) {
				assert.NotNil(t, target)
				assert.Nil(t, err)
				assert.Equal(t, "Jack", target)
			},
		},
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"children#2.age",
			func(target interface{}, err error) {
				assert.NotNil(t, target)
				assert.Nil(t, err)
				assert.Equal(t, int64(10), target)
			},
		},
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"nam",
			func(target interface{}, err error) {
				assert.Nil(t, target)
				assert.Nil(t, err)
			},
		},
		{
			`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`,
			"children#3.name",
			func(target interface{}, err error) {
				assert.Nil(t, target)
				assert.NotNil(t, err)
				t.Log(err)
			},
		},
	}
	for _, testCase := range testCases {
		kv, err := UnmarshalToMapUseInt64(testCase.name)
		if err != nil {
			t.Error(err)
		} else {
			testCase.check(MapIndex(kv, testCase.args))
		}
	}
}

func BenchmarkMapIndex(b *testing.B) {
	kv, err := UnmarshalToMapUseInt64(`{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":10}]}`)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, _ = MapIndex(kv, "children#2.age")
	}
}

func BenchmarkSonicIndex(b *testing.B) {
	str := `{"name":{"first":"Tom","last":"Anderson"},"age":37,"children":[{"name":"Sara","age":4},{"name":"Alex","age":8},{"name":"Jack","age":9223372036854775807}]}`
	root, _ := sonic.GetFromString(str)
	for i := 0; i < b.N; i++ {
		age := root.Get("children").Index(2).Get("age")
		if age != nil && age.Valid() {
			ty := age.Type()
			if ty == ast.V_NUMBER {
				ageStr, err := age.String()
				if err != nil {

				} else {
					if strings.Contains(ageStr, ".") {
						_, _ = strconv.ParseFloat(ageStr, 10)
					} else {
						_, _ = strconv.ParseInt(ageStr, 10, 64)
					}
				}
			}
		}
	}
}
