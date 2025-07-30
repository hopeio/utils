/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package ffmpeg

import (
	"fmt"
	"github.com/hopeio/gox/sdk/mp4box"
	"strings"
)

// webp 无损模式
const ImgToWebpLosslessCmd = CommonCmd + `-c:v libwebp -lossless 1 -quality 100 -compression_level 6 "%s.webp"`

// 图片转webp格式
func ImgToWebpLossless(filePath, dst string) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	return Run(fmt.Sprintf(ImgToWebpLosslessCmd, filePath, dst))
}

const ImgToWebpCmd = CommonCmd + `-c:v libwebp -quality %d "%s.webp"`

//	JPEG 采用的色彩格式是 YUVJ420P，对应的色彩区间是 0-255，而 WebP 采用的色彩格式是 YUV420P，对应的色彩区间是 16-235，也就是说如果单纯的转码，会丢失 0-15，236-255 的色彩，也就是出现了色差, 颜色空间转移：RGB < - > YUV，这会产生一些舍入误差 多次压缩后webp会出现明显色差,真的会偏绿
//
// 图片带选项转webp格式,选项目前支持质量(0-100),ffmpeg默认75
// quality推荐75
func ImgToWebp(filePath, dst string, quality int) error {
	if strings.HasSuffix(dst, ".webp") {
		dst = dst[:len(dst)-5]
	}
	return Run(fmt.Sprintf(ImgToWebpCmd, filePath, quality, dst))
}

const ImgToTAvifCmd = CommonCmd + `-c:v libaom-av1 -crf %d -cpu-used %d -row-mt 1 "%s.avif"`

// 多次压缩后avif会出现明显色差,比webp略好
// -cpu-used 3 会加速，但是图片大小会变大,质量变差,<=3比较好,推荐2
// More encoding options are available: -b 700k -tile-columns 600 -tile-rows 800 - example for the bitrate and tales.

// crf推荐18-28
func ImgToAvif(filePath, dst string, crf, cpuUsed int) error {
	if strings.HasSuffix(dst, ".avif") {
		dst = dst[:len(dst)-5]
	}
	return Run(fmt.Sprintf(ImgToTAvifCmd, filePath, crf, cpuUsed, dst))
}

const ImgToHeicCmd = CommonCmd + `-crf 20 -c:v libx265 -preset veryslow %s.mp4`
const ImgToHeicCmd2 = CommonCmd + `-hide_banner -r 1 -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2,zscale=m=170m:r=pc" -pix_fmt yuv420p -frames 1 -c:v libx265 -preset veryslow -crf 20 -x265-params range=full:colorprim=smpte170m "%s.hevc"`
const ImgToHeicCmd3 = CommonCmd + `-hide_banner -r 1 -vf "scale=trunc(iw/2)*2:trunc(ih/2)*2,zscale=m=170m:r=pc" -pix_fmt yuv420p -frames 1 -c:v libx265 -preset veryslow -crf 20 -x265-params range=full:colorprim=smpte170m:aq-strength=1.2 -deblock -2:-2 "%s.hevc"
`

func ImgToHeic(filePath, dst string) error {
	if strings.HasSuffix(dst, ".heic") {
		dst = dst[:len(dst)-5]
	}
	err := Run(fmt.Sprintf(ImgToHeicCmd, filePath, dst))
	if err != nil {
		return err
	}

	return mp4box.Heic(dst+".mp4", dst)
}

const ImgToJxlCmd = CommonCmd + `-c:v libjxl "%s.jxl"`

// 不可用,没有注明色彩空间的原因。需要显式写明 像素编码格式、色彩空间、转换色彩空间、目标色彩空间、色彩范围
/*
distance
Set the target Butteraugli distance. This is a quality setting: lower distance yields higher quality, with distance=1.0 roughly comparable to libjpeg Quality 90 for photographic content. Setting distance=0.0 yields true lossless encoding. Valid values range between 0.0 and 15.0, and sane values rarely exceed 5.0. Setting distance=0.1 usually attains transparency for most input. The default is 1.0.

effort
Set the encoding effort used. Higher effort values produce more consistent quality and usually produces a better quality/bpp curve, at the cost of more CPU time required. Valid values range from 1 to 9, and the default is 7.

modular
Force the encoder to use Modular mode instead of choosing automatically. The default is to use VarDCT for lossy encoding and Modular for lossless. VarDCT is generally superior to Modular for lossy encoding but does not support lossless encoding.
*/
func ImgToJxl(filePath, dst string) error {
	if strings.HasSuffix(dst, ".jxl") {
		dst = dst[:len(dst)-4]
	}

	return Run(fmt.Sprintf(ImgToJxlCmd, filePath, dst))
}
