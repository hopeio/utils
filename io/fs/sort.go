package fs

import "os"

type DirEntries []os.DirEntry

func (e DirEntries) Len() int {
	return len(e)
}

func (e DirEntries) Less(i, j int) bool {
	filei, _ := e[i].Info()
	filej, _ := e[j].Info()
	return filei.ModTime().After(filej.ModTime())
}

func (e DirEntries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
