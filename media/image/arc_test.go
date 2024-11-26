package image

import "testing"

func TestArc(t *testing.T) {
	arc := NewArc(5, 5, 2, 9, 8, 9)
	t.Log(arc.Bounds())
	arc = NewArc(5, 5, 8, 9, 2, 9)
	t.Log(arc.Bounds())
}
