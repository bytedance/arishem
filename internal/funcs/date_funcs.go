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

package funcs

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/arishem/tools"
	"time"
)

func GetCurrentYear(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Year()), nil
}

func GetCurrentYearDay(ctx context.Context) (interface{}, error) {
	return int64(time.Now().YearDay()), nil
}

func GetCurrentMonth(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Month()), nil
}

func GetCurrentMonthDay(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Day()), nil
}

func GetCurrentWeekDay(ctx context.Context) (interface{}, error) {
	return time.Now().Weekday().String(), nil
}

// GetCurrentHour 获取当前的小时数
func GetCurrentHour(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Hour()), nil
}

func GetCurrentMinute(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Minute()), nil
}

func GetCurrentSecond(ctx context.Context) (interface{}, error) {
	return int64(time.Now().Second()), nil
}

func GetCurrentLocation(ctx context.Context) (interface{}, error) {
	return time.Now().Location().String(), nil
}

func GetUnixSecond(ctx context.Context) (interface{}, error) {
	return time.Now().Unix(), nil
}

func DateToUnix(ctx context.Context, params map[string]interface{}) (interface{}, error) {
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
		return nil, errors.New("[DateToUnix] param: date not found")
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
			return nil, err
		}
		dateTime, err = time.ParseInLocation(timeFmt, fmt.Sprint(dateStr), loc)
	}
	if err != nil {
		return nil, err
	}
	if unit == unitSec {
		return dateTime.Unix(), nil
	} else if unit == unitMilli {
		return dateTime.UnixMilli(), nil
	} else if unit == unitMicro {
		return dateTime.UnixMicro(), nil
	} else if unit == unitNano {
		return dateTime.UnixNano(), nil
	} else {
		return time.Now().Unix(), nil
	}
}

func UnixToDate(ctx context.Context, params map[string]interface{}) (interface{}, error) {
	const (
		keyDate   = "date"
		keyFormat = "format"
	)
	dateUnix, ok := params[keyDate]
	if !ok {
		return nil, errors.New("[UnixToDate] param: date not found")
	}
	unix, err := tools.ConvToInt64(dateUnix)
	if err != nil {
		return nil, err
	}
	dateTime := time.Unix(unix, 0)
	format := tools.ReadableTimeFmtSec
	f, ok := params[keyFormat]
	if ok {
		format = fmt.Sprint(f)
	}
	return dateTime.Format(format), nil
}
