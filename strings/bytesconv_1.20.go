//go:build go1.20

package strings

import "unsafe"

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func ToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func FromBytes(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
