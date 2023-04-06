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

package parser

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/tools"
	"github.com/bytedance/arishem/internal/typedef"
	"strings"
)

//ParseArishemCondition parse condition group context from expression string, return an error if syntax is not valid.
func ParseArishemCondition(expr string, pos ...typedef.ParseObserver) (ICondEntityContext, error) {
	// parse listener
	cl := NewCachedListener()

	ruleParser, err := parseArishemRuleFromStr(expr, cl.tokenCEL)
	if err != nil {
		return nil, err
	}
	if !cl.CheckSyntaxValid() {
		return nil, cl.GenerateError()
	}
	// remove error default listeners
	ruleParser.RemoveErrorListeners()
	// add parse listener
	ruleParser.AddErrorListener(cl.parseCEL)

	// add cached parse listener first
	ruleParser.AddParseListener(cl)
	if len(pos) > 0 {
		cl.AddListenerProxy(pos...)
	}
	// parse condition
	condGroupCtx := ruleParser.CondEntity()

	if !cl.CheckSyntaxValid() {
		return nil, cl.GenerateError()
	}
	return condGroupCtx, nil
}

// ParseArishemAim parse aim context from expression string, return an error if syntax is not valid.
func ParseArishemAim(expr string, pos ...typedef.ParseObserver) (IAimEntityContext, error) {
	// parse listener
	cl := NewCachedListener()

	ruleParser, err := parseArishemRuleFromStr(expr, cl.tokenCEL)
	if err != nil {
		return nil, err
	}
	if !cl.CheckSyntaxValid() {
		return nil, cl.GenerateError()
	}
	// remove error default listeners
	ruleParser.RemoveErrorListeners()
	// add parse listener
	ruleParser.AddErrorListener(cl.parseCEL)

	// add cached parse listener first
	ruleParser.AddParseListener(cl)
	if len(pos) > 0 {
		cl.AddListenerProxy(pos...)
	}
	// parse aim
	aimCtx := ruleParser.AimEntity()

	if !cl.CheckSyntaxValid() {
		return nil, cl.GenerateError()
	}
	return aimCtx, nil
}

func parseArishemRuleFromStr(expr string, errl antlr.ErrorListener) (*arishemParser, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, errors.New("empty expr")
	}
	var err error
	expr, err = tools.Compat(expr)
	if err != nil {
		return nil, err
	}
	input := antlr.NewInputStream(expr)
	lexer := NewarishemLexer(input)

	if errl != nil {
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errl)
	}
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	return NewarishemParser(tokens), nil
}
