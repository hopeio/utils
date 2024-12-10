package geom

import "C"
import (
	"golang.org/x/exp/constraints"
	"math"
)

type LineSegment struct {
	Start Point
	End   Point
}

func NewLineSegment(start, end Point) *LineSegment {
	return &LineSegment{start, end}
}

func (l *LineSegment) Vector() Vector {
	return Vector{l.End.X - l.Start.X, l.End.Y - l.Start.Y}
}

func (l *LineSegment) StraightLine() *SlopeInterceptFormLine {
	var line SlopeInterceptFormLine
	if l.Start.X == l.End.X {
		// Vertical line
		line.Slope = math.Inf(1)   // Positive infinity to indicate vertical line
		line.Intercept = l.Start.X // The x-coordinate of the vertical line
		return &line
	}

	// Calculate slope (m)
	line.Slope = (l.End.Y - l.Start.Y) / (l.End.X - l.Start.X)

	// Calculate y-intercept (b)
	line.Intercept = l.Start.Y - line.Slope*l.Start.X
	return &line
}

func (l *LineSegment) ContainsPoint(p Point) bool {
	// Ensure the point is within the bounding box of the line segment
	if math.Min(l.Start.X, l.End.X) <= p.X && p.X <= math.Max(l.Start.X, l.End.X) &&
		math.Min(l.Start.Y, l.End.Y) <= p.Y && p.Y <= math.Max(l.Start.Y, l.End.Y) {

		// Calculate the area of the triangle formed by the three points
		area := 0.5 * math.Abs(l.Start.X*l.End.Y+p.X*l.Start.Y+l.End.X*p.Y-p.X*l.End.Y-l.Start.X*p.Y-l.End.X*l.Start.Y)

		// If the area is effectively zero, the point is on the line
		return math.Abs(area) < 1e-9
	}

	return false
}

type LineInt[T constraints.Integer] struct {
	Start PointInt[T]
	End   PointInt[T]
}

func (l *LineInt[T]) ToFloat64(factor float64) *LineSegment {
	return &LineSegment{
		Start: Point{float64(l.Start.X) / factor, float64(l.Start.Y) / factor},
		End:   Point{float64(l.End.X) / factor, float64(l.End.Y) / factor},
	}
}

func LineIntFromFloat64[T constraints.Integer](e *LineSegment, factor float64) *LineInt[T] {
	return &LineInt[T]{
		Start: PointInt[T]{T(math.Round(e.Start.X * factor)), T(math.Round(e.Start.Y * factor))},
		End:   PointInt[T]{T(math.Round(e.End.X * factor)), T(math.Round(e.End.Y * factor))},
	}
}

// y=mx+b
type SlopeInterceptFormLine struct {
	Slope     float64
	Intercept float64
}

func (l *SlopeInterceptFormLine) IsVertical() bool {
	return math.IsInf(l.Slope, 0)
}

func (l *SlopeInterceptFormLine) ToGeneralFormLine() *GeneralFormLine {
	if l.IsVertical() {
		// For vertical lines: x = k, convert to Ax + By + C = 0 where A = 1, B = 0, C = -k
		k := l.Intercept
		return &GeneralFormLine{A: 1, B: 0, C: -k}
	}

	return &GeneralFormLine{A: l.Slope, B: -1, C: l.Intercept}
}
func NewSlopeInterceptLine(m, b float64) *SlopeInterceptFormLine {
	return &SlopeInterceptFormLine{m, b}
}

// ax + by + c = 0
type GeneralFormLine struct {
	A float64
	B float64
	C float64
}

func (l *GeneralFormLine) ToSlopeInterceptLine() *SlopeInterceptFormLine {
	if l.B == 0 {
		return &SlopeInterceptFormLine{math.Inf(1), -l.C}
	}
	return &SlopeInterceptFormLine{
		Slope:     -l.A / l.B,
		Intercept: l.C / l.B,
	}
}

func NewGeneralFormLine(a, b, c float64) *GeneralFormLine {
	return &GeneralFormLine{a, b, c}
}

type StraightLine struct {
	Point
	Angle float64
}

func NewStraightLine(p Point, angle float64) *StraightLine {
	return &StraightLine{p, angle}
}

func (l *StraightLine) ToGeneralFormLine() *GeneralFormLine {
	return &GeneralFormLine{math.Cos(l.Angle), math.Sin(l.Angle), -l.X*math.Cos(l.Angle) - l.Y*math.Sin(l.Angle)}
}

func (l *StraightLine) ToSlopeInterceptLine() *SlopeInterceptFormLine {
	return l.ToGeneralFormLine().ToSlopeInterceptLine()
}

func (l *StraightLine) ContainsPoint(p Point) bool {
	// Convert angle from degrees to radians
	angleInRadians := l.Angle * math.Pi / 180

	// Handle vertical line case (angle is 90 or 270 degrees)
	if math.Mod(math.Abs(l.Angle-90), 180) < 1e-9 || math.Mod(math.Abs(l.Angle-270), 360) < 1e-9 {
		return math.Abs(p.X-l.X) < 1e-9
	}

	// Calculate slope m and intercept b
	m := math.Tan(angleInRadians)
	b := l.Y - m*l.X

	// Check if the point satisfies the line equation within a small tolerance
	tolerance := 1e-9
	return math.Abs(p.Y-(m*p.X+b)) < tolerance
}

func (l *StraightLine) IntersectStraightLine(l2 *StraightLine) bool {
	sinL := math.Sin(l.Angle)
	sinL2 := math.Sin(l2.Angle)
	return sinL == sinL2 || sinL == -sinL2
}

type Ray struct {
	Point
	Angle float64
}

func NewRay(p Point, angle float64) *Ray {
	return &Ray{p, angle}
}
