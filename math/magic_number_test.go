/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package math

import (
	"testing"
	"time"
)

func TestMagicNumber(t *testing.T) {
	key := SecondKey()
	t.Log(key)
	t.Log(ValidateSecondKey(key))
	t.Log(ValidateSecondKey(1 ^ magicNumber))
	t.Log(ValidateSecondKey(2 ^ magicNumber))
	t.Log(ValidateSecondKey(3 ^ magicNumber))
	t.Log(ValidateSecondKey(time.Now().Unix() - 1 ^ magicNumber))
}

func TestBitOperation(t *testing.T) {
	t.Log(-1 ^ (-1 << 8))
}
