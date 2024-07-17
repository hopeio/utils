package grpc

import (
	"github.com/hopeio/utils/log"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(zapgrpc.NewLogger(log.GetCallerSkipLogger(4).Logger))
}
