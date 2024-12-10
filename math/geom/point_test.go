package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMirror(t *testing.T) {
	p := Pt(1, 2)
	line := GeneralFormLine{1, 1, 0}
	assert.Equal(t, Pt(-2, -1), p.Mirror(&line))
	line = GeneralFormLine{0, 1, 0}
	assert.Equal(t, Pt(1, -2), p.Mirror(&line))
	line = GeneralFormLine{1, 0, 0}
	assert.Equal(t, Pt(-1, 2), p.Mirror(&line))
	line = GeneralFormLine{1, 0, 1}
	assert.Equal(t, Pt(-3, 2), p.Mirror(&line))
	line = GeneralFormLine{0, -1, 1}
	assert.Equal(t, Pt(1, 0), p.Mirror(&line))
}
