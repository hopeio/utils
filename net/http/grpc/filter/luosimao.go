package filter

import (
	"context"
	"github.com/hopeio/utils/sdk/luosimao"
	"google.golang.org/grpc/metadata"
)

func LuosimaoVerify(reqURL, apiKey string, ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)
	response := md.Get("luosimao")
	if len(response) == 0 || response[0] == "" {
		return luosimao.Error
	}
	return luosimao.Verify(reqURL, apiKey, response[0])
}
