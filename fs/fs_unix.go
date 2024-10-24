//go:build unix

package fs

import (
	"os"
	"syscall"
	"time"
)

func init() {
	syscall.Umask(0)
}

func CreateTime(path string) time.Time {
	fileInfo, _ := os.Stat(path)
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat_t.Ctim.Sec), 0)
}

func CreateTimeByInfo(fileInfo os.FileInfo) time.Time {
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat_t.Ctim.Sec), 0)
}

func CreateTimeByEntry(entry os.DirEntry) time.Time {
	fileInfo, _ := entry.Info()
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat_t.Ctim.Sec), 0)
}
