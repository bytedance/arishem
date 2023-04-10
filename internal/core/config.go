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

package core

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/bytedance/arishem/internal/pool"
	"github.com/bytedance/arishem/internal/typedef"
)

// Option used to config arishem engine
type Option func(cfg *Configuration)
type ExecuteOption func(rv typedef.RuleVisitor, dc typedef.DataCtx)
type ExecuteModel int32

const (
	ExecuteModelPriority ExecuteModel = iota
	ExecuteModelNoPriority
)

type RuleTree struct {
	Tree       antlr.ParseTree
	FeatParams []typedef.FeatureParam
}

type TreeCache interface {
	Put(expr string, tc *RuleTree)
	Get(expr string) (tc *RuleTree, ok bool)
}

type Configuration struct {
	Granularity func(int, ExecuteModel) int
	// prefetch feature
	Prefetch bool
	TCache   TreeCache
	// max parallel computation
	MaxParallel        int
	RuleComputePool    *pool.WorkPool
	FeatureFetchPool   *pool.WorkPool
	FeatFetcherFactory func() typedef.FeatureFetcher
	FeatVisitCache     typedef.SharedVisitCache
}

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
