/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package structtag

import (
	"log"
	"reflect"
	"testing"
)

type Bar1 struct {
	Field1 int
	Field2 string `mock:"example:'1',type:'\\w'"`
}

func TestTag(t *testing.T) {
	var bar Bar1
	typ := reflect.TypeOf(bar)
	log.Println(GetCustomTag(typ.Field(1).Tag.Get("mock"), "example"))
}
