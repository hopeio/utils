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
