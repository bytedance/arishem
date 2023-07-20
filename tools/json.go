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
	"bytes"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
	fastjson "github.com/goccy/go-json"
	"strings"
)

var sonicUseInt64 = sonic.Config{
	UseInt64: true,
}.Froze()

// ConvToJSONValidStr convert in to json valid string, basic type will use fmt.Sprint，complex type use json marshal
func ConvToJSONValidStr(in interface{}) (string, error) {
	if in == nil {
		return "null", nil
	}
	switch t := in.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool, string:
		return fmt.Sprint(t), nil
	default:
		m, err := sonic.MarshalString(t)
		if err != nil {
			return "", err
		}
		return m, nil
	}
}

// UnmarshalToSliceUseInt64 deserialize string to slice type, number use int64
func UnmarshalToSliceUseInt64(jsonStr string) ([]interface{}, error) {
	if IsBlank(jsonStr) {
		return nil, errors.New("json string is empty")
	}
	var arr []interface{}
	err := sonicUseInt64.UnmarshalFromString(jsonStr, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// UnmarshalToMapUseInt64 deserialize string to map type, number use int64
func UnmarshalToMapUseInt64(jsonStr string) (map[string]interface{}, error) {
	if IsBlank(jsonStr) {
		return nil, errors.New("json string is empty")
	}

	var kv map[string]interface{}
	err := sonicUseInt64.UnmarshalFromString(jsonStr, &kv)
	if err != nil {
		return nil, err
	}
	return kv, nil
}

// IsJSONValid validates json and returns first non-blank character position,
// if it is only one valid json value,
// Otherwise it returns invalid character position using start.
//
// Note: it does not check for the invalid UTF-8 characters.
func IsJSONValid(text string) (bool, int) {
	if IsBlank(text) {
		return false, -1
	}
	return encoder.Valid(StrToBytes(text))
}

// Compat appends to dst the JSON-encoded src with
// insignificant space characters elided.
func Compat(src string) (dst string, err error) {
	src = strings.TrimSpace(src)
	if src == "" {
		return
	}
	buf := bytes.NewBuffer(make([]byte, 0, len(src)*3/2))
	err = fastjson.Compact(buf, StrToBytes(src))
	if err != nil {
		return
	}
	dst = buf.String()
	return
}
