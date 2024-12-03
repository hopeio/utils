package svg

import (
	"github.com/hopeio/utils/encoding/gerber"
	"os"
	"testing"
)

func TestSvg(t *testing.T) {
	path := `D:\Gerber_TopLayer.GTL`
	p := NewProcessor()

	f, _ := os.Open(path)
	defer f.Close()
	err := gerber.NewParser(p).Parse(f)
	if err != nil {
		t.Error(err)
	}
	svgPath := `D:\Gerber_TopLayer.svg`
	svg, _ := os.Create(svgPath)
	defer svg.Close()
	p.Write(svg)
}
