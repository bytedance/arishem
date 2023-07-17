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

package tools

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

type Int int64

func (i Int) HashCode() string {
	return strconv.FormatInt(int64(i), 10)
}

func TestNewHashSet(t *testing.T) {
	set := NewHashSet()
	set.Add(Int(1))
	set.Add(Int(1))
	set.Add(Int(1))
	set.Add(Int(1))
	assert.True(t, set.Contains(Int(1)))
	assert.Equal(t, 1, len(set.AsList()))
	assert.Equal(t, Int(1), set.AsList()[0].(Int))
}

func TestNewConcurrentHashSet(t *testing.T) {
	set := NewConcurrentHashSet()
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			set.Add(Int(1))
			wg.Done()
		}()
	}
	wg.Wait()
	assert.True(t, set.Contains(Int(1)))
	assert.Equal(t, 1, len(set.AsList()))
	assert.Equal(t, Int(1), set.AsList()[0].(Int))
}
