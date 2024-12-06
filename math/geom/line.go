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

// ax + by + c = 0
type StraightLine struct {
	X float64
	Y float64
	C float64
}

func NewStraightLine(x, y, c float64) StraightLine {
	return StraightLine{x, y, c}
}
