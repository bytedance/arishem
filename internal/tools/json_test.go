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
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	compatJSONStr = `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "VarExpr": "upgrade_time"
            },
            "Rhs": {
                "ConstList": [
                    {
                        "StrConst": "测试中文"
                    },
                    {
                        "StrConst": "测试特殊字符   \t\n\r       ～！@#¥%……&*（）-=——+【】、「」｜；’：”，。/《》？"
                    },
                    {
                        "StrConst": "Test English"
                    },
                    {
                        "StrConst": "Test Special Characters\n\b\t\r\f~!@#$%^&*()-=_+[]\\{}|;':\",.?<>/"
                    },
                    {
                        "StrConst": "Test Special Alphabet œ∑´®†¥¨ˆøπ“‘«åß∂ƒ©˙∆˚¬…æΩ≈ç√∫˜µ≤≥÷"
                    }
                ]
            }
        }
    ]
}`
)

func TestCompat(t *testing.T) {
	// compat by standard lib
	buf := bytes.NewBuffer(make([]byte, 0, len(compatJSONStr)*3/2))
	err := json.Compact(buf, StrToBytes(compatJSONStr))
	assert.Nil(t, err)
	compatStd := buf.String()
	t.Log(compatStd)
	// compat by go-json
	for i := 0; i < 100; i++ {
		var compat string
		compat, err = Compat(compatJSONStr)
		assert.Nil(t, err)
		assert.Equal(t, compatStd, compat)
	}
}

func TestMarshal(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	for i := 0; i < 100000; i++ {
		m, err := sonic.MarshalString(&Person{
			Name: "<<>>>>=",
			Age:  100,
		})
		assert.Nil(t, err)
		assert.Equal(t, `{"Name":"<<>>>>=","Age":100}`, m)
	}
}

func TestMarshalOrderKeys(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	for i := 0; i < 1000; i++ {
		m, err := sonic.ConfigStd.MarshalToString(&Person{
			Name: "<<>>>>=",
			Age:  100,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, `{"Name":"<<>>>>=","Age":100}`, m)
	}
	sc := sonic.Config{
		EscapeHTML:  false,
		SortMapKeys: true,
	}.Froze()
	for i := 0; i < 10000; i++ {
		m, err := sc.MarshalToString(map[string]interface{}{
			"Name": "<<>>>>=",
			"Age":  100,
		})
		assert.Nil(t, err)
		assert.NotEqual(t, `{"Name":"<<>>>>=","Age":100}`, m)
		//assert.Equal(t, `{"Name":"<<>>>>=","Age":100}`, m)
		assert.Equal(t, `{"Age":100,"Name":"<<>>>>="}`, m)
	}
}
