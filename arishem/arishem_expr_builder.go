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
	"errors"
	"math"
	"strconv"
)

type Float float64

func (f Float) MarshalJSON() ([]byte, error) {
	n := float64(f)
	if math.IsInf(n, 0) || math.IsNaN(n) {
		return nil, errors.New("unsupported number")
	}
	prec := -1
	if math.Trunc(n) == n {
		prec = 1 // Force ".0" for integers.
	}
	return strconv.AppendFloat(nil, n, 'f', prec, 64), nil
}

type Const struct {
	BoolConst *bool       `json:"BoolConst,omitempty"`
	NumConst  interface{} `json:"NumConst,omitempty"`
	StrConst  *string     `json:"StrConst,omitempty"`
}

func NewBoolConst(b bool) *Const {
	return &Const{BoolConst: &b}
}

func NewNumConst(n interface{}) *Const {
	return &Const{NumConst: n}
}

func NewStrConst(s string) *Const {
	return &Const{StrConst: &s}
}

type MathExpr struct {
	OpMath    string  `json:"OpMath"`
	Lhs       *Expr   `json:"Lhs,omitempty"`
	Rhs       *Expr   `json:"Rhs,omitempty"`
	ParamList []*Expr `json:"ParamList,omitempty"`
}

func NewLRMath(opMath string, lhs, rhs *Expr) *MathExpr {
	return &MathExpr{OpMath: opMath, Lhs: lhs, Rhs: rhs}
}

func NewParamListMath(opMath string, param []*Expr) *MathExpr {
	return &MathExpr{OpMath: opMath, ParamList: param}
}

type FuncExpr struct {
	FuncName  string           `json:"FuncName"`
	ParamMap  map[string]*Expr `json:"ParamMap,omitempty"`
	ParamList []*Expr          `json:"ParamList,omitempty"`
}

func NewFunc(funcName string) *FuncExpr {
	return &FuncExpr{FuncName: funcName}
}

func NewParamMapFunc(funcName string, param map[string]*Expr) *FuncExpr {
	return &FuncExpr{FuncName: funcName, ParamMap: param}
}

func NewParamListFunc(funcName string, param []*Expr) *FuncExpr {
	return &FuncExpr{FuncName: funcName, ParamList: param}
}

type VarExpr string

type FeatureExpr struct {
	FeaturePath  string           `json:"FeaturePath"`
	BuiltinParam map[string]*Expr `json:"BuiltinParam,omitempty"`
}

func NewFeature(featurePath string) *FeatureExpr {
	return &FeatureExpr{FeaturePath: featurePath}
}

func NewFeatureWithBuiltinParam(featurePath string, param map[string]*Expr) *FeatureExpr {
	return &FeatureExpr{FeaturePath: featurePath, BuiltinParam: param}
}

type Expr struct {
	ListExpr    []*Expr          `json:"ListExpr,omitempty"`
	MapExpr     map[string]*Expr `json:"MapExpr,omitempty"`
	Const       *Const           `json:"Const,omitempty"`
	ConstList   []*Const         `json:"ConstList,omitempty"`
	MathExpr    *MathExpr        `json:"MathExpr,omitempty"`
	FuncExpr    *FuncExpr        `json:"FuncExpr,omitempty"`
	VarExpr     *VarExpr         `json:"VarExpr,omitempty"`
	FeatureExpr *FeatureExpr     `json:"FeatureExpr,omitempty"`
}

func NewListExpr(list []*Expr) *Expr {
	return &Expr{ListExpr: list}
}

func NewMapExpr(m map[string]*Expr) *Expr {
	return &Expr{MapExpr: m}
}

func NewConstExpr(c *Const) *Expr {
	return &Expr{Const: c}
}

func NewConstListExpr(cs []*Const) *Expr {
	return &Expr{ConstList: cs}
}

func NewMathExpr(mt *MathExpr) *Expr {
	return &Expr{MathExpr: mt}
}

func NewFuncExpr(f *FuncExpr) *Expr {
	return &Expr{FuncExpr: f}
}

func NewVarExpr(v *VarExpr) *Expr {
	return &Expr{VarExpr: v}
}

func NewFeatureExpr(ft *FeatureExpr) *Expr {
	return &Expr{FeatureExpr: ft}
}
