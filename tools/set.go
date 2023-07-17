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
	"github.com/bytedance/arishem/typedef"
	"sync"
)

type StringSet struct {
	strMap map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{}
}

func (ss *StringSet) Add(str string) {
	ss.strMap[str] = true
}

func (ss *StringSet) Contains(str string) bool {
	_, ok := ss.strMap[str]
	return ok
}

func (ss *StringSet) Remove(str string) {
	delete(ss.strMap, str)
}

func (ss *StringSet) AsList() []string {
	l := make([]string, 0, len(ss.strMap))
	for str := range ss.strMap {
		l = append(l, str)
	}
	return l
}

// ConcurrentHashSet contains some distinct elements and concurrent safe
type ConcurrentHashSet struct {
	*HashSet
	dataLocker *sync.RWMutex
}

func NewConcurrentHashSet() *ConcurrentHashSet {
	return &ConcurrentHashSet{HashSet: NewHashSet(), dataLocker: &sync.RWMutex{}}
}

// HashSet contains some distinct elements
type HashSet struct {
	dataMap map[string]typedef.Hashable
}

func NewHashSet() *HashSet {
	return &HashSet{dataMap: make(map[string]typedef.Hashable, 8)}
}

// Add an element into set
func (s *HashSet) Add(h typedef.Hashable) {
	if IsNil(h) {
		return
	}
	k := h.HashCode()

	s.dataMap[k] = h
}

func (s *HashSet) Contains(h typedef.Hashable) bool {
	if IsNil(h) {
		return false
	}
	flag := false
	k := h.HashCode()
	_, flag = s.dataMap[k]
	return flag
}

func (s *HashSet) Remove(h typedef.Hashable) {
	if IsNil(h) {
		return
	}
	k := h.HashCode()
	delete(s.dataMap, k)
}

func (s *HashSet) AsList() []typedef.Hashable {
	data := make([]typedef.Hashable, 0, 8)
	for _, value := range s.dataMap {
		data = append(data, value)
	}
	return data
}

// Add an element into set, no use defer to improve performance
func (s *ConcurrentHashSet) Add(h typedef.Hashable) {
	if IsNil(h) {
		return
	}
	k := h.HashCode()

	// insert element
	s.dataLocker.Lock()
	s.dataMap[k] = h
	s.dataLocker.Unlock()
}

func (s *ConcurrentHashSet) Contains(h typedef.Hashable) bool {
	if IsNil(h) {
		return false
	}
	flag := false
	k := h.HashCode()
	s.dataLocker.RLock()
	_, flag = s.dataMap[k]
	s.dataLocker.RUnlock()
	return flag
}

func (s *ConcurrentHashSet) Remove(h typedef.Hashable) {
	if IsNil(h) {
		return
	}
	k := h.HashCode()
	s.dataLocker.Lock()
	delete(s.dataMap, k)
	s.dataLocker.Unlock()
}

func (s *ConcurrentHashSet) AsList() []typedef.Hashable {
	data := make([]typedef.Hashable, 0, 8)
	s.dataLocker.RLock()
	for _, value := range s.dataMap {
		data = append(data, value)
	}
	s.dataLocker.RUnlock()
	return data
}
