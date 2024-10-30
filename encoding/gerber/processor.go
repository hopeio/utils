package gerber

import "log"

type LogProcessor struct {
}

func (l LogProcessor) Circle(circle Circle) {
	log.Println("circle", circle)
}

func (l LogProcessor) Rectangle(rectangle Rectangle) {
	log.Println("rectangle", rectangle)
}

func (l LogProcessor) Obround(obround Obround) {
	log.Println("obround", obround)
}

func (l LogProcessor) Contour(contour Contour) error {
	log.Println("contour", contour)
	return nil
}

func (l LogProcessor) Line(line Line) {
	log.Println("line", line)
}

func (l LogProcessor) Arc(arc Arc) error {
	log.Println("arc", arc)
	return nil
}

func (l LogProcessor) SetViewBox(minX, maxX, minY, maxY int) {
	log.Println("SetViewBox", minX, maxX, minY, maxY)
}

var _ Processor = (*LogProcessor)(nil)
