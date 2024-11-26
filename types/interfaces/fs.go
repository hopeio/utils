package interfaces

import (
	"io/fs"
)

type FileSystem interface {
	Mkdir(name string, perm fs.FileMode) error
	OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error)
	RemoveAll(name string) error
	Rename(oldName, newName string) error
	Stat(name string) (fs.FileInfo, error)
}
