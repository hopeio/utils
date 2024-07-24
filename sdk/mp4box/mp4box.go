package mp4box

import (
	execi "github.com/hopeio/utils/os/exec"
)

// https://www.videohelp.com/software/MP4Box
const Mp4BoxCmd = `mp4box -add-image %s.hevc:primary -ab heic -new %s.heic`

func Heic(filePath, dst string) error {
	_, err := execi.Cmd(Mp4BoxCmd)
	return err
}
