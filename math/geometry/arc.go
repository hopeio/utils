package geometry

import (
	"math"
)

// clockwise if counter clockwise ,startX,startY <->  endX, endY
type Arc struct {
	CenterX float64
	CenterY float64
	StartX  float64
	StartY  float64
	EndX    float64
	EndY    float64
}

func NewArc(centerX, centerY, startX, startY, endX, endY float64) *Arc {
	return &Arc{
		EndX:    endX,
		EndY:    endY,
		StartX:  startX,
		StartY:  startY,
		CenterX: centerX,
		CenterY: centerY,
	}
}

func (e *Arc) Bounds() *Rectangle {
	r := math.Hypot(e.StartX-e.CenterX, e.StartY-e.CenterY)
	thetaStart := math.Atan2(e.StartY-e.CenterY, e.StartX-e.CenterX)
	thetaEnd := math.Atan2(e.EndY-e.CenterY, e.EndX-e.CenterX)
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
		x := e.CenterX + r*math.Cos(theta)
		y := e.CenterY + r*math.Sin(theta)
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
	return RectNoRotate(minX, minY, maxX, maxY)
}
