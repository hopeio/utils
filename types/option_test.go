package types

import "testing"

func TestOptionP(t *testing.T) {
	v := None[int]()
	t.Log(v.IsSome())
	t.Log(v.IsNone())
	data, err := v.MarshalJSON()
	t.Log(string(data), err)
	v.IfSome(func(value int) {
		t.Log(value)
	})
	v.IfNone(func() {
		t.Log("none")
	})
}
