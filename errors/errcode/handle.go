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
