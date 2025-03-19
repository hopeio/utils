/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package grpc_gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hopeio/utils/errors/errcode"
	"github.com/hopeio/utils/net/http/consts"
	"google.golang.org/grpc/codes"

	"github.com/hopeio/utils/net/http/grpc/gateway"
	"github.com/hopeio/utils/net/http/grpc/reconn"
	stringsi "github.com/hopeio/utils/strings"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"strings"
)

func RoutingErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write(stringsi.ToBytes(http.StatusText(httpStatus)))
}

func CustomHttpError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {

	s, ok := status.FromError(err)
	if ok && s.Code() == 14 && strings.HasSuffix(s.Message(), `refused it."`) {
		//提供一个思路，这里应该是哪条连接失败重连哪条，不能这么粗暴，map的key是个关键
		if len(reconn.ReConnectMap) > 0 {
			for _, f := range reconn.ReConnectMap {
				f()
			}
		}
	}

	const fallback = `{"code": 14, "message": "failed to marshal error message"}`

	w.Header().Del(consts.HeaderTrailer)
	contentType := marshaler.ContentType(nil)
	w.Header().Set(consts.HeaderContentType, contentType)
	se, ok := err.(*errcode.ErrRep)
	if !ok {
		se = &errcode.ErrRep{Code: errcode.ErrCode(codes.Unknown), Msg: err.Error()}
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	gateway.HandleForwardResponseServerMetadata(w, md.HeaderMD)

	buf, merr := marshaler.Marshal(se)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", se, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	var wantsTrailers bool

	if te := r.Header.Get(consts.HeaderTE); strings.Contains(strings.ToLower(te), "trailers") {
		wantsTrailers = true
		gateway.HandleForwardResponseTrailerHeader(w, md.TrailerMD)
		w.Header().Set(consts.HeaderTransferEncoding, "chunked")
	}

	/*	st := HTTPStatusFromCode(se.Code)
		w.WriteHeader(st)*/
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
	if wantsTrailers {
		gateway.HandleForwardResponseTrailer(w, md.TrailerMD)
	}
}
