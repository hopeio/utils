// Package svg parses Gerber to SVG.
package svg

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/encoding/gerber"
	psvg "github.com/hopeio/utils/media/image/svg"
	"image"
	"io"
	"math"
	"strconv"
	"strings"
)

// An ElementType is a SVG element type.
type ElementType string

const (
	ElementTypeCircle    ElementType = "Circle"
	ElementTypeRectangle ElementType = "Rect"
	ElementTypePath      ElementType = "Path"
	ElementTypeLine      ElementType = "Line"
	ElementTypeArc       ElementType = "Arc"
)

// A Circle is a circle.
type Circle struct {
	Type ElementType
	gerber.Circle
	Fill string
	Attr map[string]string
}

func (e Circle) Bounds() image.Rectangle {
	radius := e.Diameter / 2
	return image.Rect(e.X-radius, e.Y-radius, e.X+radius, e.Y+radius)
}

// MarshalJSON implements json.Marshaler.
func (e Circle) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeCircle
	return json.Marshal(e)
}

func (e Circle) SetAttr(k, v string) Circle {
	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[k] = v
	return e
}

// A Rectangle is a rectangle.
type Rectangle struct {
	Type     ElementType
	Aperture string
	gerber.Rectangle
	RadiusX int
	RadiusY int
	Fill    string
	Attr    map[string]string
}

func (e Rectangle) Bounds() image.Rectangle {
	return image.Rect(e.X, e.Y, e.X+e.Width, e.Y+e.Height)
}

// MarshalJSON implements json.Marshaler.
func (e Rectangle) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeRectangle
	return json.Marshal(e)
}

func (e Rectangle) SetAttr(k, v string) Rectangle {
	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[k] = v
	return e
}

// A PathLine is a line in a SVG path.
type PathLine struct {
	Type ElementType
	X    int
	Y    int
}

// MarshalJSON implements json.Marshaler.
func (e PathLine) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeLine
	return json.Marshal(e)
}

// A PathArc is an arc in a SVG path.
type PathArc struct {
	Type     ElementType
	RadiusX  int
	RadiusY  int
	LargeArc int
	Sweep    int
	X        int
	Y        int
}

// MarshalJSON implements json.Marshaler.
func (e PathArc) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeArc
	return json.Marshal(e)
}

// A Path is a SVG path.
type Path struct {
	Type     ElementType
	Line     int
	X        int
	Y        int
	Commands []interface{}
	Fill     string
	Attr     map[string]string
}

func (e Path) Bounds() (image.Rectangle, error) {
	bounds := image.Rectangle{Min: image.Point{math.MaxInt, math.MaxInt}, Max: image.Point{-math.MaxInt, -math.MaxInt}}
	updateMinMax := func(x, y int) {
		bounds.Min.X = min(bounds.Min.X, x)
		bounds.Max.X = max(bounds.Max.X, x)
		bounds.Min.Y = min(bounds.Min.Y, y)
		bounds.Max.Y = max(bounds.Max.Y, y)
	}

	updateMinMax(e.X, e.Y)
	for _, cmd := range e.Commands {
		switch c := cmd.(type) {
		case PathLine:
			updateMinMax(c.X, c.Y)
		case PathArc:
			updateMinMax(c.X, c.Y)
		default:
			return image.Rectangle{}, fmt.Errorf("%#v", c)
		}
	}

	return bounds, nil
}

// MarshalJSON implements json.Marshaler.
func (e Path) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypePath
	return json.Marshal(e)
}

func (e Path) SetAttr(k, v string) Path {
	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[k] = v
	return e
}

// A Line is a SVG line.
type Line struct {
	Type ElementType
	gerber.Line
	Stroke string
	Attr   map[string]string
}

func (e Line) Bounds() image.Rectangle {
	return image.Rect(e.StartX, e.StartY, e.EndX, e.EndY)
}

// MarshalJSON implements json.Marshaler.
func (e Line) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeLine
	return json.Marshal(e)
}

func (e Line) SetAttr(k, v string) Line {
	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[k] = v
	return e
}

// An Arc is a SVG Arc.
type Arc struct {
	Type ElementType
	gerber.Arc
	RadiusX int
	RadiusY int

	LargeArc int
	Sweep    int

	Stroke string
	Attr   map[string]string
}

func (e Arc) Bounds() image.Rectangle {
	return image.Rect(min(e.StartX, e.EndX), max(e.StartY, e.EndY), max(e.StartX, e.EndX), max(e.StartY, e.EndY))
}

// MarshalJSON implements json.Marshaler.
func (e Arc) MarshalJSON() ([]byte, error) {
	e.Type = ElementTypeArc
	return json.Marshal(e)
}

func (e Arc) SetAttr(k, v string) Arc {
	if e.Attr == nil {
		e.Attr = make(map[string]string)
	}
	e.Attr[k] = v
	return e
}

// A Processor is a performer of Gerber graphic operations.
type Processor struct {
	// Data contains SVG elements.
	Data []interface{}

	// Viewbox of Gerber image.
	MinX int
	MaxX int
	MinY int
	MaxY int

	// Color for Gerber polarities, defaults to black and white.
	PolarityDark  string
	PolarityClear string

	// Optional scaling factor of coordinates when writing SVG image.
	Scale float64

	// Optional width and height of output SVG image.
	Width  string
	Height string

	// Whether to output javascript for interactive panning and zooming in SVG.
	PanZoom bool
}

// NewProcessor creates a Processor.
func NewProcessor() *Processor {
	p := &Processor{}
	p.Data = make([]interface{}, 0)
	p.Scale = 1
	p.PolarityDark = "white"
	p.PolarityClear = "black"
	p.PanZoom = false
	return p
}

func (p *Processor) fill(polarity bool) string {
	if polarity {
		return p.PolarityDark
	}
	return p.PolarityClear
}

func (p *Processor) Circle(circle gerber.Circle) {
	p.Data = append(p.Data, Circle{Circle: circle, Fill: p.fill(circle.Polarity)})
}

func (p *Processor) Rectangle(rectangle gerber.Rectangle) {
	rectangle.X -= rectangle.Width / 2
	rectangle.Y += rectangle.Height / 2
	p.Data = append(p.Data, Rectangle{Aperture: "R", Rectangle: rectangle, Fill: p.fill(rectangle.Polarity)})
}

func (p *Processor) Obround(obround gerber.Obround) {
	r := min(obround.Width, obround.Height) / 2
	p.Data = append(p.Data, Rectangle{Aperture: "O", Rectangle: gerber.Rectangle{Line: obround.Line,
		X: obround.X - obround.Width/2,
		Y: obround.Y + obround.Height/2, Width: obround.Width,
		Height: obround.Height, Rotation: obround.Rotation}, RadiusX: r, RadiusY: r, Fill: p.fill(obround.Polarity)})
}

func (p *Processor) Contour(contour gerber.Contour) error {
	if len(contour.Segments) == 1 {
		s := contour.Segments[0]
		if s.Interpolation == gerber.InterpolationClockwise || s.Interpolation == gerber.InterpolationCCW {
			if s.X == contour.X && s.Y == contour.Y {
				vx, vy := float64(s.X-s.XCenter), float64(s.Y-s.YCenter)
				r := int(math.Round(math.Sqrt(vx*vx + vy*vy)))
				c := Circle{Circle: gerber.Circle{Line: contour.Line, X: s.XCenter, Y: s.YCenter, Diameter: r * 2},
					Fill: p.fill(contour.Polarity)}
				p.Data = append(p.Data, c)
				return nil
			}
		}
	}

	svgPath := Path{Line: contour.Line, X: contour.X, Y: contour.Y, Fill: p.fill(contour.Polarity)}
	for i, s := range contour.Segments {
		switch s.Interpolation {
		case gerber.InterpolationLinear:
			svgPath.Commands = append(svgPath.Commands, PathLine{X: s.X, Y: s.Y})
		case gerber.InterpolationClockwise:
			fallthrough
		case gerber.InterpolationCCW:
			arc, err := calcArc(contour, i)
			if err != nil {
				return err
			}
			svgPath.Commands = append(svgPath.Commands, arc)
		default:
			return fmt.Errorf("%d %+v", i, s)
		}
	}
	p.Data = append(p.Data, svgPath)
	return nil
}

func calcArcParams(vs, ve [2]int, sweep int) (float64, int, error) {
	radiusS := math.Sqrt(math.Pow(float64(vs[0]), 2) + math.Pow(float64(vs[1]), 2))
	radiusE := math.Sqrt(math.Pow(float64(ve[0]), 2) + math.Pow(float64(ve[1]), 2))
	diff := math.Abs(radiusS - radiusE)
	diffRatio := math.Abs(radiusS/radiusE - 1)
	if diff > 3 && diffRatio > 1e-2 {
		return math.NaN(), -1, fmt.Errorf("%f %f %f %f", radiusS, radiusE, diff, diffRatio)
	}

	var largeArc int
	cross := vs[0]*ve[1] - ve[0]*vs[1]
	if (cross > 0) != (sweep == 0) {
		largeArc = 1
	}

	return radiusS, largeArc, nil
}

func calcArc(contour gerber.Contour, idx int) (PathArc, error) {
	var xs, ys int
	if idx == 0 {
		xs, ys = contour.X, contour.Y
	} else {
		prev := contour.Segments[idx-1]
		xs, ys = prev.X, prev.Y
	}

	s := contour.Segments[idx]
	arc := PathArc{X: s.X, Y: s.Y}
	switch s.Interpolation {
	case gerber.InterpolationClockwise:
		arc.Sweep = 1
	case gerber.InterpolationCCW:
		arc.Sweep = 0
	default:
		return PathArc{}, fmt.Errorf("%d", s.Interpolation)
	}

	vs := [2]int{xs - s.XCenter, ys - s.YCenter}
	ve := [2]int{s.X - s.XCenter, s.Y - s.YCenter}
	if ve == vs {
		return PathArc{}, fmt.Errorf("degenerate arc")
	}

	radius, largeArc, err := calcArcParams(vs, ve, arc.Sweep)
	if err != nil {
		return PathArc{}, fmt.Errorf("%#d %#d %#v %w", xs, ys, s, err)
	}
	arc.RadiusX, arc.RadiusY = int(math.Round(radius)), int(math.Round(radius))
	arc.LargeArc = largeArc

	return arc, nil
}

func (p *Processor) Line(gline gerber.Line) {
	line := Line{Line: gline, Stroke: p.PolarityDark}
	p.Data = append(p.Data, line)
}

func (p *Processor) Arc(garc gerber.Arc) error {
	if garc.EndX == garc.StartX && garc.EndY == garc.StartY {
		return fmt.Errorf("degenerate arc")
	}

	arc := Arc{Arc: garc, Stroke: p.PolarityDark}
	switch garc.Interpolation {
	case gerber.InterpolationClockwise:
		arc.Sweep = 1
	case gerber.InterpolationCCW:
		arc.Sweep = 0
	default:
		return fmt.Errorf("%d", garc.Interpolation)
	}

	vs := [2]int{garc.StartX - garc.CenterX, garc.StartY - garc.CenterY}
	ve := [2]int{garc.EndX - garc.CenterX, garc.EndY - garc.CenterY}

	radius, largeArc, err := calcArcParams(vs, ve, arc.Sweep)
	if err != nil {
		return err
	}
	arc.RadiusX, arc.RadiusY = int(math.Round(radius)), int(math.Round(radius))
	arc.LargeArc = largeArc

	p.Data = append(p.Data, arc)
	return nil
}

func (p *Processor) SetViewbox(minX, maxX, minY, maxY int) {
	p.MinX = minX
	p.MaxX = maxX
	p.MinY = minY
	p.MaxY = maxY
}

// Write writes Gerber graphics operations as SVG.
func (p *Processor) Write(w io.Writer) error {
	svg := "<svg "
	if p.Width != "" && p.Height != "" {
		svg += fmt.Sprintf(`width="%s" height="%s" `, p.Width, p.Height)
	}
	svg += fmt.Sprintf(`viewBox="%s %s %s %s" style="background-color: %s;" xmlns="http://www.w3.org/2000/svg">`+"\n", p.x(p.MinX), p.y(p.MaxY), p.m(p.MaxX-p.MinX), p.m(p.MaxY-p.MinY), p.PolarityClear)
	if _, err := w.Write([]byte(svg)); err != nil {
		return err
	}

	if p.PanZoom {
		if _, err := w.Write([]byte(`<script xlink:href="svgpan.js"/><g id="viewport" transform="translate(0, 0)">` + "\n")); err != nil {
			return err
		}
	}

	svgBound := image.Rect(p.MinX, p.MinY, p.MaxX, p.MaxY)
	for _, datum := range p.Data {
		bounds, err := Bounds(datum)
		if err != nil {
			return err
		}
		if bounds.Min.X > svgBound.Max.X || svgBound.Min.X > bounds.Max.X || bounds.Min.Y > svgBound.Max.Y || svgBound.Min.Y > bounds.Max.Y {
			continue
		}

		var b []byte
		switch d := datum.(type) {
		case Circle:
			b = []byte(fmt.Sprintf(`<circle cx="%s" cy="%s" r="%s" fill="%s" %s/>`, p.x(d.X), p.y(d.Y),
				p.m(d.Diameter/2),
				d.Fill, psvg.FormatAttr(d.Attr)))
		case Rectangle:
			w, h := p.m(d.Width), p.m(d.Height)
			b = []byte(fmt.Sprintf(`<rect x="%s" y="%s" width="%s" height="%s" rx="%s" ry="%s" fill="%s" transform
="rotate(%.1f, %s, %s)" %s/>`, p.x(d.X), p.y(d.Y), w, h, p.m(d.RadiusX), p.m(d.RadiusY), d.Fill, d.Rotation, p.x(d.X+d.Width/2),
				p.y(d.Y-d.Height/2), psvg.FormatAttr(d.Attr)))
		case Path:
			var err error
			b, err = p.pathBytes(d)
			if err != nil {
				return err
			}
		case Line:
			b = []byte(fmt.Sprintf(`<line x1="%s" y1="%s" x2="%s" y2="%s" stroke-width="%s" stroke-linecap="%s" stroke="%s"%s/>`, p.x(d.StartX), p.y(d.StartY), p.x(d.EndX), p.y(d.EndY), p.m(d.Width), d.Cap, d.Stroke, psvg.FormatAttr(d.Attr)))
		case Arc:
			b = []byte(fmt.Sprintf(`<path d="M %s %s A %s %s 0 %d %d %s %s" stroke-width="%s" stroke="%s" stroke
-linecap="round"%s/>`, p.x(d.StartX), p.y(d.StartY), p.m(d.RadiusX), p.m(d.RadiusY), d.LargeArc, d.Sweep, p.x(d.EndX), p.y(d.EndY), p.m(d.Width), d.Stroke, psvg.FormatAttr(d.Attr)))
		default:
			return fmt.Errorf("%+v", d)
		}
		if _, err := w.Write(append(b, '\n')); err != nil {
			return err
		}
	}

	if p.PanZoom {
		if _, err := w.Write([]byte(`</g>`)); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte(`</svg>`)); err != nil {
		return err
	}
	return nil
}

func Bounds(element interface{}) (image.Rectangle, error) {
	switch e := element.(type) {
	case Circle:
		return e.Bounds(), nil
	case Rectangle:
		return e.Bounds(), nil
	case Path:
		return e.Bounds()
	case Line:
		return e.Bounds(), nil
	case Arc:
		return e.Bounds(), nil
	default:
		return image.Rectangle{}, fmt.Errorf("%#v", e)
	}
}

func (p *Processor) pathBytes(svgp Path) ([]byte, error) {
	cmds := []string{fmt.Sprintf("M %s %s", p.x(svgp.X), p.y(svgp.Y))}
	for _, cmd := range svgp.Commands {
		var s string
		switch c := cmd.(type) {
		case PathLine:
			s = fmt.Sprintf("L %s %s", p.x(c.X), p.y(c.Y))
		case PathArc:
			s = fmt.Sprintf("A %s %s 0 %d %d %s %s", p.m(c.RadiusX), p.m(c.RadiusY), c.LargeArc, c.Sweep, p.x(c.X), p.y(c.Y))
		default:
			return nil, fmt.Errorf("%+v", c)
		}
		cmds = append(cmds, s)
	}
	b := fmt.Sprintf(`<path d="%s" fill="%s" line="%d"%s/>`, strings.Join(cmds, " "), svgp.Fill, svgp.Line, psvg.FormatAttr(svgp.Attr))
	return []byte(b), nil
}

func (p *Processor) x(x int) string {
	return strconv.FormatFloat(float64(x)*p.Scale, 'f', -1, 64)
}

func (p *Processor) y(y int) string {
	return strconv.FormatFloat(-float64(y)*p.Scale, 'f', -1, 64)
}

func (p *Processor) m(f int) string {
	return strconv.FormatFloat(float64(f)*p.Scale, 'f', -1, 64)
}

// SVG parses Gerber input into SVG.
func SVG(r io.Reader) (*Processor, error) {
	processor := NewProcessor()
	parser := gerber.NewParser(processor)
	if err := parser.Parse(r); err != nil {
		return nil, err
	}

	return processor, nil
}
