package mp4box

import osi "github.com/hopeio/utils/os"

// https://www.videohelp.com/software/MP4Box
const Mp4BoxCmd = `mp4box -add-image %s.hevc:primary -ab heic -new %s.heic`

func Heic(filePath, dst string) error {
	_, err := osi.Cmd(Mp4BoxCmd)
	return err
}
