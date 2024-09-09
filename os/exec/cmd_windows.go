//go:build windows

package exec

import (
	stringsi "github.com/hopeio/utils/strings"
	"os"
	"os/exec"
	"syscall"
)

func RunGetOutContainQuoted(s string, opts ...Option) (string, error) {
	cmd := ContainQuotedCMD(s)
	for _, opt := range opts {
		opt(cmd)
	}
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return stringsi.BytesToString(buf), err
	}
	if len(buf) == 0 {
		return "", nil
	}
	return stringsi.BytesToString(buf), nil
}

func RunContainQuoted(s string, opts ...Option) error {
	cmd := ContainQuotedCMD(s)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd.Run()
}

func ContainQuotedCMD(s string, opts ...Option) *exec.Cmd {
	exe := s
	for i, c := range s {
		if c == ' ' {
			exe = s[:i]
			break
		}
	}
	cmd := exec.Command(exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: s[len(exe):], HideWindow: true}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}
