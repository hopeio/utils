package fs

import (
	"github.com/hopeio/utils/errors/multierr"
	"io/fs"
	"os"
	"path/filepath"
)

type RangeCallback = func(dir string, entry os.DirEntry) error

// 遍历根目录中的每个文件，为每个文件调用callback,包括文件夹,与filepath.WalkDir不同的是回调函数的参数不同,filepath.WalkDir的第一个参数是文件完整路径,RangeFile是文件所在目录的路径
func Range(dir string, callback RangeCallback) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := multierr.New()
	for _, entry := range entries {
		if entry.IsDir() {
			err = RangeFile(dir+PathSeparator+entry.Name(), callback)
			if err != nil {
				errs.Append(err)
			}
		}
		err = callback(dir, entry)
		if err != nil {
			errs.Append(err)
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

// 指定遍历深度,0为只遍历一层,-1为无限遍历
func RangeDeep(dir string, callback RangeCallback, deep int) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := multierr.New()
	for _, entry := range entries {
		if entry.IsDir() && deep != 0 {
			err = RangeDeep(dir+PathSeparator+entry.Name(), callback, deep-1)
			if err != nil {
				errs.Append(err)
			}
		}
		err = callback(dir, entry)
		if err != nil {
			errs.Append(err)
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

// 遍历根目录中的每个文件，为每个文件调用callback,不包括文件夹,与filepath.WalkDir不同的是回调函数的参数不同,filepath.WalkDir的第一个参数是文件完整路径,RangeFile是文件所在目录的路径
func RangeFile(dir string, callback RangeCallback) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := multierr.New()
	for _, entry := range entries {
		if entry.IsDir() {
			err = RangeFile(dir+PathSeparator+entry.Name(), callback)
			if err != nil {
				errs.Append(err)
			}
		} else {
			err = callback(dir, entry)
			if err != nil {
				errs.Append(err)
			}
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

// 指定遍历深度,0为只遍历一层,-1为无限遍历
func RangeFileDeep(dir string, callback RangeCallback, deep int) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := multierr.New()
	for _, entry := range entries {
		if entry.IsDir() && deep != 0 {
			err = RangeFileDeep(dir+PathSeparator+entry.Name(), callback, deep-1)
			if err != nil {
				errs.Append(err)
			}
		} else {
			err = callback(dir, entry)
			if err != nil {
				errs.Append(err)
			}
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

// RangeDir 遍历根目录中的每个文件夹，为文件夹中所有文件和目录的切片(os.ReadDir的返回)调用callback
// callback 需要处理每个文件夹下的所有文件和目录,返回值为需要递归遍历的目录和error
// 几乎每个文件夹下的文件夹都会被循环两次！
func RangeDir(dir string, callback func(dir string, entries []os.DirEntry) ([]os.DirEntry, error)) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	errs := multierr.New()

	dirs, err := callback(dir, entries)
	if err != nil {
		errs.Append(err)
	}
	for _, entry := range dirs {
		if entry.IsDir() {
			err = RangeDir(dir+PathSeparator+entry.Name(), callback)
			if err != nil {
				errs.Append(err)
			}
		}
	}
	if errs.HasErrors() {
		return errs
	}
	return nil
}

func WalkDirFS(fsys fs.FS, root string, fn fs.WalkDirFunc) error {
	return fs.WalkDir(fsys, root, fn)
}

func Walk(root string, fn filepath.WalkFunc) error {
	return filepath.Walk(root, fn)
}

func WalkDir(root string, fn fs.WalkDirFunc) error {
	return filepath.WalkDir(root, fn)
}
