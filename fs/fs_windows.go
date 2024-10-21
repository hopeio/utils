//go:build windows

package fs

import (
	"os"
	"syscall"
	"time"
)

func CreateTime(path string) time.Time {
	fileInfo, _ := os.Stat(path)
	wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	tNanSeconds := wFileSys.CreationTime.Nanoseconds() /// 返回的是纳秒
	return time.Unix(0, tNanSeconds)
}
