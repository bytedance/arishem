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

package funcs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateToUnix(t *testing.T) {
	// @case 1
	params := map[string]interface{}{
		"date":      "2023-01-02",
		"precision": "day",
		"zone":      "local",
		"unit":      "sec",
	}
	unix := DateToUnix(params)
	t.Log(unix)
	// @case 2
	params["zone"] = "utc"
	unix = DateToUnix(params)
	t.Log(unix)
	// @case 3
	params["zone"] = "local"
	unix = DateToUnix(params)
	t.Log(unix)
	// @case 4
	params["unit"] = "milli"
	unix = DateToUnix(params)
	t.Log(unix)
	// @case 5
	params["unit"] = "micro"
	unix = DateToUnix(params)
	t.Log(unix)
	// @case 6
	params["unit"] = "nano"
	unix = DateToUnix(params)
	t.Log(unix)
}

func TestUnixToDate(t *testing.T) {
	// @case 1
	params := map[string]interface{}{
		"date":   "1672655888",
		"format": "2006-01-02",
	}
	unix := UnixToDate(params)
	assert.Equal(t, "2023-01-02", unix)
	// @case 2
	params["format"] = "2006-01-02 15:04"
	unix = UnixToDate(params)
	assert.Equal(t, "2023-01-02 18:38", unix)
	// @case 3
	params["format"] = "2006-01-02 15:04:05"
	unix = UnixToDate(params)
	assert.Equal(t, "2023-01-02 18:38:08", unix)
}

func TestGetCurrentYear(t *testing.T) {
	year := GetCurrentYear()
	assert.Equal(t, int64(2023), year)
}

/*func TestGetCurrentYearDay(t *testing.T) {
	day := GetCurrentYearDay()
	assert.Equal(t, int64(4), day)
}

func TestGetCurrentMonth(t *testing.T) {
	month := GetCurrentMonth()
	assert.Equal(t, int64(1), month)
}

func TestGetCurrentMonthDay(t *testing.T) {
	day := GetCurrentMonthDay()
	assert.Equal(t, int64(4), day)
}

func TestGetCurrentWeekDay(t *testing.T) {
	day := GetCurrentWeekDay()
	assert.Equal(t, "Wednesday", day)
}

func TestGetCurrentHour(t *testing.T) {
	hour := GetCurrentHour()
	assert.Equal(t, int64(18), hour)
}

func TestGetCurrentMinute(t *testing.T) {
	minute := GetCurrentMinute()
	assert.Equal(t, int64(36), minute)
}*/

func TestGetCurrentSecond(t *testing.T) {
	second := GetCurrentSecond()
	t.Log(second)
}
