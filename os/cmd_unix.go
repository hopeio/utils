//go:build unix

package os

func ContainQuotedCMD(s string) (string, error) {
	return Cmd(s)
}

func ContainQuotedStdoutCMD(s string) error {
	return StdOutCmd(s)
}
