/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package mtos

import (
	"testing"
)

func TestParseJson(t *testing.T) {
	data := `{"a":[
  "Hello",
  123,
  true,
  null,
  {"key": "value"},
  [1, 2, 3]
]}`
	t.Log(ParseJson([]byte(data)))
}
