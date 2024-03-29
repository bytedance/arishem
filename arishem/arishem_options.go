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
	"context"
	"github.com/bytedance/arishem/internal/funcs"
	"github.com/bytedance/arishem/typedef"
)

type NoParamFnPair struct {
	Name string
	Fn   func(ctx context.Context) (interface{}, error)
}

type MapParamFnPair struct {
	Name string
	Fn   func(ctx context.Context, params map[string]interface{}) (interface{}, error)
}

type ListParamFnPair struct {
	Name string
	Fn   func(ctx context.Context, params []interface{}) (interface{}, error)
}

func WithFeatureFetcherFactory(factory func() typedef.FeatureFetcher) Option {
	return func(cfg *Configuration) {
		cfg.FeatFetcherFactory = factory
	}
}

// WithVisitObserver can add visit observers to rule visitor.
func WithVisitObserver(obs ...typedef.VisitObserver) ExecuteOption {
	return func(rv typedef.RuleVisitor, dc typedef.DataCtx) {
		if rv != nil {
			rv.AddVisitObserver(obs...)
		}
	}
}

// WithFeatureFetchObserver can add feature fetch observers to feature fetcher.
func WithFeatureFetchObserver(obs ...typedef.FeatureFetchObserver) ExecuteOption {
	return func(rv typedef.RuleVisitor, dc typedef.DataCtx) {
		if dc != nil && dc.FetcherFetcher() != nil {
			dc.FetcherFetcher().AddFetchObserver(obs...)
		}
	}
}

func WithCustomNoParamFuncs(funcPairs ...NoParamFnPair) Option {
	return func(cfg *Configuration) {
		for _, pair := range funcPairs {
			funcs.RegNoPramFunc(pair.Name, pair.Fn)
		}
	}
}

func WithCustomMapParamFuncs(funcPairs ...MapParamFnPair) Option {
	return func(cfg *Configuration) {
		for _, pair := range funcPairs {
			funcs.RegMapParamFunc(pair.Name, pair.Fn)
		}
	}
}

func WithCustomListParamFuncs(funcPairs ...ListParamFnPair) Option {
	return func(cfg *Configuration) {
		for _, pair := range funcPairs {
			funcs.RegListParamFunc(pair.Name, pair.Fn)
		}
	}
}
