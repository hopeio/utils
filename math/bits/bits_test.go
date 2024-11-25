/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package bits

import (
	"math"
	"testing"
)

func TestBitNumber(t *testing.T) {
	t.Log(math.IsNaN(BaseNaN.Float()))
	t.Log(math.IsNaN((BaseNaN + 1).Float()))
	t.Log(math.IsNaN((BaseNaN + 2).Float()))
	t.Log((BaseNaN + 2).RangeInt(1, 11))
	t.Log((BaseNaN + 2).RangeUint(1, 11))
	t.Log(1 << 1)
	t.Log(uint64(math.MaxUint64))
	t.Log(uint64((1 << 64) - 1))
}
