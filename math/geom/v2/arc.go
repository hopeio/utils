package geom

import (
	"github.com/hopeio/utils/math/geom"
	"golang.org/x/exp/constraints"
	"math"
	"sort"
)

// counter clockwise if clockwise ,Start.X,Start.Y <->  endX, endY
type Arc struct {
	Center geom.Point
	Start  geom.Point
	End    geom.Point
}

func NewArc(center, start, end geom.Point) *Arc {
	return &Arc{
		Center: center,
		Start:  start,
		End:    end,
	}
}

func (a *Arc) Bounds() *Bounds {
	r := math.Hypot(a.Start.X-a.Center.X, a.Start.Y-a.Center.Y)
	thetaStart := math.Atan2(a.Start.Y-a.Center.Y, a.Start.X-a.Center.X)
	thetaEnd := math.Atan2(a.End.Y-a.Center.Y, a.End.X-a.Center.X)
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

	minX, maxX := a.Start.X, a.Start.X
	minY, maxY := a.Start.Y, a.Start.Y
	for _, theta := range angles {
		x := a.Center.X + r*math.Cos(theta)
		y := a.Center.Y + r*math.Sin(theta)
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

func (a *Arc) Sample(samples int) []geom.Point {

	r := math.Hypot(a.Start.X-a.Center.X, a.Start.Y-a.Center.Y)
	thetaStart := math.Atan2(a.Start.Y-a.Center.Y, a.Start.X-a.Center.X)
	thetaEnd := math.Atan2(a.End.Y-a.Center.Y, a.End.X-a.Center.X)
	if thetaStart > thetaEnd {
		thetaEnd += 2 * math.Pi
	}
	thetaDiff := thetaEnd - thetaStart
	points := make([]geom.Point, 0, samples)
	// Sample points along the arc
	for i := range samples {
		theta := thetaStart + thetaDiff*float64(i)/float64(samples)
		x := a.Center.X + r*math.Cos(theta)
		y := a.Center.Y + r*math.Sin(theta) // Flip y for image coordinate system
		points = append(points, geom.Point{x, y})
	}

	return points
}

// Compute the convex hull of a set of points using Andrew's monotone chain algorithm
func convexHull(points []geom.Point) []geom.Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	var lower, upper []geom.Point
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
func cross(o, a, b geom.Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

// Rotating calipers to find the minimum bounding rectangle
// untested
func minimumBoundingRectangle(hull []geom.Point) *Rectangle {
	minArea := math.MaxFloat64
	var centerX, centerY, length, width, angle float64
	for i := 0; i < len(hull); i++ {
		// Edge vector
		edge := geom.Point{hull[(i+1)%len(hull)].X - hull[i].X, hull[(i+1)%len(hull)].Y - hull[i].Y}
		edgeLength := math.Hypot(edge.X, edge.Y)
		edge.X /= edgeLength
		edge.Y /= edgeLength

		// Perpendicular vector
		perp := geom.Point{-edge.Y, edge.X}

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

	return NewRect(geom.Point{centerX, centerY}, length, width, angle)
}

// untested
func (a *Arc) minimumBoundingRectangle() *Rectangle {
	points := a.Sample(100)
	hull := convexHull(points)
	return minimumBoundingRectangle(hull)
}

type ArcInt[T constraints.Integer] struct {
	Center geom.PointInt[T]
	Start  geom.PointInt[T]
	End    geom.PointInt[T]
}

func (e *ArcInt[T]) ToFloat64(factor float64) *Arc {
	if factor == 0 {
		factor = 1
	}
	return &Arc{
		Center: geom.Point{float64(e.Center.X) / factor, float64(e.Center.Y) / factor},
		Start:  geom.Point{float64(e.Start.X) / factor, float64(e.Start.Y) / factor},
		End:    geom.Point{float64(e.End.X) / factor, float64(e.End.Y) / factor},
	}
}

func ArcIntFromFloat64[T constraints.Integer](e *Arc, factor float64) *ArcInt[T] {
	if factor == 0 {
		factor = 1
	}
	return &ArcInt[T]{
		Center: geom.PointInt[T]{T(math.Round(e.Center.X * factor)), T(math.Round(e.Center.Y * factor))},
		Start:  geom.PointInt[T]{T(math.Round(e.Start.X * factor)), T(math.Round(e.Start.Y * factor))},
		End:    geom.PointInt[T]{T(math.Round(e.End.X * factor)), T(math.Round(e.End.Y * factor))},
	}
}
