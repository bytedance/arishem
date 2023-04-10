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

package core

import (
	"errors"
	"github.com/bytedance/arishem/internal/pool"
	"github.com/bytedance/arishem/internal/tools"
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/bytedance/sonic"
	"strings"
	"sync"
)

var (
	emptyParamMap typedef.MetaType = make(map[string]interface{})
	featFetchPool *pool.WorkPool
	featFPOnce    = &sync.Once{}
)

func SetFeatureFetchPool(p *pool.WorkPool) {
	if p == nil {
		panic(errors.New("feature fetch pool is nil"))
	}
	featFPOnce.Do(func() {
		featFetchPool = p
	})
}

type featWithLatch struct {
	feat  typedef.FeatureParam
	latch *sync.WaitGroup
}

type featMetaWithParam struct {
	param typedef.MetaType
	meta  typedef.MetaType
}

type arishemDataCtx struct {
	factMeta       typedef.MetaType
	visitCache     *sync.Map
	featLatches    *tools.OnceMap
	featMap        *sync.Map
	extraMap       *sync.Map
	featureFetcher typedef.FeatureFetcher
}

func NewArishemDataCtx(factJSON string, ff typedef.FeatureFetcher) (dc typedef.DataCtx, err error) {
	if ff == nil {
		err = errors.New("feature fetcher can not be nil")
		return
	}
	var fm typedef.MetaType
	factJSON = strings.TrimSpace(factJSON)
	if factJSON == "" {
		fm = emptyParamMap
	} else {
		fm, err = tools.UnmarshalToMapUseInt64(factJSON)
		if err != nil {
			return
		}
	}
	dc = &arishemDataCtx{
		factMeta:       fm,
		visitCache:     &sync.Map{},
		featLatches:    tools.NewOnceMap(),
		featMap:        &sync.Map{},
		extraMap:       &sync.Map{},
		featureFetcher: ff,
	}
	return
}

func (a *arishemDataCtx) Get(key interface{}) (val interface{}, ok bool) {
	return a.visitCache.Load(key)
}

func (a *arishemDataCtx) Set(key interface{}, val interface{}) {
	if tools.IsNil(key) || tools.IsNil(val) {
		return
	}
	a.visitCache.Store(key, val)
}

func (a *arishemDataCtx) FetcherFetcher() typedef.FeatureFetcher {
	return a.featureFetcher
}

func (a *arishemDataCtx) PrefetchFeatures(feats []typedef.FeatureParam) {
	if a.featureFetcher == nil {
		return
	}
	for _, feat := range feats {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		// filter features which have been already submitted request
		once := a.featLatches.WriteOnce(feat.HashCode(), wg)
		if !once {
			// once indicates that a request for this feature has been already submitted if once is false
			continue
		}
		featFetchPool.Submit(func(param interface{}) {
			fwl := param.(*featWithLatch)
			defer fwl.latch.Done()

			for _, obs := range a.featureFetcher.GetFetchObservers() {
				obs.OnFeatureFetchStart(fwl.feat)
			}
			meta, err := a.featureFetcher.FetchFeature(fwl.feat, a)
			// store it into feature map
			meta = a.storeFeatureMeta(fwl.feat, meta)
			for _, obs := range a.featureFetcher.GetFetchObservers() {
				obs.OnFeatureFetchEnd(fwl.feat.HashCode(), meta, err)
			}
		}, &featWithLatch{feat: feat, latch: wg})
	}
}

func (a *arishemDataCtx) GetFactMeta() typedef.MetaType {
	return a.factMeta
}

func (a *arishemDataCtx) GetVarValue(varPath []string) (interface{}, error) {
	if len(varPath) <= 0 {
		return nil, errors.New("variable paths are empty")
	}
	return tools.MapIndexPaths(a.factMeta, varPath)
}

func (a *arishemDataCtx) GetFeatureValue(featParam typedef.FeatureParam, fieldPaths ...string) (val interface{}, err error) {
	if featParam == nil {
		err = errors.New("feature param is nil or field paths are empty")
		return
	}
	load, ok := a.featMap.Load(featParam.HashCode())
	if ok && load != nil {
		val, err = tools.MapIndexPaths(load.(*featMetaWithParam).meta, fieldPaths)
		return
	}
	// check latch
	latch := a.featLatches.Read(featParam.HashCode())
	if latch != nil {
		latch.(*sync.WaitGroup).Wait()
	}
	// check featMap again
	load, ok = a.featMap.Load(featParam.HashCode())
	if ok && load != nil {
		val, err = tools.MapIndexPaths(load.(*featMetaWithParam).meta, fieldPaths)
		return
	}
	// if still no data here, means the fetch failed,
	// we go back to fetch, no longer consider the duplication of concurrent fetch

	// call back
	for _, obs := range a.featureFetcher.GetFetchObservers() {
		obs.OnFeatureFetchStart(featParam)
	}

	// fetch here
	var feature typedef.MetaType
	feature, err = a.featureFetcher.FetchFeature(featParam, a)

	for _, obs := range a.featureFetcher.GetFetchObservers() {
		obs.OnFeatureFetchEnd(featParam.HashCode(), feature, err)
	}

	// store the data
	feature = a.storeFeatureMeta(featParam, feature)
	val, err = tools.MapIndexPaths(feature, fieldPaths)
	return
}

func (a *arishemDataCtx) GetAllFeatures() []typedef.Feature {
	var features []typedef.Feature
	a.featMap.Range(func(key, value interface{}) bool {
		fwp := value.(featMetaWithParam)
		features = append(features, newArishemFeature(key.(string), fwp.param, fwp.meta))
		return true
	})
	return features
}

func (a *arishemDataCtx) SetExtra(key interface{}, value interface{}) {
	a.extraMap.Store(key, value)
}

func (a *arishemDataCtx) GetExtra(key interface{}) (value interface{}) {
	value, _ = a.extraMap.Load(key)
	return
}

func (a *arishemDataCtx) storeFeatureMeta(fp typedef.FeatureParam, meta typedef.MetaType) (stored typedef.MetaType) {
	if len(meta) <= 0 || fp == nil {
		return
	}
	if d, ok := a.featMap.Load(fp.HashCode()); ok && d != nil {
		return
	}
	m, err := sonic.MarshalString(meta)
	if err != nil {
		return
	}
	var mt map[string]interface{}
	mt, err = tools.UnmarshalToMapUseInt64(m)
	if err != nil {
		return
	}
	a.featMap.Store(fp.HashCode(), &featMetaWithParam{param: fp.BuiltinParam(), meta: mt})
	stored = mt
	return
}
