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

type MetaType map[string]interface{}

// Feature holds the builtin parameters and metas fetched from user.
type Feature interface {
	Hashable

	BuiltinParam() MetaType
	FeatureMeta() MetaType
}

// FeatureFetcher supports fetch of feature, need user implementation.
type FeatureFetcher interface {
	ObservableFeatureFetcher

	// FetchFeature fetch one feature, return data type is map[string]interface{}
	FetchFeature(feat FeatureParam, dc DataCtx) (MetaType, error)
}

// FeatureParam is the parameter pass to FeatureFetcher to fetch a feature, which uniquely identifies a feature snapshot in a DataCtx.
type FeatureParam interface {
	Hashable

	// FeatureName is the name of one feature which is the first word separated from feature path in the Arishem rule expression.
	FeatureName() string
	// BuiltinParam is the parameter which is builtin rule expression.
	BuiltinParam() MetaType
}

// FeaturePreParser implement the ParseObserver and parse feature info in Arishem rule.
type FeaturePreParser interface {
	ParseObserver

	Data() []FeatureParam
	Error() error
}
