package errcode

const (
	// SysErr ErrCode = -1
	Success ErrCode = 0
)

var codeMap = make(map[ErrCode]string)

// 不是并发安全的，在初始化的时候做
func RegisterErrCode(code ErrCode, msg string) {
	codeMap[code] = msg
}

func init() {
	RegisterErrCode(Success, "OK")
}
