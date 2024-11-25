/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Rangeable interface {
	constraints.Ordered | time.Time | ~*time.Time | ~string
}

type Ordered interface {
	constraints.Ordered | time.Time
}
