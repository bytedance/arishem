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

package funcs

import (
	"context"
	"testing"
)

func TestMapUnion(t *testing.T) {
	// @case 1
	params := []interface{}{
		map[string]interface{}{"name": "1"},
		map[string]interface{}{"name": "3"},
		map[string]interface{}{"name": "10", "age": 18},
	}
	union, _ := MapUnion(context.Background(), params)
	t.Log(union)
}

func TestMapIntersect(t *testing.T) {
	// @case 1
	params := []interface{}{
		map[string]interface{}{"name": "1"},
		map[string]interface{}{"name": "3"},
		map[string]interface{}{"name": "10", "age": 18},
	}
	intersect, _ := MapIntersect(context.Background(), params)
	t.Log(intersect)
}
