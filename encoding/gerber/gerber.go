package gerber

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// fork: github.com/fumin/gerber
// A Processor performs Gerber graphic operations.

type Processor interface {
	// Circle draws a circle.
	Circle(Circle)

	// Rectangle draws a rectangle.
	Rectangle(Rectangle)

	// Obround draws an obround.
	Obround(Obround)

	// Contour draws a contour.
	Contour(Contour) error

	// Line draws a line.
	Line(Line)

	// Arc draws an arc.
	Arc(Arc) error

	// SetViewbox sets the viewbox of the Gerber image.
	// It is called by the Parser when parsing has completed.
	SetViewbox(minX, maxX, minY, maxY int)
}

func parseApertureID(word string) (string, error) {
	if len(word) < 3 {
		return "", fmt.Errorf("%d", len(word))
	}
	if word[0] != 'D' {
		return "", fmt.Errorf("%v", word[0])
	}

	var digits int = len(word) - 1
	for i, c := range word[1:] {
		if c >= '0' && c <= '9' {
			continue
		}
		digits = i
		break
	}

	return word[:1+digits], nil
}

type circlePrimitive struct {
	Exposure bool
	Diameter float64
	CenterX  float64
	CenterY  float64
}

type rectPrimitive struct {
	Exposure    bool
	Width       float64
	Height      float64
	CenterX     float64
	CenterY     float64
	Polarity    bool
	Rotation    float64
	SetVariable []func(p *rectPrimitive, f float64)
}

type vectorLinePrimitive struct {
	Exposure bool
	Width    float64
	StartX   float64
	StartY   float64
	EndX     float64
	EndY     float64
	Rotation float64
}

type outlinePrimitive struct {
	Exposure    bool
	NumVertices int
	Points      [][2]float64
	Rotation    float64
}

type lowerLeftLinePrimitive struct {
	Exposure bool
	Width    float64
	Height   float64
	X        float64
	Y        float64
	Rotation float64
}

type template struct {
	Line       int
	Name       string
	Primitives []interface{}
}

type LinePrimitiveNotClosedError struct {
	Line     int
	First    [2]float64
	Last     [2]float64
	FirstStr [2]string
	LastStr  [2]string
}

func (err LinePrimitiveNotClosedError) Error() string {
	return fmt.Sprintf("line primitive not closed %d %#v %#v", err.Line, err.First, err.Last)
}

func parsePrimitive(lineIdx int, word string) (interface{}, error) {
	splitted := strings.Split(word, primitiveDelimiter)
	if len(splitted) == 0 {
		return nil, fmt.Errorf("no splitted")
	}
	curLine := lineIdx
	if strings.Contains(splitted[0], "\n") {
		curLine++
	}
	code, err := strconv.Atoi(strings.TrimSpace(splitted[0]))
	if err != nil {
		return nil, err
	}
	switch code {
	case primitiveCodeCircle:
		if len(splitted) != 5 {
			return nil, fmt.Errorf("%+v", splitted)
		}
		circle := circlePrimitive{}
		exposure, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			return nil, err
		}
		if exposure == 1 {
			circle.Exposure = true
		}
		circle.Diameter, err = strconv.ParseFloat(strings.TrimSpace(splitted[2]), 64)
		if err != nil {
			return nil, err
		}
		circle.CenterX, err = strconv.ParseFloat(strings.TrimSpace(splitted[3]), 64)
		if err != nil {
			return nil, err
		}
		circle.CenterY, err = strconv.ParseFloat(strings.TrimSpace(splitted[4]), 64)
		if err != nil {
			return nil, err
		}
		return circle, nil
	case primitiveCodeVectorLine:
		if len(splitted) != 8 {
			return nil, fmt.Errorf("%+v", splitted)
		}
		line := vectorLinePrimitive{}
		exposure, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			return nil, err
		}
		if exposure == 1 {
			line.Exposure = true
		}
		line.Width, err = strconv.ParseFloat(strings.TrimSpace(splitted[2]), 64)
		if err != nil {
			return nil, err
		}
		line.StartX, err = strconv.ParseFloat(strings.TrimSpace(splitted[3]), 64)
		if err != nil {
			return nil, err
		}
		line.StartY, err = strconv.ParseFloat(strings.TrimSpace(splitted[4]), 64)
		if err != nil {
			return nil, err
		}
		line.EndX, err = strconv.ParseFloat(strings.TrimSpace(splitted[5]), 64)
		if err != nil {
			return nil, err
		}
		line.EndY, err = strconv.ParseFloat(strings.TrimSpace(splitted[6]), 64)
		if err != nil {
			return nil, err
		}
		line.Rotation, err = strconv.ParseFloat(strings.TrimSpace(splitted[7]), 64)
		if err != nil {
			return nil, err
		}
		return line, nil
	case primitiveCodeOutline:
		line := outlinePrimitive{}
		if len(splitted) < 3 {
			return nil, fmt.Errorf("%+v", splitted)
		}

		if strings.Contains(splitted[1], "\n") {
			curLine++
		}
		exposure, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			return nil, err
		}
		if exposure == 1 {
			line.Exposure = true
		}

		if strings.Contains(splitted[2], "\n") {
			curLine++
		}
		line.NumVertices, err = strconv.Atoi(strings.TrimSpace(splitted[2]))
		if err != nil {
			return nil, err
		}
		if len(splitted) != 6+2*line.NumVertices {
			return nil, fmt.Errorf("%d", len(splitted))
		}

		points := make([][2]string, 0)
		for i := 0; i < line.NumVertices+1; i++ {
			if strings.Contains(splitted[2*i+3], "\n") {
				curLine++
			}
			x, err := strconv.ParseFloat(strings.TrimSpace(splitted[2*i+3]), 64)
			if err != nil {
				return nil, fmt.Errorf("%d %w", i, err)
			}

			if strings.Contains(splitted[2*i+4], "\n") {
				curLine++
			}
			y, err := strconv.ParseFloat(strings.TrimSpace(splitted[2*i+4]), 64)
			if err != nil {
				return nil, fmt.Errorf("%d %w", i, err)
			}
			line.Points = append(line.Points, [2]float64{x, y})
			points = append(points, [2]string{strings.TrimSpace(splitted[2*i+3]), strings.TrimSpace(splitted[2*i+4])})
		}

		// The last point must be the same as the starting point.
		if line.Points[0] != line.Points[len(line.Points)-1] {
			return nil, LinePrimitiveNotClosedError{Line: curLine, First: line.Points[0], Last: line.Points[len(line.Points)-1], FirstStr: points[0], LastStr: points[len(points)-1]}
		}
		line.Points = line.Points[:len(line.Points)-1]

		line.Rotation, err = strconv.ParseFloat(strings.TrimSpace(splitted[len(splitted)-1]), 64)
		if err != nil {
			return nil, fmt.Errorf("%+v", splitted)
		}
		return line, nil
	case primitiveCodeLowerLeftLine:
		if len(splitted) != 7 {
			return nil, fmt.Errorf("%+v", splitted)
		}
		line := lowerLeftLinePrimitive{}
		exposure, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			return nil, err
		}
		if exposure == 1 {
			line.Exposure = true
		}
		line.Width, err = strconv.ParseFloat(strings.TrimSpace(splitted[2]), 64)
		if err != nil {
			return nil, err
		}
		line.Height, err = strconv.ParseFloat(strings.TrimSpace(splitted[3]), 64)
		if err != nil {
			return nil, err
		}
		line.X, err = strconv.ParseFloat(strings.TrimSpace(splitted[4]), 64)
		if err != nil {
			return nil, err
		}
		line.Y, err = strconv.ParseFloat(strings.TrimSpace(splitted[5]), 64)
		if err != nil {
			return nil, err
		}
		line.Rotation, err = strconv.ParseFloat(strings.TrimSpace(splitted[6]), 64)
		if err != nil {
			return nil, err
		}
		return line, nil
	case primitiveCodeRect:
		if len(splitted) != 7 {
			return nil, fmt.Errorf("%+v", splitted)
		}
		rect := rectPrimitive{}
		exposure, err := strconv.Atoi(strings.TrimSpace(splitted[1]))
		if err != nil {
			return nil, err
		}
		if exposure == 1 {
			rect.Exposure = true
		}
		if isVar, err := injectFloat(&rect.Width, splitted[2]); err != nil {
			return nil, err
		} else if isVar {
			rect.SetVariable = append(rect.SetVariable, func(p *rectPrimitive, f float64) {
				p.Width = f
			})
		}

		if isVar, err := injectFloat(&rect.Height, splitted[3]); err != nil {
			return nil, err
		} else if isVar {
			rect.SetVariable = append(rect.SetVariable, func(p *rectPrimitive, f float64) {
				p.Height = f
			})
		}

		if isVar, err := injectFloat(&rect.CenterX, splitted[4]); err != nil {
			return nil, err
		} else if isVar {
			rect.SetVariable = append(rect.SetVariable, func(p *rectPrimitive, f float64) {
				p.CenterX = f
			})
		}

		if isVar, err := injectFloat(&rect.CenterY, splitted[5]); err != nil {
			return nil, err
		} else if isVar {
			rect.SetVariable = append(rect.SetVariable, func(p *rectPrimitive, f float64) {
				p.CenterY = f
			})
		}

		if isVar, err := injectFloat(&rect.Rotation, splitted[6]); err != nil {
			return nil, err
		} else if isVar {
			rect.SetVariable = append(rect.SetVariable, func(p *rectPrimitive, f float64) {
				p.Rotation = f
			})
		}

		return rect, nil
	default:
		return nil, fmt.Errorf("%+v", splitted)
	}
}

type aperture struct {
	Line     int
	ID       string
	Template template
	Params   []float64
}

type regionParser struct {
	cp         *commandProcessor
	contour    Contour
	gotCommand bool
}

func newRegionParser(cp *commandProcessor, lineIdx int) *regionParser {
	p := &regionParser{}
	p.cp = cp
	p.contour = Contour{Line: lineIdx, X: cp.x, Y: cp.y, Polarity: cp.polarity}
	return p
}

func (p *regionParser) process(lineIdx int, word string) error {
	switch {
	case strings.HasPrefix(word, commandG01):
		p.cp.interpolation = InterpolationLinear
		return p.process(lineIdx, word[len(commandG01):])
	case strings.HasPrefix(word, commandG02):
		p.cp.interpolation = InterpolationClockwise
		return p.process(lineIdx, word[len(commandG02):])
	case strings.HasPrefix(word, commandG03):
		p.cp.interpolation = InterpolationCCW
		return p.process(lineIdx, word[len(commandG03):])
	case strings.HasSuffix(word, commandD01):
		return p.processD01(lineIdx, word[:len(word)-len(commandD01)])
	case strings.HasSuffix(word, commandD02):
		return p.processD02(lineIdx, word[:len(word)-len(commandD02)])
	case word == commandG37:
		return p.cp.pc.Contour(p.contour)
	case strings.HasPrefix(word, "X"):
		return p.processModalD01(lineIdx, word)
	default:
		return fmt.Errorf("unknown command")
	}
}

func (p *regionParser) processModalD01(lineIdx int, word string) error {
	if !p.cp.modalD01 {
		return fmt.Errorf("not in modal D01 mode")
	}
	return p.processD01(lineIdx, word)
}

func (p *regionParser) processD01(lineIdx int, word string) error {
	coords, err := parseCoord(word)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("\"%s\"", word))
	}
	x, y := p.cp.findXY(coords)

	s := Segment{Interpolation: p.cp.interpolation, X: x, Y: y}
	switch s.Interpolation {
	case InterpolationClockwise:
		fallthrough
	case InterpolationCCW:
		i, j, err := p.cp.findIJ(coords)
		if err != nil {
			return fmt.Errorf("%+v %w", coords, err)
		}
		s.XCenter, s.YCenter = p.cp.x+i, p.cp.y+j
	}
	p.contour.Segments = append(p.contour.Segments, s)

	p.cp.setXY(x, y)
	p.cp.modalD01 = true
	return nil
}

func (p *regionParser) processD02(lineIdx int, word string) error {
	if p.gotCommand {
		if err := p.cp.pc.Contour(p.contour); err != nil {
			return err
		}
	}
	p.gotCommand = true

	coords, err := parseCoord(word)
	if err != nil {
		return fmt.Errorf("\"%s\" %w", word, err)
	}
	x, y := p.cp.findXY(coords)
	p.cp.setXY(x, y)

	p.contour = Contour{Line: lineIdx, X: x, Y: y, Polarity: p.cp.polarity}
	p.cp.modalD01 = false
	return nil
}

// An Interpolation is a Gerber interpolation method.
type Interpolation int

// A LineCap is the shape at the endpoints of a line.
type LineCap string

const (
	// Linear interpolation.
	InterpolationLinear Interpolation = iota
	// Clockwise arc interpolation
	InterpolationClockwise
	// Counter clockwise arc interpolation.
	InterpolationCCW

	// LineCapButt strokes do not extend beyond a line's two endpoints.
	LineCapButt LineCap = "butt"
	// LineCapRound strokes will be extended by a half circle with a diameter equal to the stroke width.
	LineCapRound LineCap = "round"

	primitiveCodeCircle        = 1
	primitiveCodeVectorLine    = 20
	primitiveCodeRect          = 21
	primitiveCodeOutline       = 4
	primitiveCodeLowerLeftLine = 22

	primitiveDelimiter = ","

	commandFS  = "FS"
	commandMO  = "MO"
	commandG04 = "G04"
	commandIP  = "IP"
	commandLN  = "LN"
	commandLP  = "LP"
	commandG74 = "G74"
	commandG75 = "G75"
	commandAD  = "AD"
	commandAM  = "AM"
	commandG36 = "G36"
	commandG37 = "G37"
	commandG54 = "G54"
	commandG01 = "G01"
	commandG02 = "G02"
	commandG03 = "G03"
	commandD01 = "D01"
	commandD02 = "D02"
	commandD03 = "D03"
	commandSR  = "SR"
	commandM02 = "M02"

	templateNameCircle    = "C"
	templateNameRectangle = "R"
	templateNameObround   = "O"
)

type Unit string

const (
	UnitMillimeter Unit = "mm"
	UnitInch       Unit = "inch"
)

type commandProcessor struct {
	pc        Processor
	templates map[string]template
	apertures map[string]aperture
	rp        *regionParser

	decimal       float64
	unit          Unit
	interpolation Interpolation
	x             int
	y             int
	ap            aperture
	polarity      bool

	minX int
	maxX int
	minY int
	maxY int

	modalD01 bool
}

func newCommandProcessor(pc Processor) *commandProcessor {
	p := &commandProcessor{}
	p.pc = pc
	p.templates = make(map[string]template)
	p.apertures = make(map[string]aperture)

	// Gerber RS-274X Format User guide.
	// Part Number 414 100 014 Rev D March, 2001.
	// Quote:
	// When a new layer is generated, interpolation will be reset to linear (G01).
	p.interpolation = InterpolationLinear

	// Default polarity is dark.
	p.polarity = true

	p.minX = math.MaxInt
	p.maxX = -math.MaxInt
	p.minY = math.MaxInt
	p.maxY = -math.MaxInt

	return p
}

func (p *commandProcessor) processWord(lineIdx int, word string) error {
	switch {
	case p.rp != nil:
		if err := p.rp.process(lineIdx, word); err != nil {
			return fmt.Errorf("%+v %w", p.rp, err)
		}
		if word == commandG37 {
			p.rp = nil
		}
		return nil
	}

	switch {
	case word == "":
		return nil
	case strings.HasPrefix(word, commandFS):
		return p.processFS(lineIdx, word)
	case strings.HasPrefix(word, commandMO):
		return p.processMO(lineIdx, word)
	case strings.HasPrefix(word, commandAD):
		return p.parseAD(lineIdx, word[len(commandAD):])
	case strings.HasPrefix(word, commandG04):
		return nil
	case strings.HasPrefix(word, commandLP):
		return p.processLP(lineIdx, word)
	case strings.HasPrefix(word, commandIP):
		return nil
	case strings.HasPrefix(word, commandLN):
		return nil
	case strings.HasPrefix(word, commandG74):
		return nil
	case strings.HasPrefix(word, commandG75):
		return nil
	case strings.HasPrefix(word, commandG54):
		return p.processDnn(lineIdx, word[len(commandG54):])
	case word == commandG36:
		p.rp = newRegionParser(p, lineIdx)
		return nil
	case strings.HasPrefix(word, commandG01):
		p.interpolation = InterpolationLinear
		return p.processWord(lineIdx, word[len(commandG01):])
	case strings.HasPrefix(word, commandG02):
		p.interpolation = InterpolationClockwise
		return p.processWord(lineIdx, word[len(commandG02):])
	case strings.HasPrefix(word, commandG03):
		p.interpolation = InterpolationCCW
		return p.processWord(lineIdx, word[len(commandG03):])
	case strings.HasSuffix(word, commandD01):
		return p.processD01(lineIdx, word[:len(word)-len(commandD01)])
	case strings.HasSuffix(word, commandD02):
		return p.processD02(lineIdx, word[:len(word)-len(commandD02)])
	case strings.HasSuffix(word, commandD03):
		return p.processD03(lineIdx, word[:len(word)-len(commandD03)])
	case strings.HasPrefix(word, "D"):
		return p.processDnn(lineIdx, word)
	case strings.HasPrefix(word, commandSR):
		return p.processSR(lineIdx, word)
	case strings.HasPrefix(word, "X"):
		return p.processModalD01(lineIdx, word)
	case strings.HasPrefix(word, commandAM):
		words := strings.Split(word, wordTerminator)
		return p.processExtended(lineIdx, words)
	case word == commandM02:
		p.pc.SetViewbox(p.minX, p.maxX, p.minY, p.maxY)
		return nil
	default:
		return fmt.Errorf("unknown command")
	}
}

func (p *commandProcessor) setXY(x, y int) {
	p.x = x
	if x < p.minX {
		p.minX = x
	}
	if x > p.maxX {
		p.maxX = x
	}

	p.y = y
	if y < p.minY {
		p.minY = y
	}
	if y > p.maxY {
		p.maxY = y
	}
}

func (p *commandProcessor) processModalD01(lineIdx int, word string) error {
	if !p.modalD01 {
		return fmt.Errorf("not in modal D01 mode")
	}
	return p.processD01(lineIdx, word)
}

func (p *commandProcessor) processD01(lineIdx int, word string) error {
	coords, err := parseCoord(word)
	if err != nil {
		return fmt.Errorf("\"%s\" %w", word, err)
	}
	x, y := p.findXY(coords)

	var diameter int
	switch p.ap.Template.Name {
	case templateNameCircle:
		diameter = p.u(p.ap.Params[0])
	case templateNameRectangle:
		if p.ap.Params[0] != p.ap.Params[1] {
			return fmt.Errorf("%+v", p.ap)
		}
		diameter = p.u(p.ap.Params[0])
	default:
		return fmt.Errorf("%+v", p.ap)
	}

	switch p.interpolation {
	case InterpolationLinear:
		p.pc.Line(Line{lineIdx, p.x, p.y, x, y, diameter, LineCapRound, 0})
	case InterpolationClockwise:
		fallthrough
	case InterpolationCCW:
		i, j, err := p.findIJ(coords)
		if err != nil {
			return fmt.Errorf("%+v", coords)
		}
		xc, yc := p.x+i, p.y+j
		if err := p.pc.Arc(Arc{lineIdx, p.x, p.y, x, y, xc, yc, diameter, p.interpolation}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%d", p.interpolation)
	}

	p.setXY(x, y)
	p.modalD01 = true
	return nil
}

func (p *commandProcessor) processD02(lineIdx int, word string) error {
	coords, err := parseCoord(word)
	if err != nil {
		return fmt.Errorf("\"%s\", %w", word, err)
	}
	x, y := p.findXY(coords)
	p.setXY(x, y)
	p.modalD01 = false
	return nil
}

func (p *commandProcessor) processD03(lineIdx int, word string) error {
	coords, err := parseCoord(word)
	if err != nil {
		return fmt.Errorf("\"%s\" %w", word, err)
	}
	x, y := p.findXY(coords)
	p.setXY(x, y)

	if err := p.flash(lineIdx); err != nil {
		return err
	}
	p.modalD01 = false
	return nil
}

func (p *commandProcessor) flash(lineIdx int) error {
	params := p.ap.Params
	switch p.ap.Template.Name {
	case templateNameCircle:
		p.pc.Circle(Circle{lineIdx, p.x, p.y, p.u(params[0]), p.polarity})
	case templateNameRectangle:
		p.pc.Rectangle(Rectangle{Line: lineIdx, X: p.x, Y: p.y, Width: p.u(params[0]), Height: p.u(params[1]),
			Polarity: p.polarity})
	case templateNameObround:
		p.pc.Obround(Obround{lineIdx, p.x, p.y, p.u(params[0]), p.u(params[1]), p.polarity, 0})
	default:
		return p.flashUserDefinedTmpl(lineIdx)
	}
	return nil
}

func (p *commandProcessor) flashUserDefinedTmpl(lineIdx int) error {
	if !p.polarity {
		return fmt.Errorf("%v", p.polarity)
	}
	for i, primitive := range p.ap.Template.Primitives {
		switch pm := primitive.(type) {
		case circlePrimitive:
			if !pm.Exposure {
				return fmt.Errorf("%d %+v", i, pm)
			}
			p.pc.Circle(Circle{lineIdx, p.x + p.u(pm.CenterX), p.y + p.u(pm.CenterY), p.u(pm.Diameter), p.polarity})
		case vectorLinePrimitive:
			if !pm.Exposure {
				return fmt.Errorf("%d %+v", i, pm)
			}
			if pm.Rotation != 0 {
				return fmt.Errorf("%d %+v", i, pm)
			}
			p.pc.Line(Line{lineIdx, p.x + p.u(pm.StartX), p.y + p.u(pm.StartY), p.x + p.u(pm.EndX), p.y + p.u(pm.EndY),
				p.u(pm.Width), LineCapButt, 0})
		case outlinePrimitive:
			if !pm.Exposure {
				return fmt.Errorf("%d %+v", i, pm)
			}
			if pm.Rotation != 0 {
				return fmt.Errorf("%d %+v", i, pm)
			}
			contour, err := p.contourFromOutline(lineIdx, pm)
			if err != nil {
				return fmt.Errorf("%d %+v %w", i, pm, err)
			}
			if err := p.pc.Contour(contour); err != nil {
				return fmt.Errorf("%d %+v %w", i, pm, err)
			}
		case lowerLeftLinePrimitive:
			if !pm.Exposure {
				return fmt.Errorf("%d %+v", i, pm)
			}
			if pm.Rotation != 0 {
				return fmt.Errorf("%d %+v", i, pm)
			}
			p.pc.Rectangle(Rectangle{Line: lineIdx, X: p.x + p.u(pm.X+pm.Width/2), Y: p.y + p.u(pm.Y+pm.Height/2),
				Width:  p.u(pm.Width),
				Height: p.u(pm.Height), Polarity: p.polarity, Rotation: pm.Rotation})
		case rectPrimitive:
			if !pm.Exposure {
				return fmt.Errorf("%d %+v", i, pm)
			}
			if pm.SetVariable != nil {
				for i, f := range pm.SetVariable {
					f(&pm, p.ap.Params[i])
				}
			}
			p.pc.Rectangle(Rectangle{Line: lineIdx, X: p.x + p.u(pm.CenterX), Y: p.y + p.u(pm.CenterY), Width: p.u(pm.Width),
				Height: p.u(pm.Height), Polarity: pm.Polarity, Rotation: pm.Rotation})
		default:
			return fmt.Errorf("%d %+v", i, p)
		}
	}
	return nil
}

func (p *commandProcessor) contourFromOutline(lineIdx int, outline outlinePrimitive) (Contour, error) {
	contour := Contour{Line: lineIdx, Polarity: p.polarity}
	if len(outline.Points) < 3 {
		return Contour{}, fmt.Errorf("%+v", outline.Points)
	}
	contour.X = p.x + p.u(outline.Points[0][0])
	contour.Y = p.y + p.u(outline.Points[0][1])

	for _, pt := range outline.Points[1:] {
		s := Segment{Interpolation: InterpolationLinear, X: p.x + p.u(pt[0]), Y: p.y + p.u(pt[1])}
		contour.Segments = append(contour.Segments, s)
	}
	return contour, nil
}

type coord struct {
	S byte
	I int
}

func parseCoord(word string) ([]coord, error) {
	if word == "" {
		return nil, nil
	}

	coords := make([]coord, 0)
	cur := coord{S: word[0]}
	var digits []byte
	var err error
	for i, c := range []byte(word[1:]) {
		switch {
		case c == '+' || c == '-' || (c >= '0' && c <= '9'):
			digits = append(digits, c)
		default:
			cur.I, err = strconv.Atoi(string(digits))
			if err != nil {
				return nil, fmt.Errorf("%d \"%s\" %w", i, digits, err)
			}
			coords = append(coords, cur)
			cur = coord{}
			cur.S = c
			digits = digits[:0]
		}
	}

	cur.I, err = strconv.Atoi(string(digits))
	if err != nil {
		return nil, fmt.Errorf("invalid digits \"%s\" %w", digits, err)
	}
	coords = append(coords, cur)

	return coords, nil
}

func (p *commandProcessor) findXY(coords []coord) (int, int) {
	x := p.x
	for _, c := range coords {
		if c.S == 'X' {
			x = c.I
			break
		}
	}

	y := p.y
	for _, c := range coords {
		if c.S == 'Y' {
			y = c.I
			break
		}
	}

	return x, y
}

func (p *commandProcessor) findIJ(coords []coord) (int, int, error) {
	var i int
	var got bool
	for _, c := range coords {
		if c.S == 'I' {
			i = c.I
			got = true
			break
		}
	}
	if !got {
		return -math.MaxInt, -math.MaxInt, fmt.Errorf("no i")
	}

	got = false
	var j int
	for _, c := range coords {
		if c.S == 'J' {
			j = c.I
			got = true
			break
		}
	}
	if !got {
		return -math.MaxInt, -math.MaxInt, fmt.Errorf("no j")
	}

	return i, j, nil
}

func (p *commandProcessor) parseAD(lineIdx int, word string) error {
	aperture := aperture{Line: lineIdx}
	var err error
	aperture.ID, err = parseApertureID(word)
	if err != nil {
		return err
	}
	afterAID := word[len(aperture.ID):]

	var tmplName string
	commaIdx := strings.Index(afterAID, ",")
	if commaIdx == -1 {
		tmplName = afterAID
	} else {
		tmplName = afterAID[:commaIdx]
	}

	switch tmplName {
	case templateNameCircle:
		aperture.Template = template{Name: templateNameCircle}
	case templateNameRectangle:
		aperture.Template = template{Name: templateNameRectangle}
	case templateNameObround:
		aperture.Template = template{Name: templateNameObround}
	default:
		var ok bool
		aperture.Template, ok = p.templates[tmplName]
		if !ok {
			tmpls := make([]string, 0, len(p.templates))
			for k := range p.templates {
				tmpls = append(tmpls, k)
			}
			return fmt.Errorf("%s %+v", tmplName, tmpls)
		}
	}

	if commaIdx != -1 {
		if commaIdx+1 > len(afterAID) {
			return fmt.Errorf("%d %s", commaIdx, afterAID)
		}
		params := strings.Split(afterAID[commaIdx+1:], "X")
		for i, pStr := range params {
			p, err := strconv.ParseFloat(pStr, 64)
			if err != nil {
				return fmt.Errorf("%d %w", i, err)
			}
			aperture.Params = append(aperture.Params, p)
		}
	}
	var expectedParams int = len(aperture.Params)
	switch tmplName {
	case templateNameCircle:
		expectedParams = 1
	case templateNameRectangle:
		expectedParams = 2
	case templateNameObround:
		expectedParams = 2
	}
	if expectedParams != len(aperture.Params) {
		return fmt.Errorf("%d %+v", expectedParams, aperture.Params)
	}

	if prev, ok := p.apertures[aperture.ID]; ok {
		return fmt.Errorf("%+v", prev)
	}
	p.apertures[aperture.ID] = aperture

	return nil
}

func (p *commandProcessor) processFS(lineIdx int, word string) error {
	if len(word) < 7 {
		return fmt.Errorf("%d", len(word))
	}
	decimal, err := strconv.Atoi(word[6:7])
	if err != nil {
		return err
	}
	p.decimal = math.Pow(10, float64(decimal))
	return nil
}

func (p *commandProcessor) u(f float64) int {
	return int(math.Round(f * p.decimal))
}

func (p *commandProcessor) processMO(lineIdx int, word string) error {
	if len(word) != 4 {
		return fmt.Errorf("%d", len(word))
	}
	unit := word[2:]
	switch unit {
	case "MM":
		p.unit = UnitMillimeter
		return nil
	case "IN":
		p.unit = UnitInch
		return nil
	default:
		return fmt.Errorf("%s", unit)
	}
}

func (p *commandProcessor) processSR(lineIdx int, word string) error {
	if word != "SRX1Y1I0J0" {
		return fmt.Errorf("unsupported SR")
	}
	return nil
}

func (p *commandProcessor) processLP(lineIdx int, word string) error {
	if len(word) != 3 {
		return fmt.Errorf("%d", len(word))
	}
	switch word[2] {
	case 'D':
		p.polarity = true
	case 'C':
		p.polarity = false
	default:
		return fmt.Errorf("%s", word)
	}
	return nil
}

func (p *commandProcessor) processDnn(lineIdx int, word string) error {
	var ok bool
	p.ap, ok = p.apertures[word]
	if !ok {
		aps := make([]string, 0, len(p.apertures))
		for k := range p.apertures {
			aps = append(aps, k)
		}
		return fmt.Errorf("%+v", aps)
	}
	p.modalD01 = false
	return nil
}

func (p *commandProcessor) processExtended(lineIdx int, words []string) error {
	if len(words) == 0 {
		return fmt.Errorf("no words")
	}

	switch {
	case strings.HasPrefix(words[0], commandAM):
		tmpl := template{Line: lineIdx, Name: words[0][len(commandAM):]}
		for _, w := range words[1:] {
			primitive, err := parsePrimitive(lineIdx, w)
			if err != nil {
				return err
			}
			tmpl.Primitives = append(tmpl.Primitives, primitive)
		}
		p.templates[tmpl.Name] = tmpl
	default:
		return fmt.Errorf("unknown command")
	}

	return nil
}

// A Parser is a Gerber format parser.
// For each graphical operation parsed from an input stream,
// Parser calls the corresponding method of its Processor.
type Parser struct {
	cmdStart     int
	cmdLines     []string
	cmdProcessor *commandProcessor
}

const (
	variableKey              = "$"
	extendedCommandDelimiter = "%"
	wordTerminator           = "*"
	wordCommand              = -1
)

// NewParser creates a Parser.
func NewParser(pc Processor) *Parser {
	p := &Parser{}
	p.cmdStart = wordCommand
	p.cmdProcessor = newCommandProcessor(pc)
	return p
}

func (p *Parser) parse(lineIdx int, line string) error {
	if p.cmdStart != wordCommand {
		if !strings.HasSuffix(line, extendedCommandDelimiter) {
			p.cmdLines = append(p.cmdLines, line)
			return nil
		}
		remainder := len(line) - len(extendedCommandDelimiter)
		if remainder > 0 {
			p.cmdLines = append(p.cmdLines, line[:remainder])
		}

		// Split by *
		joined := strings.Join(p.cmdLines, "\n")
		if len(joined) == 0 {
			return fmt.Errorf("%d", p.cmdStart)
		}
		if !strings.HasSuffix(joined, wordTerminator) {
			return fmt.Errorf("%s", joined)
		}
		joined = joined[:len(joined)-len(wordTerminator)]
		words := strings.Split(joined, wordTerminator)

		cmdStart := p.cmdStart
		p.cmdStart = wordCommand
		return p.cmdProcessor.processExtended(cmdStart, words)
	}

	if strings.HasPrefix(line, extendedCommandDelimiter) {
		if strings.HasSuffix(line, extendedCommandDelimiter) {
			word := line[len(extendedCommandDelimiter) : len(line)-len(extendedCommandDelimiter)]
			if !strings.HasSuffix(word, wordTerminator) {
				return fmt.Errorf("%s", word)
			}
			return p.cmdProcessor.processWord(lineIdx, word[:len(word)-len(wordTerminator)])
		}

		p.cmdStart = lineIdx
		p.cmdLines = p.cmdLines[:0]
		p.cmdLines = append(p.cmdLines, line[len(extendedCommandDelimiter):])
		return nil
	}

	if !strings.HasSuffix(line, wordTerminator) {
		return fmt.Errorf("%s", line)
	}
	word := line[:len(line)-len(wordTerminator)]
	if err := p.cmdProcessor.processWord(lineIdx, word); err != nil {
		return fmt.Errorf("unable to parse word \"%s\" %w", word, err)
	}

	return nil
}

// Parse parses the Gerber format stream.
func (parser *Parser) Parse(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	var lineIdx int

	for scanner.Scan() {
		lineIdx++
		line := scanner.Text()
		if line == "" {
			continue
		}
		if err := parser.parse(lineIdx, line); err != nil {
			return fmt.Errorf("at line %d: \"%s\" %w", lineIdx, line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func injectFloat(v *float64, s string) (isVariable bool, err error) {
	if strings.HasPrefix(s, variableKey) {
		return true, nil
	}
	*v, err = strconv.ParseFloat(strings.TrimSpace(s), 64)
	return
}

type Loayer string

const (
	GTL Loayer = "GTL" //顶层走线
	GBL Loayer = "GBL" //底层走线
	GTO Loayer = "GTO" //顶层丝印
	GBO Loayer = "GBO" //底层丝印
	GTS Loayer = "GTS" // 顶层阻焊
	GBS Loayer = "GBS" //底层阻焊
	GPT Loayer = "GPT" //顶层主焊盘
	GPB Loayer = "GPB" //底层主焊盘
	G1  Loayer = "G1"  //内部走线层1
	G2  Loayer = "G2"  //内部走线层2
	G3  Loayer = "G3"  //内部走线层3
	G4  Loayer = "G4"  //内部走线层4
	GP1 Loayer = "GP1" //内平面1(负片)
	GP2 Loayer = "GP2" //内平面2(负片)
	GM1 Loayer = "GM1" //机械层1
	GM2 Loayer = "GM2" //机械层2
	GM3 Loayer = "GM3" //机械层3
	GM4 Loayer = "GM4" //机械层4
	GKO Loayer = "GKO" //禁止布线层(可做板子外形)
)
