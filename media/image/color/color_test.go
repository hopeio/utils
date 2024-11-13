package color

import (
	"image/color"
	"testing"
)

func TestColor(t *testing.T) {
	c := color.RGBA{
		R: 155,
		G: 233,
		B: 65,
		A: 255,
	}
	t.Log(ColorToGray(c))
	t.Log(ColorToGray2(c))
	t.Log(RGBAToGray(c))
}

func TestRGB(t *testing.T) {
	r := uint8(22)
	r32 := uint32(r)
	r32 |= r32 << 8
	t.Log(r32)
	t.Log(uint8(r32))
}
