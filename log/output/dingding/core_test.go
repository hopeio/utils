package dingding

import (
	"github.com/hopeio/utils/log"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDingDing(t *testing.T) {

	log.SetDefaultLogger(&log.Config{
		Development: false,
		Level:       zapcore.DebugLevel,
		OutputPaths: log.OutPutPaths{},
		Name:        "",
	}, NewCore("", "", zapcore.DebugLevel))
	log.Info("测试")
}
