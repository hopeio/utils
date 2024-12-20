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
func RegisterErrCode(code ErrCode, msg string) {
	codeMap[code] = msg
}

func init() {
	RegisterErrCode(Canceled, "Canceled")
	RegisterErrCode(Unknown, "Unknown")
	RegisterErrCode(InvalidArgument, "InvalidArgument")
	RegisterErrCode(DeadlineExceeded, "DeadlineExceeded")
	RegisterErrCode(NotFound, "NotFound")
	RegisterErrCode(AlreadyExists, "AlreadyExists")
	RegisterErrCode(PermissionDenied, "PermissionDenied")
	RegisterErrCode(ResourceExhausted, "ResourceExhausted")
	RegisterErrCode(FailedPrecondition, "FailedPrecondition")
	RegisterErrCode(Aborted, "Aborted")
	RegisterErrCode(OutOfRange, "OutOfRange")
	RegisterErrCode(Unimplemented, "Unimplemented")
	RegisterErrCode(Internal, "Internal")
	RegisterErrCode(Unavailable, "Unavailable")
	RegisterErrCode(DataLoss, "DataLoss")
	RegisterErrCode(Unauthenticated, "Unauthenticated")
}
