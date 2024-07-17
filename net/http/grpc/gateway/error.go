package gateway

import (
	"fmt"
	"github.com/hopeio/protobuf/errcode"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/textproto"
)

func HttpStatusFromCode(code errcode.ErrCode) int {
	switch code {
	case errcode.Success:
		return http.StatusOK
	case errcode.Canceled:
		return http.StatusRequestTimeout
	case errcode.Unknown:
		return http.StatusInternalServerError
	case errcode.InvalidArgument:
		return http.StatusBadRequest
	case errcode.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case errcode.NotFound:
		return http.StatusNotFound
	case errcode.AlreadyExists:
		return http.StatusConflict
	case errcode.PermissionDenied:
		return http.StatusForbidden
	case errcode.Unauthenticated:
		return http.StatusUnauthorized
	case errcode.ResourceExhausted:
		return http.StatusTooManyRequests
	case errcode.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case errcode.Aborted:
		return http.StatusConflict
	case errcode.OutOfRange:
		return http.StatusBadRequest
	case errcode.Unimplemented:
		return http.StatusNotImplemented
	case errcode.Internal:
		return http.StatusInternalServerError
	case errcode.Unavailable:
		return http.StatusServiceUnavailable
	case errcode.DataLoss:
		return http.StatusInternalServerError
	}

	grpclog.Infof("Unknown gRPC error code: %v", code)
	return http.StatusInternalServerError
}

func OutgoingHeaderMatcher(key string) (string, bool) {
	switch key {
	case
		httpi.HeaderSetCookie:
		return key, true
	}
	return "", false
}

var headerMatcher = []string{httpi.HeaderSetCookie}

func HandleForwardResponseServerMetadata(w http.ResponseWriter, md metadata.MD) {
	for _, k := range headerMatcher {
		if vs, ok := md[k]; ok {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
	}
}

func HandleForwardResponseTrailerHeader(w http.ResponseWriter, md metadata.MD) {
	for k := range md {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", grpc.MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func HandleForwardResponseTrailer(w http.ResponseWriter, md metadata.MD) {
	for k, vs := range md {
		tKey := fmt.Sprintf("%s%s", grpc.MetadataTrailerPrefix, k)
		for _, v := range vs {
			w.Header().Add(tKey, v)
		}
	}
}
