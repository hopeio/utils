package geom

type Polygon []Point

type RegularPolygon struct {
	Centre Point
	Radius float64
	Sides  int
}
