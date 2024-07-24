package ffmpeg

import (
	osi "github.com/hopeio/utils/os"
	"log"
)

// doc: https://ffmpeg.org/ffmpeg-codecs.html
// https://ffmpeg.org/download.html

const CommonCmd = ` -i "%s" -y `

type Config struct {
	Path string
}

var execPath = "ffmpeg"

func SetExecPath(path string) {
	execPath = path
}

func ffmpegCmd(cmd string) error {
	cmd = execPath
	log.Println(cmd)
	err := osi.ContainQuotedStdoutCMD(cmd)
	if err != nil {
		log.Println(err)
		return err
	}
	//log.Println(res)
	return nil
}
