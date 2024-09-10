package image

import (
	"image"
	"io"
	"os"
)

type Decode func(r io.Reader) (image.Image, error)
type Encode func(io.Writer, image.Image) error

func Transfer(src, dst string, decode Decode, encode Encode) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	img, err := decode(file)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	dstImg, err := os.Create(dst)
	if err != nil {
		return err
	}
	err = encode(dstImg, img)
	if err != nil {
		return err
	}
	return dstImg.Close()
}
