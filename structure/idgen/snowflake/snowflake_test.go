/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package snowflake

import "testing"

func TestSnowFlake(t *testing.T) {
	node := NewNode(1, 10, 12)
	for i := 0; i < 100; i++ {
		id := node.Generate()
		t.Log(id)
		t.Log(id.Base32())
		t.Log(id.Base36())
		t.Log(id.Base58())
		t.Log(id.Base64())
	}
}
