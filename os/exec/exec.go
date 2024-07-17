package exec

import (
	"log"
	"os"
	"os/exec"

	osi "github.com/hopeio/utils/os"
)

func Run(arg string) error {
	words := osi.Split(arg)
	cmd := exec.Command(words[0], words[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.String())
	return cmd.Run()
}
