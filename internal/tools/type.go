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
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"math"
	"reflect"
	"strconv"
	"strings"
)

var (
	ErrConvType = errors.New("convert type not support")
	ErrConvNil  = errors.New("convert type is nil")
)

// ConvToUnifiedStringType convert in parameter into unified string type, bool type will be converted into "0" or "1",
// basic type will simply convert to string, other type will try to use json encoding to do converting.
func ConvToUnifiedStringType(in interface{}) (out string) {
	if IsNil(in) {
		return "null"
	}
	switch t := in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		out = fmt.Sprint(t)
	case float32, float64:
		f64, _ := ConvToFloat64(t)
		modf, frac := math.Modf(f64)
		if math.Abs(frac) < floatEpsilon {
			out = fmt.Sprint(int64(modf))
		} else {
			out = fmt.Sprint(t)
		}
	case bool:
		if t {
			out = "1"
		} else {
			out = "0"
		}
	case string:
		switch t {
		case "1.0", "+1.0", "-1.0":
			out = "1"
		case "0.0", "+0.0", "-0.0":
			out = "0"
		default:
			out = t
		}
	default:
		var err error
		out, err = sonic.MarshalString(t)
		if err != nil {
			out = fmt.Sprint(t)
		}
	}
	return
}

// ConvToInt64 convert in parameter to type int64, better performance than runtime reflection
func ConvToInt64(in interface{}) (out int64, err error) {
	switch tVal := in.(type) {
	case nil:
		err = ErrConvType
	case int:
		out = int64(tVal)
	case int8:
		out = int64(tVal)
	case int16:
		out = int64(tVal)
	case int32:
		out = int64(tVal)
	case int64:
		out = tVal
	case uint:
		out = int64(tVal)
	case uint8:
		out = int64(tVal)
	case uint16:
		out = int64(tVal)
	case uint32:
		out = int64(tVal)
	case uint64:
		out = int64(tVal)
	case float32:
		out = int64(tVal)
	case float64:
		out = int64(tVal)
	case bool:
		if tVal {
			out = 1
		} else {
			out = 0
		}
	case string:
		tVal = strings.TrimSpace(tVal)
		var num interface{}
		num, err = StringToNumber(tVal)
		if err == nil {
			switch n := num.(type) {
			case int64:
				out = n
			case uint64:
				out = int64(n)
			case float64:
				out = int64(n)
			}
		}
	default:
		err = ErrConvType
	}
	return
}

// ConvToUint64 convert in parameter to type uint64, better performance than runtime reflection
func ConvToUint64(in interface{}) (out uint64, err error) {
	switch tVal := in.(type) {
	case nil:
		err = ErrConvNil
	case int:
		out = uint64(tVal)
	case int8:
		out = uint64(tVal)
	case int16:
		out = uint64(tVal)
	case int32:
		out = uint64(tVal)
	case int64:
		out = uint64(tVal)
	case uint:
		out = uint64(tVal)
	case uint8:
		out = uint64(tVal)
	case uint16:
		out = uint64(tVal)
	case uint32:
		out = uint64(tVal)
	case uint64:
		out = tVal
	case float32:
		out = uint64(tVal)
	case float64:
		out = uint64(tVal)
	case bool:
		if tVal {
			out = 1
		} else {
			out = 0
		}
	case string:
		tVal = strings.TrimSpace(tVal)
		var num interface{}
		num, err = StringToNumber(tVal)
		if err == nil {
			switch n := num.(type) {
			case int64:
				out = uint64(n)
			case uint64:
				out = n
			case float64:
				out = uint64(n)
			}
		}
	default:
		err = ErrConvType
	}
	return
}

// ConvToFloat64 convert in parameter to type float64, better performance than runtime reflection
func ConvToFloat64(in interface{}) (out float64, err error) {
	switch tVal := in.(type) {
	case nil:
		err = ErrConvNil
	case int:
		out = float64(tVal)
	case int8:
		out = float64(tVal)
	case int16:
		out = float64(tVal)
	case int32:
		out = float64(tVal)
	case int64:
		out = float64(tVal)
	case uint:
		out = float64(tVal)
	case uint8:
		out = float64(tVal)
	case uint16:
		out = float64(tVal)
	case uint32:
		out = float64(tVal)
	case uint64:
		out = float64(tVal)
	case float32:
		out = float64(tVal)
	case float64:
		out = tVal
	case bool:
		if tVal {
			out = 1
		} else {
			out = 0
		}
	case string:
		tVal = strings.TrimSpace(tVal)
		var num interface{}
		num, err = StringToNumber(tVal)
		if err == nil {
			switch n := num.(type) {
			case int64:
				out = float64(n)
			case uint64:
				out = float64(n)
			case float64:
				out = n
			}
		}
	default:
		err = ErrConvType
	}
	return
}

// ConvToBool convert in parameter to type boolean, better performance than runtime reflection
func ConvToBool(in interface{}) (out bool, err error) {
	switch tVal := in.(type) {
	case nil:
		err = ErrConvNil
	case int:
		out = tVal == 1
	case int8:
		out = tVal == 1
	case int16:
		out = tVal == 1
	case int32:
		out = tVal == 1
	case int64:
		out = tVal == 1
	case uint:
		out = tVal == 1
	case uint8:
		out = tVal == 1
	case uint16:
		out = tVal == 1
	case uint32:
		out = tVal == 1
	case uint64:
		out = tVal == 1
	case float32:
		out = math.Abs(float64(1-tVal)) < floatEpsilon
	case float64:
		out = math.Abs(1-tVal) < floatEpsilon
	case bool:
		out = tVal
	case string:
		tVal = strings.TrimSpace(tVal)
		switch tVal {
		case "+1", "-1", "1.0", "+1.0", "-1.0":
			out = true
		case "+0", "-0", "0.0", "+0.0", "-0.0":
			out = false
		default:
			out, err = strconv.ParseBool(tVal)
		}
	default:
		err = ErrConvType
	}
	return
}

// ConvToSliceUnifyType convert in parameter to type slice, better performance than runtime reflection
// will unify element type to int64, float64 when in type is integer or float
func ConvToSliceUnifyType(in interface{}) (out []interface{}, err error) {
	switch in.(type) {
	case nil:
		out = nil
	case int, int8, int16, int32, int64:
		t, _ := ConvToInt64(in)
		out = []interface{}{t}
	case uint, uint8, uint16, uint32, uint64:
		t, _ := ConvToUint64(in)
		out = []interface{}{t}
	case float32, float64:
		t, _ := ConvToFloat64(in)
		out = []interface{}{t}
	case bool:
		out = []interface{}{in.(bool)}
	case string:
		s, ue := UnmarshalToSliceUseInt64(in.(string))
		if ue != nil {
			out = []interface{}{in.(string)}
		} else {
			out = s
		}
	case []int:
		inSlice := in.([]int)
		for _, o := range inSlice {
			out = append(out, int64(o))
		}
	case []int8:
		inSlice := in.([]int8)
		for _, o := range inSlice {
			out = append(out, int64(o))
		}
	case []int16:
		inSlice := in.([]int16)
		for _, o := range inSlice {
			out = append(out, int64(o))
		}
	case []int32:
		inSlice := in.([]int32)
		for _, o := range inSlice {
			out = append(out, int64(o))
		}
	case []int64:
		inSlice := in.([]int64)
		for _, o := range inSlice {
			out = append(out, o)
		}
	case []uint:
		inSlice := in.([]uint)
		for _, o := range inSlice {
			out = append(out, uint64(o))
		}
	case []uint8:
		inSlice := in.([]uint8)
		for _, o := range inSlice {
			out = append(out, uint64(o))
		}
	case []uint16:
		inSlice := in.([]uint16)
		for _, o := range inSlice {
			out = append(out, uint64(o))
		}
	case []uint32:
		inSlice := in.([]uint32)
		for _, o := range inSlice {
			out = append(out, uint64(o))
		}
	case []uint64:
		inSlice := in.([]uint64)
		for _, o := range inSlice {
			out = append(out, o)
		}
	case []float32:
		inSlice := in.([]float32)
		for _, o := range inSlice {
			out = append(out, float64(o))
		}
	case []float64:
		inSlice := in.([]float64)
		for _, o := range inSlice {
			out = append(out, o)
		}
	case []bool:
		inSlice := in.([]bool)
		for _, o := range inSlice {
			out = append(out, o)
		}
	case []string:
		inSlice := in.([]string)
		for _, o := range inSlice {
			out = append(out, o)
		}
	case []interface{}:
		out = in.([]interface{})
	default:
		out, err = convToSliceUnifyTypeByReflect(in)
	}
	return
}

func convToSliceUnifyTypeByReflect(in interface{}) (out []interface{}, err error) {
	inValue := reflect.ValueOf(in)
	if inValue.Kind() != reflect.Slice || inValue.Kind() != reflect.Array {
		err = ErrConvType
		return
	}
	out = make([]interface{}, 0, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		out = append(out, inValue.Index(i).Interface())
	}
	return
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Itob(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func U8tob(u uint8) bool {
	if u == 1 {
		return true
	}
	return false
}

func I32tob(u int32) bool {
	if u == 1 {
		return true
	}
	return false
}
