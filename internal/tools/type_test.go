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

func TestConvToBool(t *testing.T) {
	testCases := []struct {
		in    interface{}
		check func(out interface{}, err error)
	}{
		{
			"1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.True(t, out.(bool))
			},
		},
		{
			"+1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.True(t, out.(bool))
			},
		},
		{
			"-1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.True(t, out.(bool))
			},
		},
		{
			"1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.True(t, out.(bool))
			},
		},
		{
			"+1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.True(t, out.(bool))
			},
		},
		{
			"0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.False(t, out.(bool))
			},
		},
		{
			"+0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.False(t, out.(bool))
			},
		},
		{
			"-0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.False(t, out.(bool))
			},
		},
		{
			"0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.False(t, out.(bool))
			},
		},
		{
			"+0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.False(t, out.(bool))
			},
		},
	}

	for _, tc := range testCases {
		tc.check(ConvToBool(tc.in))
	}
}

func TestConvToInt64(t *testing.T) {
	testCases := []struct {
		in    interface{}
		check func(out interface{}, err error)
	}{
		{
			"1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(1), out)
			},
		},
		{
			"+1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(1), out)
			},
		},
		{
			"-1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(-1), out)
			},
		},
		{
			"1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(1), out)
			},
		},
		{
			"+1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(1), out)
			},
		},
		{
			"0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(0), out)
			},
		},
		{
			"+0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(0), out)
			},
		},
		{
			"-0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(0), out)
			},
		},
		{
			"0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(0), out)
			},
		},
		{
			"+0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(0), out)
			},
		},
		{
			"9223372036854775807", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(9223372036854775807), out)
			},
		},
		{
			"-9223372036854775808", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(-9223372036854775808), out)
			},
		},
		{
			"1111.808", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(1111), out)
			},
		},
		{
			"-1111.808", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(-1111), out)
			},
		},
		{
			"-1.111808E3", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, int64(-1111), out)
			},
		},
	}

	for _, tc := range testCases {
		tc.check(ConvToInt64(tc.in))
	}
}

func TestConvToFloat64(t *testing.T) {
	testCases := []struct {
		in    interface{}
		check func(out interface{}, err error)
	}{
		{
			"1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(1), out)
			},
		},
		{
			"+1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(1), out)
			},
		},
		{
			"-1", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(-1), out)
			},
		},
		{
			"1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(1), out)
			},
		},
		{
			"+1.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(1), out)
			},
		},
		{
			"0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(0), out)
			},
		},
		{
			"+0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(0), out)
			},
		},
		{
			"-0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(0), out)
			},
		},
		{
			"0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(0), out)
			},
		},
		{
			"+0.0", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, float64(0), out)
			},
		},
		{
			"9223372036854775807", func(out interface{}, err error) {
				assert.Nil(t, err)
				// precision lost
				assert.Equal(t, 9223372036854775808.0, out)
			},
		},
		{
			"-9223372036854775808", func(out interface{}, err error) {
				assert.Nil(t, err)
				// precision lost
				assert.Equal(t, -9223372036854776000.0, out)
			},
		},
		{
			"1111.808", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 1111.808, out)
			},
		},
		{
			"-1111.808", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, -1111.808, out)
			},
		},
		{
			"-1.111808E3", func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, -1111.808, out)
			},
		},
	}

	for _, tc := range testCases {
		tc.check(ConvToFloat64(tc.in))
	}
}

func TestConvToSliceUnifyType(t *testing.T) {
	testCases := []struct {
		in    interface{}
		check func(out interface{}, err error)
	}{
		{
			1,
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{int64(1)}, out)
			},
		},
		{
			[]uint32{1},
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{uint64(1)}, out)
			},
		},
		{
			[]int32{-1},
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{int64(-1)}, out)
			},
		},
		{
			1.0,
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{1.0}, out)
			},
		},
		{
			[]float32{1.0},
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{1.0}, out)
			},
		},
		{
			"Jack",
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{"Jack"}, out)
			},
		},
		{
			`["1","2","3"]`,
			func(out interface{}, err error) {
				assert.Nil(t, err)
				assert.Equal(t, []interface{}{"1", "2", "3"}, out)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(ConvToSliceUnifyType(tc.in))
	}
}

func TestConvToUnifiedStringType(t *testing.T) {
	testCases := []struct {
		in    interface{}
		check func(out string)
	}{
		{
			123456789,
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "123456789", out)
			},
		},
		{
			1.2345,
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "1.2345", out)
			},
		},
		{
			1.0000000000000001,
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "1", out)
			},
		},
		{
			true,
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "1", out)
			},
		},
		{
			"1.0",
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "1", out)
			},
		},
		{
			[]interface{}{1, 2, 3},
			func(out string) {
				assert.NotEmpty(t, out)
				assert.Equal(t, "[1,2,3]", out)
			},
		},
	}
	for _, tc := range testCases {
		tc.check(ConvToUnifiedStringType(tc.in))
	}
}
