package parallel

import (
	"github.com/hopeio/utils/errors/multierr"
	"github.com/hopeio/utils/types/funcs"
	"golang.org/x/sync/errgroup"
)

func RunIgnoreError(tasks []funcs.FuncReturnErr) error {
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

func Run(tasks []funcs.FuncReturnErr) error {
	var group errgroup.Group
	for _, task := range tasks {
		group.Go(task)
	}
	return group.Wait()
}

func RunReturnData[T any](tasks []funcs.FuncReturnDataOrErr[T]) ([]T, error) {
	var group errgroup.Group
	ret := make([]T, len(tasks))
	for i, task := range tasks {
		group.Go(func() error {
			data, err := task()
			ret[i] = data
			return err
		})
	}
	return ret, group.Wait()
}
