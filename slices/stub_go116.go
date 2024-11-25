/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

//go:build go1.16 && !go1.20
// +build go1.16,!go1.20

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
	_ "unsafe"
)

//go:linkname GrowSlice runtime.growslice
//goland:noinspection GoUnusedParameter
func GrowSlice(et *reflect.Type, old reflect.Slice, cap int) reflect.Slice
