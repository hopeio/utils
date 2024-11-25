/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

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
