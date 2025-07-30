/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package retry

import "github.com/hopeio/gox/errors/multierr"

func RunTimes(times int, f func(int) error) error {
	var err error
	for i := 0; i < times; i++ {
		err1 := f(i)
		if err1 == nil {
			return nil
		}
		err = multierr.Append(err, err1)
	}

	return err
}

func Run(f func(int) bool) {
	for i := 0; ; i++ {
		if !f(i) {
			break
		}
	}
}
