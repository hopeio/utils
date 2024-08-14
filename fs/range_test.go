package fs

import (
	"testing"
)

func TestRange(t *testing.T) {
	it, err := All("D:\\data")
	if err.HasErrors() {
		t.Error(err)
	}

	for ent := range it {
		t.Log(ent.Name())
	}
	if err.HasErrors() {
		t.Error(err)
	}
}
