package gateway

import (
	"context"
	"github.com/hopeio/protobuf/response"
	"google.golang.org/protobuf/proto"
	"net/http"
)

func ResponseHook(ctx context.Context, writer http.ResponseWriter, message proto.Message) error {
	if res, ok := message.(*response.HttpResponse); ok {
		hlen := len(res.Header)
		for i := 0; i < hlen && i+1 < hlen; i += 2 {
			writer.Header().Set(res.Header[i], res.Header[i+1])
		}
		writer.WriteHeader(int(res.StatusCode))
		writer.Write(res.Body)
	}
	/*	if message == nil{
		*(&message) = &response.TinyRep{Message: "OK"}
	}*/
	return nil
}
