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

package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsNumber(t *testing.T) {
	testCases := []struct {
		args  interface{}
		check func(is bool)
	}{
		{
			1000,
			func(is bool) {
				assert.True(t, is)
			},
		},
		{
			1000.10223,
			func(is bool) {
				assert.True(t, is)
			},
		},
		{
			"1000",
			func(is bool) {
				assert.False(t, is)
			},
		},
		{
			true,
			func(is bool) {
				assert.False(t, is)
			},
		},
		{
			struct{}{},
			func(is bool) {
				assert.False(t, is)
			},
		},
	}
	for _, testCase := range testCases {
		testCase.check(IsNumber(testCase.args))
	}
}

func TestStringToNumber(t *testing.T) {
	testCases := []struct {
		args  string
		check func(num interface{}, err error)
	}{
		{
			"1000",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, uint64(1000), num)
			},
		},
		{
			"-1000",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, int64(-1000), num)
			},
		},
		{
			"0",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, uint64(0), num)
			},
		},
		{
			"0.0",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, 0.0, num)
			},
		},
		{
			"1000.10223",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, 1000.10223, num)
			},
		},
		{
			"1.1022e3",
			func(num interface{}, err error) {
				assert.Nil(t, err)
				assert.IsType(t, 1102.2, num)
			},
		},
	}
	for _, testCase := range testCases {
		testCase.check(StringToNumber(testCase.args))
	}
}
