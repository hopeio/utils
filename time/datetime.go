/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package time

import (
	"database/sql"
	"database/sql/driver"
	stringsi "github.com/hopeio/utils/strings"
	"time"
)

type Date int32

func (d Date) Time() time.Time {
	return time.Unix(int64(d)*int64(DaySecond), 0)
}

// Scan scan time.
func (d *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*d = Date(nullTime.Time.Unix() / int64(DaySecond))
	return
}

// Value get time value.
func (d Date) Value() (driver.Value, error) {
	return []byte(time.Unix(int64(d)*int64(DaySecond), 0).Format(time.DateOnly)), nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, 12)
	b = append(b, '"')
	b = append(b, stringsi.ToBytes(d.Time().Format(time.DateOnly))...)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(data) == 0 || str == "null" {
		return nil
	}

	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
		t, err := time.ParseInLocation(time.DateOnly, str, time.UTC)
		if err != nil {
			return err
		}
		*d = Date(t.Unix() / int64(DaySecond))
		return nil
	}
	return nil
}

func (ts Date) GormDataType() string {
	return "time"
}

type DateTime int64

func (d DateTime) Time() time.Time {
	return time.Unix(int64(d), 0)
}

// Scan scan time.
func (d *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*d = DateTime(nullTime.Time.Unix())
	return
}

// Value get time value.
func (d DateTime) Value() (driver.Value, error) {
	return time.Unix(int64(d), 0), nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = append(b, stringsi.ToBytes(time.Unix(int64(d), 0).Format(time.DateTime))...)
	b = append(b, '"')
	return b, nil
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(data) == 0 || str == "null" {
		return nil
	}

	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
		t, err := time.ParseInLocation(time.DateTime, str, time.Local)
		if err != nil {
			return err
		}
		*d = DateTime(t.Unix())
		return nil
	}
	return nil
}

func (ts DateTime) GormDataType() string {
	return "time"
}
