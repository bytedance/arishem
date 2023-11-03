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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// ArishemParserRuleContext implement GET/SET altNum function and cache the text.
type ArishemParserRuleContext struct {
	*antlr.BaseParserRuleContext
	//altNum int
	text string
}

func NewArishemParserRuleContext(parent antlr.ParserRuleContext, invokingState int) *ArishemParserRuleContext {
	cp := &ArishemParserRuleContext{
		BaseParserRuleContext: antlr.NewBaseParserRuleContext(parent, invokingState),
		//altNum:                antlr.ATNInvalidAltNumber,
	}
	cp.SetParent(parent)
	cp.SetInvokingState(invokingState)
	return cp
}

// GetAltNumber get the alternative number of context.
//func (c *ArishemParserRuleContext) GetAltNumber() int {
//	return c.altNum
//}

// SetAltNumber set the alternative number of context.
//func (c *ArishemParserRuleContext) SetAltNumber(altNumber int) {
//	c.altNum = altNumber
//}

// GetText original use DFS to generate text without cache storage, here to store it in first invoking.
func (c *ArishemParserRuleContext) GetText() string {
	if len(c.text) <= 0 {
		c.text = c.BaseParserRuleContext.GetText()
	}
	return c.text
}
