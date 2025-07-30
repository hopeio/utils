/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package strings

import (
	reflect2 "github.com/hopeio/gox/reflect"
	"unsafe"
)

//go:nosplit
func Mem2Str(v []byte) (s string) {
	(*reflect2.String)(unsafe.Pointer(&s)).Len = (*reflect2.Slice)(unsafe.Pointer(&v)).Len
	(*reflect2.String)(unsafe.Pointer(&s)).Ptr = (*reflect2.Slice)(unsafe.Pointer(&v)).Ptr
	return
}

//go:nosplit
func Str2Mem(s string) (v []byte) {
	(*reflect2.Slice)(unsafe.Pointer(&v)).Cap = (*reflect2.String)(unsafe.Pointer(&s)).Len
	(*reflect2.Slice)(unsafe.Pointer(&v)).Len = (*reflect2.String)(unsafe.Pointer(&s)).Len
	(*reflect2.Slice)(unsafe.Pointer(&v)).Ptr = (*reflect2.String)(unsafe.Pointer(&s)).Ptr
	return
}

func BytesFrom(p unsafe.Pointer, n int, c int) (r []byte) {
	(*reflect2.Slice)(unsafe.Pointer(&r)).Ptr = p
	(*reflect2.Slice)(unsafe.Pointer(&r)).Len = n
	(*reflect2.Slice)(unsafe.Pointer(&r)).Cap = c
	return
}

//go:nocheckptr
func IndexChar(src string, index int) unsafe.Pointer {
	return unsafe.Pointer(uintptr((*reflect2.String)(unsafe.Pointer(&src)).Ptr) + uintptr(index))
}

//go:nocheckptr
func IndexByte(ptr []byte, index int) unsafe.Pointer {
	return unsafe.Pointer(uintptr((*reflect2.Slice)(unsafe.Pointer(&ptr)).Ptr) + uintptr(index))
}

//go:nosplit
func StrPtr(s string) unsafe.Pointer {
	return (*reflect2.String)(unsafe.Pointer(&s)).Ptr
}

//go:nosplit
func StrFrom(p unsafe.Pointer, n int64) (s string) {
	(*reflect2.String)(unsafe.Pointer(&s)).Ptr = p
	(*reflect2.String)(unsafe.Pointer(&s)).Len = int(n)
	return
}
