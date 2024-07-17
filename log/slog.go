package log

import (
	"context"
	"github.com/hopeio/utils/slices"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log/slog"
)

var _ slog.Handler = &Logger{}

func (l *Logger) NewSLogger() *slog.Logger {
	return slog.New(l)
}

func (l *Logger) Enabled(ctx context.Context, level slog.Level) bool {
	return l.Logger.Core().Enabled(zapcore.Level(level / 4))
}

func (l *Logger) Handle(ctx context.Context, record slog.Record) error {
	if ce := l.Check(zapcore.Level(record.Level/4), record.Message); ce != nil {
		ce.Write()
	}
	return nil
}

func (l *Logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return l.With(slices.Map(attrs, func(attr slog.Attr) zap.Field {
		return zap.String(attr.Key, attr.Value.String())
	})...)
}

func (l *Logger) WithGroup(name string) slog.Handler {
	return l.Named(name)
}
