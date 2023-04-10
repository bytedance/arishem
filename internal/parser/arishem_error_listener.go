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

package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

//CollectErrorListener collect syntax errors and indicate expression whether valid.
type CollectErrorListener struct {
	*antlr.DefaultErrorListener
	SyntaxValid bool
	ErrorMsgs   []string
}

//NewCollectErrorListener construct a new CollectErrorListener
func NewCollectErrorListener() *CollectErrorListener {
	rel := new(CollectErrorListener)
	rel.SyntaxValid = true
	return rel
}

//SyntaxError collect formatted error message, set valid false.
func (r *CollectErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	r.ErrorMsgs = append(r.ErrorMsgs, fmt.Sprintf("line@%d,%d:%s", line, column, msg))
	r.SyntaxValid = false
}
