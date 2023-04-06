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
	"bytes"
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

func BenchmarkCompat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Compat(compatJSONStr)
	}
}

func BenchmarkCompatByStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, len(compatJSONStr)*3/2))
		_ = json.Compact(buf, StrToBytes(compatJSONStr))
	}
}

func BenchmarkCompatAlready(b *testing.B) {
	src := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ConstList":[{"StrConst":"测试中文"},{"StrConst":"测试特殊字符          ～！@#¥%……&*（）-=——+【】、「」｜；’：”，。/《》？"},{"StrConst":"Test English"},{"StrConst":"Test Special Characters\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/"},{"StrConst":"Test Special Alphabet œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"}]}}]}`
	for i := 0; i < b.N; i++ {
		_, _ = Compat(src)
	}
}
func BenchmarkCompatAlreadyByStd(b *testing.B) {
	src := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ConstList":[{"StrConst":"测试中文"},{"StrConst":"测试特殊字符          ～！@#¥%……&*（）-=——+【】、「」｜；’：”，。/《》？"},{"StrConst":"Test English"},{"StrConst":"Test Special Characters\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/"},{"StrConst":"Test Special Alphabet œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"}]}}]}`
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, len(src)*3/2))
		_ = json.Compact(buf, StrToBytes(src))
	}
}

func BenchmarkMarshal(b *testing.B) {
	type Person struct {
		Name string
		Age  int
	}
	for i := 0; i < b.N; i++ {
		_, _ = sonic.MarshalString(&Person{
			Name: "<<>>>>=",
			Age:  100,
		})
	}
}
