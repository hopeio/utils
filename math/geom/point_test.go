package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMirror(t *testing.T) {
	p := Pt(1, 2)
	line := StraightLine{1, 1, 0}
	assert.Equal(t, Pt(-2, -1), p.Mirror(&line))
	line = StraightLine{0, 1, 0}
	assert.Equal(t, Pt(1, -2), p.Mirror(&line))
	line = StraightLine{1, 0, 0}
	assert.Equal(t, Pt(-1, 2), p.Mirror(&line))
	line = StraightLine{1, 0, 1}
	assert.Equal(t, Pt(-3, 2), p.Mirror(&line))
	line = StraightLine{0, -1, 1}
	assert.Equal(t, Pt(1, 0), p.Mirror(&line))
}
