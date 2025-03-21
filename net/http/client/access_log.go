/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"github.com/hopeio/utils/log"
	stringsi "github.com/hopeio/utils/strings"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LogLevel int8

const (
	LogLevelSilent LogLevel = iota
	LogLevelError
	LogLevelInfo
)

type Body struct {
	Data        []byte
	ContentType ContentType
}

func NewBody(data []byte, contentType ContentType) *Body {
	return &Body{Data: data, ContentType: contentType}
}

func (b *Body) IsJson() bool {
	return b.ContentType == ContentTypeJson
}

func (b *Body) IsProtobuf() bool {
	return b.ContentType == ContentTypeGrpc
}

type AccessLogParam struct {
	Method, Url       string
	Request           *http.Request
	Response          *http.Response
	ReqBody, RespBody *Body
	ProcessTime       time.Duration
}
type AccessLog func(param *AccessLogParam, err error)

func DefaultLogger(param *AccessLogParam, err error) {
	reqField, respField, statusField := zap.Skip(), zap.Skip(), zap.Skip()
	if param.ReqBody != nil {
		key := "body"
		if param.ReqBody.IsJson() {
			reqField = zap.Reflect(key, log.RawJson(param.ReqBody.Data))
		} else if param.ReqBody.IsProtobuf() {
			reqField = zap.Binary(key, param.ReqBody.Data)
		} else {
			reqField = zap.String(key, stringsi.BytesToString(param.ReqBody.Data))
		}
	}
	if param.RespBody != nil && param.RespBody.Data != nil {
		key := "result"
		if param.RespBody.IsJson() {
			respField = zap.Reflect(key, log.RawJson(param.RespBody.Data))
		} else if param.RespBody.IsProtobuf() {
			respField = zap.Binary(key, param.RespBody.Data)
		} else {
			if len(param.RespBody.Data) > 500 {
				respField = zap.String(key, "result is too long")
			} else {
				respField = zap.String(key, stringsi.BytesToString(param.RespBody.Data))
			}
		}
	}
	if param.Response != nil {
		statusField = zap.Int("status", param.Response.StatusCode)
	}

	log.Default().Logger.Info("http request", zap.String("url", param.Url),
		zap.String("method", param.Method),
		reqField,
		zap.Duration("duration", param.ProcessTime),
		respField,
		statusField,
		zap.Error(err),
	)
}
