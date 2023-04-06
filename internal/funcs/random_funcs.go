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

package funcs

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func RandomUUID() interface{} {
	return uuid.NewV4().String()
}

func RandomUUIDWithReplacer(params map[string]interface{}) interface{} {
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
	u := uuid.NewV4().String()
	if repHyphen && replacer != "-" {
		u = strings.ReplaceAll(u, "-", replacer)
	}
	return u
}
