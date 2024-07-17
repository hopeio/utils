package errors

type Unwrap interface {
	Unwrap(err error) error
}

type Is interface {
	Is(err error) bool
}

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

type wrapErrors struct {
	msg  string
	errs []error
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrors) Unwrap() []error {
	return e.errs
}
