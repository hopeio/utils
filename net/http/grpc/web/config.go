/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package web

import (
	"google.golang.org/grpc"
	"net/http"
)

// Deprecated
type GrpcWebServerConfig struct {
	WithOriginFunc                     func(origin string) bool
	WithEndpointsFunc                  func() []string
	WithCorsForRegisteredEndpointsOnly bool
	WithAllowedRequestHeaders          []string
	WithWebsockets                     bool
	WithWebsocketOriginFunc            func(req *http.Request) bool
	WithWebsocketsMessageReadLimit     bool
	WithAllowNonRootResource           bool
}

func DefaultGrpcWebServer(grpcServer *grpc.Server) *WrappedGrpcServer {
	return WrapServer(grpcServer, WithAllowedRequestHeaders([]string{"*"}), WithWebsockets(true), WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}), WithOriginFunc(func(origin string) bool {
		return true
	}))
}
