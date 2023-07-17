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
	"github.com/bytedance/arishem/typedef"
	"sync"
)

type ObjPool struct {
	pool *sync.Pool
}

// NewObjPool construct a new object pool
func NewObjPool(new func() typedef.Poolable) *ObjPool {
	return &ObjPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new()
			},
		},
	}
}

func (o *ObjPool) Ask() typedef.Poolable {
	obj := o.pool.Get()
	if obj == nil {
		return nil
	}
	pb, ok := obj.(typedef.Poolable)
	if !ok {
		return nil
	}
	return pb
}

func (o *ObjPool) Giveback(p typedef.Poolable) {
	if p == nil {
		return
	}
	p.Reset()
	o.pool.Put(p)
}
