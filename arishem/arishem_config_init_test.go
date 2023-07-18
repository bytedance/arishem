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
	"github.com/bytedance/arishem/typedef"
)

func init() {
	Initialize(
		DefaultConfiguration(),
		WithFeatureFetcherFactory(func() typedef.FeatureFetcher {
			return NewMyFeatureFetcher()
		}),
		WithCustomNoParamFuncs(NoParamFnPair{name: "NPHello", fn: MyCustomNoFuncHelloArishem}),
		WithCustomMapParamFuncs(MapParamFnPair{name: "MPHello", fn: MyCustomMapFuncHelloArishem}),
		WithCustomListParamFuncs(ListParamFnPair{name: "LPHello", fn: MyCustomListFuncHelloArishem}),
	)
}

func MyCustomNoFuncHelloArishem() interface{} {
	return "Hello Arishem"
}

func MyCustomMapFuncHelloArishem(param map[string]interface{}) interface{} {
	return "Hello Arishem"
}

func MyCustomListFuncHelloArishem(param []interface{}) interface{} {
	return "Hello Arishem"
}
