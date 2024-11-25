/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package time

import "time"

type Layout interface {
	Layout() string
}

type TimeStr[T Layout] string

func (t TimeStr[T]) Time() (time.Time, error) {
	var v T
	parse, err := time.Parse(v.Layout(), string(t))
	if err != nil {
		return time.Time{}, err
	}
	return parse, nil
}

func (dt TimeStr[T]) MarshalJSON() ([]byte, error) {
	return []byte(`"` + dt + `"`), nil
}

func (dt *TimeStr[T]) UnmarshalJSON(data []byte) error {
	str := string(data)
	// Ignore null, like in the main JSON package.
	if str == "null" {
		return nil
	}
	*dt = TimeStr[T](str[1 : len(str)-1])
	return nil
}

// time.DateTime
type DateTimeStr = TimeStr[EdateTime]

func NewDateTimeStr(t time.Time) DateTimeStr {
	return DateTimeStr(t.Format(time.DateTime))
}

type DateStr = TimeStr[Edate]

func NewDateStr(t time.Time) DateStr {
	return DateStr(t.Format(time.DateOnly))
}
