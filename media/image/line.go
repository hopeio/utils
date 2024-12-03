package image

import "image"

type LineCap int

const (
	// LineCapButt strokes do not extend beyond a line's two endpoints.
	LineCapButt LineCap = iota
	// LineCapRound strokes will be extended by a half circle with a diameter equal to the stroke width.
	LineCapRound
)

type Line struct {
	Start image.Point
	End   image.Point
}
