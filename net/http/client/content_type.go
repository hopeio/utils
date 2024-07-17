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
	return httpi.ContentBinaryHeaderValue + ";charset=UTF-8"
}

func (c *ContentType) Decode(contentType string) {
	if strings.HasPrefix(contentType, httpi.ContentJsonHeaderValue) {
		*c = ContentTypeJson
	} else if strings.HasPrefix(contentType, httpi.ContentFormHeaderValue) {
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
	httpi.ContentJsonHeaderValue,
	httpi.ContentFormHeaderValue,
	httpi.ContentFormMultipartHeaderValue,
	httpi.ContentGrpcHeaderValue,
	httpi.ContentGrpcWebHeaderValue,
	httpi.ContentXmlUnreadableHeaderValue,
	httpi.ContentTextHeaderValue,
	httpi.ContentBinaryHeaderValue,
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
