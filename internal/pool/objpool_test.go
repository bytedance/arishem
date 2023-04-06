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

package pool

import (
	"github.com/bytedance/arishem/internal/typedef"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type poolableT struct {
}

func (p *poolableT) Reset() {}

func TestNewObjPool(t *testing.T) {
	const getNum = 1 << 12
	var newCount int32 = 0
	pool := NewObjPool(func() typedef.Poolable {
		atomic.AddInt32(&newCount, 1)
		return new(poolableT)
	})
	wg := sync.WaitGroup{}
	wg.Add(getNum)
	for i := 0; i < getNum; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3)))

			p := pool.Ask().(*poolableT)
			defer pool.Giveback(p)

			time.Sleep(time.Duration(rand.Intn(3)))
		}()
	}
	wg.Wait()
	assert.Less(t, newCount, int32(getNum))
	t.Log(newCount)
}

func BenchmarkObjPool(b *testing.B) {
	pool := NewObjPool(func() typedef.Poolable {
		return new(poolableT)
	})
	for i := 0; i < b.N; i++ {
		ask := pool.Ask()
		pool.Giveback(ask)
	}
}

func BenchmarkNoObjPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = new(poolableT)
	}
}
