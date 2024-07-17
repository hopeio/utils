package log

import (
	"context"
	contexti "github.com/hopeio/utils/context"
	"go.uber.org/zap"
)

func TraceIdField(ctx context.Context) zap.Field {
	return zap.String(FieldTraceId, contexti.TraceId(ctx))
}
