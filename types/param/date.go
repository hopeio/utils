/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

import (
	timei "github.com/hopeio/utils/time"
	"strconv"
	"time"
)

type DateFilter struct {
	DateBegin string `json:"dateBegin" comment:"起始时间"`
	DateEnd   string `json:"dateEnd" comment:"结束时间"`
	RangeEnum int    `json:"rangeEnum" comment:"1-今天,2-本周，3-本月，4-今年"`
}

// 赋值本周期，并返回下周期日期
func (d *DateFilter) Scope() (time.Time, time.Time) {
	beginStr, endStr := d.scope()
	begin, _ := time.Parse(timei.LayoutDateTime, beginStr)
	end, _ := time.Parse(timei.LayoutDateTime, endStr)
	return begin, end
}

func (d *DateFilter) scope() (string, string) {
	if d.DateBegin != "" && d.DateEnd != "" {
		begin := d.DateBegin + timei.DayBeginTimeWithSpace
		end := d.DateEnd + timei.DayEndTimeWithSpace
		return begin, end
	}
	//如果传的是RangeEnum，截止日期都是这一天
	now := time.Now()
	d.DateEnd = now.Format(timei.LayoutDate) + timei.DayEndTimeWithSpace
	switch d.RangeEnum {
	case 1:
		beginStr := now.Format(timei.LayoutDate)
		d.DateBegin = beginStr
	case 2:
		weekday := now.Weekday()
		if weekday == time.Sunday {
			weekday = 6
		} else {
			weekday -= 1
		}
		begin := now.AddDate(0, 0, -int(weekday))
		d.DateBegin = begin.Format("2006-01-02") + timei.DayBeginTimeWithSpace

	case 3:
		d.DateBegin = now.Format("2006-01") + "-01 00:00:00"
	case 4:
		d.DateBegin = strconv.Itoa(now.Year()) + "-01-01 00:00:00"
	}
	return d.DateBegin, d.DateEnd
}
