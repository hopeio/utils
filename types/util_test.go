/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

import "testing"

type Foo struct {
	A int
}

func TestNilValue(t *testing.T) {
	t.Log(Nil[Foo]())
}
