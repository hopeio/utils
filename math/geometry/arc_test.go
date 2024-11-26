package geometry

import "testing"

func TestArc(t *testing.T) {
	arc := NewArc(0, 2, -3, -2, 3, -2)
	t.Log(arc.Bounds())
	arc = NewArc(0, 2, 3, -2, -3, -2)
	t.Log(arc.Bounds())
}
