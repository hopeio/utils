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
		return &errs
	}
	return nil
}

type Parallel struct {
	taskCh     chan funcs.FuncWithErr
	workNum    int
	wg         sync.WaitGroup
	retryTimes int
}

func New(workNum int, opts ...Option) *Parallel {
	return &Parallel{taskCh: make(chan funcs.FuncWithErr, workNum), workNum: workNum}
}

func (p *Parallel) RetryTimes(retryTimes int) *Parallel {
	p.retryTimes = retryTimes
	return p
}

func (p *Parallel) Run() {
	for _ = range p.workNum {
		go func() {
			for task := range p.taskCh {
				err := task()
				if err != nil {
					if p.retryTimes > 0 {
						for _ = range p.retryTimes - 1 {
							err = task()
							if err == nil {
								break
							}
						}
					}
					log.Error(err)
				}
				p.wg.Done()
			}
		}()
	}
}

func (p *Parallel) AddTask(task funcs.FuncWithErr) {
	p.wg.Add(1)
	p.taskCh <- task
}

func (p *Parallel) Stop() {
	p.wg.Wait()
	close(p.taskCh)
}

type Option func(p *Parallel)

func RetryTimes(retryTimes int) Option {
	return func(p *Parallel) {
		p.retryTimes = retryTimes
	}
}
