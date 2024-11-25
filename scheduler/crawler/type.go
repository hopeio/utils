/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package crawler

import (
	"context"
	"github.com/hopeio/utils/scheduler/engine"
)

type Request = engine.Task[string]
type TaskFunc = engine.TaskFunc[string]

func NewRequest(key string, kind engine.Kind, taskFunc TaskFunc) *Request {
	return &Request{
		Key:  key,
		Kind: kind,
		Run:  taskFunc,
	}
}

type Config = engine.Config[string]
type Engine = engine.Engine[string]

func NewEngine(workerCount uint64) *engine.Engine[string] {
	return engine.NewEngine[string](workerCount)
}

type HandlerFunc func(ctx context.Context, url string) ([]*Request, error)

func NewUrlRequest(url string, handler HandlerFunc) *Request {
	if handler == nil {
		return nil
	}
	return &Request{Key: url, Run: func(ctx context.Context) ([]*Request, error) {
		return handler(ctx, url)
	}}
}

func NewUrlKindRequest(url string, kind engine.Kind, handleFunc HandlerFunc) *Request {
	if handleFunc == nil {
		return nil
	}
	req := NewUrlRequest(url, handleFunc)
	req.SetKind(kind)
	return req
}
