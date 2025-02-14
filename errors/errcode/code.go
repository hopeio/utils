/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package errcode

const (
	// SysErr ErrCode = -1
	Success            ErrCode = 0
	Canceled           ErrCode = 1
	Unknown            ErrCode = 2
	InvalidArgument    ErrCode = 3
	DeadlineExceeded   ErrCode = 4
	NotFound           ErrCode = 5
	AlreadyExists      ErrCode = 6
	PermissionDenied   ErrCode = 7
	ResourceExhausted  ErrCode = 8
	FailedPrecondition ErrCode = 9
	Aborted            ErrCode = 10
	OutOfRange         ErrCode = 11
	Unimplemented      ErrCode = 12
	Internal           ErrCode = 13
	Unavailable        ErrCode = 14
	DataLoss           ErrCode = 15
	Unauthenticated    ErrCode = 16
)

var codeMap = make(map[ErrCode]string)

// 不是并发安全的，在初始化的时候做
func Register(code ErrCode, msg string) {
	codeMap[code] = msg
}

func init() {
	Register(Canceled, "Canceled")
	Register(Unknown, "Unknown")
	Register(InvalidArgument, "InvalidArgument")
	Register(DeadlineExceeded, "DeadlineExceeded")
	Register(NotFound, "NotFound")
	Register(AlreadyExists, "AlreadyExists")
	Register(PermissionDenied, "PermissionDenied")
	Register(ResourceExhausted, "ResourceExhausted")
	Register(FailedPrecondition, "FailedPrecondition")
	Register(Aborted, "Aborted")
	Register(OutOfRange, "OutOfRange")
	Register(Unimplemented, "Unimplemented")
	Register(Internal, "Internal")
	Register(Unavailable, "Unavailable")
	Register(DataLoss, "DataLoss")
	Register(Unauthenticated, "Unauthenticated")
}
