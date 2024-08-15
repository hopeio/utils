package gerber

import (
	imagei "github.com/hopeio/utils/media/image"
	"image"
	"math"
)

// A Segment is a stroked line.
type Segment struct {
	Interpolation Interpolation
	X             int
	Y             int
	XCenter       int
	YCenter       int
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
	for _, s := range e.Segments {
		bounds = imagei.Union(bounds, image.Point{X: s.X, Y: s.Y})
	}
	return imagei.Union(bounds, image.Point{X: e.X, Y: e.Y})
}

type Rectangle struct {
	Line     int
	X        int
	Y        int
	Width    int
	Height   int
	XCenter  int
	YCenter  int
	Polarity bool
	Rotation float64
}

func (e *Rectangle) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{X: e.X - e.Width/2, Y: e.Y - e.Height/2}, Max: image.Point{X: e.X + e.Width/2,
		Y: e.Y + e.Height/2}}
}

type Obround struct {
	Line     int
	X        int
	Y        int
	Width    int
	Height   int
	Polarity bool
	Rotation float64
}

func (e *Obround) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{X: e.X - e.Width/2, Y: e.Y - e.Height/2}, Max: image.Point{X: e.X + e.Width/2,
		Y: e.Y + e.Height/2}}
}

type Circle struct {
	Line     int
	X        int
	Y        int
	Diameter int
	Polarity bool
}

func (e *Circle) Bounds() image.Rectangle {
	return image.Rectangle{Min: image.Point{X: e.X - e.Diameter/2, Y: e.Y - e.Diameter/2}, Max: image.Point{X: e.X + e.Diameter/2,
		Y: e.Y + e.Diameter/2}}
}

type Arc struct {
	Line    int
	XEnd    int
	YEnd    int
	XStart  int
	YStart  int
	XCenter int
	YCenter int
	Width   int
	Interpolation
}

func (e *Arc) Bounds() image.Rectangle {
	return image.Rect(e.XStart, e.YStart, e.XEnd, e.YEnd)
}

type Line struct {
	Line     int
	XStart   int
	YStart   int
	XEnd     int
	YEnd     int
	Width    int
	Cap      LineCap
	Rotation float64
}

func (e *Line) Bounds() image.Rectangle {
	return image.Rect(e.XStart, e.YStart, e.XEnd, e.YEnd)
}
