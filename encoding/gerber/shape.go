package gerber

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
type Obround struct {
	Line     int
	X        int
	Y        int
	Width    int
	Height   int
	Polarity bool
	Rotation float64
}

type Circle struct {
	Line     int
	X        int
	Y        int
	Diameter int
	Polarity bool
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
