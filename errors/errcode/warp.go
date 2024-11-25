/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package errcode

type WarpErrCode struct {
	ErrCode
	err error
}

func (x *WarpErrCode) Error() string {
	return x.ErrCode.Error()
}

func (x *WarpErrCode) Unwrap() error {
	return x.err
}

type WrapErrRep struct {
	ErrRep
	err error
}

func (e *WrapErrRep) Error() string {
	return e.Msg
}

func (e *WrapErrRep) Unwrap() error {
	return e.err
}

type WarpError struct {
	Message string
	err     error
}

func (e *WarpError) Error() string {
	return e.Message
}

func (e *WarpError) Unwrap() error {
	return e.err
}
