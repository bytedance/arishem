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
	"fmt"
	"github.com/bytedance/arishem/internal/core"
	"github.com/bytedance/arishem/typedef"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type MyFeatureFetcher struct {
	lock      sync.RWMutex
	observers map[string]typedef.FeatureFetchObserver
}

func NewMyFeatureFetcher() *MyFeatureFetcher {
	return &MyFeatureFetcher{lock: sync.RWMutex{}, observers: make(map[string]typedef.FeatureFetchObserver)}
}

func (m *MyFeatureFetcher) AddFetchObserver(v ...typedef.FeatureFetchObserver) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, observer := range v {
		m.observers[observer.HashCode()] = observer
	}
}

func (m *MyFeatureFetcher) GetFetchObservers() []typedef.FeatureFetchObserver {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var obs []typedef.FeatureFetchObserver
	for _, observer := range m.observers {
		obs = append(obs, observer)
	}
	return obs
}

func (m *MyFeatureFetcher) ClearFetchObservers() {
	m.observers = make(map[string]typedef.FeatureFetchObserver)
}

func (m *MyFeatureFetcher) FetchFeature(ctx context.Context, feat typedef.FeatureParam, dc typedef.DataCtx) (typedef.MetaType, error) {
	switch feat.FeatureName() {
	case "user":
		if feat.BuiltinParam() != nil {
			time.Sleep(100 * time.Millisecond)
			if feat.BuiltinParam()["user_id"] != uint64(100) {
				panic("error param")
			}
			name, _ := dc.GetVarValue([]string{"user1", "name"})
			age, _ := dc.GetVarValue([]string{"user1", "age"})
			return map[string]interface{}{"name": name, "age": age}, nil
		} else {
			time.Sleep(200 * time.Millisecond)
			name, _ := dc.GetVarValue([]string{"user2", "name"})
			age, _ := dc.GetVarValue([]string{"user2", "age"})
			return map[string]interface{}{"name": name, "age": age}, nil
		}
	case "payment_info":
		time.Sleep(150 * time.Millisecond)
		age, _ := dc.GetFeatureValue(core.NewDefaultFeatureParamWithNoBuiltinParam("user"), "age")
		rate, _ := dc.GetVarValue([]string{"rate#2"})
		return map[string]interface{}{"account": float64(age.(int64)) * rate.(float64)}, nil
	}
	return nil, nil
}

func buildFeaturesArishemRules() []RuleTarget {
	var rt []RuleTarget
	pr, err := NewPriorityRule("p-rule3", 3, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "==",
            "Lhs": {
                "FeatureExpr": {
                    "FeaturePath": "user.age"
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 24
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule3 passed!"
    }
}`)
	if err == nil {
		rt = append(rt, pr)
	}
	pr, err = NewPriorityRule("p-rule5", 5, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
                "FeatureExpr": {
                    "FeaturePath": "user.name"
                }
            },
            "Rhs": {
                "Const": {
                    "StrConst": "Jack"
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule5 passed!"
    }
}`)
	if err == nil {
		rt = append(rt, pr)
	}
	pr, err = NewPriorityRule("p-rule4", 4, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": ">",
            "Lhs": {
				"FeatureExpr": {
                    "FeaturePath": "user.name",
					"BuiltinParam": {
						"user_id": {
							"Const": {
								"NumConst": 100
							}
						}
					}
                }
            },
            "Rhs": {
                "Const": {
                    "StrConst": "John"
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule4 passed!"
    }
}`)
	if err == nil {
		rt = append(rt, pr)
	}
	pr, err = NewPriorityRule("p-rule2", 2, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "LIST_IN",
            "Lhs": {
                "FeatureExpr": {
                    "FeaturePath": "payment_info.account"
                }
            },
            "Rhs": {
                "ConstList": [
                    {
                        "NumConst": 10
                    },
                    {
                        "NumConst": 98
                    },
                    {
                        "NumConst": 101
                    },
                    {
                        "NumConst": -1.12e3
                    }
                ]
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule2 passed!"
    }
}`)
	if err == nil {
		rt = append(rt, pr)
	}
	pr, err = NewPriorityRule("p-rule1", 1, `{
    "OpLogic": "&&",
    "Conditions": [
        {
            "Operator": "<",
            "Lhs": {
                "FeatureExpr": {
                    "FeaturePath": "payment_info.account"
                }
            },
            "Rhs": {
                "Const": {
                    "NumConst": 90
                }
            }
        }
    ]
}`, `{
    "Const": {
        "StrConst": "p-rule1 passed!"
    }
}`)
	if err == nil {
		rt = append(rt, pr)
	}
	npr, err := NewNoPriorityRule("np-rule1", `
	{
	    "OpLogic": "&&",
	    "Conditions": [
	        {
	            "Operator": "==",
	            "Lhs": {
	                "FeatureExpr": {
	                    "FeaturePath": "user.name"
	                }
	            },
	            "Rhs": {
	                "Const": {
	                    "NumConst": 1
	                }
	            }
	        }
	    ]
	}`, `{
	    "Const": {
	        "StrConst": "np-rule1 passed!"
	    }
	}`)
	if err == nil {
		rt = append(rt, npr)
	}
	npr, err = NewNoPriorityRule("np-rule2", `
	{
	    "OpLogic": "&&",
	    "Conditions": [
	        {
	            "Operator": "==",
	            "Lhs": {
	                "FeatureExpr": {
	                    "FeaturePath": "payment_info.account"
	                }
	            },
	            "Rhs": {
	                "Const": {
	                    "NumConst": 90
	                }
	            }
	        }
	    ]
	}`, `{
	    "Const": {
	        "StrConst": "np-rule2 passed!"
	    }
	}`)
	if err == nil {
		rt = append(rt, npr)
	}
	npr, err = NewNoPriorityRule("np-rule3", `
	{
	    "OpLogic": "&&",
	    "Conditions": [
	        {
	            "Operator": "==",
	            "Lhs": {
	                "FeatureExpr": {
	                    "FeaturePath": "user.name"
	                }
	            },
	            "Rhs": {
	                "Const": {
	                    "StrConst": "John"
	                }
	            }
	        }
	    ]
	}`, `{
	    "Const": {
	        "StrConst": "np-rule3 passed!"
	    }
	}`)
	if err == nil {
		rt = append(rt, npr)
	}
	return rt
}

type MyObserver struct {
	name    string
	lock    *sync.RWMutex
	timeMap map[string]time.Time
}

func NewMyObserver(name string) *MyObserver {
	return &MyObserver{name: name, lock: &sync.RWMutex{}, timeMap: make(map[string]time.Time)}
}

func (m *MyObserver) OnFeatureFetchStart(ctx context.Context, feat typedef.FeatureParam) {
	now := time.Now()
	fmt.Printf("%v feature=>%s start fetch, paramHash=>%s\n", now, feat.HashCode(), core.GenBuiltinParamHash(feat.BuiltinParam()))
	m.lock.Lock()
	defer m.lock.Unlock()
	m.timeMap[feat.HashCode()] = now
}

func (m *MyObserver) OnFeatureFetchEnd(ctx context.Context, featureHash string, featureValue typedef.MetaType, err error) {
	now := time.Now()
	fmt.Printf("%v feature=>%s end fetch\n", now, featureHash)

	m.lock.RLock()
	defer m.lock.RLock()
	fmt.Printf("feature=>%s end, cost time=>%v\n", featureHash, time.Since(m.timeMap[featureHash]))
}

func (m *MyObserver) HashCode() string {
	return m.name
}

func (m *MyObserver) OnJudgeNodeVisitEnd(ctx context.Context, info typedef.JudgeNode, vt typedef.VisitTarget) {
	fmt.Printf("rule:%s judge node passed: %v, left(%v) %s right(%v)\n", vt.Identifier(), info.Passed(), info.Left(), info.Operator(), info.Right())
}

func (m *MyObserver) OnVisitError(ctx context.Context, node, errMsg string, vt typedef.VisitTarget) {
	logger.Infof("error=>%v during node=>%v\n", errMsg, node)
}

func TestFeatureFetcher(t *testing.T) {
	rules := buildFeaturesArishemRules()
	st := time.Now()
	dc, _ := DataContext(context.Background(), `{"user1":{"name":"John","age":30},"user2":{"name":"Jane","age":18},"rate":[1.0,2.0,5.0]}`)
	observer := NewMyObserver("my observer")
	rrs := ExecuteRules(rules, dc, WithVisitObserver(observer), WithFeatureFetchObserver(observer))
	t.Logf("execute rules total cost=>%v", time.Since(st))
	ellipse := time.Since(st).Milliseconds()
	assert.NotEmpty(t, rrs)
	assert.Len(t, rrs, 2)
	for _, rr := range rrs {
		assert.Contains(t, []string{"p-rule5", "np-rule2"}, rr.Identifier())
		if rr.Identifier() == "p-rule5" {
			assert.Equal(t, "p-rule5 passed!", rr.Aim().AsExpr())
		} else {
			assert.Equal(t, "np-rule2 passed!", rr.Aim().AsExpr())
		}
	}
	// all features fetcher less than 200~210 ms
	assert.Less(t, ellipse, 210*time.Millisecond)
}
