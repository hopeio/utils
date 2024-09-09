package _go

import (
	execi "github.com/hopeio/utils/os/exec"
	"os"
	"strings"
)

const GoListDir = `go list -m -f {{.Dir}} `
const GOPATHKey = "GOPATH"

var gopath, modPath string

func init() {
	if gopath == "" {
		gopath = os.Getenv(GOPATHKey)
	}
	if gopath != "" && !strings.HasSuffix(gopath, "/") {
		gopath = gopath + "/"
	}
	modPath = gopath + "pkg/mod/"
}

func GetDepDir(dep string) string {
	if !strings.Contains(dep, "@") {
		return modDepDir(dep)
	}
	depPath := modPath + dep
	_, err := os.Stat(depPath)
	if os.IsNotExist(err) {
		depPath = modDepDir(dep)
	}
	return depPath
}

func modDepDir(dep string) string {
	depPath, err := execi.RunGetOut(GoListDir + dep)
	if err != nil || depPath == "" {
		execi.RunGetOut("go get " + dep)
		depPath, _ = execi.RunGetOut(GoListDir + dep)
	}
	return depPath
}
