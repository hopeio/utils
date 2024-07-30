package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"

	"github.com/hopeio/utils/encoding/protobuf/jsonpb"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/grpc/reconn"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"strings"
)

func HttpError(ctx *gin.Context, err error) {

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

	delete(ctx.Request.Header, httpi.HeaderTrailer)
	contentType := jsonpb.JsonPb.ContentType(nil)
	ctx.Header(httpi.HeaderContentType, contentType)

	se := &errcode.ErrRep{Code: errcode.ErrCode(s.Code()), Msg: s.Message()}

	buf, merr := jsonpb.JsonPb.Marshal(se)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", se, merr)
		ctx.Status(http.StatusInternalServerError)
		if _, err := io.WriteString(ctx.Writer, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	if _, err := ctx.Writer.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

}
