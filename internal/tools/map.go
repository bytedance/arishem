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

package tools

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// OnceMap will only insert the same key once in a heavy concurrent environment.
type OnceMap struct {
	storage map[string]interface{}
	locker  sync.RWMutex
}

func NewOnceMap() *OnceMap {
	return &OnceMap{storage: make(map[string]interface{}, 8)}
}

func (o *OnceMap) WriteOnce(key string, value interface{}) bool {
	if len(key) <= 0 || value == nil {
		return false
	}
	// try reading it first
	v := o.Read(key)
	if v != nil {
		return false
	}
	// try to write data
	o.locker.Lock()
	// try reading it again
	v = o.storage[key]
	if v != nil {
		o.locker.Unlock()
		return false
	}
	// write it into map
	o.storage[key] = value
	// release lock
	o.locker.Unlock()
	return true
}

func (o *OnceMap) Read(key string) interface{} {
	if len(key) <= 0 {
		return nil
	}
	o.locker.RLock()
	val := o.storage[key]
	o.locker.RUnlock()
	return val
}

// MapIndex by Arishem JSON PATH, which type is: a.b.c and use '#' to indicate array index, array index must start from 0.
// The key type of MAP can only be string type, and the value type can only be type-map[string]interface{} or type-interface{} or type-[]interface{}.
// example
// {
//    "name": {
//        "first": "Tom",
//        "last": "Anderson"
//    },
//    "age": 37,
//    "children": [
//        {
//            "name": "Sara",
//            "age": 4
//        },
//        {
//            "name": "Alex",
//            "age": 8
//        },
//        {
//            "name": "Jack",
//            "age": 10
//        }
//    ]
// }
// if you want to get first name, PATH will be 'name.first', and if you want to get his children Sara's name
// PATH will be 'children#0.name'.
// benchmark: half performance of sonic Ask, but can get error message and with int64 type correctly
func MapIndex(kv map[string]interface{}, path string) (interface{}, error) {
	return MapIndexPaths(kv, strings.Split(path, "."))
}

func MapIndexPaths(kv map[string]interface{}, paths []string) (v interface{}, err error) {
	if len(kv) <= 0 {
		return
	}
	if len(paths) <= 0 { // get itself
		v = kv
		return
	}
	iterVal := kv
	for i, size := 0, len(paths)-1; i < size; i++ {
		path := paths[i]
		var val interface{}
		val, err = innerMapIndex(iterVal, path)
		if err != nil {
			return
		}
		if val == nil {
			err = fmt.Errorf("can not find the previous collection by key=>%s, path=>%s", path, strings.Join(paths, "."))
			return
		}
		switch valT := val.(type) {
		case map[string]interface{}:
			iterVal = valT
		default:
			err = fmt.Errorf("previous collection type not map, previous key=>%s, path=>%s", path, strings.Join(paths, "."))
			return
		}
	}
	v, err = innerMapIndex(iterVal, paths[len(paths)-1])
	return
}

func innerMapIndex(d map[string]interface{}, path string) (v interface{}, err error) {
	if len(d) <= 0 {
		err = errors.New("read data from empty map")
		return
	}
	if IsBlank(path) {
		err = errors.New("path is empty")
		return
	}
	// whether array indexing
	arrIdx := -1
	sIdx := strings.Index(path, "#")
	if sIdx > 0 {
		arrIdx, err = strconv.Atoi(path[sIdx+1:])
		if err != nil {
			err = fmt.Errorf("parse index error by key=>%s, err=>%s", path, err.Error())
			return
		}
		path = path[:sIdx]
	}
	target, ok := d[path]
	if !ok {
		return
	}
	v = target
	if arrIdx < 0 {
		return
	}
	// use array indexing
	switch val := target.(type) {
	case []interface{}:
		if len(val) <= arrIdx {
			err = fmt.Errorf("array indexing out of bound, key=>%s", path)
		} else {
			v = val[arrIdx]
		}
	default:
		err = fmt.Errorf("wrong indexing: array index used for type=>%T by key=>%s", target, path)
	}
	return
}
