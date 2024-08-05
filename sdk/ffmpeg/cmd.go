package ffmpeg

import (
	"github.com/hopeio/utils/log"
	execi "github.com/hopeio/utils/os/exec"
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
	cmd = execPath + " " + cmd
	log.Debug("exec:", cmd)
	err := execi.ContainQuotedStdoutCMD(cmd)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
