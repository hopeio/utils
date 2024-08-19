package log

import (
	"fmt"
	"go.uber.org/zap"
)

// with stack
func StackError(args ...any) {
	if ce := stackLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func StackErrorf(template string, args ...any) {
	if ce := stackLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func StackErrorw(msg string, fields ...zap.Field) {
	if ce := stackLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// no caller

func NoCallerDebug(args ...any) {
	if ce := noCallerLogger.Check(zap.DebugLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerInfo(args ...any) {
	if ce := noCallerLogger.Check(zap.InfoLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerWarn(args ...any) {
	if ce := noCallerLogger.Check(zap.WarnLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerError(args ...any) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerPanic(args ...any) {
	if ce := noCallerLogger.Check(zap.PanicLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerFatal(args ...any) {
	if ce := noCallerLogger.Check(zap.FatalLevel, trimLineBreak(fmt.Sprintln(args...))); ce != nil {
		ce.Write()
	}
}

func NoCallerErrorf(template string, args ...any) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func NoCallerFatalf(template string, args ...any) {
	if ce := noCallerLogger.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func NoCallerErrorw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoCallerPanicw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func NoCallerFatalw(msg string, fields ...zap.Field) {
	if ce := noCallerLogger.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}
