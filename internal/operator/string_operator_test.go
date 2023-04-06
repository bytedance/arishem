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

func TestStringContainsOperate(t *testing.T) {
	const OPERATOR = "STRING_CONTAINS"
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			"123",
			"1",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"123",
			" ",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]int{1, 2, 3},
			"3",
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

func TestStringEndWithOperate(t *testing.T) {
	const OPERATOR = "STRING_END_WITH"
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			[]int{1, 2, 3},
			"3",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"123",
			"3",
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

func TestStringStartWithOperate(t *testing.T) {
	const OPERATOR = "STRING_START_WITH"
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			[]int{1, 2, 3},
			"1",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"123",
			"1",
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

func TestContainRegexpOperate(t *testing.T) {
	const OPERATOR = "CONTAIN_REGULAR"
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			"123",
			`.*`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"123",
			`1234`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"123",
			`[0-9]+`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"",
			`[0-9]+`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"",
			`[a-zA-Z_]*`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"Mike",
			`[a-zA-Z_]*`,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"9Mike",
			`^[a-zA-Z_][a-zA-Z_]*`,
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
