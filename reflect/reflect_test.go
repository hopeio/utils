/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package reflect

import (
	"reflect"
	"testing"
)

type Foo struct {
	A int
	B string
	C *int
	D *map[int]any
	E *[]int
}
type Bar struct {
	Foo Foo
	C   string
}

func TestGetExpectTypeValue(t *testing.T) {
	a := Bar{Foo: Foo{A: 1}}
	b := Foo{}
	v := CopyFieldValueByType(&a, &b)
	if v {
		t.Log(b)
	}
}

func TestInitStruct(t *testing.T) {
	var f *Foo
	v := reflect.ValueOf(&f)
	InitStruct(v)
	t.Log(*f)
	t.Log(*f.C)
	t.Log(*f.D)
	t.Log(*f.E)
	t.Log(v.Elem().Interface())
}
