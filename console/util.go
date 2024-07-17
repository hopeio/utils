package console

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func IsTerminal() bool {
	if _, exists := os.LookupEnv("TERM"); exists {
		return true
	}
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		return true
	}
	if terminal.IsTerminal(int(os.Stderr.Fd())) {
		return true
	}
	return false
}
