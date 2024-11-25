/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

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
	clf := *lf
	clf.SkipLineEnding = true
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

func Log(lvl zapcore.Level, args ...any) {
	if ce := defaultLogger.Check(lvl, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
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

func DebugEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func InfoEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func WarnEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func ErrorEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func DPanicEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.DPanicLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func PanicEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func FatalEntry(args ...any) *zapcore.CheckedEntry {
	return defaultLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...)))
}

func Println(args ...any) {
	if ce := defaultLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}
