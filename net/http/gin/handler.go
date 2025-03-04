/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/gin/binding"
	"github.com/hopeio/utils/types"
	"net/http"
)

// only example

type GinService[REQ, RES any] func(*gin.Context, REQ) (RES, *httpi.ErrRep)

func HandlerWrap[REQ, RES any](service GinService[*REQ, *RES]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(REQ)
		err := binding.Bind(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
			return
		}
		res, reserr := service(ctx, req)
		if reserr != nil {
			ctx.JSON(http.StatusOK, reserr)
			return
		}
		if httpres, ok := any(res).(httpi.IHttpResponse); ok {
			httpi.ResponseWrite(ctx.Writer, httpres)
			return
		}
		ctx.JSON(http.StatusOK, httpi.NewSuccessResData(res))
	}
}

func HandlerWrapCompatibleGRPC[REQ, RES any](service types.GrpcServiceMethod[*REQ, *RES]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(REQ)
		err := binding.Bind(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
			return
		}
		res, err := service(httpi.WarpContext(ctx), req)
		if err != nil {
			ctx.JSON(http.StatusOK, httpi.ErrRepFrom(err))
			return
		}
		if httpres, ok := any(res).(httpi.IHttpResponse); ok {
			httpi.ResponseWrite(ctx.Writer, httpres)
			return
		}
		ctx.JSON(http.StatusOK, httpi.NewSuccessResData(res))
	}
}
