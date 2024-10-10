package monitor

import (
	"context"
	"sync"
	"sync/atomic"
)

type Monitor struct {
	ctx      context.Context
	cancel   context.CancelFunc
	num, end atomic.Int32
	running  atomic.Bool
	callback func()
}

func New(ctx context.Context, callback func()) *Monitor {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	return &Monitor{
		ctx:      ctx,
		cancel:   cancel,
		callback: sync.OnceFunc(callback),
	}
}

type Task func() []Task

// 有没有可能父协程return后，子协程还没开始执行
// 这里有问题，存在协程快速执行完后直接执行回调的情况，此时并非所有任务完成
func (ng *Monitor) Run(fns ...func()) error {
	/*	if !ng.running.CompareAndSwap(false, true) {
		return errors.New("monitor is running")
	}*/
	ng.num.Add(int32(len(fns)))
	for _, fn := range fns {
		go func() {
			fn()
			ng.end.Add(1)
			if ng.num.Load() == ng.end.Load() {
				//ng.running.Store(false)
				ng.callback()
			}
		}()
	}
	return nil
}

func (ng *Monitor) Context() context.Context {
	return ng.ctx
}

func (ng *Monitor) Cancel() {
	ng.cancel()
}
