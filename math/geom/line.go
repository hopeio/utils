package geom

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Line struct {
	Start Point
	End   Point
}

func (l *Line) Vector() Vector {
	return Vector{l.End.X - l.Start.X, l.End.Y - l.Start.Y}
}

func (l *Line) StraightLine() *SlopeInterceptLine {
	var line SlopeInterceptLine
	if l.Start.X == l.End.X {
		// Vertical line
		line.M = math.Inf(1) // Positive infinity to indicate vertical line
		line.B = l.Start.X   // The x-coordinate of the vertical line
		return &line
	}

	// Calculate slope (m)
	line.M = (l.End.Y - l.Start.Y) / (l.End.X - l.Start.X)

	// Calculate y-intercept (b)
	line.B = l.Start.Y - line.M*l.Start.X
	return &line
}

type LineInt[T constraints.Integer] struct {
	Start PointInt[T]
	End   PointInt[T]
}

func (l *LineInt[T]) ToFloat64(factor float64) *Line {
	return &Line{
		Start: Point{float64(l.Start.X) / factor, float64(l.Start.Y) / factor},
		End:   Point{float64(l.End.X) / factor, float64(l.End.Y) / factor},
	}
}

func LineIntFromFloat64[T constraints.Integer](e *Line, factor float64) *LineInt[T] {
	return &LineInt[T]{
		Start: PointInt[T]{T(math.Round(e.Start.X * factor)), T(math.Round(e.Start.Y * factor))},
		End:   PointInt[T]{T(math.Round(e.End.X * factor)), T(math.Round(e.End.Y * factor))},
	}
}

// y=mx+b
type SlopeInterceptLine struct {
	M float64
	B float64
}

func (l *SlopeInterceptLine) IsVertical() bool {
	return math.IsInf(l.M, 0)
}

func (l *SlopeInterceptLine) ToGeneralFormLine() *StraightLine {
	if l.IsVertical() {
		// For vertical lines: x = k, convert to Ax + By + C = 0 where A = 1, B = 0, C = -k
		k := l.B
		return &StraightLine{A: 1, B: 0, C: -k}
	}

	return &StraightLine{A: l.M, B: -1, C: l.B}
}
func NewSlopeInterceptLine(m, b float64) *SlopeInterceptLine {
	return &SlopeInterceptLine{m, b}
}

// ax + by + c = 0
type StraightLine struct {
	A float64
	B float64
	C float64
}

func (l *StraightLine) ToSlopeInterceptLine() *SlopeInterceptLine {
	if l.B == 0 {
		return &SlopeInterceptLine{math.Inf(1), -l.C}
	}
	return &SlopeInterceptLine{
		M: -l.A / l.B,
		B: l.C / l.B,
	}
}

func NewGeneralFormLine(a, b, c float64) *StraightLine {
	return &StraightLine{a, b, c}
}
