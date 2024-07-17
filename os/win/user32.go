package win

import (
	"syscall"
)

var (
	user32          = syscall.NewLazyDLL("User32.dll")
	procEnumWindows = user32.NewProc("EnumWindows")
	findWindow      = user32.NewProc("FindWindowW")
)
