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

func TestBetweenAllCloseOperate(t *testing.T) {
	const OpId = BetweenAllClose
	// @case 1
	judge, err := Judge(1, 2, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 2}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 5
	judge, err = Judge(0, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 6
	judge, err = Judge(0, []string{"0", "1", "2"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
}

func TestNotBetweenAllCloseOperate(t *testing.T) {
	const OpId = "!" + BetweenAllClose
	// @case 1
	judge, err := Judge(1, 2, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 2}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 5
	judge, err = Judge(0, []string{"1", "2"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 6
	judge, err = Judge(0, []string{"0", "1", "2"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
}

func TestBetweenAllOpenOperate(t *testing.T) {
	const OpId = BetweenAllOpen
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 5
	judge, err = Judge(0, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 6
	judge, err = Judge(2, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
}

func TestNotBetweenAllOpenOperate(t *testing.T) {
	const OpId = "!" + BetweenAllOpen
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 5
	judge, err = Judge(0, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 6
	judge, err = Judge(2, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
}

func TestBetweenLeftOpenRightCloseOperate(t *testing.T) {
	const OpId = BetweenLeftOpenRightClose
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 5
	judge, err = Judge(5, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 6
	judge, err = Judge(3, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
}

func TestNotBetweenLeftOpenRightCloseOperate(t *testing.T) {
	const OpId = "!" + BetweenLeftOpenRightClose
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 5
	judge, err = Judge(5, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 6
	judge, err = Judge(3, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
}

func TestBetweenLeftCloseRightOpenOperate(t *testing.T) {
	const OpId = BetweenLeftCloseRightOpen
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 5
	judge, err = Judge(5, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 6
	judge, err = Judge(3, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
}

func TestNotBetweenLeftCloseRightOpenOperate(t *testing.T) {
	const OpId = "!" + BetweenLeftCloseRightOpen
	// @case 1
	judge, err := Judge(1, 5, OpId)
	assert.NotNil(t, err, "err=>%v", err)
	assert.False(t, judge)
	// @case 2
	judge, err = Judge(1, []int{1, 5}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 3
	judge, err = Judge(1, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 4
	judge, err = Judge(3, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
	// @case 5
	judge, err = Judge(5, []string{"1", "5"}, OpId)
	assert.Nil(t, err)
	assert.True(t, judge)
	// @case 6
	judge, err = Judge(3, []string{"1", "4", "8"}, OpId)
	assert.Nil(t, err)
	assert.False(t, judge)
}

func TestListInOperate(t *testing.T) {
	const OPERATOR = ListIn
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			"true",
			"false",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"true",
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"true",
			[]int{1, 2, 3},
			func(judge bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, judge)
				t.Log(err.Error())
			},
		},
		{
			true,
			[]int{1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"faLse",
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, judge)
				t.Log(err.Error())
			},
		},
		{
			[]string{"false"},
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, judge)
				t.Log(err.Error())
			},
		},
		{
			"true",
			[]float32{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.NotNil(t, err)
				assert.False(t, judge)
				t.Log(err.Error())
			},
		},
		{
			false,
			[]float32{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"jack",
			[]string{"jack", "mike"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			1,
			`[1,2,3]`,
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

func TestListContainsOperate(t *testing.T) {
	const OPERATOR = ListContains
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{

			"false",
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"true",
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			[]int{1, 2, 3},
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]int{1, 2, 3},
			true,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			[]int{0, 1, 2, 3},
			"faLse",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]int{0, 1, 2, 3},
			[]string{"false"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]float32{0, 1, 2, 3},
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]float32{0, 1, 2, 3},
			true,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			[]string{"jack", "mike"},
			"jack",
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

func TestListRetainOperate(t *testing.T) {
	const OPERATOR = ListRetain
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			"false",
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"true",
			"true",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			true,
			[]int{1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"faLse",
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]interface{}{false},
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"true",
			[]float32{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"jack",
			[]string{"jack", "mike"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			`[1,2,3]`,
			`["1","2","3"]`,
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

func TestSubListInOperate(t *testing.T) {
	const OPERATOR = SubListIn
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{

			[]string{"false"},
			[]string{"true"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]string{"true"},
			[]string{"true"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			[]string{"true"},
			[]int{1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]string{"faLse"},
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			[]string{"false"},
			[]int{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"true",
			[]float32{0, 1, 2, 3},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"jack",
			[]string{"jack", "mike"},
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			`[1,2,3]`,
			`["1","2","3"]`,
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
