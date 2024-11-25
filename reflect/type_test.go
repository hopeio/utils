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

func TestDerefInterfaceType(t *testing.T) {
	var a any
	a = 1
	v := reflect.TypeOf(&a)
	t.Log(v.Kind())
	v1 := v.Elem()
	t.Log(v1.Kind())
	v2 := v1.Elem()
	t.Log(v2.Kind())
}
