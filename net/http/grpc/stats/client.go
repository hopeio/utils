package stats

import (
	"context"
	"github.com/hopeio/context/httpctx"
	httpi "github.com/hopeio/utils/net/http"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/stats"
)

type InternalClientHandler struct {
}

// HandleConn exists to satisfy gRPC stats.Handler.
func (c *InternalClientHandler) HandleConn(ctx context.Context, cs stats.ConnStats) {
	// no-op
}

// TagConn exists to satisfy gRPC stats.Handler.
func (c *InternalClientHandler) TagConn(ctx context.Context, cti *stats.ConnTagInfo) context.Context {
	// no-op
	return ctx
}

// HandleRPC implements per-RPC tracing and stats instrumentation.
func (c *InternalClientHandler) HandleRPC(ctx context.Context, rs stats.RPCStats) {
}

// TagRPC implements per-RPC context management.
func (c *InternalClientHandler) TagRPC(ctx context.Context, rti *stats.RPCTagInfo) context.Context {
	ctxi := httpctx.FromContextValue(ctx)
	return metadata.AppendToOutgoingContext(ctx, httpi.HeaderTraceID,
		ctxi.TraceID,
		httpi.HeaderGrpcInternal, httpi.HeaderGrpcInternal)
}
