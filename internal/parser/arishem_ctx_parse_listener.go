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
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/bytedance/gopkg/util/xxhash3"
	"strings"
)

// CachedListener implement rule exit and trigger rule's GetText() manually and generate alternative number.
type CachedListener struct {
	tokenCEL *CollectErrorListener
	parseCEL *CollectErrorListener
	proxy    []typedef.ParseObserver
}

func NewCachedListener() *CachedListener {
	return &CachedListener{
		tokenCEL: NewCollectErrorListener(),
		parseCEL: NewCollectErrorListener(),
	}
}

// VisitTerminal no need for implementation
func (c *CachedListener) VisitTerminal(node antlr.TerminalNode) {
	for _, p := range c.proxy {
		p.VisitTerminal(node, c.tokenCEL.ErrorMsgs, c.parseCEL.ErrorMsgs)
	}
}

// VisitErrorNode no need for implementation
func (c *CachedListener) VisitErrorNode(node antlr.ErrorNode) {
	for _, p := range c.proxy {
		p.VisitErrorNode(node, c.tokenCEL.ErrorMsgs, c.parseCEL.ErrorMsgs)
	}
}

// EnterEveryRule no need for implementation
func (c *CachedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	for _, p := range c.proxy {
		p.EnterEveryRule(ctx, c.tokenCEL.ErrorMsgs, c.parseCEL.ErrorMsgs)
	}
}

// ExitEveryRule Set the outer alternative number for this context node.
// Using hashString to generate alt num, 32-bit(normal 64-bit) number is enough for texts
func (c *CachedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// trigger the ArishemParserRuleContext GetText() function manually to store text, in case of frequency GetText() invoking
	text := ctx.GetText()
	altNum := int(xxhash3.HashString(text))
	ctx.SetAltNumber(altNum)
	// callback proxy listeners
	for _, p := range c.proxy {
		p.ExitEveryRule(ctx, c.tokenCEL.ErrorMsgs, c.parseCEL.ErrorMsgs)
	}
}

func (c *CachedListener) AddListenerProxy(pos ...typedef.ParseObserver) {
	if len(pos) <= 0 {
		return
	}
	if pos == nil {
		return
	}
	if c.proxy == nil {
		c.proxy = make([]typedef.ParseObserver, 0, 4)
	}
	for _, po := range pos {
		if po == nil {
			continue
		}
		c.proxy = append(c.proxy, pos...)
	}
}

func (c *CachedListener) CheckSyntaxValid() bool {
	return c.tokenCEL.SyntaxValid && c.parseCEL.SyntaxValid
}

func (c *CachedListener) GenerateError() error {
	if len(c.tokenCEL.ErrorMsgs) <= 0 && len(c.parseCEL.ErrorMsgs) <= 0 {
		return nil
	}
	errMsg := ""
	if len(c.tokenCEL.ErrorMsgs) > 0 {
		errMsg += "lexer errors: " + strings.Join(c.tokenCEL.ErrorMsgs, ", ")
	}
	if len(c.parseCEL.ErrorMsgs) > 0 {
		if errMsg != "" {
			errMsg += "\n"
		}
		errMsg += "syntax errors: " + strings.Join(c.parseCEL.ErrorMsgs, ", ")
	}
	return errors.New(errMsg)
}
