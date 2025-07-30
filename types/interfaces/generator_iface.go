/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package interfaces

import (
	"github.com/hopeio/gox/types/constraints"
	"time"
)

type IdGenerator[T constraints.ID] interface {
	Id() T
}

type DurationGenerator interface {
	Duration() time.Duration
}
