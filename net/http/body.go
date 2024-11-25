/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

import (
	"io"
	"net/http"
)

func NewNoCloseRequest(req *http.Request, s []byte) {
	v := NewNoCloseBody(s)
	req.ContentLength = int64(v.Len())
	req.Body = v
	req.GetBody = func() (io.ReadCloser, error) {
		return v, nil
	}
}

func NewNoCloseStrRequest(req *http.Request, s string) {
	v := NewNoCloseStrBody(s)
	req.ContentLength = int64(v.Len())
	req.Body = v
	req.GetBody = func() (io.ReadCloser, error) {
		return v, nil
	}
}

type NoCloseBody struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *NoCloseBody) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

func (r *NoCloseBody) Size() int64 { return int64(len(r.s)) }

func (r *NoCloseBody) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *NoCloseBody) Close() error {
	r.i = 0
	return nil
}

// 适用于轮询
func NewNoCloseBody(s []byte) *NoCloseBody { return &NoCloseBody{s, 0, -1} }

type NoCloseStrBody struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *NoCloseStrBody) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

func (r *NoCloseStrBody) Size() int64 { return int64(len(r.s)) }

func (r *NoCloseStrBody) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *NoCloseStrBody) Close() error {
	r.i = 0
	return nil
}

// 适用于轮询
func NewNoCloseStrBody(s string) *NoCloseStrBody { return &NoCloseStrBody{s, 0, -1} }
