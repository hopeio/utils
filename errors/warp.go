/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package errors

type Unwrapper interface {
	Unwrap() error
}

type Is interface {
	Is(err error) bool
}

// fmt
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrapError{
		msg: msg,
		err: err,
	}
}
