//go:build windows

package os

import (
	stringsi "github.com/hopeio/utils/strings"
	"os"
	"os/exec"
	"syscall"
)

func ContainQuotedCMD(s string) (string, error) {
	exe := s
	for i, c := range s {
		if c == ' ' {
			exe = s[:i]
			break
		}
	}
	cmd := exec.Command(exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: s[len(exe):], HideWindow: true}
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return stringsi.BytesToString(buf), err
	}
	if len(buf) == 0 {
		return "", nil
	}
	return stringsi.BytesToString(buf), nil
}

func ContainQuotedStdoutCMD(s string) error {
	exe := s
	for i, c := range s {
		if c == ' ' {
			exe = s[:i]
			break
		}
	}
	cmd := exec.Command(exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: s[len(exe):], HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
