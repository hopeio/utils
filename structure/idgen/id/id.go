/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package id

import (
	"sync/atomic"
)

var currentID uint64 = 0

// 单机顺序id
func NewOrderID() uint64 {
	return atomic.AddUint64(&currentID, 1)
}
