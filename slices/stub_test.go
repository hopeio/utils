// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slices

import (
	"github.com/hopeio/utils/reflect"
	"testing"
	"unsafe"
)

func TestGrowSlic(t *testing.T) {
	et := reflect.UnpackEface(byte(0)).Type
	n := make([]byte, 1, 1024)
	println("n len:", len(n), "cap:", cap(n))
	v := GrowSlice(et, *(*reflect.GoSlice)(unsafe.Pointer(&n)), 1025)
	println("v len:", v.Len, "cap:", v.Cap)
	// according to go growslice rule, the next cap must be at most 1.5x of old.Cap
	if v.Cap > 1536 {
		t.Fatal(v.Cap)
	}
}
