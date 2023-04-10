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
	"sync"
)

var (
	arishemConfiguration *core.Configuration
	once                 = &sync.Once{}
)

func DefaultConfiguration() *core.Configuration {
	c := new(core.Configuration)
	core.ApplyOptions(c,
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
func Initialize(cfg *core.Configuration, opts ...core.Option) {
	if cfg == nil {
		panic("arishem initialize by nil configuration")
	}
	once.Do(func() {
		arishemConfiguration = cfg
		core.ApplyOptions(arishemConfiguration, opts...)
		core.SetFeaturePreParseCache(arishemConfiguration.FeatVisitCache)
		core.SetFeatureFetchPool(arishemConfiguration.FeatureFetchPool)
	})
}

// FeatureFetcherFactory provides the ability to create a feature fetcher.
func FeatureFetcherFactory() typedef.FeatureFetcher {
	return arishemConfiguration.FeatFetcherFactory()
}
