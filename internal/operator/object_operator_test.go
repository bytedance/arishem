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

package operator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNull(t *testing.T) {
	const OPERATOR = "IS_NULL"
	var intNil *int = nil
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			nil,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			intNil,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
				assert.False(t, isNil(intNil))
			},
		},
		{
			1,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			struct{}{},
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(Judge(tc.left, tc.right, OPERATOR))
	}
}

func TestNotNull(t *testing.T) {
	const OPERATOR = "NOT IS_NULL"
	var intNil *int = nil
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			nil,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			intNil,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
				assert.True(t, !isNil(intNil))
			},
		},
		{
			1,
			nil,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(Judge(tc.left, tc.right, OPERATOR))
	}
}

func isNil(x interface{}) bool {
	return x == nil
}
