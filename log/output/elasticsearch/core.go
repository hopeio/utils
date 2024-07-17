package elasticsearch

import "go.uber.org/zap/zapcore"

func NewCore() zapcore.Core { return &core{} }

type core struct {
	zapcore.Level
}

func (c *core) Enabled(lvl zapcore.Level) bool    { return lvl > c.Level }
func (n *core) With([]zapcore.Field) zapcore.Core { return n }
func (c *core) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	return ce.AddCore(ent, c)
}
func (*core) Write(zapcore.Entry, []zapcore.Field) error { return nil }
func (*core) Sync() error                                { return nil }
