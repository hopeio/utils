package raw

import (
	"fmt"
	colori "github.com/hopeio/utils/media/image/color"
	"image"
	"image/color"
)

type BGR struct {
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int // 3 * r.Dx()
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (r *BGR) ColorModel() color.Model {
	return colori.RGBModel
}

func (r *BGR) Bounds() image.Rectangle {
	return r.Rect
}

func (p *BGR) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (r *BGR) At(x, y int) color.Color {
	if !(image.Point{X: x, Y: y}.In(r.Rect)) {
		return colori.RGB{}
	}
	i := r.PixOffset(x, y)
	b, g, cr := r.Pix[i], r.Pix[i+1], r.Pix[i+2]
	return colori.RGB{R: cr, G: g, B: b}
}

func NewBGR(rawValues []byte, width, height int) (*BGR, error) {
	if len(rawValues) != width*height*3 {
		return nil, fmt.Errorf("invalid image raw data")
	}
	return &BGR{
		Pix:    rawValues,
		Stride: width * 3,
		Rect:   image.Rect(0, 0, width, height),
	}, nil
}
