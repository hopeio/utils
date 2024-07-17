package slices

import (
	constraints2 "github.com/hopeio/utils/cmp"
	"golang.org/x/exp/constraints"
)

// BinarySearch 二分查找
func BinarySearch[V constraints.Ordered](arr []constraints2.CompareKey[V], x constraints2.CompareKey[V]) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid].CompareKey() == x.CompareKey() {
			return mid
		} else if x.CompareKey() > arr[mid].CompareKey() {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
