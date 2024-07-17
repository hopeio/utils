// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sync

import (
	"sync/atomic"
	"unsafe"
)

type DirectItem struct {
	Next unsafe.Pointer
	V    interface{}
}

func LoadItem(p *unsafe.Pointer) *DirectItem {
	return (*DirectItem)(atomic.LoadPointer(p))
}
func CasItem(p *unsafe.Pointer, old, new *DirectItem) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
