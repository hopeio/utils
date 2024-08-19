package parallel

import (
	"github.com/hopeio/utils/errors/multierr"
	"github.com/hopeio/utils/log"
	"github.com/hopeio/utils/types/funcs"
	"sync"
)

func Run(tasks []funcs.FuncWithErr) error {
	ch := make(chan error)
	for _, task := range tasks {
		task := task // 兼容!go1.22
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
	taskCh  chan funcs.FuncContinue
	workNum uint
	wg      sync.WaitGroup
}

func New(workNum uint, opts ...Option) *Parallel {
	return &Parallel{taskCh: make(chan funcs.FuncContinue, workNum), workNum: workNum}
}

func (p *Parallel) Run() {
	g := func() {
		defer func() {
			if err := recover(); err != nil {
				log.ErrorS(err)
			}
		}()
		for task := range p.taskCh {
			var times = uint(1)
			for task(times) {
				times++
			}
			p.wg.Done()
		}
	}
	for _ = range p.workNum {
		go g()
	}
}

func (p *Parallel) AddTask(task funcs.FuncContinue) {
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
