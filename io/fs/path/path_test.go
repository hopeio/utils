package path

import (
	"fmt"
	"path"
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	dir := "https://a\\video/a.jpg"
	t.Log(path.Split(dir))
	t.Log(filepath.Split(dir))
	t.Log(filepath.Dir(dir), filepath.Base(dir))
}

func TestClean(t *testing.T) {
	s := `......`
	fmt.Println(len(s))
	r := FileCleanse(s)
	fmt.Println(r)
	fmt.Println(len(r))
}

func TestRune(t *testing.T) {
	t.Log('，')
}
