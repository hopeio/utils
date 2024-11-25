/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package errcode

func ErrHandle(err any) error {
	if e, ok := err.(*ErrRep); ok {
		return e
	}
	if e, ok := err.(ErrCode); ok {
		return e.ErrRep()
	}
	if e, ok := err.(error); ok {
		return Unknown.Msg(e.Error())
	}
	return Unknown.ErrRep()
}
