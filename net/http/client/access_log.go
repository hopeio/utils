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

type AccessLogParam struct {
	Method, Url       string
	Request           *http.Request
	Response          *http.Response
	ReqBody, RespBody []byte
	ProcessTime       time.Duration
}
type AccessLog func(param *AccessLogParam, err error)

func DefaultLogger(param *AccessLogParam, err error) {
	reqField, respField, statusField := zap.Skip(), zap.Skip(), zap.Skip()
	if len(param.ReqBody) > 0 {
		key := "body"
		if param.ReqBody.IsJson() {
			reqField = zap.Reflect(key, log.RawJson(param.ReqBody))
		} else {
			reqField = zap.String(key, stringsi.BytesToString(param.ReqBody))
		}
	}
	if len(param.RespBody) > 0 {
		key := "result"
		if len(param.RespBody) > 500 {
			respField = zap.String(key, "result is too long")
		} else {
			respField = zap.Reflect(key, log.RawJson(param.RespBody))
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
