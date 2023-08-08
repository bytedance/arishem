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

func SingleFilterStr(arr []string, filter string) []string {
	if arr == nil {
		return nil
	}
	res := make([]string, 0, len(arr))
	for _, s := range arr {
		if s == filter {
			continue
		}
		res = append(res, s)
	}
	return res
}

func MultiFilterStr(arr []string, filters ...string) []string {
	if arr == nil {
		return nil
	}
	res := make([]string, 0, len(arr))
	for _, s := range arr {
		if ContainsStr(filters, s) {
			continue
		}
		res = append(res, s)
	}
	return res
}

func ContainsNum(num []int64, target int64) bool {
	if num == nil {
		return false
	}
	for i := range num {
		if num[i] == target {
			return true
		}
	}
	return false
}

func ContainsStr(strs []string, target string) bool {
	if strs == nil {
		return false
	}
	for i := range strs {
		if strs[i] == target {
			return true
		}
	}
	return false
}

func DistinctStr(elements []string) []string {
	if len(elements) <= 0 {
		return elements
	}
	recorder := map[string]bool{}
	for _, element := range elements {
		recorder[element] = true
	}
	elements = make([]string, 0, len(recorder))
	for k := range recorder {
		elements = append(elements, k)
	}
	return elements
}

func SliceFilterUint64(arr []uint64, filter []uint64) []uint64 {
	if arr == nil {
		return nil
	}
	if filter == nil {
		return arr
	}
	record := make(map[uint64]bool, len(filter))
	for _, u := range filter {
		record[u] = true
	}
	res := make([]uint64, 0, len(arr))
	for _, u := range arr {
		if record[u] {
			continue
		}
		res = append(res, u)
	}
	return res
}

func BatchSliceUint64(arr []uint64, batchSize int) [][]uint64 {
	if len(arr) <= 0 || batchSize <= 0 {
		return [][]uint64{arr}
	}
	var batches [][]uint64
	for batchSize < len(arr) {
		arr, batches = arr[batchSize:], append(batches, arr[0:batchSize:batchSize])
	}
	// elements left
	batches = append(batches, arr)
	return batches
}

func BatchSliceInt64(arr []int64, batchSize int) [][]int64 {
	if len(arr) <= 0 || batchSize <= 0 {
		return [][]int64{arr}
	}
	var batches [][]int64
	for batchSize < len(arr) {
		arr, batches = arr[batchSize:], append(batches, arr[0:batchSize:batchSize])
	}
	// elements left
	batches = append(batches, arr)
	return batches
}
