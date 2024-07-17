// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package encoder

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
	"time"
)

// DefaultLineEnding defines the default line ending when writing logs.
// Alternate line endings specified in EncoderConfig can override this
// behavior.
const DefaultLineEnding = "\n"

// OmitKey defines the key to use when callers want to remove a key from log output.
const OmitKey = ""

// A TimeEncoder serializes a time.Time to a primitive type.
type TimeEncoder func(time.Time, zapcore.PrimitiveArrayEncoder)

// EpochTimeEncoder serializes a time.Time to a floating-point number of seconds
// since the Unix epoch.
func EpochTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	sec := float64(nanos) / float64(time.Second)
	enc.AppendFloat64(sec)
}

// EpochMillisTimeEncoder serializes a time.Time to a floating-point number of
// milliseconds since the Unix epoch.
func EpochMillisTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	millis := float64(nanos) / float64(time.Millisecond)
	enc.AppendFloat64(millis)
}

// EpochNanosTimeEncoder serializes a time.Time to an integer number of
// nanoseconds since the Unix epoch.
func EpochNanosTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano())
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}

// ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string
// with millisecond precision.
//
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
// instead of appending a pre-formatted string value.
func ISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, "2006-01-02T15:04:05.000Z0700", enc)
}

// RFC3339TimeEncoder serializes a time.Time to an RFC3339-formatted string.
//
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
// instead of appending a pre-formatted string value.
func RFC3339TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, time.RFC3339, enc)
}

// RFC3339NanoTimeEncoder serializes a time.Time to an RFC3339-formatted string
// with nanosecond precision.
//
// If enc supports AppendTimeLayout(t time.Time,layout string), it's used
// instead of appending a pre-formatted string value.
func RFC3339NanoTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	encodeTimeLayout(t, time.RFC3339Nano, enc)
}

// TimeEncoderOfLayout returns TimeEncoder which serializes a time.Time using
// given layout.
func TimeEncoderOfLayout(layout string) TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		encodeTimeLayout(t, layout, enc)
	}
}

// UnmarshalText unmarshals text to a TimeEncoder.
// "rfc3339nano" and "RFC3339Nano" are unmarshaled to RFC3339NanoTimeEncoder.
// "rfc3339" and "RFC3339" are unmarshaled to RFC3339TimeEncoder.
// "iso8601" and "ISO8601" are unmarshaled to ISO8601TimeEncoder.
// "millis" is unmarshaled to EpochMillisTimeEncoder.
// "nanos" is unmarshaled to EpochNanosEncoder.
// Anything else is unmarshaled to EpochTimeEncoder.
func (e *TimeEncoder) UnmarshalText(text []byte) error {
	switch string(text) {
	case "rfc3339nano", "RFC3339Nano":
		*e = RFC3339NanoTimeEncoder
	case "rfc3339", "RFC3339":
		*e = RFC3339TimeEncoder
	case "iso8601", "ISO8601":
		*e = ISO8601TimeEncoder
	case "millis":
		*e = EpochMillisTimeEncoder
	case "nanos":
		*e = EpochNanosTimeEncoder
	default:
		*e = EpochTimeEncoder
	}
	return nil
}

// UnmarshalYAML unmarshals YAML to a TimeEncoder.
// If value is an object with a "layout" field, it will be unmarshaled to  TimeEncoder with given layout.
//
//	timeEncoder:
//	  layout: 06/01/02 03:04pm
//
// If value is string, it uses UnmarshalText.
//
//	timeEncoder: iso8601
func (e *TimeEncoder) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var o struct {
		Layout string `json:"layout" yaml:"layout"`
	}
	if err := unmarshal(&o); err == nil {
		*e = TimeEncoderOfLayout(o.Layout)
		return nil
	}

	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return e.UnmarshalText([]byte(s))
}

// UnmarshalJSON unmarshals JSON to a TimeEncoder as same way UnmarshalYAML does.
func (e *TimeEncoder) UnmarshalJSON(data []byte) error {
	return e.UnmarshalYAML(func(v interface{}) error {
		return json.Unmarshal(data, v)
	})
}

// A DurationEncoder serializes a time.Duration to a primitive type.
type DurationEncoder func(time.Duration, zapcore.PrimitiveArrayEncoder)

// SecondsDurationEncoder serializes a time.Duration to a floating-point number of seconds elapsed.
func SecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Second))
}

// NanosDurationEncoder serializes a time.Duration to an integer number of
// nanoseconds elapsed.
func NanosDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(int64(d))
}

// MillisDurationEncoder serializes a time.Duration to an integer number of
// milliseconds elapsed.
func MillisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Nanoseconds() / 1e6)
}

// StringDurationEncoder serializes a time.Duration using its built-in String
// method.
func StringDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(d.String())
}

// UnmarshalText unmarshals text to a DurationEncoder. "string" is unmarshaled
// to StringDurationEncoder, and anything else is unmarshaled to
// NanosDurationEncoder.
func (e *DurationEncoder) UnmarshalText(text []byte) error {
	switch string(text) {
	case "string":
		*e = StringDurationEncoder
	case "nanos":
		*e = NanosDurationEncoder
	case "ms":
		*e = MillisDurationEncoder
	default:
		*e = SecondsDurationEncoder
	}
	return nil
}

// A NameEncoder serializes a period-separated logger name to a primitive
// type.
type NameEncoder func(string, zapcore.PrimitiveArrayEncoder)

// FullNameEncoder serializes the logger name as-is.
func FullNameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(loggerName)
}

// UnmarshalText unmarshals text to a NameEncoder. Currently, everything is
// unmarshaled to FullNameEncoder.
func (e *NameEncoder) UnmarshalText(text []byte) error {
	switch string(text) {
	case "full":
		*e = FullNameEncoder
	default:
		*e = FullNameEncoder
	}
	return nil
}

// An EncoderConfig allows users to configure the concrete encoders supplied by
// zapcore.
type EncoderConfig struct {
	// Set the keys used for each log entry. If any key is empty, that portion
	// of the entry is omitted.
	MessageKey    string `json:"messageKey" yaml:"messageKey"`
	LevelKey      string `json:"levelKey" yaml:"levelKey"`
	TimeKey       string `json:"timeKey" yaml:"timeKey"`
	NameKey       string `json:"nameKey" yaml:"nameKey"`
	CallerKey     string `json:"callerKey" yaml:"callerKey"`
	FunctionKey   string `json:"functionKey" yaml:"functionKey"`
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktraceKey"`
	LineEnding    string `json:"lineEnding" yaml:"lineEnding"`
	// Configure the primitive representations of common complex types. For
	// example, some users may want all time.Times serialized as floating-point
	// seconds since epoch, while others may prefer ISO8601 strings.
	EncodeTime     TimeEncoder     `json:"timeEncoder" yaml:"timeEncoder"`
	EncodeDuration DurationEncoder `json:"durationEncoder" yaml:"durationEncoder"`

	// Unlike the other primitive type encoders, EncodeName is optional. The
	// zero value falls back to FullNameEncoder.
	EncodeName NameEncoder `json:"nameEncoder" yaml:"nameEncoder"`
	// Configures the field separator used by the console encoder. Defaults
	// to tab.
	ConsoleSeparator string `json:"consoleSeparator" yaml:"consoleSeparator"`
}
