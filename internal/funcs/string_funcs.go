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

package funcs

import (
	"context"
	"errors"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/encoder"
)

func MarshalString(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	const (
		keyTarget     = "target"
		keyEscapeHTML = "escape_html"
	)
	target, ok := params[keyTarget]
	if !ok {
		return nil, errors.New("[MarshalString] param: target not found")
	}
	var escapeHTML bool
	espI, ok := params[keyEscapeHTML]
	if ok {
		escapeHTML, _ = tools.ConvToBool(espI)
	}
	var str string
	if escapeHTML {
		encode, err := encoder.Encode(target, encoder.EscapeHTML)
		if err == nil {
			str = tools.BytesToStr(encode)
		}
	} else {
		m, err := sonic.MarshalString(target)
		if err == nil {
			str = m
		}
	}
	return str, nil
}
