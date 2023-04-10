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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinFlatWithInt(t *testing.T) {
	// @case 1
	minval := MinFlatWithInt([]interface{}{1, 2, 3})
	assert.Equal(t, int64(1), minval)
	// @case 2
	minval = MinFlatWithInt([]interface{}{[]int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(1), minval)
	// @case 3
	minval = MinFlatWithInt([]interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(0), minval)
}

func TestMinFlatWithFloat(t *testing.T) {
	minval := MinFlatWithFloat([]interface{}{1, 2, 3})
	assert.Equal(t, float64(1), minval)
	// @case 2
	minval = MinFlatWithFloat([]interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, float64(0.1), minval)
}

func TestMaxFlatWithInt(t *testing.T) {
	maxval := MaxFlatWithInt([]interface{}{1, 2, 3})
	assert.Equal(t, int64(3), maxval)
	// @case 2
	maxval = MaxFlatWithInt([]interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(3), maxval)
}

func TestMaxFlatWithFloat(t *testing.T) {
	maxval := MaxFlatWithFloat([]interface{}{1, 2, 3})
	assert.Equal(t, float64(3), maxval)
	// @case 2
	maxval = MaxFlatWithFloat([]interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3.5})
	t.Log(maxval)
	assert.Equal(t, float64(3.5), maxval)
}

func TestListFlatJoin(t *testing.T) {
	// @case 1
	join := ListFlatJoin([]interface{}{"-", 1, 2, 3, []float64{1.23, 2.45}, "zhangsan"})
	assert.Equal(t, "1-2-3-1.23-2.45-zhangsan", join)
	// @case 2
}

func TestListJoin(t *testing.T) {
	params := map[string]interface{}{
		"list": []float64{1.23, 2.45},
		"sep":  ",",
	}
	// @case 1
	join := ListJoin(params)
	assert.Equal(t, "1.23,2.45", join)
	// @case 2
	params["list"] = []int64{1, 23, 45}
	join = ListJoin(params)
	assert.Equal(t, "1,23,45", join)
	// @case 3
	params["sep"] = "-"
	join = ListJoin(params)
	assert.Equal(t, "1-23-45", join)
}

func TestListRemove(t *testing.T) {
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	remove := ListRemove(params)
	t.Log(remove)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	remove = ListRemove(params)
	t.Log(remove)
}

func TestListUnion(t *testing.T) {
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	union := ListUnion(params)
	t.Log(union)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	union = ListUnion(params)
	t.Log(union)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
		[]string{"zhang", "san"},
	}
	union = ListUnion(params)
	t.Log(union)
}

func TestListIntersect(t *testing.T) {
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	intersect := ListIntersect(params)
	t.Log(intersect)
	// @case 2
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	intersect = ListIntersect(params)
	t.Log(intersect)
	// @case 3
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
		[]string{"1", "zhang", "san"},
		[]string{"1", "zhang", "gou"},
	}
	intersect = ListIntersect(params)
	t.Log(intersect)
	// @case 4
	params = []interface{}{
		[]string{"1", "zhang", "san"},
		[]string{"1", "zhang", "gou"},
	}
	intersect = ListIntersect(params)
	t.Log(intersect)
	// @case 5
	type nameAge struct {
		Name string
		Age  int
	}
	params = []interface{}{
		[]*nameAge{{Name: "zhangsan", Age: 18}, {Name: "lisi", Age: 20}},
		[]*nameAge{{Name: "zhangsa", Age: 18}, {Name: "lisi", Age: 20}},
	}
	intersect = ListIntersect(params)
	t.Log(intersect)
}
