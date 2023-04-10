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

package pool

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewWorkPool(t *testing.T) {
	const (
		num       = 64
		submitNum = num * 10
	)

	var peekGrtCount int32 = 0
	pool := NewWorkPool(num, 30*time.Second)
	wg := sync.WaitGroup{}
	wg.Add(submitNum)
	for i := 0; i < submitNum; i++ {
		time.Sleep(time.Duration(rand.Intn(3)))
		pool.Submit(func(param interface{}) {
			wg.Done()
		}, nil)
		peekGrtCount = int32(math.Max(float64(peekGrtCount), float64(atomic.LoadInt32(&pool.size))))
	}
	wg.Wait()
	assert.Less(t, peekGrtCount, int32(submitNum))
	t.Log(peekGrtCount)
}
