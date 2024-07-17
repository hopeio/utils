package monitor

import (
	"context"

	"go.uber.org/atomic"
)

// 监测子协程是否跑完，需要代码层面的配合
// 一点都不优雅
// 没什么优不优雅，这就像wg.Add(1)
type Monitor struct {
	context.Context
	context.CancelFunc
	run, end *atomic.Int32
	callback func()
}

func New(ctx context.Context, callback func()) *Monitor {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	return &Monitor{
		Context:    ctx,
		CancelFunc: cancel,
		run:        atomic.NewInt32(0),
		end:        atomic.NewInt32(0),
		callback:   callback,
	}
}

// 有没有可能父协程return后，子协程还没开始执行
func (ng *Monitor) Run(fn func()) {
	ng.run.Add(1)
	go func() {
		fn()
		ng.end.Add(1)
		if ng.run.Load() == ng.end.Load() {
			ng.callback()
		}
	}()
}

func (ng *Monitor) Cancel() {
	ng.CancelFunc()
}
