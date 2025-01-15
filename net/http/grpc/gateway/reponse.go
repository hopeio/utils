package gateway

import (
	"context"
	"github.com/hopeio/utils/encoding/protobuf/jsonpb"
	httpi "github.com/hopeio/utils/net/http"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func Response(ctx context.Context, writer http.ResponseWriter, message proto.Message) error {
	if v, ok := message.(httpi.IHttpResponse); ok {
		_, err := httpi.ResponseWrite(writer, v)
		return err
	}
	var buf []byte
	var err error
	switch rb := message.(type) {
	case responseBody:
		buf, err = jsonpb.JsonPb.Marshal(rb.ResponseBody())
	case xxxResponseBody:
		buf, err = jsonpb.JsonPb.Marshal(rb.XXX_ResponseBody())
	default:
		buf, err = jsonpb.JsonPb.Marshal(message)
	}

	if err != nil {
		grpclog.Infof("Marshal error: %v", err)
		return err
	}

	if _, err = writer.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
	return nil
}

type xxxResponseBody interface {
	XXX_ResponseBody() interface{}
}

type responseBody interface {
	ResponseBody() interface{}
}
