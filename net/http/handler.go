/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

import (
	"encoding/json"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/types"
	"net/http"
)

type Handlers []http.Handler

func (hs Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range hs {
		handler.ServeHTTP(w, r)
	}
}

type HandlerFuncs []http.HandlerFunc

func (hs HandlerFuncs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range hs {
		handler(w, r)
	}
}

func (hs HandlerFuncs) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range hs {
			handler(w, r)
		}
	}
}

func (hs *HandlerFuncs) Add(handler http.HandlerFunc) {
	*hs = append(*hs, handler)
}

type ReqResp struct {
	*http.Request
	http.ResponseWriter
}
type Service[REQ, RES any] func(ctx ReqResp, req REQ) (RES, error)

func HandlerWrap[REQ, RES any](service Service[*REQ, *RES]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(REQ)
		err := binding.Bind(r, req)
		if err != nil {
			return
		}
		res, err := service(ReqResp{r, w}, req)
		if err != nil {
			return
		}
		if httpres, ok := any(res).(IHttpResponse); ok {
			ResponseWrite(w, httpres)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set(HeaderContentType, ContentTypeJsonUtf8)
		json.NewEncoder(w).Encode(res)
	})
}
func HandlerWrapCompatibleGRPC[REQ, RES any](method types.GrpcServiceMethod[*REQ, *RES]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(REQ)
		err := binding.Bind(r, req)
		if err != nil {
			return
		}
		res, err := method(r.Context(), req)
		if err != nil {
			return
		}
		if httpres, ok := any(res).(IHttpResponse); ok {
			ResponseWrite(w, httpres)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set(HeaderContentType, ContentTypeJsonUtf8)
		json.NewEncoder(w).Encode(res)
	})
}
