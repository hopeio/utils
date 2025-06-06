package handlerwrap

import (
	"context"
	"encoding/json"
	"github.com/hopeio/utils/errors/errcode"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/types"
	"net/http"
)

type Service[REQ, RES any] func(ctx ReqResp, req REQ) (RES, *httpi.ErrRep)

type warpKey struct{}

var warpContextKey = warpKey{}

func WarpContext(v any) context.Context {
	return context.WithValue(context.Background(), warpContextKey, v)
}

func UnWarpContext(ctx context.Context) any {
	return ctx.Value(warpContextKey)
}

type ReqResp struct {
	*http.Request
	http.ResponseWriter
}

func HandlerWrap[REQ, RES any](service Service[*REQ, *RES]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(REQ)
		err := binding.Bind(r, req)
		if err != nil {
			httpi.RespErrCodeMsg(w, errcode.InvalidArgument, err.Error())
			return
		}
		res, errRep := service(ReqResp{r, w}, req)
		if err != nil {
			errRep.Response(w)
			return
		}
		anyres := any(res)
		if httpres, ok := anyres.(httpi.ICommonResponseTo); ok {
			httpres.CommonResponse(httpi.CommonResponseWriter{ResponseWriter: w})
			return
		}
		if httpres, ok := anyres.(httpi.IHttpResponseTo); ok {
			httpres.Response(w)
			return
		}
		json.NewEncoder(w).Encode(res)
	})
}
func HandlerWrapCompatibleGRPC[REQ, RES any](method types.GrpcService[*REQ, *RES]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(REQ)
		err := binding.Bind(r, req)
		if err != nil {
			httpi.RespSuccessData(w, errcode.InvalidArgument.Wrap(err))
			return
		}
		res, err := method(WarpContext(ReqResp{r, w}), req)
		if err != nil {
			httpi.ErrRepFrom(err).Response(w)
			return
		}
		anyres := any(res)
		if httpres, ok := anyres.(httpi.ICommonResponseTo); ok {
			httpres.CommonResponse(httpi.CommonResponseWriter{w})
			return
		}
		if httpres, ok := anyres.(httpi.IHttpResponseTo); ok {
			httpres.Response(w)
			return
		}
		w.Header().Set(consts.HeaderContentType, consts.ContentTypeJsonUtf8)
		json.NewEncoder(w).Encode(res)
	})
}
