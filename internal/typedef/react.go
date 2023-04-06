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

package typedef

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// ParseObserver is the observer that can listen to rule parse events with token and parse error messages.
type ParseObserver interface {
	VisitTerminal(node antlr.TerminalNode, tokenErrMsgs, parseErrMsgs []string)
	VisitErrorNode(node antlr.ErrorNode, tokenErrMsgs, parseErrMsgs []string)
	EnterEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string)
	ExitEveryRule(ctx antlr.ParserRuleContext, tokenErrMsgs, parseErrMsgs []string)
}

// VisitObserver is the observer that can listen to rule visit events.
type VisitObserver interface {
	Hashable

	// OnJudgeNodeVisitEnd passes the condition node
	OnJudgeNodeVisitEnd(info JudgeNode, vt VisitTarget)
	OnVisitError(node, errMsg string, vt VisitTarget)
}

// ObservableVisitor supports rule events callback behaviour.
type ObservableVisitor interface {
	AddVisitObserver(v ...VisitObserver)
	GetVisitObservers() []VisitObserver
	ClearVisitObservers()
}

// FeatureFetchObserver can listen to feature fetch events.
type FeatureFetchObserver interface {
	Hashable

	OnFeaturesFetchStart(feat FeatureParam)
	OnFeaturesFetchEnd(featureHash string, featureValue MetaType, err error)
}

// ObservableFeatureFetcher supports feature fetch events callback behaviour.
type ObservableFeatureFetcher interface {
	AddFetchObserver(v ...FeatureFetchObserver)
	GetFetchObservers() []FeatureFetchObserver
	ClearFetchObservers()
}
