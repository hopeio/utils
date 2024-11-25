/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package time

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
	"time"
)

// 毫秒
type Timestamp = timestamp[milliTime]

func NewTimeStamp(t time.Time) Timestamp {
	return Timestamp(t.UnixMilli())
}

type SecondTimestamp = timestamp[secondTime]
type MilliTimestamp = timestamp[milliTime]
type MicroTimestamp = timestamp[microTime]
type NanoTimestamp = timestamp[nanoTime]

type TimestampConstraints interface {
	Timestamp(time.Time) int64
	Time(int64) time.Time
}

type timestamp[T TimestampConstraints] int64

func (t timestamp[T]) Time() time.Time {
	var v T
	return v.Time(int64(t))
}

// Scan scan time.
func (t *timestamp[T]) Scan(value interface{}) (err error) {
	var v T
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = timestamp[T](v.Timestamp(nullTime.Time))
	return
}

// Value get time value.
func (t timestamp[T]) Value() (driver.Value, error) {
	var v T
	return v.Time(int64(t)), nil
}

func (t timestamp[T]) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, int64(t), 10), nil
}

func (t *timestamp[T]) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(data) == 0 || str == "null" {
		return nil
	}
	parseInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*t = timestamp[T](parseInt)
	return nil
}

// go生成的纳米级时间戳最后两位恒为0
type nanoTime struct{}

func (nanoTime) Encoding() *Encoding {
	return &Encoding{
		EncodeType: EncodeTypeUnixNanoseconds,
	}
}

func (nanoTime) Timestamp(t time.Time) int64 {
	return t.UnixNano()
}
func (nanoTime) Time(t int64) time.Time {
	return time.Unix(0, t)
}

type microTime struct{}

func (microTime) Encoding() *Encoding {
	return &Encoding{
		EncodeType: EncodeTypeUnixMicroseconds,
	}
}
func (microTime) Timestamp(t time.Time) int64 {
	return t.UnixMicro()
}
func (microTime) Time(t int64) time.Time {
	return time.UnixMicro(t)
}

type milliTime struct{}

func (milliTime) Encoding() *Encoding {
	return &Encoding{
		EncodeType: EncodeTypeUnixMilliseconds,
	}
}

func (milliTime) Timestamp(t time.Time) int64 {
	return t.UnixMilli()
}
func (milliTime) Time(t int64) time.Time {
	return time.UnixMilli(t)
}

type secondTime struct{}

func (secondTime) Encoding() *Encoding {
	return &Encoding{
		EncodeType: EncodeTypeUnixSeconds,
	}
}

func (secondTime) Timestamp(t time.Time) int64 {
	return t.Unix()
}
func (secondTime) Time(t int64) time.Time {
	return time.Unix(t, 0)
}
