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

package typedef

// Identifiable is the interface that can be identified.
type Identifiable interface {
	Identifier() string
}

// Comparable is the interface that can compare with other Comparable instance.
type Comparable interface {
	Compare(other Comparable) int
}

// Poolable is the interface that can be recycled in an object pool.
type Poolable interface {
	Reset()
}

// Pool is the object pool interface.
type Pool interface {
	Ask() Poolable
	Giveback(p Poolable)
}
