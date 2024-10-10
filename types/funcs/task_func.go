package funcs

import (
	"context"
)

type GrpcServiceMethod[REQ, RES any] func(context.Context, REQ) (RES, error)

type Func func()

type FuncWithErr func() error
type FuncContinue func(times uint) (retry bool)

func (f FuncContinue) Do(times uint) (retry bool) {
	return f(times)
}

type Task func(context.Context)
type TaskWithErr func(context.Context) error
