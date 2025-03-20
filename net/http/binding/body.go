/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"fmt"
	"github.com/hopeio/utils/encoding"
	"io"
	"net/http"
)

// 避免代码膨胀,避免为每种编码写一个实现,事实上,一个服务器几乎只确定一种交互格式,所以做最小化实现,虽然引入gin不可避免的仍然会把所有包引入

type bodyBinding struct {
	name         string
	unmarshaller func([]byte, any) error
	decoder      func(io.Reader) encoding.Decoder
}

func (b *bodyBinding) Name() string {
	return b.name
}

func (b *bodyBinding) Bind(req *http.Request, obj interface{}) error {
	if b.decoder != nil {
		return b.decoder(req.Body).Decode(obj)
	}
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	return b.unmarshaller(data, obj)
}

func (b *bodyBinding) RegisterUnmarshaller(name string, unmarshaller func(data []byte, obj any) error) {
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func (b *bodyBinding) RegisterDecoder(name string, decoder func(io.Reader) encoding.Decoder) {
	CustomBody.name = name
	CustomBody.decoder = decoder
}
