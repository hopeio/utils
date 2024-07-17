package errors

type ErrMsg string

func (e ErrMsg) Error() string {
	return string(e)
}
