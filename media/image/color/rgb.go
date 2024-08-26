package color

type RGB struct {
	R, G, B uint16
}

func NewRGB(r, g, b uint16) RGB {
	return RGB{R: r, G: g, B: b}
}
