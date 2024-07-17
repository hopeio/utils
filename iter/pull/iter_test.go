package pull

import "testing"

func TestIter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	SliceStreamOf(s).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		return i + 10
	}).ForEach(func(i int) {
		t.Log(i)
	})
}
