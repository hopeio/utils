package image

import (
	"image"
	"math"
)

// clockwise if counter clockwise ,startX,startY <->  endX, endY
type Arc struct {
	CenterX int
	CenterY int
	StartX  int
	StartY  int
	EndX    int
	EndY    int
}

func NewArc(centerX, centerY, startX, startY, endX, endY int) *Arc {
	return &Arc{
		EndX:    endX,
		EndY:    endY,
		StartX:  startX,
		StartY:  startY,
		CenterX: centerX,
		CenterY: centerY,
	}
}

// 与平面坐标系的区别就是-y (平移到第四象限或者直接镜像过去)
func (e *Arc) Bounds() image.Rectangle {
	r := math.Hypot(float64(e.StartX-e.CenterX), float64(e.StartY-e.CenterY))
	thetaStart := math.Atan2(float64(e.CenterY-e.StartY), float64(e.StartX-e.CenterX))
	thetaEnd := math.Atan2(float64(e.CenterY-e.EndY), float64(e.EndX-e.CenterX))
	if thetaStart < 0 {
		thetaStart += 2 * math.Pi
	}
	if thetaEnd < 0 {
		thetaEnd += 2 * math.Pi
	}

	angles := []float64{thetaStart, thetaEnd}
	for _, a := range []float64{math.Pi / 2, math.Pi, 3 * math.Pi / 2, math.Pi * 2} {
		if thetaStart > thetaEnd {
			if a >= thetaStart || a <= thetaEnd {
				angles = append(angles, a)
			}
		} else {
			if thetaStart <= a && a <= thetaEnd {
				angles = append(angles, a)
			}
		}
	}

	minX, maxX := e.StartX, e.StartX
	minY, maxY := e.StartY, e.StartY
	for _, theta := range angles {
		x := e.CenterX + int(math.Round(r*math.Cos(theta)))
		y := e.CenterY - int(math.Round(r*math.Sin(theta)))
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	return image.Rect(minX, minY, maxX, maxY)
}
