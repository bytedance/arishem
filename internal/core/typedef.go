/*
 *  Copyright Bytedance Ltd. and/or its affiliates.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package core

import (
	"strings"
)

const (
	logicAnd   = "and"
	logicOr    = "or"
	logicNot   = "not"
	mathDivide = "/"
	mathMod    = "%"
	strQuot    = `"`
	strDot     = "."
)

func toStdLogic(logicStr string) string {
	logicStr = strings.ToLower(logicStr)
	if logicStr == "and" || logicStr == "&&" {
		return logicAnd
	} else if logicStr == "or" || logicStr == "||" {
		return logicOr
	} else if logicStr == "not" || logicStr == "!" {
		return logicNot
	}
	return ""
}

func StrippingQuotation(text string) string {
	return strings.TrimPrefix(strings.TrimSuffix(text, strQuot), strQuot)
}
