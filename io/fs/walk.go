package fs

import "os"

type Visitor interface {
	Visit(dir string, file os.DirEntry) (w Visitor)
}

type inspector func(string, os.DirEntry) bool

func (f inspector) Visit(dir string, file os.DirEntry) Visitor {
	if f == nil {
		return f
	}
	if f(dir, file) {
		return f
	}
	return nil
}

// Inspect
func Inspect(dir string, file func(string, os.DirEntry) bool) error {
	return walk(inspector(file), dir, nil)
}

func walk(v Visitor, dir string, file os.DirEntry) error {
	if v = v.Visit(dir, file); v == nil {
		return nil
	}
	if file.IsDir() {
		dirs, err := os.ReadDir(dir + PathSeparator)
		if err != nil {
			return err
		}
		for _, file := range dirs {
			err = walk(v, dir+PathSeparator+file.Name(), file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
