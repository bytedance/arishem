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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/typedef"
	"sync"
)

type Option func(cfg *Configuration)
type ExecuteOption func(rv typedef.RuleVisitor, dc typedef.DataCtx)
type ExecuteMode int32

type RuleTree struct {
	Tree       antlr.ParseTree
	FeatParams []typedef.FeatureParam
}

type TreeCache interface {
	Put(expr string, tc *RuleTree)
	Get(expr string) (tc *RuleTree, ok bool)
}

type Configuration struct {
	Granularity func(int, ExecuteMode) int
	// prefetch feature
	Prefetch bool
	TCache   TreeCache
	// max parallel computation
	MaxParallel        int
	RuleComputePool    typedef.ConcurrentPool
	FeatureFetchPool   typedef.ConcurrentPool
	FeatFetcherFactory func() typedef.FeatureFetcher
	FeatVisitCache     typedef.SharedVisitCache
}

const (
	ExecuteModePriority ExecuteMode = iota
	ExecuteModeNoPriority
)

var (
	arishemConfiguration *Configuration
	once                 = &sync.Once{}
)

func ApplyOptions(c *Configuration, opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

func ApplyExecuteOptions(rv typedef.RuleVisitor, dc typedef.DataCtx, opts ...ExecuteOption) {
	for _, opt := range opts {
		opt(rv, dc)
	}
}

func DefaultConfiguration() *Configuration {
	c := new(Configuration)
	c.Prefetch = true
	ApplyOptions(c,
		WithDefTreeCache(),
		WithDefMaxParallels(),
		WithDefGranularity(),
		WithDefRuleComputePool(),
		WithDefFeatureFetchPool(),
		WithDefFeatFetcherFactory(),
		WithDefFeatVisitCache(),
	)
	return c
}

// Initialize arishem, this func should be called only once in your init function.
func Initialize(cfg *Configuration, opts ...Option) {
	if cfg == nil {
		panic("arishem initialize by nil configuration")
	}
	once.Do(func() {
		arishemConfiguration = cfg
		ApplyOptions(arishemConfiguration, opts...)
		core.SetFeaturePreParseCache(arishemConfiguration.FeatVisitCache)
		core.SetFeatureFetchPool(arishemConfiguration.FeatureFetchPool)
	})
}

// FeatureFetcherFactory provides the ability to create a feature fetcher.
func FeatureFetcherFactory() typedef.FeatureFetcher {
	return arishemConfiguration.FeatFetcherFactory()
}
