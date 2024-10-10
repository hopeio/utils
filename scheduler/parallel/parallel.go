package parallel

import (
	"github.com/hopeio/utils/errors/multierr"
	"github.com/hopeio/utils/log"
	"github.com/hopeio/utils/types/funcs"
	"github.com/hopeio/utils/types/interfaces"
	"sync"
)

func Run(tasks []funcs.FuncWithErr) error {
	ch := make(chan error)
	for _, task := range tasks {
		go func() {
			ch <- task()
		}()
	}
	var errs multierr.MultiError
	for err := range ch {
		if err != nil {
			errs.Append(err)
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

type Parallel struct {
	taskCh chan interfaces.FuncContinue
	wg     sync.WaitGroup
}

func New(workNum uint, opts ...Option) *Parallel {
	taskCh := make(chan interfaces.FuncContinue, workNum)
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

func (p *Parallel) AddFunc(task funcs.FuncContinue) {
	p.wg.Add(1)
	p.taskCh <- task
}

func (p *Parallel) AddTask(task interfaces.FuncContinue) {
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
