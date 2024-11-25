/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package grpc_gateway

import (
	"github.com/hopeio/utils/encoding/json"
	responsei "github.com/hopeio/utils/net/http"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var JsonPb = &JSONPb{}

type JSONPb struct {
}

func (*JSONPb) ContentType(_ interface{}) string {
	return "application/json"
}

func (j *JSONPb) Marshal(v any) ([]byte, error) {
	if _, ok := v.(error); ok {
		return json.Marshal(v)
	}
	if msg, ok := v.(*wrapperspb.StringValue); ok {
		v = msg.Value
	}
	return json.Marshal(&responsei.ResAnyData{
		Data: v,
	})
}

func (j *JSONPb) Name() string {
	return "jsonpb"
}

func (j *JSONPb) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *JSONPb) Delimiter() []byte {
	return []byte("\n")
}

// NewDecoder returns a runtime.Decoder which reads JSON stream from "r".
func (j *JSONPb) NewDecoder(r io.Reader) runtime.Decoder {
	return json.NewDecoder(r)
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (j *JSONPb) NewEncoder(w io.Writer) runtime.Encoder {
	return json.NewEncoder(w)
}

func (j *JSONPb) ContentTypeFromMessage(v interface{}) string {
	return j.ContentType(v)
}
