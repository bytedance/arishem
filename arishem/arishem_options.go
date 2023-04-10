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

package arishem

import (
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/typedef"
)

func WithFeatureFetcherFactory(factory func() typedef.FeatureFetcher) core.Option {
	return func(cfg *core.Configuration) {
		cfg.FeatFetcherFactory = factory
	}
}

// WithVisitObserver can add visit observers to rule visitor.
func WithVisitObserver(obs ...typedef.VisitObserver) core.ExecuteOption {
	return func(rv typedef.RuleVisitor, dc typedef.DataCtx) {
		if rv != nil {
			rv.AddVisitObserver(obs...)
		}
	}
}

// WithFeatureFetchObserver can add feature fetch observers to feature fetcher.
func WithFeatureFetchObserver(obs ...typedef.FeatureFetchObserver) core.ExecuteOption {
	return func(rv typedef.RuleVisitor, dc typedef.DataCtx) {
		if dc != nil && dc.FetcherFetcher() != nil {
			dc.FetcherFetcher().AddFetchObserver(obs...)
		}
	}
}
