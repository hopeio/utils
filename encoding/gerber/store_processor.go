package gerber

type StoreProcessor struct {
	Circles  []*Circle
	Rects    []*Rectangle
	Obrounds []*Obround
	Lines    []*Line
	Contours []*Contour
	Arcs     []*Arc
	ViewBox  *ViewBox
}

func (s *StoreProcessor) Circle(circle *Circle) {
	s.Circles = append(s.Circles, circle)
}

func (s *StoreProcessor) Rectangle(rectangle *Rectangle) {
	s.Rects = append(s.Rects, rectangle)
}

func (s *StoreProcessor) Obround(obround *Obround) {
	s.Obrounds = append(s.Obrounds, obround)
}

func (s *StoreProcessor) Contour(contour *Contour) {
	s.Contours = append(s.Contours, contour)
}

func (s *StoreProcessor) Line(line *Line) {
	s.Lines = append(s.Lines, line)
}

func (s *StoreProcessor) Arc(arc *Arc) {
	s.Arcs = append(s.Arcs, arc)
}

func (s *StoreProcessor) SetViewBox(box *ViewBox) {
	s.ViewBox = box
}

var _ Processor = (*StoreProcessor)(nil)
