package gerber

import (
	"github.com/hopeio/utils/math/geometry"
	imagei "github.com/hopeio/utils/media/image"
	"image"
)

// A Segment is a stroked line.
type Segment struct {
	Interpolation Interpolation
	X             float64
	Y             float64
	CenterX       float64
	CenterY       float64
}

// A Contour is a closed sequence of connected linear or circular segments.
type Contour struct {
	Line     int
	X        float64
	Y        float64
	Segments []Segment
	Polarity bool
}

func (e *Contour) Bounds() *geometry.Bounds {
	bounds := image.Rectangle{Min: image.Point{X: int(e.X), Y: int(e.Y)}, Max: image.Point{int(e.X), int(e.Y)}}
	lastPoint := image.Point{X: int(e.X), Y: int(e.Y)}
	for _, s := range e.Segments {
		if s.Interpolation == InterpolationLinear {
			bounds = imagei.RectUnionPoint(bounds, image.Point{X: int(s.X), Y: int(s.Y)})
		} else {
			// 粗暴解决
			bounds = bounds.Union(image.Rect(lastPoint.X, lastPoint.Y, int(s.X), int(s.Y)).Add(image.Pt(int(s.X), int(s.Y))))
		}

	}
	return geometry.BoundsFromImageRect(bounds)
}

type Rectangle struct {
	Line     int
	Polarity bool
	geometry.Rectangle
}

func (e *Rectangle) Bounds() *geometry.Bounds {
	return e.Rectangle.Bounds()
}

type Obround struct {
	Line     int
	Polarity bool
	geometry.Rectangle
}

func (e *Obround) Bounds() *geometry.Bounds {
	return e.Rectangle.Bounds()
}

type Circle struct {
	Line     int
	Polarity bool
	geometry.Circle
}

func (e *Circle) Bounds() *geometry.Bounds {
	return e.Circle.Bounds()
}

type Arc struct {
	Line int
	geometry.Arc
	StrokeWidth float64
	Interpolation
}

func (e *Arc) Bounds() *geometry.Bounds {
	//TODO
	return nil
}

type Line struct {
	LineNo int
	geometry.Line
	StrokeWidth float64
	Cap         LineCap
}

func (e *Line) Bounds() *geometry.Bounds {
	if e.Cap == LineCapButt {
		vector := geometry.NewVector(geometry.Point{e.StartX, e.StartY}, geometry.Point{e.EndX, e.EndY})
		return geometry.NewRect((e.StartX+e.EndX)/2, (e.StartY+e.EndY)/2, vector.Length(), e.StrokeWidth, vector.Angle()).Bounds()
	}
	//TODO
	return nil
}
