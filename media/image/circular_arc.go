package image

import "image"

// counter clockwise if clockwise ,startX,startY <->  endX, endY
type CircularArc struct {
	Center image.Point
	Start  image.Point
	End    image.Point
}
