/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

import "golang.org/x/exp/constraints"

type Enum[T constraints.Unsigned | ~string] struct {
	Value T
}
