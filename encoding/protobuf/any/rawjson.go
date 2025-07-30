/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package any

import (
	"github.com/hopeio/gox/encoding/json"
)

type RawJson []byte

func NewAny(v interface{}) (RawJson, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *RawJson) MarshalJSON() ([]byte, error) {
	if len(*a) == 0 {
		return []byte("null"), nil
	}
	return *a, nil
}

func (a *RawJson) Size() int {
	return len(*a)
}

func (a *RawJson) MarshalTo(b []byte) (int, error) {
	return copy(b, *a), nil
}

func (a *RawJson) Unmarshal(b []byte) error {
	*a = b
	return nil
}

func (a *RawJson) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(*a)
	copy(dAtA[i:], *a)
	return len(*a), nil
}

type randyRawJson interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func NewPopulatedRawJson(r randyRawJson, easy bool) *RawJson {
	if !easy && r.Intn(10) != 0 {
	}
	any := RawJson("{}")
	return &any
}
