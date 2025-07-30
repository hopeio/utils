/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package interfaces

import "github.com/hopeio/gox/types/constraints"

type Key[T constraints.Key] interface {
	Key() T
}

type Collector[S any, T any, R any] interface {
	Builder() S
	Append(builder S, element T)
	Finish(builder S) R
}
