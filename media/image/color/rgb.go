package color

import "image/color"

type RGB64 struct {
	R, G, B uint16
}

func (c RGB64) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xffff
	return
}

func NewRGB64(r, g, b uint16) RGB64 {
	return RGB64{R: r, G: g, B: b}
}

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xffff
	return
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
}

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

var (
	RGBModel = color.ModelFunc(rgbModel)
)
