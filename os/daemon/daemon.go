/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package daemon

import (
	"flag"
	"os"
	"os/exec"

	"github.com/hopeio/utils/log"
)

func init() {
	var d bool
	flag.BoolVar(&d, "d", false, "守护进程")
	if !flag.Parsed() {
		flag.Parse()
	}

	if d {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-d=true" {
				os.Args[i] = "-d=false"
			}
		}
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		_ = cmd.Start()
		log.Info("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}
