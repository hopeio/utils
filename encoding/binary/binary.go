package binary

import (
	"golang.org/x/exp/constraints"
	"unsafe"
)

func ToInt64(b []byte) int64 {
	return int64(b[7]) | int64(b[6])<<8 | int64(b[5])<<16 | int64(b[4])<<24 |
		int64(b[3])<<32 | int64(b[2])<<40 | int64(b[1])<<48 | int64(b[0])<<56
}

func Int64To(i int64) []byte {
	return []byte{
		byte(i >> 56),
		byte(i >> 48),
		byte(i >> 40),
		byte(i >> 32),
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}

func BinaryTo[T constraints.Integer](b []byte) T {
	var v T
	byteNum := unsafe.Sizeof(v)
	for i := byteNum - 1; i > 0; i-- {
		v |= T(b[i]) << (8 * (byteNum - i - 1))

	}
	return v
}

func ToBinary[T constraints.Integer](v T) []byte {
	byteNum := unsafe.Sizeof(v)
	b := make([]byte, byteNum)
	for i := range byteNum {
		b[i] = byte(v >> (8 * (byteNum - i - 1)))

	}
	return b
}

func BTo[T constraints.Integer](b []byte) T {
	return *(*T)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b))))
}

func ToB[T constraints.Integer](v T) []byte {
	byteNum := unsafe.Sizeof(v)
	b := make([]byte, byteNum)
	*(*T)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b)))) = v
	return b
}

func ToInt(b []byte) int {
	return *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b))))
}

func IntTo(i int) []byte {
	b := make([]byte, 8)
	*(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b)))) = i
	return b
}

// 比标准库慢很多,10倍左右，string和bytes互转只是节省复制内存，unsafe操作有很多检测
// binary.LittleEndian.Uint64(b)
func ToUint(b []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b))))
}

// binary.LittleEndian.PutUint64(b)
func UintTo(i uint64) []byte {
	b := make([]byte, 8)
	*(*uint64)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b)))) = i
	return b
}
