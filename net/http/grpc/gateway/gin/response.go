/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/consts"

	"github.com/hopeio/utils/encoding/protobuf/jsonpb"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/grpc"
	"github.com/hopeio/utils/net/http/grpc/gateway"
	"google.golang.org/protobuf/proto"
)

func ForwardResponseMessage(ctx *gin.Context, md grpc.ServerMetadata, message proto.Message) {
	if res, ok := message.(httpi.IHttpResponse); ok {
		httpi.RespWrite(ctx.Writer, res)
		return
	}
	gateway.HandleForwardResponseServerMetadata(ctx.Writer, md.HeaderMD)
	gateway.HandleForwardResponseTrailerHeader(ctx.Writer, md.TrailerMD)

	contentType := jsonpb.JsonPb.ContentType(message)
	ctx.Header(consts.HeaderContentType, contentType)

	if !message.ProtoReflect().IsValid() {
		ctx.Writer.Write(httpi.ResponseOk)
		return
	}
	gateway.HandleForwardResponseTrailer(ctx.Writer, md.TrailerMD)
	err := gateway.Response(ctx, ctx.Writer, message)
	if err != nil {
		HttpError(ctx, err)
		return
	}
}
