/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

import (
	"context"
)

type GrpcServiceMethod[REQ, RES any] func(context.Context, REQ) (RES, error)

type Func func()

type FuncReturnErr func() error
type FuncReturnDataOrErr[T any] func() (T, error)
type FuncRetry func(times uint) (retry bool)

func (f FuncRetry) Do(times uint) (retry bool) {
	return f(times)
}

type Task func(context.Context)
type TaskWithErr func(context.Context) error
