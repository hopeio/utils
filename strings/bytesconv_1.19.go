//go:build !go1.20

package strings

import "unsafe"

// 这个方式好一点,新建一个结构体承载
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 这个直接用slice类型强转,会丢失cap信息
func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
