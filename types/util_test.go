package types

import "testing"

type Foo struct {
	A int
}

func TestNilValue(t *testing.T) {
	t.Log(Nil[Foo]())
}
