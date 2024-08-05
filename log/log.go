package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

func init() {
	SetDefaultLogger(&Config{Development: true, Level: zapcore.DebugLevel})
}

type skipLogger struct {
	*Logger
	needUpdate bool
}

var (
	defaultLogger  *Logger
	stackLogger    *Logger
	noCallerLogger *Logger
	skipLoggers    = make([]skipLogger, 10)
	mu             sync.Mutex
)

//go:nosplit
func Default() *Logger {
	return defaultLogger
}

func SetDefaultLogger(lf *Config, cores ...zapcore.Core) {
	mu.Lock()
	defer mu.Unlock()

	defaultLogger = lf.NewLogger(cores...)
	stackLogger = defaultLogger.WithOptions(zap.WithCaller(true), zap.AddStacktrace(zapcore.ErrorLevel))
	noCallerLogger = defaultLogger.WithOptions(zap.WithCaller(false))
	for i := 0; i < len(skipLoggers); i++ {
		if skipLoggers[i].Logger != nil {
			skipLoggers[i].needUpdate = true
		}
	}
}

// range -3~6
func GetCallerSkipLogger(skip int) *Logger {
	if skip < -3 {
		panic("skip不小于-3")
	}
	if skip > 6 {
		panic("skip不大于6")
	}
	idx := skip + 3
	if skipLoggers[idx].needUpdate || skipLoggers[idx].Logger == nil {
		mu.Lock()
		skipLoggers[idx].Logger = defaultLogger.AddSkip(skip)
		skipLoggers[idx].needUpdate = false
		mu.Unlock()
	}
	return skipLoggers[idx].Logger
}

func GetNoCallerLogger() *Logger {
	return noCallerLogger
}
func GetStackLogger() *Logger {
	return stackLogger
}

func Sync() error {
	return defaultLogger.Sync()
}

func Debug(args ...any) {
	if ce := defaultLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Info(args ...any) {
	if ce := defaultLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Warn(args ...any) {
	if ce := defaultLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Error(args ...any) {
	if ce := defaultLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Panic(args ...any) {
	if ce := defaultLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Fatal(args ...any) {
	if ce := defaultLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func Printf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Debugf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.DebugLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Infof(template string, args ...any) {
	if ce := defaultLogger.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Warnf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.WarnLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Errorf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Panicf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.PanicLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Fatalf(template string, args ...any) {
	if ce := defaultLogger.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Debugw(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Infow(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Warnw(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Errorw(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Panicw(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Fatalw(msg string, fields ...zap.Field) {
	if ce := defaultLogger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// with stack
func ErrorS(args ...any) {
	if ce := stackLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func ErrorSf(template string, args ...any) {
	if ce := stackLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func ErrorSw(msg string, fields ...zap.Field) {
	if ce := stackLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// no caller

func DebugNC(args ...any) {
	if ce := noCallerLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func InfoNC(args ...any) {
	if ce := noCallerLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func WarnNC(args ...any) {
	if ce := noCallerLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func ErrorNC(args ...any) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func PanicNC(args ...any) {
	if ce := noCallerLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func FatalNC(args ...any) {
	if ce := noCallerLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func ErrorNCf(template string, args ...any) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func FatalNCf(template string, args ...any) {
	if ce := noCallerLogger.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func ErrorNCw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func PanicNCw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func FatalNCw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Log(lvl zapcore.Level, msg string) {
	if ce := defaultLogger.Check(lvl, msg); ce != nil {
		ce.Write()
	}
}

func Logf(lvl zapcore.Level, msg string, args ...any) {
	if ce := defaultLogger.Check(lvl, fmt.Sprintf(msg, args...)); ce != nil {
		ce.Write()
	}
}

func Logw(lvl zapcore.Level, msg string, fields ...zapcore.Field) {
	if ce := defaultLogger.Check(lvl, msg); ce != nil {
		ce.Write(fields...)
	}
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return defaultLogger.Check(lvl, msg)
}

func DebugCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func InfoCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func WarnCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func ErrorCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func DPanicCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.DPanicLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func PanicCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func FatalCE(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...)))
}
