package geometry

import (
	"golang.org/x/exp/constraints"
	"math"
	"sort"
)

// counter clockwise if clockwise ,startX,startY <->  endX, endY
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

func (a *Arc) Bounds() *Bounds {
	r := math.Hypot(a.StartX-a.CenterX, a.StartY-a.CenterY)
	thetaStart := math.Atan2(a.StartY-a.CenterY, a.StartX-a.CenterX)
	thetaEnd := math.Atan2(a.EndY-a.CenterY, a.EndX-a.CenterX)
	if thetaStart < 0 {
		thetaStart += 2 * math.Pi
	}
	if thetaEnd < 0 {
		thetaEnd += 2 * math.Pi
	}

	angles := []float64{thetaStart, thetaEnd}
	for _, a := range []float64{math.Pi / 2, math.Pi, 3 * math.Pi / 2, math.Pi * 2} {
		if thetaStart < thetaEnd {
			if thetaStart <= a && a <= thetaEnd {
				angles = append(angles, a)
			}
		} else {
			if a >= thetaStart || a <= thetaEnd {
				angles = append(angles, a)
			}
		}
	}

	minX, maxX := a.StartX, a.StartX
	minY, maxY := a.StartY, a.StartY
	for _, theta := range angles {
		x := a.CenterX + r*math.Cos(theta)
		y := a.CenterY + r*math.Sin(theta)
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
	return NewBounds(minX, minY, maxX, maxY)
}

func (a *Arc) Sample(samples int) []Point {

	r := math.Hypot(a.StartX-a.CenterX, a.StartY-a.CenterY)
	thetaStart := math.Atan2(a.StartY-a.CenterY, a.StartX-a.CenterX)
	thetaEnd := math.Atan2(a.EndY-a.CenterY, a.EndX-a.CenterX)
	if thetaStart > thetaEnd {
		thetaEnd += 2 * math.Pi
	}
	thetaDiff := thetaEnd - thetaStart
	points := make([]Point, 0, samples)
	// Sample points along the arc
	for i := range samples {
		theta := thetaStart + thetaDiff*float64(i)/float64(samples)
		x := a.CenterX + r*math.Cos(theta)
		y := a.CenterY + r*math.Sin(theta) // Flip y for image coordinate system
		points = append(points, Point{x, y})
	}

	return points
}

// Compute the convex hull of a set of points using Andrew's monotone chain algorithm
func convexHull(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	var lower, upper []Point
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}

// Cross product of vectors OA and OB
func cross(o, a, b Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

// Rotating calipers to find the minimum bounding rectangle
// untested
func minimumBoundingRectangle(hull []Point) *Rectangle {
	minArea := math.MaxFloat64
	var centerX, centerY, length, width, angle float64
	for i := 0; i < len(hull); i++ {
		// Edge vector
		edge := Point{hull[(i+1)%len(hull)].X - hull[i].X, hull[(i+1)%len(hull)].Y - hull[i].Y}
		edgeLength := math.Hypot(edge.X, edge.Y)
		edge.X /= edgeLength
		edge.Y /= edgeLength

		// Perpendicular vector
		perp := Point{-edge.Y, edge.X}

		// Project points onto edge and perpendicular
		var minProj, maxProj float64
		var minPerp, maxPerp float64
		for j, p := range hull {
			proj := edge.X*(p.X-hull[i].X) + edge.Y*(p.Y-hull[i].Y)
			perpProj := perp.X*(p.X-hull[i].X) + perp.Y*(p.Y-hull[i].Y)
			if j == 0 || proj < minProj {
				minProj = proj
			}
			if j == 0 || proj > maxProj {
				maxProj = proj
			}
			if j == 0 || perpProj < minPerp {
				minPerp = perpProj
			}
			if j == 0 || perpProj > maxPerp {
				maxPerp = perpProj
			}
		}

		// Compute area of rectangle
		widthTemp := maxPerp - minPerp
		lengthTemp := maxProj - minProj
		area := widthTemp * lengthTemp
		if area < minArea {
			minArea = area
			centerX = hull[i].X + (minProj+maxProj)/2*edge.X
			centerY = hull[i].Y + (minProj+maxProj)/2*edge.Y
			length = lengthTemp
			width = widthTemp
			angle = math.Atan2(edge.Y, edge.X)
		}
	}

	return NewRect(centerX, centerY, length, width, angle)
}

// untested
func (a *Arc) minimumBoundingRectangle() *Rectangle {
	points := a.Sample(100)
	hull := convexHull(points)
	return minimumBoundingRectangle(hull)
}

type ArcInt[T constraints.Integer] struct {
	CenterX T
	CenterY T
	StartX  T
	StartY  T
	EndX    T
	EndY    T
}

func (e *ArcInt[T]) ToFloat64(factor float64) *Arc {
	if factor == 0 {
		factor = 1
	}
	return &Arc{
		CenterX: float64(e.CenterX) / factor,
		CenterY: float64(e.CenterY) / factor,
		StartX:  float64(e.StartX) / factor,
		StartY:  float64(e.StartY) / factor,
		EndX:    float64(e.EndX) / factor,
		EndY:    float64(e.EndY) / factor,
	}
}

func ArcIntFromFloat64[T constraints.Integer](e *Arc, factor float64) *ArcInt[T] {
	if factor == 0 {
		factor = 1
	}
	return &ArcInt[T]{
		CenterX: T(math.Round(e.CenterX * factor)),
		CenterY: T(math.Round(e.CenterY * factor)),
		StartX:  T(math.Round(e.StartX * factor)),
		StartY:  T(math.Round(e.StartY * factor)),
		EndX:    T(math.Round(e.EndX * factor)),
		EndY:    T(math.Round(e.EndY * factor)),
	}
}
