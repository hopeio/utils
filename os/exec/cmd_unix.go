//go:build unix

/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package exec

import (
	osi "github.com/hopeio/gox/os"
	"os/exec"
)

func RunGetOutContainQuoted(s string) (string, error) {
	return RunGetOut(s)
}

func RunContainQuoted(s string) error {
	return Run(s)
}

func ContainQuotedCMD(s string) *exec.Cmd {
	words := osi.Split(s)
	return exec.Command(words[0], words[1:]...)
}
