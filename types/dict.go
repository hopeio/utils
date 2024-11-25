/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

type Dict[K comparable, V any] struct {
	Key   K
	Value V
}
