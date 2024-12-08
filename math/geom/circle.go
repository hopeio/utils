package geom

import "golang.org/x/exp/constraints"

type Circle struct {
	Centre   Point
	Diameter float64
}

func (c *Circle) Bounds() *Bounds {
	return NewBounds(c.Centre.X, c.Centre.Y, c.Diameter, c.Diameter)
}

type CircleInt[T constraints.Integer] struct {
	Center   PointInt[T]
	Diameter T
}

func (e *CircleInt[T]) ToFloat64(factor float64) *Circle {
	if factor == 0 {
		factor = 1
	}
	return &Circle{
		Centre:   Point{float64(e.Center.X) / factor, float64(e.Center.Y) / factor},
		Diameter: float64(e.Diameter) / factor,
	}
}
