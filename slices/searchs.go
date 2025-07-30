/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package slices

import (
	"github.com/hopeio/gox/cmp"
)

// BinarySearch 二分查找
func BinarySearch[T any, S ~[]cmp.Comparable[T]](arr S, x T) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid].Compare(x) == 0 {
			return mid
		} else if arr[mid].Compare(x) < 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
