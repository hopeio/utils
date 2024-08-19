package log

import (
	"fmt"
	"go.uber.org/zap"
)

// no line ending

func NoLEDebug(args ...any) {
	if ce := noLineEndingLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEInfo(args ...any) {
	if ce := noLineEndingLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEWarn(args ...any) {
	if ce := noLineEndingLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEError(args ...any) {
	if ce := noLineEndingLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEPanic(args ...any) {
	if ce := noCallerLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEFatal(args ...any) {
	if ce := noLineEndingLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoLEDebugf(template string, args ...any) {
	if ce := noLineEndingLogger.Check(zap.DebugLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}
func NoLEInfof(template string, args ...any) {
	if ce := noLineEndingLogger.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}
func NoLEErrorf(template string, args ...any) {
	if ce := noLineEndingLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func NoLEFatalf(template string, args ...any) {
	if ce := noLineEndingLogger.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func NoLEDebugw(msg string, fields ...zap.Field) {
	if ce := noLineEndingLogger.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoLEInfow(msg string, fields ...zap.Field) {
	if ce := noLineEndingLogger.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoLEErrorw(msg string, fields ...zap.Field) {
	if ce := noLineEndingLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoLEPanicw(msg string, fields ...zap.Field) {
	if ce := noLineEndingLogger.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoLEFatalw(msg string, fields ...zap.Field) {
	if ce := noLineEndingLogger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}
