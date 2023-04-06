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

package arishem

// NewParamMapAction create a new action aim with param map.
func NewParamMapAction(name string, param map[string]*Expr) *ActionAim {
	return &ActionAim{ActionName: name, ParamMap: param, aaType: actionAimTypeParamMap}
}

// NewParamListAction create a new action aim with param list.
func NewParamListAction(name string, param []*Expr) *ActionAim {
	return &ActionAim{ActionName: name, ParamList: param, aaType: actionAimTypeParamList}
}

type ActionAim struct {
	aaType     actionAimType
	ActionName string           `json:"ActionName"`
	ParamMap   map[string]*Expr `json:"ParamMap,omitempty"`
	ParamList  []*Expr          `json:"ParamList,omitempty"`
}

func (a *ActionAim) AddKeyParam(key string, val *Expr) {
	if key == "" {
		return
	}
	if a.aaType != actionAimTypeParamMap {
		return
	}
	a.ParamMap[key] = val
}

func (a *ActionAim) AddParam(val *Expr) {
	if a.aaType != actionAimTypeParamList {
		return
	}
	a.ParamList = append(a.ParamList, val)
}
