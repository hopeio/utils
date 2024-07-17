package dingding

import (
	"github.com/hopeio/utils/sdk/dingding"
	"go.uber.org/zap/zapcore"
	"strconv"
	"time"
)

func NewCore(token, secret string, level zapcore.Level) zapcore.Core {
	return &core{
		RobotConfig: dingding.RobotConfig{
			Token:  token,
			Secret: secret,
		},
		Level: level,
	}
}

type core struct {
	dingding.RobotConfig
	zapcore.Level
	fields []zapcore.Field
}

func (c *core) Enabled(lvl zapcore.Level) bool { return lvl > c.Level }
func (c *core) With(fields []zapcore.Field) zapcore.Core {
	c.fields = fields
	return c
}
func (c *core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return ce.AddCore(ent, c)
}
func (c *core) Write(e zapcore.Entry, fields []zapcore.Field) error {

	enc := NewDingEncoder(&zapcore.EncoderConfig{
		MessageKey:     "信息",
		LevelKey:       "级别",
		TimeKey:        "时间",
		CallerKey:      "调用行",
		FunctionKey:    "函数",
		SkipLineEnding: true,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
		},
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(strconv.FormatInt(d.Nanoseconds()/1e6, 10) + "ms")
		},
		EncodeCaller: func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(caller.TrimmedPath())
		},
		ConsoleSeparator: "",
	})

	buffer, err := enc.EncodeEntry(e, append(fields, c.fields...))
	if err != nil {
		return err
	}

	return dingding.SendRobotMarkDownMessageWithSecret(c.Token, c.Secret, "日志", buffer.String())
}
func (*core) Sync() error { return nil }
