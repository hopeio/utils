package retry

import "github.com/hopeio/utils/errors/multierr"

func ReTry(times int, f func() error) error {
	var err error
	for i := 0; i < times; i++ {
		err1 := f()
		if err1 == nil {
			return nil
		}
		err = multierr.Append(err, err1)
	}

	return err
}
