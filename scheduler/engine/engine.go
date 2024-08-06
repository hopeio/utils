package engine

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgraph-io/ristretto"
	"github.com/hopeio/utils/io/fs"
	"github.com/hopeio/utils/log"
	"github.com/hopeio/utils/slices"
	"github.com/hopeio/utils/structure/heap"
	time2 "github.com/hopeio/utils/time"
	"golang.org/x/time/rate"
	"sync"
	"sync/atomic"
	"time"
)

// 目前受限于ristretto.Cache的泛型限制,考虑移除并引入lru或boolong filter
type Key interface {
	uint64 | string | byte | int | int32 | uint32 | int64
}

type Config[KEY Key] struct {
	WorkerCount     uint64
	MonitorInterval time.Duration // 全局检测定时器间隔时间，任务的卡住检测，worker panic recover都可以用这个检测
	DoneCache       ristretto.Config[KEY, struct{}]
	EnableTelemetry bool
}

func (c *Config[KEY]) NewEngine() *Engine[KEY] {
	return c.NewEngineWithContext(context.Background())
}

func (c *Config[KEY]) NewEngineWithContext(ctx context.Context) *Engine[KEY] {
	c.Init()
	ctx, cancel := context.WithCancel(ctx)
	cache, _ := ristretto.NewCache(&ristretto.Config[KEY, struct{}]{
		NumCounters:        1e4,   // number of keys to track frequency of (10M).
		MaxCost:            1e3,   // maximum cost of cache (MaxCost * 1MB).
		BufferItems:        64,    // number of keys per Get buffer.
		Metrics:            false, // number of keys per Get buffer.
		IgnoreInternalCost: true,
	})
	engine := &Engine[KEY]{
		workerCount:      c.WorkerCount,
		ctx:              ctx,
		cancel:           cancel,
		taskChanConsumer: make(chan *Task[KEY]),
		errTaskChan:      make(chan *Task[KEY]),
		readyTaskHeap:    heap.Heap[*Task[KEY]]{},
		monitorInterval:  c.MonitorInterval,
		done:             cache,
		errHandler:       func(task *Task[KEY]) { task.ErrLog() },
		mu:               sync.RWMutex{},
		zeroKey:          *new(KEY),
	}
	return engine
}

func (c *Config[KEY]) Init() {
	if c.WorkerCount == 0 {
		c.WorkerCount = 10
	}
	if c.MonitorInterval == 0 {
		c.MonitorInterval = 5 * time.Second
	}
	if c.DoneCache.NumCounters == 0 {
		c.DoneCache.NumCounters = 1e4
	}
	if c.DoneCache.MaxCost == 0 {
		c.DoneCache.MaxCost = 1e3
	}
	if c.DoneCache.BufferItems == 0 {
		c.DoneCache.BufferItems = 64
	}

}

func NewConfig[KEY Key]() *Config[KEY] {
	return &Config[KEY]{
		WorkerCount:     10,
		MonitorInterval: 5 * time.Second,
		DoneCache: ristretto.Config[KEY, struct{}]{
			NumCounters:        1e4,   // number of keys to track frequency of (10M).
			MaxCost:            1e3,   // maximum cost of cache (MaxCost * 1MB).
			BufferItems:        64,    // number of keys per Get buffer.
			Metrics:            false, // number of keys per Get buffer.
			IgnoreInternalCost: true,
		},
	}
}

type Engine[KEY Key] struct {
	workerCount, currentWorkerCount, workingWorkerCount uint64
	waitTaskCount                                       uint64
	workers                                             []*Worker[KEY]
	// workerGroup [][]*Worker[KEY] //TODO 工作组概念
	taskChanConsumer chan *Task[KEY]
	errTaskChan      chan *Task[KEY]
	readyTaskHeap    heap.Heap[*Task[KEY]]
	readyTaskMu      sync.RWMutex
	ctx              context.Context
	cancel           context.CancelFunc // 手动停止执行
	wg               sync.WaitGroup     // 控制确保所有任务执行完
	mu               sync.RWMutex
	speedLimit       time2.Ticker
	rateLimiter      *rate.Limiter
	//TODO
	monitorInterval      time.Duration // 全局检测定时器间隔时间，任务的卡住检测，worker panic recover都可以用这个检测
	workerFactoryRunning atomic.Bool
	errHandlerRunning    bool
	enableTelemetry      bool
	isRunning, isStopped bool
	EngineStatistics
	done         *ristretto.Cache[KEY, struct{}]
	kindHandlers []*KindHandler[KEY]
	errHandler   func(task *Task[KEY])
	onStop       []func(context.Context)
	zeroKey      KEY // 泛型不够强大,又为了性能妥协的字段
}

type KindHandler[KEY Key] struct {
	Skip        bool
	speedLimit  time2.Ticker
	rateLimiter *rate.Limiter
	// TODO 指定Kind的Handler
	Handler TaskFunc[KEY]
}

func NewEngine[KEY Key](workerCount uint64) *Engine[KEY] {
	return NewEngineWithContext[KEY](workerCount, context.Background())
}

func NewEngineWithContext[KEY Key](workerCount uint64, ctx context.Context) *Engine[KEY] {
	conf := NewConfig[KEY]()
	conf.WorkerCount = workerCount
	return conf.NewEngineWithContext(ctx)
}

func (e *Engine[KEY]) SkipKind(kinds ...Kind) *Engine[KEY] {
	length := slices.Max(kinds) + 1
	if e.kindHandlers == nil {
		e.kindHandlers = make([]*KindHandler[KEY], length)
	}
	if int(length) > len(e.kindHandlers) {
		e.kindHandlers = append(e.kindHandlers, make([]*KindHandler[KEY], int(length)-len(e.kindHandlers))...)
	}
	for _, kind := range kinds {
		if e.kindHandlers[kind] == nil {
			e.kindHandlers[kind] = &KindHandler[KEY]{Skip: true}
		} else {
			e.kindHandlers[kind].Skip = true
		}

	}
	return e
}

func (e *Engine[KEY]) MonitorInterval(interval time.Duration) {
	if interval < time.Second {
		log.Warn("monitor interval min one second")
		interval = time.Second
	}

	e.monitorInterval = interval
}

func (e *Engine[KEY]) ErrHandler(errHandler func(task *Task[KEY])) *Engine[KEY] {
	e.errHandler = errHandler
	return e
}

func (e *Engine[KEY]) ErrHandlerUtilSuccess() *Engine[KEY] {
	log.Warn("ErrHandlerUtilSuccess will clear history exec log contains err")
	return e.ErrHandler(func(task *Task[KEY]) {
		task.errTimes = 0
		task.reExecLogs = task.reExecLogs[:0]
		e.AddOptionTasks(task.ctx, task.Priority, task)
	})
}

func (e *Engine[KEY]) ErrHandlerRetryTimes(times int) *Engine[KEY] {
	return e.ErrHandler(func(task *Task[KEY]) {
		if task.reExecTimes < times {
			task.errTimes = 0
			task.reExecLogs = task.reExecLogs[:0]
			e.AddOptionTasks(task.ctx, task.Priority, task)
		} else {
			task.ErrLog()
		}

	})
}

func (e *Engine[KEY]) ErrHandlerWriteToFile(path string) *Engine[KEY] {
	file, err := fs.Create(path)
	if err != nil {
		panic(err)
	}
	e.OnStop(func(context.Context) {
		file.Close()
	})
	return e.ErrHandler(func(task *Task[KEY]) {
		spew.Fdump(file, task)
	})
}

func (e *Engine[KEY]) OnStop(callBack func(context.Context)) *Engine[KEY] {
	e.onStop = append(e.onStop, callBack)
	return e
}

func (e *Engine[KEY]) SpeedLimited(interval time.Duration) *Engine[KEY] {
	e.speedLimit = time2.NewTicker(interval)
	return e
}

func (e *Engine[KEY]) RandSpeedLimited(minInterval, maxInterval time.Duration) *Engine[KEY] {
	e.speedLimit = time2.NewRandTicker(minInterval, maxInterval)
	return e
}

func (e *Engine[KEY]) KindSpeedLimit(kind Kind, interval time.Duration) *Engine[KEY] {
	limiter := time2.NewRandTicker(interval, interval)
	e.kindSpeedLimit(kind, limiter)
	return e
}

func (e *Engine[KEY]) KindRandSpeedLimit(kind Kind, minInterval, maxInterval time.Duration) *Engine[KEY] {
	limiter := time2.NewRandTicker(minInterval, maxInterval)
	e.kindSpeedLimit(kind, limiter)
	return e
}

func (e *Engine[KEY]) kindSpeedLimit(kind Kind, limiter time2.Ticker) *Engine[KEY] {
	if e.kindHandlers == nil {
		e.kindHandlers = make([]*KindHandler[KEY], int(kind)+1)
	}
	if int(kind)+1 > len(e.kindHandlers) {
		e.kindHandlers = append(e.kindHandlers, make([]*KindHandler[KEY], int(kind)+1-len(e.kindHandlers))...)
	}
	if e.kindHandlers[kind] == nil {
		e.kindHandlers[kind] = &KindHandler[KEY]{speedLimit: limiter}
	} else {
		e.kindHandlers[kind].speedLimit = limiter
	}
	return e
}

// 多个kind共用一个timer
func (e *Engine[KEY]) KindGroupSpeedLimit(interval time.Duration, kinds ...Kind) *Engine[KEY] {
	limiter := time2.NewRandTicker(interval, interval)
	for _, kind := range kinds {
		e.kindSpeedLimit(kind, limiter)
	}
	return e
}

func (e *Engine[KEY]) KindGroupRandSpeedLimit(minInterval, maxInterval time.Duration, kinds ...Kind) *Engine[KEY] {
	limiter := time2.NewRandTicker(minInterval, maxInterval)
	for _, kind := range kinds {
		e.kindSpeedLimit(kind, limiter)
	}
	return e
}

func (e *Engine[KEY]) Limiter(r rate.Limit, b int) *Engine[KEY] {
	e.rateLimiter = rate.NewLimiter(r, b)
	return e
}

func (e *Engine[KEY]) KindLimiter(kind Kind, r rate.Limit, b int) *Engine[KEY] {
	e.kindLimiter(kind, r, b)
	return e
}

func (e *Engine[KEY]) kindLimiter(kind Kind, r rate.Limit, b int) {
	if e.kindHandlers == nil {
		e.kindHandlers = make([]*KindHandler[KEY], int(kind)+1)
	}
	if int(kind)+1 > len(e.kindHandlers) {
		e.kindHandlers = append(e.kindHandlers, make([]*KindHandler[KEY], int(kind)+1-len(e.kindHandlers))...)
	}
	if e.kindHandlers[kind] == nil {
		e.kindHandlers[kind] = &KindHandler[KEY]{rateLimiter: rate.NewLimiter(r, b)}
	} else {
		e.kindHandlers[kind].rateLimiter = rate.NewLimiter(r, b)
	}
}

type AddTask[KEY Key] func(ctx context.Context, priority int, task ...*Task[KEY])

// TaskSource,参数为添加任务的函数，直到该函数运行结束，任务引擎才会检测任务是否结束
func (e *Engine[KEY]) TaskSource(taskSource func(addTask *Engine[KEY])) {
	e.wg.Add(1)
	go func() {
		taskSource(e)
		e.wg.Done()
	}()
}
