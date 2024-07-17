package log

import "testing"

func TestLog(t *testing.T) {
	Info("test")
}

func TestLogStack(t *testing.T) {
	ErrorS("test")
}

func TestLogNoCaller(t *testing.T) {
	noCallerLogger.Debug("test")
}
