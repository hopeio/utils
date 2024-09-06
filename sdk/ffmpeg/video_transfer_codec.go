package ffmpeg

import (
	"fmt"
)

type PerSet string

const (
	Ultrafast PerSet = "ultrafast"
	SuperFast PerSet = "superfast"
	VeryFast  PerSet = "veryfast"
	Faster    PerSet = "faster"
	Fast      PerSet = "fast"
	Medium    PerSet = "medium"
	Slow      PerSet = "slow"
	Slower    PerSet = "slower"
	VerySlow  PerSet = "veryslow"
	Placebo   PerSet = "placebo"
)

const param = "-global_quality 20"

const H264ToH265ByIntelGPUCmd = `ffmpeg -hwaccel qsv -c:v h264_qsv -i %s -c:v hevc_qsv -preset veryslow -g 60 -gpu_copy 1 -c:a copy "%s"`

const cmd1 = `preset=veryslow,profile=main,look_ahead=1,global_quality=18`

func H264ToH265ByIntelGPU(filePath, dst string) error {
	return Run(fmt.Sprintf(H264ToH265ByIntelGPUCmd, filePath, dst))
}

// libaom-av1
const ToAv1Libaomav1Cmd = CommonCmd + `-c:v libaom-av1 -crf %d -cpu-used %d -row-mt 1 -y "%s"`

// cpu-used
// Set the quality/encoding speed tradeoff. Valid range is from 0 to 8, higher numbers indicating greater speed and lower quality. The default value is 1, which will be slow and high quality.
// row-mt 是否多线程 0否,1是
// tiles 图块数,配合row-mt, axb 猜测是一帧，分成几成几的图片,如2x2就是2行2列4张图分别编码,默认为输入视频大小所需的最小图块数（对于最大 4K 和 4K 的大小，这是 1x1（即单个图块）。
// 很慢,cpu-used调高质量差,推荐3
// crf推荐18-28
func ToAV1ByLibaomav1(filePath, dst string, crf, cpuUsed int) error {
	return Run(fmt.Sprintf(ToAv1Libaomav1Cmd, filePath, crf, cpuUsed, dst))
}

// libsvtav1
// librav1e

// libx264
const ToH264Cmd = CommonCmd + `-c:v libx264 -profile high -preset %s -crf %d -y "%s"`

// crf推荐18
func ToH264ByXlib264(filePath, dst string, crf int, perset PerSet) error {
	return Run(fmt.Sprintf(ToH264Cmd, filePath, perset, crf, dst))
}

// libvpx

// libx265
const ToH265Cmd = CommonCmd + `-c:v libx265 -preset %s -crf %d -y "%s"`

// crf推荐23
func ToH265ByXlib265(filePath, dst string, crf int, perset PerSet) error {
	return Run(fmt.Sprintf(ToH265Cmd, filePath, perset, crf, dst))
}

// libvpx
