package ffmpeg

import (
	"fmt"
	"github.com/hopeio/utils/io/fs"
	"github.com/hopeio/utils/io/fs/path"
	osi "github.com/hopeio/utils/os"
)

type Frame int

func (f Frame) String() string {
	switch f {
	case I:
		return "I"
	case P:
		return "P"
	case B:
		return "B"
	}
	return "unknown"
}

const (
	I Frame = iota
	P
	B
)

const GetFrameCmd = CommonCmd + `-vf "select=eq(pict_type\,%s)" -fps_mode vfr -qscale:v 2 -f image2 %s/%%03d.jpg`

func GetFrame(src string, f Frame) error {
	//cmd := `ffmpeg -i ` + src + ` -vf "select=eq(pict_type\,` + f.String() + `)" -vsync vfr -qscale:v 2 -f image2 ` + dst + `/%03d.jpg`
	dst := path.CleanDir(src) + f.String() + "Frame"
	fs.Mkdir(dst)
	cmd := fmt.Sprintf(GetFrameCmd, src, f.String(), dst)
	_, err := osi.ContainQuotedCMD(cmd)
	return err
}
