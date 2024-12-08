package geom

type Interpolation int

const (
	InterpolationCCW Interpolation = iota
	InterpolationClockwise
)

// A Segment is a stroked line.
type Segment struct {
	Interpolation Interpolation
	End           Point
	Centre        Point
}

// A Contour is a closed sequence of connected linear or circular segments.
type Contour struct {
	Start    Point
	Segments []Segment
}
