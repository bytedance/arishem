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
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/internal/pool"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"github.com/dgraph-io/ristretto"
	"math"
	"runtime"
	"strings"
	"time"
)

var (
	defTreeSz      int64 = 1 << 8
	defCmpPoolSz         = 1 << 7
	defIOPoolSz          = 1 << 5
	defCmpWaitTime       = 5 * time.Second
	defIOWaitTime        = 30 * time.Second
	dCpu                 = dualCpuNum()
	factor               = int(math.Pow(float64(dCpu*4), 0.5))
)

func WithDefMaxParallels() core.Option {
	return func(cfg *core.Configuration) {
		cfg.MaxParallel = 10 * (1 << 12)
	}
}

func WithDefGranularity() core.Option {
	return func(cfg *core.Configuration) {
		cfg.Granularity = granularity
	}
}

func WithDefTreeCache() core.Option {
	return func(cfg *core.Configuration) {
		c, _ := ristretto.NewCache(&ristretto.Config{
			NumCounters:        10 * defTreeSz,
			MaxCost:            defTreeSz,
			BufferItems:        64,
			IgnoreInternalCost: true,
		})
		cfg.TCache = &defaultTreeCache{
			storage: c,
		}
	}
}

func WithDefFeatureFetchPool() core.Option {
	return func(cfg *core.Configuration) {
		cfg.FeatureFetchPool = pool.NewWorkPool(defIOPoolSz, defIOWaitTime)
	}
}

func WithDefRuleComputePool() core.Option {
	return func(cfg *core.Configuration) {
		cfg.RuleComputePool = pool.NewWorkPool(defCmpPoolSz, defCmpWaitTime)
	}
}

func WithDefFeatFetcherFactory() core.Option {
	return func(cfg *core.Configuration) {
		cfg.FeatFetcherFactory = func() typedef.FeatureFetcher {
			return &defaultFeatureFeature{}
		}
	}
}

func WithDefFeatVisitCache() core.Option {
	return func(cfg *core.Configuration) {
		c, _ := ristretto.NewCache(&ristretto.Config{
			NumCounters:        10 * defTreeSz,
			MaxCost:            defTreeSz,
			BufferItems:        64,
			IgnoreInternalCost: true,
		})
		cfg.FeatVisitCache = &defaultFeatCache{c: c}
	}
}

type defaultFeatureFeature struct{}

func (d *defaultFeatureFeature) AddFetchObserver(v ...typedef.FeatureFetchObserver) {}

func (d *defaultFeatureFeature) GetFetchObservers() []typedef.FeatureFetchObserver { return nil }

func (d *defaultFeatureFeature) ClearFetchObservers() {}

func (d *defaultFeatureFeature) FetchFeature(ctx context.Context, feat typedef.FeatureParam, dc typedef.DataCtx) (typedef.MetaType, error) {
	return nil, nil
}

type defaultFeatCache struct {
	c *ristretto.Cache
}

func (d *defaultFeatCache) Get(key interface{}) (val interface{}, ok bool) {
	return d.c.Get(key)
}

func (d *defaultFeatCache) Set(key interface{}, val interface{}) {
	if tools.IsNil(key) {
		return
	}
	d.c.Set(key, val, 1)
}

type defaultTreeCache struct {
	storage *ristretto.Cache
}

func (d *defaultTreeCache) Put(expr string, tc *core.RuleTree) {
	expr = strings.TrimSpace(expr)
	if expr == "" || tools.IsNil(tc) {
		return
	}
	d.storage.Set(expr, tc, 1)
}

func (d *defaultTreeCache) Get(expr string) (tc *core.RuleTree, ok bool) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return
	}
	var out interface{}
	out, ok = d.storage.Get(expr)
	if ok && out != nil {
		tc, ok = out.(*core.RuleTree)
	}
	return
}

func granularity(ruleGroupSize int, model core.ExecuteModel) int {
	if model == core.ExecuteModelNoPriority {
		if ruleGroupSize <= arishemConfiguration.MaxParallel {
			return ruleGroupSize
		}
		return arishemConfiguration.MaxParallel
	}
	if ruleGroupSize < dCpu {
		return ruleGroupSize
	}
	if ruleGroupSize < dCpu*2 {
		return ruleGroupSize / 2
	}
	if dCpu <= 1 { // when x = 1，2*√x intersects with 2*x
		return 1
	}
	return (dCpu*2 - factor) * 2
}

// dualCpuNum returns twice the number of CPU cores in the current OS system, platform dynamics
func dualCpuNum() int {
	ncpu := runtime.NumCPU()
	switch runtime.GOOS {
	case "darwin":
		return ncpu
	case "linux", "freebsd", "netbsd", "android":
		return ncpu * 2
	default:
		return ncpu
	}
}
