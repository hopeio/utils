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

type WarpErrRep struct {
	ErrRep
	err error
}

func (e *WarpErrRep) Error() string {
	return e.Message
}

func (e *WarpErrRep) Unwrap() error {
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
