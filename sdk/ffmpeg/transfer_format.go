/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package ffmpeg

import (
	"bytes"
	"fmt"
	execi "github.com/hopeio/utils/os/exec"
	fs2 "github.com/hopeio/utils/os/fs"
	"log"
	"os"
)

const TransferFormatGPUCmd = ` -hwaccel qsv -i "%s" -c copy -y "%s"`

func TransferFormatGPU(filePath, dst string) error {
	command := fmt.Sprintf(execPath+TransferFormatGPUCmd, filePath, dst)
	log.Println(command)
	_, err := execi.RunGetOut(command)
	return err
}

const TransferFormatCmd = CommonCmd + ` -c copy -y "%s"`

func TransferFormat(filePath, dst string) error {
	return Run(fmt.Sprintf(TransferFormatCmd, filePath, dst))
}

const ConcatCmd = ` -f concat -safe 0  -i "%s" -c copy -y "%s"`

func ConcatByFile(filePath, dst string) error {
	return Run(fmt.Sprintf(ConcatCmd, filePath, dst))
}

func Concat(dir, dst string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var data bytes.Buffer
	for _, file := range files {
		data.WriteString(`file '` + dir + fs2.PathSeparator + file.Name() + "\n")
	}
	ffmpegFilePath := dir + fs2.PathSeparator + "file.txt"

	file, err := os.Create(ffmpegFilePath)
	if err != nil {
		return fmt.Errorf("create ffmpeg file failed：%s", err.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer file.Close()
	_, err = file.Write(data.Bytes())
	if err != nil {
		return fmt.Errorf("write to %s: %s", ffmpegFilePath, err.Error())
	}
	return ConcatByFile(ffmpegFilePath, dst)
}
