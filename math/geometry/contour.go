package geometry

type Interpolation int

const (
	InterpolationCCW Interpolation = iota
	InterpolationClockwise
)

// A Segment is a stroked line.
type Segment struct {
	Interpolation Interpolation
	X             float64
	Y             float64
	XCenter       float64
	YCenter       float64
}

// A Contour is a closed sequence of connected linear or circular segments.
type Contour struct {
	X        float64
	Y        float64
	Segments []Segment
}
