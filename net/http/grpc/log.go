/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package grpc

import (
	"github.com/hopeio/gox/log"
	"go.uber.org/zap/zapgrpc"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLoggerV2(zapgrpc.NewLogger(log.CallerSkipLogger(4).Logger))
}
