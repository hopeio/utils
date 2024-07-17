package context

import (
	"context"
)

func TraceId(ctx context.Context) string {
	if traceId, ok := ctx.Value(traceIdKey{}).(string); ok {
		return traceId
	}
	return "unknown"
}

type traceIdKey struct{}

func SetTranceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}
