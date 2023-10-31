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
	"fmt"

	"github.com/google/uuid"
	"strings"
)

func RandomUUID(ctx context.Context) (interface{}, error) {
	nUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return nUUID.String(), nil
}

func RandomUUIDWithReplacer(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	const (
		keyReplacer = "replacer"
	)
	repHyphen := false
	replacer := "-"
	repHyI, ok := params[keyReplacer]
	if ok {
		repHyphen = true
		replacer = fmt.Sprint(repHyI)
	}
	nUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	u := nUUID.String()
	if repHyphen && replacer != "-" {
		u = strings.ReplaceAll(u, "-", replacer)
	}
	return u, nil
}
