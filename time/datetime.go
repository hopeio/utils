/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package time

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/hopeio/utils/encoding/binary"
	stringsi "github.com/hopeio/utils/strings"
	"io"
	"strconv"
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
	str := stringsi.BytesToString(data)
	if len(data) == 0 || str == "null" {
		return nil
	}

	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
		t, err := time.Parse(time.DateOnly, str)
		if err != nil {
			return err
		}
		*d = Date(t.Unix() / int64(DaySecond))
		return nil
	} else {
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		if len(str) == 13 {
			*d = Date(v / 1000 / int64(DaySecond))
		} else {
			*d = Date(v / int64(DaySecond))
		}
	}
	return nil
}

func (d Date) MarshalText() ([]byte, error) {
	return stringsi.ToBytes(d.Time().Format(time.DateOnly)), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	str := stringsi.BytesToString(data)
	t, err := time.Parse(time.DateOnly, str)
	if err != nil {
		return err
	}
	*d = Date(t.Unix() / int64(DaySecond))
	return nil
}

func (ts Date) GormDataType() string {
	return "time"
}

func (ts Date) MarshalBinary() ([]byte, error) {
	return binary.ToBinary(ts), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (ts *Date) UnmarshalBinary(data []byte) error {
	*ts = binary.BinaryTo[Date](data)
	return nil
}

func (ts Date) GobEncode() ([]byte, error) {
	return ts.MarshalBinary()
}

func (ts *Date) GobDecode(data []byte) error {
	return ts.UnmarshalBinary(data)
}

func (x Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(x.Time().Format(time.DateOnly)))
}

func (x *Date) UnmarshalGQL(v interface{}) error {
	if i, ok := v.(string); ok {
		t, err := time.Parse(time.DateOnly, i)
		if err != nil {
			return err
		}
		*x = Date(t.Unix() / int64(DaySecond))
		return nil
	}
	return errors.New("enum need integer type")
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
	str := stringsi.BytesToString(data)
	if len(data) == 0 || str == "null" {
		return nil
	}

	if len(str) > 1 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
		t, err := time.Parse(time.DateTime, str)
		if err != nil {
			return err
		}
		*d = DateTime(t.Unix())
		return nil
	} else {
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		if len(str) == 13 {
			*d = DateTime(v / 1000)
		} else {
			*d = DateTime(v)
		}
	}
	return nil
}

func (d DateTime) MarshalText() ([]byte, error) {
	return stringsi.ToBytes(d.Time().Format(time.DateTime)), nil
}

func (d *DateTime) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	str := stringsi.BytesToString(data)
	t, err := time.Parse(time.DateTime, str)
	if err != nil {
		return err
	}
	*d = DateTime(t.Unix())
	return nil
}

func (ts DateTime) GormDataType() string {
	return "time"
}
