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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNoPriorityRule(t *testing.T) {
	rule, err := NewNoPriorityRule("no-pri 1",
		`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}}]}}]}`,
		`{
    "ActionName": "Greeting",
    "ParamMap": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`)
	assert.Nil(t, err)
	assert.NotNil(t, rule)
	assert.NotNil(t, rule.ConditionPTree())
	assert.NotNil(t, rule.AimPTree())
	assert.NotEmpty(t, rule.CdtFeatParams())
	assert.NotEmpty(t, rule.AimFeatParams())
}

func TestNewNoPriorityRuleByWrongParseSyntax(t *testing.T) {
	rule, err := NewNoPriorityRule("no-pri 1",
		`{"OpLogic":"&&","Condition":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}}]}}]}`,
		`{
    "ActionName": "Greeting",
    "BuiltinParam": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`)
	assert.NotNil(t, err)
	assert.Nil(t, rule)
}

func TestTryParse(t *testing.T) {
	cdtTree, aimTree, err := tryParse(
		`{"OpLogic":"&&","Condition":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}}]}}]}`,
		`{
    "ActionName": "Greeting",
    "ParamMap": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`)
	assert.NotNil(t, err)
	t.Log(err.Error())
	assert.Nil(t, cdtTree)
	assert.NotNil(t, aimTree)
}

func TestTryParse2(t *testing.T) {
	cdtTree, aimTree, err := tryParse(
		`{"OpLogic":"&&","Condition":[{"Operator":">","Lhs":{"VarExpr":"upgrade_time"},"Rhs":{"ListExpr":[{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_profile.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}},{"FeatureExpr":{"FeaturePath":"user_history.username"}}]}}]}`,
		`{
    "ActionName": "Greeting",
    "ParamMaps": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`)
	assert.NotNil(t, err)
	assert.Nil(t, cdtTree)
	assert.Nil(t, aimTree)
	t.Log(err.Error())
}

func TestNewNoPriorityRuleCacheHit(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, _ = NewNoPriorityRule("no-pri 1",
			`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
			`{
    "ActionName": "Greeting",
    "BuiltinParam": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`)
	}
}

func TestNewPriorityRule(t *testing.T) {
	rule, err := NewPriorityRule("pri 1", 100,
		`{"OpLogic":"&&","Conditions":[{"Operator":">","Lhs":{"Const":{"NumConst":100}},"Rhs":{"Const":{"NumConst":10}}}]}`,
		`{
    "ActionName": "Greeting",
    "ParamMap": {
        "UserName": {
            "FeatureExpr": {
                "FeaturePath": "user_profile.username"
            }
        }
    }
}`,
	)
	assert.Nil(t, err)
	assert.NotNil(t, rule)
	assert.NotNil(t, rule.ConditionPTree())
	assert.NotNil(t, rule.AimPTree())
	assert.NotEmpty(t, rule.AimFeatParams())
}
