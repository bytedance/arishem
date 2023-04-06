/*
 *  Copyright 2023 ByteDance Inc.
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
	"github.com/bytedance/arishem/internal/tools"
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/bytedance/gopkg/util/xxhash3"
	"github.com/bytedance/sonic"
	"strconv"
	"strings"
)

var (
	emptyParamMapHash = strconv.FormatUint(xxhash3.HashString("{}"), 10)
	featHashSplit     = "$$"
)

func NewDefaultFeatureParamWithNoBuiltinParam(featureName string) typedef.FeatureParam {
	return newArishemFeatureParam(featureName, nil)
}

func NewDefaultFeatureParamWithBuiltinParam(featureName string, builtinParam typedef.MetaType) typedef.FeatureParam {
	return newArishemFeatureParam(featureName, builtinParam)
}

func DecodeFeatureHash(featureHash string) (featureName string, builtinParamHash string) {
	keys := strings.Split(featureHash, featHashSplit)
	switch {
	case len(keys) == 1:
		featureName = keys[0]
	case len(keys) > 1:
		featureName = keys[0]
		builtinParamHash = keys[1]
	}
	return
}

func GenBuiltinParamHash(builtinParam typedef.MetaType) string {
	builtinParamStr := emptyParamMapHash
	if !tools.IsNil(builtinParam) {
		builtinParamStr, _ = sonic.MarshalString(builtinParamStr)
		if builtinParamStr == "" {
			return emptyParamMapHash
		}
	}
	switch builtinParamStr {
	case "", "{}", "null":
		return emptyParamMapHash
	}
	return strconv.FormatUint(xxhash3.HashString(builtinParamStr), 10)
}

type arishemFeature struct {
	hashCode     string
	builtinParam typedef.MetaType
	meta         typedef.MetaType
}

func newArishemFeature(hashCode string, builtinParam typedef.MetaType, meta typedef.MetaType) *arishemFeature {
	return &arishemFeature{hashCode: hashCode, builtinParam: builtinParam, meta: meta}
}

func (a *arishemFeature) FeatureName() string {
	name, _ := DecodeFeatureHash(a.hashCode)
	return name
}

func (a *arishemFeature) HashCode() string {
	return a.hashCode
}

func (a *arishemFeature) BuiltinParam() typedef.MetaType {
	return a.builtinParam
}

func (a *arishemFeature) FeatureMeta() typedef.MetaType {
	return a.meta
}

type arishemFeatureParam struct {
	hashCode         string
	featureName      string
	builtinParam     typedef.MetaType
	builtinParamHash string
}

func newArishemFeatureParam(featureName string, builtinParam typedef.MetaType) typedef.FeatureParam {
	afp := new(arishemFeatureParam)
	afp.featureName = featureName
	afp.builtinParam = builtinParam
	afp.builtinParamHash = GenBuiltinParamHash(builtinParam)
	afp.hashCode = afp.featureName + featHashSplit + afp.builtinParamHash
	return afp
}

func (a *arishemFeatureParam) HashCode() string {
	return a.hashCode
}

func (a *arishemFeatureParam) FeatureName() string {
	return a.featureName
}

func (a *arishemFeatureParam) BuiltinParam() typedef.MetaType {
	return a.builtinParam
}
