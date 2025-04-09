/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"fmt"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/reflect/mtos"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, obj any) error {
	return binding.CommonBind(RequestSource{ctx}, obj)
}

type RequestSource struct {
	*gin.Context
}

func (s RequestSource) Uri() mtos.Setter {
	return (uriSource)(s.Params)
}

func (s RequestSource) Query() mtos.Setter {
	return (mtos.KVsSource)(s.Request.URL.Query())
}

func (s RequestSource) Header() mtos.Setter {
	return (binding.HeaderSource)(s.Request.Header)
}

func (s RequestSource) Form() mtos.Setter {
	contentType := s.Request.Header.Get(consts.HeaderContentType)
	if contentType == consts.ContentTypeForm {
		err := s.Request.ParseForm()
		if err != nil {
			return nil
		}
		return (mtos.KVsSource)(s.Request.PostForm)
	}
	if contentType == consts.ContentTypeMultipart {
		err := s.Request.ParseMultipartForm(binding.DefaultMemory)
		if err != nil {
			return nil
		}
		return (*binding.MultipartSource)(s.Request.MultipartForm)
	}
	return nil
}

func (s RequestSource) BodyBind(obj any) error {
	if s.Request.Method == http.MethodGet {
		return nil
	}
	data, err := io.ReadAll(s.Request.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	return binding.BodyUnmarshaller(data, obj)
}
