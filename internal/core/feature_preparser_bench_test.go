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
	"github.com/stretchr/testify/assert"
	"testing"
)

// parse function benchmark is to contrast feature pre parse time cost, the two time cost are almost equal
func BenchmarkCdtParse(b *testing.B) {
	cdtExpr := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}}]}}]}`
	for i := 0; i < b.N; i++ {
		_, _ = parser.ParseArishemCondition(cdtExpr)
	}
}

func BenchmarkFeaturePreParse(b *testing.B) {
	cdtExpr := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}}]}}]}`
	for i := 0; i < b.N; i++ {
		featL := NewFeaturePreParseListener()
		_, _ = parser.ParseArishemCondition(cdtExpr, featL)
		assert.Len(b, featL.Data(), 6)
	}
}
