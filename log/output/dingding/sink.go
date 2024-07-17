package dingding

import (
	"github.com/hopeio/utils/sdk/dingding"
	"go.uber.org/zap"
	"net/url"
)

type sink dingding.RobotConfig

// TODO
func (th *sink) Write(b []byte) (n int, err error) {
	return
}

func (th *sink) Sync() error {
	return nil
}

func (th *sink) Close() error {
	return nil
}

// dingding://${token}?sercret=${sercret}
func RegisterSink() {
	_ = zap.RegisterSink("dingding", func(url *url.URL) (sinkv zap.Sink, e error) {
		th := new(sink)
		return th, nil
	})
}

func NewSink(token, secret string) zap.Sink {
	th := new(sink)
	th.Token = token
	th.Secret = secret
	return th
}
