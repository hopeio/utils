/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package parallel

import (
	"github.com/hopeio/gox/log"
	"github.com/hopeio/gox/types"
	"github.com/hopeio/gox/types/interfaces"
	"sync"
)

type Parallel struct {
	taskCh chan interfaces.TaskRetry
	wg     sync.WaitGroup
}

func New(workNum uint, opts ...Option) *Parallel {
	taskCh := make(chan interfaces.TaskRetry, workNum)
	p := &Parallel{taskCh: taskCh}
	g := func() {
		defer func() {
			if err := recover(); err != nil {
				log.StackError(err)
			}
		}()
		for task := range taskCh {
			var times = uint(1)
			for task.Do(times) {
				times++
			}
			p.wg.Done()
		}
	}
	for range workNum {
		go g()
	}
	return p
}

func (p *Parallel) AddFunc(task types.FuncRetry) {
	p.wg.Add(1)
	p.taskCh <- task
}

func (p *Parallel) AddTask(task interfaces.TaskRetry) {
	p.wg.Add(1)
	p.taskCh <- task
}

func (p *Parallel) Wait() {
	p.wg.Wait()
}

func (p *Parallel) Stop() {
	p.wg.Wait()
	close(p.taskCh)
}

type Option func(p *Parallel)

type TaskChain []func() error

func (t *TaskChain) Do(times uint) bool {
	taskChain := *t
	for i := 0; i < len(taskChain); i++ {
		err := taskChain[i]()
		if err != nil {
			*t = taskChain[i:]
			return true
		}
	}
	return false
}
