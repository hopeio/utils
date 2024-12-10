package geom

import "testing"

func TestLine(t *testing.T) {
	line := LineSegment{Start: Pt(5, 0), End: Pt(5, 1)}
	t.Log(line.ToSlopeInterceptFormLine())
}
