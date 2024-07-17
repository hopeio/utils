package time

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type SecondTime = Time[secondTime]
type MilliTime = Time[milliTime]
type MicroTime = Time[microTime]
type NanoTime = Time[nanoTime]

type Encode interface {
	Encoding() *Encoding
}

type Time[T Encode] time.Time

func NewTime[T Encode](t time.Time) Time[T] {
	return Time[T](t)
}

func (dt Time[T]) Time() time.Time {
	return time.Time(dt)
}

func (dt *Time[T]) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*dt = Time[T](nullTime.Time)
	return
}

func (dt Time[T]) Value() (driver.Value, error) {
	return time.Time(dt), nil
}

func (dt Time[T]) Format(format string) string {
	return time.Time(dt).Format(format)
}

func (dt Time[T]) GormDataType() string {
	return "time"
}

func (dt Time[T]) MarshalBinary() ([]byte, error) {
	return time.Time(dt).MarshalBinary()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (dt *Time[T]) UnmarshalBinary(data []byte) error {
	return (*time.Time)(dt).UnmarshalBinary(data)
}

func (dt Time[T]) GobEncode() ([]byte, error) {
	return dt.MarshalBinary()
}

func (dt *Time[T]) GobDecode(data []byte) error {
	return dt.UnmarshalBinary(data)
}

func (dt Time[T]) MarshalJSON() ([]byte, error) {
	var v T
	return v.Encoding().marshalJSON(dt.Time())
}

func (dt *Time[T]) UnmarshalJSON(data []byte) error {
	var v T
	return v.Encoding().unmarshalJSON((*time.Time)(dt), data)
}

type Edate struct{}

func (Edate) Layout() string {
	return time.DateOnly
}

func (Edate) Encoding() *Encoding {
	return &Encoding{
		Layout: time.DateOnly,
	}
}

type EdateTime struct{}

func (EdateTime) Layout() string {
	return time.DateTime
}

func (EdateTime) Encoding() *Encoding {
	return &Encoding{
		Layout: time.DateTime,
	}
}

type Etime struct{}

func (Etime) Encoding() *Encoding {
	return &Encoding{
		Layout: time.RFC3339,
	}
}

func (Etime) Layout() string {
	return time.RFC3339
}
