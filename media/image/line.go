package image

type LineCap int

const (
	// LineCapButt strokes do not extend beyond a line's two endpoints.
	LineCapButt LineCap = iota
	// LineCapRound strokes will be extended by a half circle with a diameter equal to the stroke width.
	LineCapRound
)

type Line struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}
