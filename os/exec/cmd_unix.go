//go:build unix

package exec

func RunGetOutContainQuoted(s string) (string, error) {
	return RunGetOut(s)
}

func RunContainQuoted(s string) error {
	return Run(s)
}

func ContainQuotedCMD(s string) *exec.Cmd {
	words := Split(s)
	return exec.Command(words[0], words[1:]...)
}
