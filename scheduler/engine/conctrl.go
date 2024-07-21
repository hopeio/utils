package engine

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hopeio/utils/log"
	id2 "github.com/hopeio/utils/structure/idgen/id"
	synci "github.com/hopeio/utils/sync"
	"sync/atomic"
	"time"
)

func (e *Engine[KEY]) Run(tasks ...*Task[KEY]) {
	e.lock.Lock()
	if e.isRunning {
		if len(tasks) > 0 {
			e.AddTasks(tasks...)
		}
		e.lock.Unlock()
		return
	}
	if !e.errHandleRunning {
		go func() {
			for {
				select {
				case <-e.ctx.Done():
					return
				case task := <-e.taskErrChan:
					e.taskErrHandleCount++
					e.errHandler(task)
					e.wg.Done()
				}
			}
		}()
		e.errHandleRunning = true
	}
	e.addWorker()
	if !e.isRunning {
		e.isRunning = true
		e.wg.Add(1)
		go func() {
			timer := time.NewTimer(5 * time.Second)
			defer timer.Stop()
			var emptyTimes uint
			var readyTaskCh chan *Task[KEY]
			var readyTask *Task[KEY]
			taskChan := e.taskChanProducer
		loop:
			for {
				if len(e.taskReadyHeap) > 0 && readyTask == nil {
					readyTask, _ = e.taskReadyHeap.Pop()
					readyTaskCh = e.taskChanConsumer
				}

				// go 无法动态的设置case,但是可以动态的把channel置为nil
				if len(e.taskReadyHeap) >= int(e.limitWaitTaskCount) {
					taskChan = nil
				} else {
					taskChan = e.taskChanProducer
				}
				select {
				case readyTaskTmp := <-taskChan:
					e.taskReadyHeap.Push(readyTaskTmp)
				case readyTaskCh <- readyTask:
					readyTaskCh = nil
					readyTask = nil
				case <-timer.C:
					//检测任务是否已空
					if uint(e.workingWorkerCount) == 0 && len(e.taskReadyHeap) == 0 {
						e.lock.Lock()
						counter, _ := synci.WaitGroupState(&e.wg)
						if counter == 1 {
							emptyTimes++
							if emptyTimes > 2 {
								log.GetNoCallerLogger().Debug("the task is about to end.")
								e.wg.Done()
								e.isRunning = false
								e.lock.Unlock()
								break loop
							}
						}
						e.lock.Unlock()
					}
					fmt.Printf("\r[Running] task:D:%d/T:%d/S:%d/H:%d/F:%d/E:%d,worker: %d/%d", e.taskDoneCount, e.taskTotalCount, e.taskSkipCount, e.taskErrHandleCount, e.taskFailedCount, e.taskErrorTimes, e.workingWorkerCount, e.currentWorkerCount)
					timer.Reset(e.monitorInterval)
				case <-e.ctx.Done():
					if err := e.ctx.Err(); err != nil {
						log.Error(err)
					}
					break loop
				}

			}
		}()
	}

	e.lock.Unlock()
	if len(tasks) > 0 {
		e.AddTasks(tasks...)
	}
	e.wg.Wait()
	log.GetNoCallerLogger().Infof("[END] task:D:%d/T:%d/S:%d/H:%d/F:%d/E:%d", e.taskDoneCount, e.taskTotalCount, e.taskSkipCount, e.taskErrHandleCount, e.taskFailedCount, e.taskErrorTimes)
}

func (e *Engine[KEY]) newWorker(readyTask *Task[KEY]) {
	atomic.AddUint64(&e.currentWorkerCount, 1)
	//id := c.currentWorkerCount
	// 这里考虑回复多channel,worker数量多起来的时候,channel维护的goroutine数量太多
	worker := &Worker[KEY]{id: uint(e.currentWorkerCount)}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				worker.canExecute = false
				log.ErrorS(r, spew.Sdump(readyTask))
				atomic.AddUint64(&e.taskFailedCount, 1)
				e.wg.Done()
				// 创建一个新的
				e.newWorker(nil)
			}
			atomic.AddUint64(&e.currentWorkerCount, ^uint64(0))
		}()
		worker.canExecute = true
		if readyTask != nil {
			e.ExecTask(worker, readyTask)
		}
		for {
			select {
			case readyTask = <-e.taskChanConsumer:
				e.ExecTask(worker, readyTask)
			case <-e.ctx.Done():
				worker.canExecute = false
				return
			}
		}
	}()
	e.workers = append(e.workers, worker)
}

func (e *Engine[KEY]) addWorker() {
	if atomic.LoadUint64(&e.currentWorkerCount) == 0 {
		e.newWorker(nil)
	}
	if e.workerFactoryRunning.Load() {
		return
	}
	go func() {
		e.workerFactoryRunning.Store(true)
		for {
			select {
			case readyTask := <-e.taskChanConsumer:
				if atomic.LoadUint64(&e.currentWorkerCount) < atomic.LoadUint64(&e.limitWorkerCount) {
					e.newWorker(readyTask)
				} else {
					log.Info("worker count is full")
					e.taskChanProducer <- readyTask
					e.workerFactoryRunning.Store(false)
					return
				}
			case <-e.ctx.Done():
				return
			}
		}
	}()

}

func (e *Engine[KEY]) addTasks(ctx context.Context, priority int, tasks ...*Task[KEY]) {
	l := len(tasks)
	atomic.AddUint64(&e.taskTotalCount, uint64(l))
	e.wg.Add(l)
	for _, task := range tasks {
		if task == nil || task.TaskFunc == nil {
			atomic.AddUint64(&e.taskTotalCount, ^uint64(0))
			continue
		}
		if ctx != nil {
			task.Context = ctx
		}
		task.Priority += priority
		task.id = id2.NewOrderID()
		e.taskChanProducer <- task
	}
}

func (e *Engine[KEY]) AddOptionTasks(ctx context.Context, priority int, tasks ...*Task[KEY]) {
	e.addTasks(ctx, priority, tasks...)
}

func (e *Engine[KEY]) AddTasks(tasks ...*Task[KEY]) {
	e.addTasks(nil, 0, tasks...)
}

func (e *Engine[KEY]) AsyncAddTasks(tasks ...*Task[KEY]) {
	if len(tasks) > 0 {
		go e.addTasks(nil, 0, tasks...)
	}
}

func (e *Engine[KEY]) AsyncAddOptionTasks(ctx context.Context, priority int, tasks ...*Task[KEY]) {
	if len(tasks) > 0 {
		go e.addTasks(ctx, priority, tasks...)
	}
}

func (e *Engine[KEY]) reTryTasks(tasks ...*Task[KEY]) {
	if len(tasks) > 0 {
		go func() {
			for _, task := range tasks {
				task.Priority++
				e.taskChanProducer <- task
			}
		}()
	}
}

func (e *Engine[KEY]) AddWorker(num int) {
	atomic.AddUint64(&e.limitWorkerCount, uint64(num))
	e.addWorker()
}

func (e *Engine[KEY]) NewFixedWorker(interval time.Duration) int {
	taskChan := make(chan *Task[KEY])
	worker := &Worker[KEY]{id: uint(e.currentWorkerCount), typ: fixedType, taskCh: taskChan}
	e.workers = append(e.workers, worker)
	e.newFixedWorker(worker, interval)
	return len(e.workers) - 1
}

func (e *Engine[KEY]) newFixedWorker(worker *Worker[KEY], interval time.Duration) {
	go func() {
		var task *Task[KEY]
		defer func() {
			if r := recover(); r != nil {
				worker.canExecute = false
				log.ErrorS(r, spew.Sdump(task))
				atomic.AddUint64(&e.taskFailedCount, 1)
				e.wg.Done()
				// 创建一个新的
				e.newFixedWorker(worker, interval)
			}
		}()
		var timer *time.Ticker
		// 如果有任务时间间隔,
		if interval > 0 {
			timer = time.NewTicker(interval)
		}
		worker.canExecute = true
		for task = range worker.taskCh {
			if interval > 0 {
				<-timer.C
			}
			e.ExecTask(worker, task)
		}
	}()
}

func (e *Engine[KEY]) AddFixedTasks(workerId int, generation int, tasks ...*Task[KEY]) error {
	err := fmt.Errorf("不存在workId为%d的fixed worker,请调用NewFixedWorker添加", workerId)
	if workerId > len(e.workers)-1 {
		return err
	}
	worker := e.workers[workerId]
	if worker.typ != fixedType {
		return err
	}
	l := len(tasks)
	atomic.AddUint64(&e.taskTotalCount, uint64(l))
	e.wg.Add(l)
	for _, task := range tasks {
		if task == nil || task.TaskFunc == nil {
			atomic.AddUint64(&e.taskTotalCount, ^uint64(0))
			continue
		}
		task.Priority += generation
		task.id = id2.NewOrderID()
		worker.taskCh <- task
	}
	return nil
}

func (e *Engine[KEY]) RunSingleWorker(tasks ...*Task[KEY]) {
	e.limitWorkerCount = 1
	e.Run(tasks...)
}

func (e *Engine[KEY]) ExecTask(worker *Worker[KEY], task *Task[KEY]) {
	atomic.AddUint64(&e.workingWorkerCount, 1)
	worker.isExecuting = true
	worker.currentTask = task
	if e.execTask(task) {
		e.wg.Done()
	}
	atomic.AddUint64(&e.workingWorkerCount, ^uint64(0))
	worker.isExecuting = false
}

func (e *Engine[KEY]) execTask(task *Task[KEY]) bool {
	if task.Key != e.zeroKey {
		if _, ok := e.done.Get(task.Key); ok {
			atomic.AddUint64(&e.taskSkipCount, 1)
			return true
		}
	}

	if e.speedLimit != nil {
		e.speedLimit.Wait()
	}

	if e.rateLimiter != nil {
		e.rateLimiter.Wait(task.Context)
	}

	var kindHandler *KindHandler[KEY]
	if e.kindHandlers != nil && int(task.Kind) < len(e.kindHandlers) {
		kindHandler = e.kindHandlers[task.Kind]
	}

	if kindHandler != nil {
		if kindHandler.Skip {
			atomic.AddUint64(&e.taskSkipCount, 1)
			return true
		}

		if kindHandler.speedLimit != nil {
			kindHandler.speedLimit.Wait()
		}
		if kindHandler.rateLimiter != nil {
			_ = kindHandler.rateLimiter.Wait(task.Context)
		}
	}

	if task.ReExecTimes > 0 {
		task.reExecLogs = append(task.reExecLogs, &execLog{
			execBeginAt: time.Now(),
		})
	} else {
		task.execBeginAt = time.Now()
	}
	tasks, err := task.TaskFunc.Do(task.Context)
	if task.ReExecTimes > 0 {
		task.reExecLogs[len(task.reExecLogs)-1].execEndAt = time.Now()
	} else {
		task.execEndAt = time.Now()
	}

	if err != nil {
		atomic.AddUint64(&e.taskErrorTimes, 1)
		task.ErrTimes++
		if task.ReExecTimes > 0 {
			task.reExecLogs[len(task.reExecLogs)-1].err = err
		} else {
			task.err = err
		}

		if task.ErrTimes < 5 {
			task.ReExecTimes++
			log.Warnf("%v执行失败:%v,将第%d次执行", task.Key, err, task.ReExecTimes)
			e.reTryTasks(task)
		} else {
			log.Warn(task.Key, "多次执行失败:", err, "将执行错误处理")
			e.taskErrChan <- task
		}

		return false
	}
	if task.Key != e.zeroKey {
		e.done.SetWithTTL(task.Key, struct{}{}, 1, time.Hour)
	}
	if len(tasks) > 0 {
		e.AsyncAddOptionTasks(task.Context, task.Priority+1, tasks...)
	}
	atomic.AddUint64(&e.taskDoneCount, 1)
	return true
}

func (e *Engine[KEY]) Stop() {
	e.cancel()
	if e.speedLimit != nil {
		e.speedLimit.Stop()
	}
	e.done.Close()
	for _, kindHandler := range e.kindHandlers {
		if kindHandler != nil {
			if kindHandler.speedLimit != nil {
				kindHandler.speedLimit.Stop()
			}
			if kindHandler.rateLimiter != nil {
				kindHandler.rateLimiter = nil
			}
		}
	}

	for _, callback := range e.stopCallBack {
		callback()
	}
	e.isStopped = true
}

func (e *Engine[KEY]) StopAfter(interval time.Duration) *Engine[KEY] {
	time.AfterFunc(interval, e.Stop)
	return e
}
