/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	httpi "github.com/hopeio/utils/net/http"
	"strings"
)

type ContentType uint8

func (c ContentType) String() string {
	if c < ContentTypeApplication {
		return contentTypes[c] + ";charset=UTF-8"
	}
	return httpi.ContentTypeOctetStream + ";charset=UTF-8"
}

func (c *ContentType) Decode(contentType string) {
	if strings.HasPrefix(contentType, httpi.ContentTypeJson) {
		*c = ContentTypeJson
	} else if strings.HasPrefix(contentType, httpi.ContentTypeForm) {
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
	httpi.ContentTypeJson,
	httpi.ContentTypeForm,
	httpi.ContentTypeMultipart,
	httpi.ContentTypeGrpc,
	httpi.ContentTypeGrpcWeb,
	httpi.ContentTypeXmlUnreadable,
	httpi.ContentTypeText,
	httpi.ContentTypeOctetStream,
	/*	httpi.ContentImagePngHeaderValue,
		httpi.ContentImageJpegHeaderValue,
		httpi.ContentImageGifHeaderValue,
		httpi.ContentImageBmpHeaderValue,
		httpi.ContentImageWebpHeaderValue,
		httpi.ContentImageAvifHeaderValue,
		httpi.ContentImageTiffHeaderValue,
		httpi.ContentImageXIconHeaderValue,
		httpi.ContentImageVndMicrosoftIconHeaderValue,*/
}
