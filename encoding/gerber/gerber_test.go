package gerber

import (
	"os"
	"testing"
)

func TestGerber(t *testing.T) {
	path := `XXX`
	p := LogProcessor{}

	f, _ := os.Open(path)
	defer f.Close()
	err := NewParser(p).Parse(f)
	if err != nil {
		t.Error(err)
	}
}
