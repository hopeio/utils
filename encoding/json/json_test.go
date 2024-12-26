/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Foo struct {
	a int
	b string
	c json.RawMessage
}

func TestJson(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		foo := Foo{a: 1, b: "str"}
		data, _ := Marshal(foo)
		t.Log(string(data))
		var f Foo
		Unmarshal(data, &f)
		t.Log(f)
		reflect.DeepEqual(string(data), `{"a":1,"b":"str","c":null}`)
	})
}
