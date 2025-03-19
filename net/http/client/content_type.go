/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"github.com/hopeio/utils/net/http/consts"
	"strings"
)

type ContentType uint8

func (c ContentType) String() string {
	if c < ContentTypeApplication {
		return contentTypes[c] + ";charset=UTF-8"
	}
	return consts.ContentTypeOctetStream + ";charset=UTF-8"
}

func (c *ContentType) Decode(contentType string) {
	if strings.HasPrefix(contentType, consts.ContentTypeJson) {
		*c = ContentTypeJson
	} else if strings.HasPrefix(contentType, consts.ContentTypeForm) {
		*c = ContentTypeForm
	} else if strings.HasPrefix(contentType, "text") {
		*c = ContentTypeText
	} else if strings.HasPrefix(contentType, "image") {
		*c = ContentTypeImage
	} else if strings.HasPrefix(contentType, "video") {
		*c = ContentTypeVideo
	} else if strings.HasPrefix(contentType, "audio") {
		*c = ContentTypeAudio
	} else if strings.HasPrefix(contentType, "application") {
		*c = ContentTypeApplication
	} else {
		*c = ContentTypeJson
	}
}

const (
	ContentTypeJson ContentType = iota
	ContentTypeForm
	ContentTypeFormData
	ContentTypeGrpc
	ContentTypeGrpcWeb
	ContentTypeXml
	ContentTypeText
	ContentTypeBinary
	ContentTypeApplication
	ContentTypeImage
	ContentTypeAudio
	ContentTypeVideo
	contentTypeUnSupport
)

var contentTypes = []string{
	consts.ContentTypeJson,
	consts.ContentTypeForm,
	consts.ContentTypeMultipart,
	consts.ContentTypeGrpc,
	consts.ContentTypeGrpcWeb,
	consts.ContentTypeXmlUnreadable,
	consts.ContentTypeText,
	consts.ContentTypeOctetStream,
	/*	consts.ContentImagePngHeaderValue,
		consts.ContentImageJpegHeaderValue,
		consts.ContentImageGifHeaderValue,
		consts.ContentImageBmpHeaderValue,
		consts.ContentImageWebpHeaderValue,
		consts.ContentImageAvifHeaderValue,
		consts.ContentImageTiffHeaderValue,
		consts.ContentImageXIconHeaderValue,
		consts.ContentImageVndMicrosoftIconHeaderValue,*/
}
