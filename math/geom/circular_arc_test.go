package geom

import "testing"

func TestArc(t *testing.T) {
	arc := CircularArcFromPoints(Pt(0, 2), Pt(-3, -2), Pt(3, -2))
	t.Log(arc.Bounds())
	t.Log(arc.minimumBoundingRectangle())
	arc = CircularArcFromPoints(Pt(0, 2), Pt(3, -2), Pt(-3, -2))
	t.Log(arc.Bounds())
	t.Log(arc.minimumBoundingRectangle())
}
