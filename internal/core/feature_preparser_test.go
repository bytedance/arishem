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
	"github.com/bytedance/arishem/internal/parser"
	"github.com/bytedance/arishem/tools"
	"github.com/bytedance/arishem/typedef"
	"github.com/dgraph-io/ristretto"
	"github.com/stretchr/testify/assert"
	"testing"
)

type exprType int

const (
	exprTypeCondition = iota
	exprTypeAim
)

type fc struct {
	c *ristretto.Cache
}

func (f *fc) Get(key interface{}) (val interface{}, ok bool) {
	return f.c.Get(key)
}

func (f *fc) Set(key interface{}, val interface{}) {
	if tools.IsNil(key) {
		return
	}
	f.c.Set(key, val, 1)
}

func init() {
	c, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters:        10 * 256,
		MaxCost:            256,
		BufferItems:        64,
		IgnoreInternalCost: true,
	})
	SetFeaturePreParseCache(&fc{c: c})
}

func TestFeatPreParser(t *testing.T) {
	testCases := []struct {
		name  string
		et    exprType
		args  string
		check func(featParams []typedef.FeatureParam)
	}{
		{
			"CONDITION: no feature",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"VarExpr":"user.username"}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 0)
			},
		},
		{
			"CONDITION: no param map single feature",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"FeatureExpr":{"FeaturePath":"user.username"}}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 1)
				fp := featParams[0]
				assert.NotNil(t, featParams[0])
				assert.Equal(t, "user", fp.FeatureName())
				assert.Nil(t, fp.BuiltinParam())
			},
		},
		{
			"CONDITION: no param map multi same features",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}}]}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 1)
				fp := featParams[0]
				assert.NotNil(t, featParams[0])
				assert.Equal(t, "user", fp.FeatureName())
				assert.Nil(t, fp.BuiltinParam())
			},
		},
		{
			"CONDITION: no param map multi mix features",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}}]}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 3)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					assert.Contains(t, []string{"user", "user_profile", "user_history"}, fp.FeatureName())
					assert.Empty(t, fp.BuiltinParam())
					assert.Nil(t, fp.BuiltinParam())
				}
			},
		},
		{
			"CONDITION: param map single feature",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":"LIST_IN","Lhs":{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},"Rhs":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 1)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					assert.Equal(t, "user", fp.FeatureName())
					assert.NotEmpty(t, fp.BuiltinParam())
					assert.NotNil(t, fp.BuiltinParam())
					assert.Len(t, fp.BuiltinParam(), 2)
					assert.Equal(t, []interface{}{"Mike", "Jack", "Andrew", "John"}, fp.BuiltinParam()["QueryNameRange"])
					assert.Equal(t, float64(20), fp.BuiltinParam()["AgeCalculated"])
				}
			},
		},
		{
			"CONDITION: param map multi mix features",
			exprTypeCondition,
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}}]}}]}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 6)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					if fp.BuiltinParam() != nil {
						assert.Contains(t, []string{"user", "user_profile", "user_history"}, fp.FeatureName())
						assert.NotNil(t, fp.BuiltinParam())
						assert.Len(t, fp.BuiltinParam(), 2)
						assert.Equal(t, []interface{}{"Mike", "Jack", "Andrew", "John"}, fp.BuiltinParam()["QueryNameRange"])
						assert.Equal(t, float64(20), fp.BuiltinParam()["AgeCalculated"])
						t.Logf("%v", fp)
					}
				}
			},
		},
		{
			"AIM: action aim param map with single no param feature",
			exprTypeAim,
			`{"ActionName":"Greeting","ParamMap":{"UserName":{"FeatureExpr":{"FeaturePath":"user_profile.username"}}}}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 1)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					assert.Contains(t, "user_profile", fp.FeatureName())
					assert.Empty(t, fp.BuiltinParam())
				}
			},
		},
		{
			"AIM: multi action aim, one with param map feature",
			exprTypeAim,
			`[{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]},{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username","BuiltinParam":{"user_id":{"Const":{"NumConst":1234567}}}}}]}]`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 2)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					if fp.BuiltinParam() != nil {
						assert.Equal(t, "user_profile", fp.FeatureName())
						assert.NotNil(t, fp.BuiltinParam())
						assert.Len(t, fp.BuiltinParam(), 1)
						assert.Equal(t, uint64(1234567), fp.BuiltinParam()["user_id"])
					}
				}
			},
		},
		{
			"AIM: expr aim with MapExpr",
			exprTypeAim,
			`{"MapExpr":{"UserName":{"VarExpr":"user.username"},"UserAge":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":8}}}},"UserWeight":{"MathExpr":{"OpMath":"+","ParamList":[{"Const":{"NumConst":100}},{"Const":{"NumConst":20}}]}},"UserPersonality":{"FeatureExpr":{"FeaturePath":"user_center.user_personality","BuiltinParam":{"user_id":{"Const":{"NumConst":1234567}}}}}}}`,
			func(featParams []typedef.FeatureParam) {
				assert.Len(t, featParams, 1)
				for _, param := range featParams {
					assert.NotNil(t, param)
					fp := param

					assert.Equal(t, "user_center", fp.FeatureName())
					assert.NotNil(t, fp.BuiltinParam())
					assert.Len(t, fp.BuiltinParam(), 1)
					assert.Equal(t, uint64(1234567), fp.BuiltinParam()["user_id"])
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Log(tc.name)
		fppl := NewFeaturePreParseListener()
		if tc.et == exprTypeCondition {
			_, err := parser.ParseArishemCondition(tc.args, fppl)
			assert.Nil(t, err)
			tc.check(fppl.Data())
		} else if tc.et == exprTypeAim {
			_, err := parser.ParseArishemAim(tc.args, fppl)
			assert.Nil(t, err)
			tc.check(fppl.Data())
		}
	}
}
