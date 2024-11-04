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
