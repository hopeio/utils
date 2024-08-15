// Package svg parses Gerber to SVG.
package svg

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/encoding/gerber"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func (p *Processor) UnmarshalJSON(b []byte) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 || bytes.Equal(residue, []byte{'\n'}) {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\" %#v", residue, residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Data":
			residue, err = p.decodeData(residue)
		case "MinX":
			p.MinX, residue, err = decodeInt(residue)
		case "MaxX":
			p.MaxX, residue, err = decodeInt(residue)
		case "MinY":
			p.MinY, residue, err = decodeInt(residue)
		case "MaxY":
			p.MaxY, residue, err = decodeInt(residue)
		case "PolarityDark":
			p.PolarityDark, residue, err = decodeString(residue)
		case "PolarityClear":
			p.PolarityClear, residue, err = decodeString(residue)
		case "Scale":
			p.Scale, residue, err = decodeFloat(residue)
		case "Width":
			p.Width, residue, err = decodeString(residue)
		case "Height":
			p.Height, residue, err = decodeString(residue)
		case "PanZoom":
			p.PanZoom, residue, err = decodeBool(residue)
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Processor) UnmarshalJSON_1(b []byte) error {
	pmap := make(map[string]interface{})
	if err := json.Unmarshal(b, &pmap); err != nil {
		return err
	}

	if err := mapstructure.Decode(pmap, p); err != nil {
		return err
	}

	data := make([]interface{}, 0, len(p.Data))
	for i, d := range p.Data {
		m, ok := d.(map[string]interface{})
		if !ok {
			return fmt.Errorf("%d %+v", i, d)
		}
		eType, ok := m["Type"].(string)
		if !ok {
			return fmt.Errorf("%d %+v", i, d)
		}
		switch ElementType(eType) {
		case ElementTypeCircle:
			e := Circle{}
			if err := mapstructure.Decode(m, &e); err != nil {
				return fmt.Errorf("%d %+v %w", i, d, err)
			}
			data = append(data, e)
		case ElementTypeRectangle:
			e := Rectangle{}
			if err := mapstructure.Decode(m, &e); err != nil {
				return fmt.Errorf("%d %+v %w", i, d, err)
			}
			data = append(data, e)
		case ElementTypePath:
			e := Path{}
			if err := mapstructure.Decode(m, &e); err != nil {
				return fmt.Errorf("%d %+v %w", i, d, err)
			}
			cmds := make([]interface{}, 0, len(e.Commands))
			for j, cj := range e.Commands {
				c, ok := cj.(map[string]interface{})
				if !ok {
					return fmt.Errorf("%d %d %+v %+v", i, j, cj, d)
				}
				cType, ok := c["Type"].(string)
				if !ok {
					return fmt.Errorf("%d %d %+v %+v", i, j, cj, d)
				}
				switch ElementType(cType) {
				case ElementTypeLine:
					cmd := PathLine{}
					if err := mapstructure.Decode(c, &cmd); err != nil {
						return fmt.Errorf("%d %d %+v %+v", i, j, c, d)
					}
					cmds = append(cmds, cmd)
				case ElementTypeArc:
					cmd := PathArc{}
					if err := mapstructure.Decode(c, &cmd); err != nil {
						return fmt.Errorf("%d %d %+v %+v", i, j, c, d)
					}
					cmds = append(cmds, cmd)
				default:
					return fmt.Errorf("%d %d %+v %+v", i, j, c, d)
				}
			}
			e.Commands = cmds
			data = append(data, e)
		case ElementTypeLine:
			e := Line{}
			if err := mapstructure.Decode(m, &e); err != nil {
				return fmt.Errorf("%d %+v %w", i, d, err)
			}
			data = append(data, e)
		case ElementTypeArc:
			e := Arc{}
			if err := mapstructure.Decode(m, &e); err != nil {
				return fmt.Errorf("%d %+v %w", i, d, err)
			}
			data = append(data, e)
		default:
			return fmt.Errorf("%d %+v", i, m)
		}
	}
	p.Data = data

	return nil
}

func (p *Processor) decodeData(b []byte) ([]byte, error) {
	// Empty array.
	if bytes.HasPrefix(b, []byte("null")) {
		b = b[4:]
		if b[0] == ',' {
			b = b[1:]
		}
		return b, nil
	}

	// Opening '['.
	if len(b) < 1 {
		return b, fmt.Errorf("%d", len(b))
	}
	b = b[1:]

	for {
		segment, err := findSegment(b)
		if err != nil {
			return b, err
		}
		if err := p.decodeSegment(segment); err != nil {
			return b, err
		}
		if b[len(segment)] == ']' {
			b = b[len(segment)+1:]
			if b[0] == ',' {
				b = b[1:]
			}
			break
		}
		if len(segment)+1 >= len(b) {
			return b, fmt.Errorf("%d %d \"%s\"", len(segment)+1, len(b), segment)
		}
		b = b[len(segment)+1:]
	}

	return b, nil
}

func findSegment(bs []byte) ([]byte, error) {
	level := 0
	end := -1
Loop:
	for i, b := range bs {
		switch b {
		case '}':
			level--
			if level == 0 {
				end = i
				break Loop
			}
		case '{':
			level++
		}
	}
	if end == -1 {
		return nil, fmt.Errorf("not closed")
	}
	return bs[:end+1], nil
}

func (p *Processor) decodeSegment(b []byte) error {
	elmType, err := findElementType(b)
	if err != nil {
		return err
	}
	switch elmType {
	case ElementTypeCircle:
		var c Circle
		if err := decodeCircle(b, &c); err != nil {
			return err
		}
		p.Data = append(p.Data, c)
	case ElementTypeRectangle:
		var r Rectangle
		if err := decodeRectangle(b, &r); err != nil {
			return err
		}
		p.Data = append(p.Data, r)
	case ElementTypePath:
		var ph Path
		if err := decodePath(b, &ph); err != nil {
			return err
		}
		p.Data = append(p.Data, ph)
	case ElementTypeLine:
		var l Line
		if err := decodeLine(b, &l); err != nil {
			return err
		}
		p.Data = append(p.Data, l)
	case ElementTypeArc:
		var a Arc
		if err := decodeArc(b, &a); err != nil {
			return err
		}
		p.Data = append(p.Data, a)
	default:
		return fmt.Errorf("\"%s\"", elmType)
	}
	return nil
}

func findElementType(bs []byte) (ElementType, error) {
	level := -1
	for i, b := range bs {
		switch b {
		case '}':
			level--
		case '{':
			level++
		case ':':
			if level != 0 {
				break
			}

			if i-5 < 0 {
				break
			}
			if !bytes.Equal(bs[i-5:i-1], []byte("Type")) {
				break
			}
			endIdx := bytes.IndexByte(bs[i+2:], '"')
			if endIdx == -1 {
				return "", fmt.Errorf("not closed")
			}
			return ElementType(bs[i+2 : i+2+endIdx]), nil
		}
	}
	return "", fmt.Errorf("not found")
}

func decodeCircle(b []byte, elm *Circle) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "Line":
			elm.Line, residue, err = decodeInt(residue)
		case "X":
			elm.X, residue, err = decodeInt(residue)
		case "Y":
			elm.Y, residue, err = decodeInt(residue)
		case "Diameter":
			elm.Diameter, residue, err = decodeInt(residue)
		case "Fill":
			elm.Fill, residue, err = decodeString(residue)
		case "Attr":
			// Attr is expected to be always null.
			residue = residue[5:]
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeRectangle(b []byte, elm *Rectangle) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "Line":
			elm.Line, residue, err = decodeInt(residue)
		case "Aperture":
			elm.Aperture, residue, err = decodeString(residue)
		case "X":
			elm.X, residue, err = decodeInt(residue)
		case "Y":
			elm.Y, residue, err = decodeInt(residue)
		case "Width":
			elm.Width, residue, err = decodeInt(residue)
		case "Height":
			elm.Height, residue, err = decodeInt(residue)
		case "RadiusX":
			elm.RadiusX, residue, err = decodeInt(residue)
		case "RadiusY":
			elm.RadiusY, residue, err = decodeInt(residue)
		case "Fill":
			elm.Fill, residue, err = decodeString(residue)
		case "Attr":
			// Attr is expected to be always null.
			residue = residue[5:]
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodePath(b []byte, elm *Path) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "Line":
			elm.Line, residue, err = decodeInt(residue)
		case "X":
			elm.X, residue, err = decodeInt(residue)
		case "Y":
			elm.Y, residue, err = decodeInt(residue)
		case "Commands":
			residue, err = decodePathCommands(elm, residue)
		case "Fill":
			elm.Fill, residue, err = decodeString(residue)
		case "Attr":
			// Attr is expected to be always null.
			residue = residue[5:]
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeLine(b []byte, elm *Line) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "Line":
			elm.Line.Line, residue, err = decodeInt(residue)
		case "StartX":
			elm.StartX, residue, err = decodeInt(residue)
		case "StartY":
			elm.StartY, residue, err = decodeInt(residue)
		case "EndX":
			elm.EndX, residue, err = decodeInt(residue)
		case "EndY":
			elm.EndY, residue, err = decodeInt(residue)
		case "Width":
			elm.Width, residue, err = decodeInt(residue)
		case "Cap":
			var lineCap string
			lineCap, residue, err = decodeString(residue)
			elm.Cap = gerber.LineCap(lineCap)
		case "Stroke":
			elm.Stroke, residue, err = decodeString(residue)
		case "Attr":
			// Attr is expected to be always null.
			residue = residue[5:]
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodeArc(b []byte, elm *Arc) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "Line":
			elm.Line, residue, err = decodeInt(residue)
		case "StartX":
			elm.StartX, residue, err = decodeInt(residue)
		case "StartY":
			elm.StartY, residue, err = decodeInt(residue)
		case "RadiusX":
			elm.RadiusX, residue, err = decodeInt(residue)
		case "RadiusY":
			elm.RadiusY, residue, err = decodeInt(residue)
		case "LargeArc":
			elm.LargeArc, residue, err = decodeInt(residue)
		case "Sweep":
			elm.Sweep, residue, err = decodeInt(residue)
		case "EndX":
			elm.EndX, residue, err = decodeInt(residue)
		case "EndY":
			elm.EndY, residue, err = decodeInt(residue)
		case "Width":
			elm.Width, residue, err = decodeInt(residue)
		case "CenterX":
			elm.CenterX, residue, err = decodeInt(residue)
		case "CenterY":
			elm.CenterY, residue, err = decodeInt(residue)
		case "Stroke":
			elm.Stroke, residue, err = decodeString(residue)
		case "Attr":
			// Attr is expected to be always null.
			residue = residue[5:]
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodePathLine(b []byte, elm *PathLine) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "X":
			elm.X, residue, err = decodeInt(residue)
		case "Y":
			elm.Y, residue, err = decodeInt(residue)
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodePathArc(b []byte, elm *PathArc) error {
	// Skip '{'.
	if len(b) < 1 {
		return fmt.Errorf("%d", len(b))
	}
	residue := b[1:]

	for {
		if len(residue) == 0 {
			break
		}
		closingK := bytes.Index(residue[1:], []byte(`":`))
		if closingK == -1 {
			return fmt.Errorf("\"%s\"", residue)
		}
		key := string(residue[1 : 1+closingK])
		residue = residue[1+closingK+2:]
		var err error
		switch key {
		case "Type":
			var elmType string
			elmType, residue, err = decodeString(residue)
			elm.Type = ElementType(elmType)
		case "RadiusX":
			elm.RadiusX, residue, err = decodeInt(residue)
		case "RadiusY":
			elm.RadiusY, residue, err = decodeInt(residue)
		case "LargeArc":
			elm.LargeArc, residue, err = decodeInt(residue)
		case "Sweep":
			elm.Sweep, residue, err = decodeInt(residue)
		case "X":
			elm.X, residue, err = decodeInt(residue)
		case "Y":
			elm.Y, residue, err = decodeInt(residue)
		default:
			err = fmt.Errorf("\"%s\"", key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func decodePathCommands(ph *Path, b []byte) ([]byte, error) {
	// Opening '['.
	if len(b) < 1 {
		return b, fmt.Errorf("%d", len(b))
	}
	b = b[1:]

	for {
		segment, err := findSegment(b)
		if err != nil {
			return b, err
		}
		if err := decodePathSegment(ph, segment); err != nil {
			return b, err
		}
		if b[len(segment)] == ']' {
			b = b[len(segment)+1:]
			if b[0] == ',' {
				b = b[1:]
			}
			break
		}
		if len(segment)+1 >= len(b) {
			return b, fmt.Errorf("%d %d \"%s\"", len(segment)+1, len(b), segment)
		}
		b = b[len(segment)+1:]
	}

	return b, nil
}

func decodePathSegment(ph *Path, b []byte) error {
	elmType, err := findElementType(b)
	if err != nil {
		return err
	}
	switch elmType {
	case ElementTypeLine:
		var l PathLine
		if err := decodePathLine(b, &l); err != nil {
			return err
		}
		ph.Commands = append(ph.Commands, l)
	case ElementTypeArc:
		var a PathArc
		if err := decodePathArc(b, &a); err != nil {
			return err
		}
		ph.Commands = append(ph.Commands, a)
	default:
		return fmt.Errorf("\"%s\"", elmType)
	}
	return nil
}

func decodeInt(b []byte) (int, []byte, error) {
	idx := bytes.IndexByte(b, ',')
	if idx == -1 {
		idx = bytes.IndexByte(b, '}')
		if idx == -1 {
			return -1, b, fmt.Errorf("no comma")
		}
	}
	i, err := strconv.Atoi(string(b[:idx]))
	if err != nil {
		return -1, b, err
	}
	return i, b[idx+1:], nil
}

func decodeFloat(b []byte) (float64, []byte, error) {
	idx := bytes.IndexByte(b, ',')
	if idx == -1 {
		idx = bytes.IndexByte(b, '}')
		if idx == -1 {
			return -1, b, fmt.Errorf("no comma")
		}
	}
	f, err := strconv.ParseFloat(string(b[:idx]), 64)
	if err != nil {
		return -1, b, err
	}
	return f, b[idx+1:], nil
}

func decodeString(b []byte) (string, []byte, error) {
	idx := bytes.Index(b, []byte(`",`))
	if idx == -1 {
		idx = bytes.IndexByte(b, '}')
		if idx == -1 {
			return "", b, fmt.Errorf("no comma")
		}
	}
	// Opening '"'
	if len(b) < 1 {
		return "", b, fmt.Errorf("%d", len(b))
	}
	s := string(b[1:idx])
	return s, b[idx+2:], nil
}

func decodeBool(b []byte) (bool, []byte, error) {
	idx := bytes.IndexByte(b, ',')
	if idx == -1 {
		idx = bytes.IndexByte(b, '}')
		if idx == -1 {
			return false, b, fmt.Errorf("no comma")
		}
	}
	bol, err := strconv.ParseBool(string(b[:idx]))
	if err != nil {
		return false, b, err
	}
	return bol, b[idx+1:], nil
}
