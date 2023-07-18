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

package typedef

import "context"

// SharedVisitCache is the interface that can pass data among multi visitors.
type SharedVisitCache interface {
	// Get will get cached node visit result by its alternative number, Set and Get must concurrent safe.
	Get(key interface{}) (val interface{}, ok bool)
	// Set will set a visit result by its alternative number.
	Set(key interface{}, val interface{})
}

// DataCtx is the data context during the rule visitation.
type DataCtx interface {
	SharedVisitCache

	Context() context.Context

	FetcherFetcher() FeatureFetcher
	PrefetchFeatures(feats []FeatureParam)

	GetFactMeta() MetaType
	GetVarValue(varPath []string) (interface{}, error)
	GetFeatureValue(featParam FeatureParam, fieldPaths ...string) (interface{}, error)

	GetAllFeatures() []Feature

	SetExtra(key interface{}, value interface{})
	GetExtra(key interface{}) (value interface{})
}
