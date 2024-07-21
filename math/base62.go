package math

import (
	"github.com/hopeio/utils/types/constraints"
	"math"
	"strconv"
	"unsafe"
)

const digits string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func FormatUint(num, base int64) string {
	if base < 37 {
		return strconv.FormatInt(num, int(base))
	}
	var bytes []byte
	for num > 0 {
		bytes = append(bytes, digits[num%base])
		num = num / base
	}
	reverse(bytes)
	return string(bytes)
}

func ParseUint(str string, base float64) (uint64, error) {
	if base < 37 {
		return strconv.ParseUint(str, int(base), strconv.IntSize)
	}
	var num uint64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := findIndex(str[i])
		num += uint64(math.Pow(base, float64(n-i-1)) * float64(pos))
	}
	return num, nil
}

func findIndex(b byte) int {
	if b < 'A' {
		return int(b - '0')
	} else if b > 'Z' {
		return 10 + int(b-'a')
	}
	return 36 + int(b-'A')
}
func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func FormatFloat(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func ToBytes[T constraints.Number](t T) []byte {
	size := unsafe.Sizeof(t)
	return unsafe.Slice((*byte)(unsafe.Pointer(&t)), size)
}
