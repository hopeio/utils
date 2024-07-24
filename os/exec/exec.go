package exec

import (
	"log"
	"os"
	"os/exec"
)

func Run(arg string) error {
	words := Split(arg)
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.String())
	return cmd.Run()
}
