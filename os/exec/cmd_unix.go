//go:build unix

package exec

func ContainQuotedCMD(s string) (string, error) {
	return Cmd(s)
}

func ContainQuotedStdoutCMD(s string) error {
	return StdOutCmd(s)
}
