package geom

import "golang.org/x/exp/constraints"

type Circle struct {
	Center   Point
	Diameter float64
}

func (c *Circle) Bounds() *Bounds {
	return NewBounds(c.Center.X, c.Center.Y, c.Diameter, c.Diameter)
}

type CircleInt[T constraints.Integer] struct {
	Center   PointInt[T]
	Diameter T
}
