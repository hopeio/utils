package funcs

import (
	"context"
)

type GrpcServiceMethod[REQ, RES any] func(context.Context, REQ) (RES, error)

type Func func()

type FuncWithErr func() error
type FuncContinue func(times uint) (retry bool)

type Task func(context.Context)
type TaskWithErr func(context.Context) error
