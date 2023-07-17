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

import "time"

// TimeString2UnixLocal Calculate the unix value of timeStr in the local time zone.
func TimeString2UnixLocal(timeStr string) (int64, error) {
	parse, err := time.ParseInLocation(ReadableTimeFmtSec, timeStr, time.Local)
	if err != nil {
		parse, err = time.ParseInLocation(ReadableTimeFmtTCST, timeStr, time.Local)
		if err != nil {
			return 0, err
		}
	}
	return parse.Unix(), nil
}

// TimeString2UnixLocalLayout Calculate the unix value of timeStr in the local time zone, the format of timeStr is layout.
func TimeString2UnixLocalLayout(layout, timeStr string) (int64, error) {
	parse, err := time.ParseInLocation(layout, timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return parse.Unix(), nil
}

// TimeString2Local Calculate the unix value of timeStr in the local time zone.
func TimeString2Local(timeStr string) (time.Time, error) {
	parse, err := time.ParseInLocation(ReadableTimeFmtSec, timeStr, time.Local)
	if err != nil {
		parse, err = time.ParseInLocation(ReadableTimeFmtTCST, timeStr, time.Local)
		if err != nil {
			parse, err = time.ParseInLocation(ReadableTimeFmtT, timeStr, time.Local)
			if err != nil {
				return time.Now(), err
			}
		}
	}
	return parse, nil
}
