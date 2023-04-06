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

package parser

import "testing"

func BenchmarkParseArishemCondition(b *testing.B) {
	cdtExpr := `{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username","BuiltinParam":{"QueryNameRange":{"ConstList":[{"StrConst":"Mike"},{"StrConst":"Jack"},{"StrConst":"Andrew"},{"StrConst":"John"}]},"AgeCalculated":{"MathExpr":{"OpMath":"+","Lhs":{"Const":{"NumConst":10}},"Rhs":{"Const":{"NumConst":10}}}}}}}]}}]}`
	for i := 0; i < b.N; i++ {
		_, _ = ParseArishemCondition(cdtExpr)
	}
}

func BenchmarkParseArishemAim(b *testing.B) {
	aimExpr := `[{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]},{"ActionName":"Greeting","ParamList":[{"FeatureExpr":{"FeaturePath":"user_profile.username"}}]}]`
	for i := 0; i < b.N; i++ {
		_, _ = ParseArishemAim(aimExpr)
	}
}
