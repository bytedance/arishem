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
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

var unescaper = strings.NewReplacer(`\,`, `,`, `\"`, `"`, `\ `, ` `, `\=`, `=`)

// StrToBytes by using point operation, not value copy.
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// BytesToStr by using point operation, not value copy.
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) <= 0
}

func IsString(target interface{}) bool {
	if target == nil {
		return false
	}
	switch target.(type) {
	case string:
		return true
	default:
		return false
	}
}

func HasChineseChar(str string) bool {
	var flag bool
	for _, v := range str {
		flag = unicode.Is(unicode.Han, v)
		if flag {
			break
		}
	}
	return flag
}

func JoinString(value interface{}, sep string) string {
	res := ""
	if value == nil {
		return res
	}
	objList := reflect.ValueOf(value)
	if objList.Kind() == reflect.Slice && objList.Len() > 0 {
		v := objList.Index(0).Interface()
		if v != nil && v != "" {
			res += fmt.Sprint(v)
		}
		for i, siz := 1, objList.Len(); i < siz; i++ {
			res += sep
			val := objList.Index(i).Interface()
			if val == nil || val == "" {
				continue
			}
			res += fmt.Sprint(val)
		}
	} else {
		v := objList.Interface()
		if v != nil && v != "" {
			res += fmt.Sprint(v)
		}
	}
	return res
}

func TruncateStringWithEllipsis(text string, maxSiz int) string {
	const truncateSuffix = `......[truncated]`
	left, _ := TruncateString(text, maxSiz, truncateSuffix)
	return left
}

func TruncateString(text string, maxSiz int, suffix string) (string, bool) {
	left, truncated := TruncateBytes(StrToBytes(text), maxSiz, StrToBytes(suffix))
	return BytesToStr(left), truncated
}

// TruncateBytes godoc
// @text param: check should be truncated
// @maxSiz param: truncate size whether @text should be truncated
// @suffix param: suffix add to text when @text is truncated
// @[]byte bytes calculate left
// @bool bytes whether truncated
// this method designed to prevent huge bytes allocate and performance cost
func TruncateBytes(text []byte, maxSiz int, suffix []byte) ([]byte, bool) {
	textLen := len(text)
	if textLen <= 0 {
		return text, false
	}
	suffixLen := len(suffix)
	if textLen <= maxSiz+suffixLen {
		return text, false
	}
	if suffix == nil {
		return text[:maxSiz], true
	}
	buf := make([]byte, maxSiz+suffixLen)
	copy(buf, text[:maxSiz])
	copy(buf[maxSiz:], suffix)
	return buf, true
}

// UnescapeString returns unescaped version of in.
func UnescapeString(in string) string {
	if strings.IndexByte(in, '\\') == -1 { // has no '\'
		return in
	}
	return unescaper.Replace(in)
}
