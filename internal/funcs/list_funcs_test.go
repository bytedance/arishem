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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinFlatWithInt(t *testing.T) {
	ctx := context.Background()
	// @case 1
	minval, _ := MinFlatWithInt(ctx, []interface{}{1, 2, 3})
	assert.Equal(t, int64(1), minval)
	// @case 2
	minval, _ = MinFlatWithInt(ctx, []interface{}{[]int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(1), minval)
	// @case 3
	minval, _ = MinFlatWithInt(ctx, []interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(0), minval)
}

func TestMinFlatWithFloat(t *testing.T) {
	ctx := context.Background()
	minval, _ := MinFlatWithFloat(ctx, []interface{}{1, 2, 3})
	assert.Equal(t, float64(1), minval)
	// @case 2
	minval, _ = MinFlatWithFloat(ctx, []interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, float64(0.1), minval)
}

func TestMaxFlatWithInt(t *testing.T) {
	ctx := context.Background()
	maxval, _ := MaxFlatWithInt(ctx, []interface{}{1, 2, 3})
	assert.Equal(t, int64(3), maxval)
	// @case 2
	maxval, _ = MaxFlatWithInt(ctx, []interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3})
	assert.Equal(t, int64(3), maxval)
}

func TestMaxFlatWithFloat(t *testing.T) {
	ctx := context.Background()
	maxval, _ := MaxFlatWithFloat(ctx, []interface{}{1, 2, 3})
	assert.Equal(t, float64(3), maxval)
	// @case 2
	maxval, _ = MaxFlatWithFloat(ctx, []interface{}{[]float64{0.1}, []int64{1}, []float64{2}, 3.5})
	t.Log(maxval)
	assert.Equal(t, float64(3.5), maxval)
}

func TestListFlatJoin(t *testing.T) {
	ctx := context.Background()
	// @case 1
	join, _ := ListFlatJoin(ctx, []interface{}{"-", 1, 2, 3, []float64{1.23, 2.45}, "zhangsan"})
	assert.Equal(t, "1-2-3-1.23-2.45-zhangsan", join)
	// @case 2
}

func TestListJoin(t *testing.T) {
	ctx := context.Background()
	params := map[string]interface{}{
		"list": []float64{1.23, 2.45},
		"sep":  ",",
	}
	// @case 1
	join, _ := ListJoin(ctx, params)
	assert.Equal(t, "1.23,2.45", join)
	// @case 2
	params["list"] = []int64{1, 23, 45}
	join, _ = ListJoin(ctx, params)
	assert.Equal(t, "1,23,45", join)
	// @case 3
	params["sep"] = "-"
	join, _ = ListJoin(ctx, params)
	assert.Equal(t, "1-23-45", join)
}

func TestListRemove(t *testing.T) {
	ctx := context.Background()
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	remove, _ := ListRemove(ctx, params)
	t.Log(remove)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	remove, _ = ListRemove(ctx, params)
	t.Log(remove)
}

func TestListUnion(t *testing.T) {
	ctx := context.Background()
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	union, _ := ListUnion(ctx, params)
	t.Log(union)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	union, _ = ListUnion(ctx, params)
	t.Log(union)
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
		[]string{"zhang", "san"},
	}
	union, _ = ListUnion(ctx, params)
	t.Log(union)
}

func TestListIntersect(t *testing.T) {
	ctx := context.Background()
	// @case 1
	params := []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0},
	}
	intersect, _ := ListIntersect(ctx, params)
	t.Log(intersect)
	// @case 2
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
	}
	intersect, _ = ListIntersect(ctx, params)
	t.Log(intersect)
	// @case 3
	params = []interface{}{
		[]int64{1, 2, 3},
		[]float64{1.0, 2.0, 3.1},
		[]string{"1", "zhang", "san"},
		[]string{"1", "zhang", "gou"},
	}
	intersect, _ = ListIntersect(ctx, params)
	t.Log(intersect)
	// @case 4
	params = []interface{}{
		[]string{"1", "zhang", "san"},
		[]string{"1", "zhang", "gou"},
	}
	intersect, _ = ListIntersect(ctx, params)
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
	intersect, _ = ListIntersect(ctx, params)
	t.Log(intersect)
}
