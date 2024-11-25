/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package poller

import (
	"context"
	"testing"
	"time"
)

func TestPollerRand(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	poller := NewPoller()
	poller.RandRun(ctx, time.Second, time.Second*2, func(ctx context.Context) {
		t.Log("hello")
	})
}

func TestPoller(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	poller := NewPoller()
	poller.Run(ctx, time.Second, func(ctx context.Context) {
		t.Log("hello")
	})
}
