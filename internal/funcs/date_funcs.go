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
	"fmt"
	"github.com/bytedance/arishem/internal/tools"
	"time"
)

func GetCurrentYear() interface{} {
	return int64(time.Now().Year())
}

func GetCurrentYearDay() interface{} {
	return int64(time.Now().YearDay())
}

func GetCurrentMonth() interface{} {
	return int64(time.Now().Month())
}

func GetCurrentMonthDay() interface{} {
	return int64(time.Now().Day())
}

func GetCurrentWeekDay() interface{} {
	return time.Now().Weekday().String()
}

// GetCurrentHour 获取当前的小时数
func GetCurrentHour() interface{} {
	return int64(time.Now().Hour())
}

func GetCurrentMinute() interface{} {
	return int64(time.Now().Minute())
}

func GetCurrentSecond() interface{} {
	return int64(time.Now().Second())
}

func GetCurrentLocation() interface{} {
	return time.Now().Location().String()
}

func GetUnixSecond() interface{} {
	return time.Now().Unix()
}

func DateToUnix(params map[string]interface{}) interface{} {
	const (
		keyDate      = "date"
		keyPrecision = "precision"
		keyZone      = "zone"
		keyUnit      = "unit"
	)
	const (
		datePrecisionDay   = "day"
		datePrecisionMin   = "min"
		datePrecisionSec   = "sec"
		datePrecisionMilli = "milli"
	)
	const (
		zoneTypeLocal    = "local"
		zoneTypeUTC      = "utc"
		zoneTypeShanghai = "shanghai"
	)
	const (
		unitSec   = "sec"
		unitMilli = "milli"
		unitMicro = "micro"
		unitNano  = "nano"
	)
	dateStr, ok := params[keyDate]
	if !ok {
		return time.Now().Unix()
	}
	precision := "sec"
	p, ok := params[keyPrecision]
	if ok {
		precision = fmt.Sprint(p)
	}
	timeFmt := tools.ReadableTimeFmtSec
	if precision == datePrecisionDay {
		timeFmt = tools.ReadableTimeFmtDay
	} else if precision == datePrecisionMin {
		timeFmt = tools.ReadableTimeFmtMin
	} else if precision == datePrecisionSec {
		timeFmt = tools.ReadableTimeFmtSec
	} else if precision == datePrecisionMilli {
		timeFmt = tools.ReadableTimeFmtMilli
	}
	zone := "local"
	z, ok := params[keyZone]
	if ok {
		zone = fmt.Sprint(z)
	}
	unit := "sec"
	u, ok := params[keyUnit]
	if ok {
		unit = fmt.Sprint(u)
	}
	var dateTime time.Time
	var err error
	if zone == zoneTypeLocal {
		dateTime, err = time.ParseInLocation(timeFmt, fmt.Sprint(dateStr), time.Local)
	} else if zone == zoneTypeUTC {
		dateTime, err = time.ParseInLocation(timeFmt, fmt.Sprint(dateStr), time.UTC)
	} else if zone == zoneTypeShanghai {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		dateTime, err = time.ParseInLocation(timeFmt, fmt.Sprint(dateStr), loc)
	} else {
		var loc *time.Location
		loc, err = time.LoadLocation(zone)
		if err != nil {
			return time.Now().Unix()
		}
		dateTime, err = time.ParseInLocation(timeFmt, fmt.Sprint(dateStr), loc)
	}
	if err != nil {
		return time.Now().Unix()
	}
	if unit == unitSec {
		return dateTime.Unix()
	} else if unit == unitMilli {
		return dateTime.UnixMilli()
	} else if unit == unitMicro {
		return dateTime.UnixMicro()
	} else if unit == unitNano {
		return dateTime.UnixNano()
	} else {
		return time.Now().Unix()
	}
}

func UnixToDate(params map[string]interface{}) interface{} {
	const (
		keyDate   = "date"
		keyFormat = "format"
	)
	dateUnix, ok := params[keyDate]
	if !ok {
		return time.Now().Format(tools.ReadableTimeFmtSec)
	}
	unix, err := tools.ConvToInt64(dateUnix)
	if err != nil {
		return time.Now().Format(tools.ReadableTimeFmtSec)
	}
	dateTime := time.Unix(unix, 0)
	format := tools.ReadableTimeFmtSec
	f, ok := params[keyFormat]
	if ok {
		format = fmt.Sprint(f)
	}
	return dateTime.Format(format)
}
