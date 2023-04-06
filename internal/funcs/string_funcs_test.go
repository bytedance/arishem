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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalString(t *testing.T) {
	testCases := []struct {
		param map[string]interface{}
		check func(interface{})
	}{
		{
			map[string]interface{}{
				"target":      false,
				"escape_html": true,
			},
			func(i interface{}) {
				assert.NotNil(t, i)
				assert.NotEmpty(t, i.(string))
				assert.Equal(t, "false", i.(string))
			},
		},
		{
			map[string]interface{}{
				"target": struct {
					Name string
					Age  int
				}{
					"<Thunder>",
					18,
				},
				"escape_html": true,
			},
			func(i interface{}) {
				assert.NotNil(t, i)
				assert.NotEmpty(t, i.(string))
				assert.Equal(t, `{"Name":"\u003cThunder\u003e","Age":18}`, i.(string))
			},
		},
		{
			map[string]interface{}{
				"target": struct {
					Name string
					Age  int
				}{
					"<Thunder>",
					18,
				},
			},
			func(i interface{}) {
				assert.NotNil(t, i)
				assert.NotEmpty(t, i.(string))
				assert.Equal(t, `{"Name":"<Thunder>","Age":18}`, i.(string))
			},
		},
	}
	for _, tc := range testCases {
		tc.check(MarshalString(tc.param))
	}
}
