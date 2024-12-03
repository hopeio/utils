package geom

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Line struct {
	StartX float64
	StartY float64
	EndX   float64
	EndY   float64
}

func (l *Line) Vector() Vector {
	return Vector{l.EndX - l.StartX, l.EndY - l.StartY}
}

type LineInt[T constraints.Integer] struct {
	StartX T
	StartY T
	EndX   T
	EndY   T
}

func (l *LineInt[T]) ToFloat64(factor float64) *Line {
	return &Line{
		StartX: float64(l.StartX) / factor,
		StartY: float64(l.StartY) / factor,
		EndX:   float64(l.EndX) / factor,
		EndY:   float64(l.EndY) / factor,
	}
}

func LineIntFromFloat64[T constraints.Integer](e *Line, factor float64) *LineInt[T] {
	return &LineInt[T]{
		StartX: T(math.Round(e.StartX * factor)),
		StartY: T(math.Round(e.StartY * factor)),
		EndX:   T(math.Round(e.EndX * factor)),
		EndY:   T(math.Round(e.EndY * factor)),
	}
}
