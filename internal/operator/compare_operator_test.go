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

package operator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualOperate(t *testing.T) {
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			1,
			"1",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			1,
			"1.0",
			func(judge bool, err error) {
				assert.Nil(t, err)
				// "1.0" != "1"
				assert.False(t, judge)
			},
		},
		{
			"1",
			"1.0",
			func(judge bool, err error) {
				assert.Nil(t, err)
				// "1.0" != "1"
				assert.False(t, judge)
			},
		},
		{
			"1",
			1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"1",
			-1.0,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"1.1",
			-1.1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"-1.1",
			-1.1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"1.11e3",
			-1110,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"-1.11e3",
			-1110,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"-12341234",
			-12341234,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"12341234",
			-12341234,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(Judge(tc.left, tc.right, Equal))
	}
}

func TestLessOperate(t *testing.T) {
	testCases := []struct {
		left  interface{}
		right interface{}
		check func(judge bool, err error)
	}{
		{
			1,
			"1",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			1,
			"1.0",
			func(judge bool, err error) {
				assert.Nil(t, err)
				// "1.0" < "1"
				assert.True(t, judge)
			},
		},
		{
			"1",
			"1.0",
			func(judge bool, err error) {
				assert.Nil(t, err)
				// "1.0" < "1"
				assert.True(t, judge)
			},
		},
		{
			"1",
			1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"1",
			-1.0,
			func(judge bool, err error) {
				assert.Nil(t, err)
				// 1 > -1
				assert.False(t, judge)
			},
		},
		{
			"1.1",
			-1.1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				// 1.1 > -1.1
				assert.False(t, judge)
			},
		},
		{
			"-1.1",
			-1.1,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"1.11e3",
			-1110,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"-1.11e3",
			1110,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"-12341234",
			12341234,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"12341234",
			-12341234,
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"2023-02-13",
			"2023-02-14",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"2023-02-14",
			"2023-02-14",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"2023-02-14 15:06:00",
			"2023-02-14",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.False(t, judge)
			},
		},
		{
			"-0123",
			"0123",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
		{
			"10",
			"2",
			func(judge bool, err error) {
				assert.Nil(t, err)
				assert.True(t, judge)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(Judge(tc.left, tc.right, Less))
	}
}
