package gerber

import (
	"github.com/hopeio/utils/math/geometry"
	imagei "github.com/hopeio/utils/media/image"
	"image"
	"math"
)

// A Segment is a stroked line.
type Segment struct {
	Interpolation Interpolation
	X             int
	Y             int
	CenterX       int
	CenterY       int
}

// A Contour is a closed sequence of connected linear or circular segments.
type Contour struct {
	Line     int
	X        int
	Y        int
	Segments []Segment
	Polarity bool
}

func (e *Contour) Bounds() image.Rectangle {
	bounds := image.Rectangle{Min: image.Point{X: math.MaxInt, Y: math.MaxInt}, Max: image.Point{-math.MaxInt, -math.MaxInt}}
	lastPoint := image.Point{X: e.X, Y: e.Y}
	for _, s := range e.Segments {
		if s.Interpolation == InterpolationLinear {
			bounds = imagei.RectUnionPoint(bounds, image.Point{X: s.X, Y: s.Y})
		} else {
			// 粗暴解决
			bounds = bounds.Union(image.Rect(lastPoint.X, lastPoint.Y, s.X, s.Y).Add(image.Pt(s.X, s.Y)))
		}

	}
	return bounds
}

type Rectangle struct {
	Line     int
	Polarity bool
	geometry.Rectangle
}

func (e *Rectangle) Bounds() *geometry.Rectangle {
	return &e.Rectangle
}

type Obround struct {
	Line     int
	Polarity bool
	geometry.Rectangle
}

func (e *Obround) Bounds() *geometry.Rectangle {
	return &e.Rectangle
}

type Circle struct {
	Line     int
	Polarity bool
	geometry.Circle
}

func (e *Circle) Bounds() *geometry.Rectangle {
	return e.Circle.Bounds()
}

type Arc struct {
	Line int
	geometry.Arc
	StrokeWidth float64
	Interpolation
}

func (e *Arc) Bounds() *geometry.Rectangle {
	// 粗暴解决
	return image.Rect(e.StartX, e.StartY, e.EndX, e.EndY).Add(image.Pt(e.CenterX, e.CenterY))
}

type Line struct {
	LineNo int
	geometry.Line
	StrokeWidth float64
	Cap         LineCap
}

func (e *Line) Bounds() *geometry.Rectangle {
	return image.Rect(e.StartX-e.StrokeWidth/2, e.StartY-e.StrokeWidth/2, e.EndX+e.StrokeWidth/2, e.EndY+e.StrokeWidth/2)
}
