/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"net/url"
)

type Elasticsearch struct {
	Es    *elasticsearch.Client
	Index string
}

// TODO
func (k *Elasticsearch) Write(b []byte) (n int, err error) {
	return
}

func (k *Elasticsearch) Sync() error {
	return nil
}

func (k *Elasticsearch) Close() error {
	return nil
}

// kibana://${token}?index=${index}
func RegisterSink() {
	_ = zap.RegisterSink("elastic", func(url *url.URL) (sink zap.Sink, e error) {
		k := new(Elasticsearch)
		k.Es, e = elasticsearch.NewClient(elasticsearch.Config{})
		k.Index = url.Query().Get("index")
		return k, e
	})
}

func New(es *elasticsearch.Client, index string) zap.Sink {
	return &Elasticsearch{es, index}
}
