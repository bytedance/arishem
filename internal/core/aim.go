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
	"github.com/bytedance/arishem/typedef"
)

type aimType int32
type paramType int32

const (
	aimTypeActionList aimType = iota
	aimTypeAction
	aimTypeExpr
)

const (
	paramTypeMap paramType = iota
	paramTypeList
)

type arishemActionAim struct {
	name  string
	param interface{}
	pType paramType
}

func (a *arishemActionAim) ActionName() string {
	return a.name
}

func (a *arishemActionAim) Parameter() interface{} {
	return a.param
}

func newArishemActionAim(actionName string, parameter interface{}, pType paramType) typedef.ActionAim {
	return &arishemActionAim{name: actionName, param: parameter, pType: pType}
}

func (a *arishemActionAim) ParamAsMap() (paramMap map[string]interface{}) {
	if a.pType == paramTypeMap && a.param != nil {
		paramMap = a.param.(map[string]interface{})
	}
	return
}

func (a *arishemActionAim) ParamAsList() (paramList []interface{}) {
	if a.pType == paramTypeList && a.param != nil {
		paramList = a.param.([]interface{})
	}
	return
}

type arishemAim struct {
	aim   interface{}
	aType aimType
}

func newArishemAim(aim interface{}, aType aimType) typedef.Aim {
	return &arishemAim{aim: aim, aType: aType}
}

func (a *arishemAim) AsExpr() interface{} {
	if a.aType == aimTypeExpr && a.aim != nil {
		return a.aim
	}
	return nil
}

func (a *arishemAim) AsActionList() []typedef.ActionAim {
	if a.aType == aimTypeActionList && a.aim != nil {
		return a.aim.([]typedef.ActionAim)
	}
	return nil
}

func (a *arishemAim) AsAction() typedef.ActionAim {
	if a.aType == aimTypeAction && a.aim != nil {
		return a.aim.(typedef.ActionAim)
	}
	return nil
}
