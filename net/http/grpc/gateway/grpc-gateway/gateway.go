/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package grpc_gateway

import (
	"context"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/net/http/grpc/gateway"
	"net/http"
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
)

type GatewayHandler func(context.Context, *runtime.ServeMux)

func New() *runtime.ServeMux {

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &JSONPb{}),

		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			area, err := url.PathUnescape(req.Header.Get(consts.HeaderArea))
			if err != nil {
				area = ""
			}
			var token = httpi.GetToken(req)

			return map[string][]string{
				consts.HeaderArea:          {area},
				consts.HeaderDeviceInfo:    {req.Header.Get(consts.HeaderDeviceInfo)},
				consts.HeaderLocation:      {req.Header.Get(consts.HeaderLocation)},
				consts.HeaderAuthorization: {token},
			}
		}),
		runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case
				"Accept",
				"Accept-Charset",
				"Accept-Language",
				"Accept-Ranges",
				//"Token",
				"Cache-Control",
				"Content-Type",
				//"Cookie",
				"Date",
				"Expect",
				"From",
				"Host",
				"If-Match",
				"If-Modified-Since",
				"If-None-Match",
				"If-Schedule-Tag-Match",
				"If-Unmodified-Since",
				"Max-Forwards",
				"Origin",
				"Pragma",
				"Referer",
				"User-Agent",
				"Via",
				"Warning":
				return key, true
			}
			return "", false
		}),
		runtime.WithOutgoingHeaderMatcher(gateway.OutgoingHeaderMatcher))

	runtime.WithForwardResponseOption(gateway.Response)(gwmux)
	runtime.WithRoutingErrorHandler(RoutingErrorHandler)(gwmux)
	runtime.WithErrorHandler(CustomHttpError)(gwmux)

	return gwmux
}
