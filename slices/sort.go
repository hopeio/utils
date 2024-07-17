package slices

import (
	cmpi "github.com/hopeio/utils/cmp"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[T constraints.Ordered](s []T) {
	slices.Sort(s)
}

func SortByKey[K constraints.Ordered, T cmpi.CompareKey[K]](s []T) {
	slices.SortFunc(s, cmpi.CompareByKey[K, T])
}

func quickSort[T constraints.Ordered](array []T, left, right int) {
	if left < right {
		x, i := array[right], left-1
		var temp T
		for j := left; j <= right; j++ {
			if array[j] <= x {
				i++
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
		quickSort(array, left, i-1)
		quickSort(array, i+1, right)
	}
}

func bubbleSort[T constraints.Ordered](array []T) {
	low, high := 0, len(array)-1

	var tmp T
	var j int
	for low < high {
		for j = low; j < high; j++ {
			//正向冒泡,找到最大者
			if array[j] > array[j+1] {
				tmp = array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
			if array[low] > array[j] {
				tmp = array[low]
				array[low] = array[j]
				array[j] = tmp
			}
		}
		high--
		low++
	}
}
