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
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
)

func init() {
	Initialize(
		DefaultConfiguration(),
		WithFeatureFetcherFactory(func() typedef.FeatureFetcher {
			return NewMyFeatureFetcher()
		}),
		WithCustomNoParamFuncs(NoParamFnPair{Name: "NPHello", Fn: MyCustomNoFuncHelloArishem}),
		WithCustomMapParamFuncs(MapParamFnPair{Name: "MPHello", Fn: MyCustomMapFuncHelloArishem}),
		WithCustomListParamFuncs(ListParamFnPair{Name: "LPHello", Fn: MyCustomListFuncHelloArishem}),
		WithEnableSubCondition("OnlyPriceMatch", `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"price"},"Rhs":{"Const":{"NumConst":10}}}]}`),
	)
}

func MyCustomNoFuncHelloArishem(ctx context.Context) (interface{}, error) {
	return "Hello Arishem", nil
}

func MyCustomMapFuncHelloArishem(ctx context.Context, param map[string]interface{}) (interface{}, error) {
	return "Hello " + tools.ConvToUnifiedStringType(param["name"]), nil
}

func MyCustomListFuncHelloArishem(ctx context.Context, param []interface{}) (interface{}, error) {
	return "Hello " + tools.ConvToUnifiedStringType(param[0]), nil
}
