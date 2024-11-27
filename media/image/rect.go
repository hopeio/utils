package image

import (
	"image"
)

func RectUnionPoint(rect image.Rectangle, p image.Point) image.Rectangle {
	if p.X < rect.Min.X {
		rect.Min.X = p.X
	}
	if p.X > rect.Max.X {
		rect.Max.X = p.X
	}
	if p.Y < rect.Min.Y {
		rect.Min.Y = p.Y
	}
	if p.Y > rect.Max.Y {
		rect.Max.Y = p.Y
	}
	return rect
}
