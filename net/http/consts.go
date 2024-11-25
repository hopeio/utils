/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

const (
	// ContentTypeJavascript header value for JSONP & Javascript data.
	ContentTypeJavascript = "text/javascript"
	// ContentTypeHtml is the  string of text/html response header's content type value.
	ContentTypeHtml = "text/html"
	ContentTypeCss  = "text/css"
	// ContentTypeText header value for Text data.
	ContentTypeText = "text/plain"
	// ContentTypeXml header value for XML data.
	ContentTypeXml = "text/xml"
	// ContentTypeMarkdown custom key/content type, the real is the text/html.
	ContentTypeMarkdown = "text/markdown"
	// ContentTypeYamlText header value for YAML plain text.
	ContentTypeYamlText = "text/yaml"

	// ContentTypeMultipart header value for post multipart form data.
	ContentTypeMultipart = "multipart/form-data"

	// ContentTypeOctetStream header value for binary data.
	ContentTypeOctetStream = "application/octet-stream"
	// ContentTypeWebassembly header value for web assembly files.
	ContentTypeWebassembly = "application/wasm"
	// ContentTypeJson header value for JSON data.
	ContentTypeJson = "application/json"
	// ContentTypeJsonProblem header value for JSON API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentTypeJsonProblem = "application/problem+json"
	// ContentTypeXmlProblem header value for XML API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentTypeXmlProblem           = "application/problem+xml"
	ContentTypeJavascriptUnreadable = "application/javascript"
	// ContentTypeXmlUnreadable obsolete header value for XML.
	ContentTypeXmlUnreadable = "application/xml"
	// ContentTypeYaml header value for YAML data.
	ContentTypeYaml = "application/x-yaml"
	// ContentTypeProtobuf header value for Protobuf messages data.
	ContentTypeProtobuf = "application/x-protobuf"
	// ContentTypeMsgPack header value for MsgPack data.
	ContentTypeMsgPack = "application/msgpack"
	// ContentTypeMsgPack2 alternative header value for MsgPack data.
	ContentTypeMsgPack2 = "application/x-msgpack"
	// ContentTypeForm header value for post form data.
	ContentTypeForm = "application/x-www-form-urlencoded"

	// ContentTypeGrpc Content-Type header value for gRPC.
	ContentTypeGrpc    = "application/grpc"
	ContentTypeGrpcWeb = "application/grpc-web"
	ContentPdf         = "application/pdf"
	ContentJsonUtf8    = "application/json;charset=utf-8"
	ContentFormParam   = "application/x-www-form-urlencoded;param=value"

	ContentImagePng                  = "image/png"
	ContentImageJpeg                 = "image/jpeg"
	ContentImageGif                  = "image/gif"
	ContentImageBmp                  = "image/bmp"
	ContentImageWebp                 = "image/webp"
	ContentImageAvif                 = "image/avif"
	ContentTypeImageHeif             = "image/heif"
	ContentTypeImageSvg              = "image/svg+xml"
	ContentTypeImageTiff             = "image/tiff"
	ContentTypeImageXIcon            = "image/x-icon"
	ContentTypeImageVndMicrosoftIcon = "image/vnd.microsoft.icon"

	ContentTypeCharsetUtf8 = "charset=UTF-8"
)

const (
	HeaderDeviceInfo = "Device-AuthInfo"
	HeaderLocation   = "Location"
	HeaderArea       = "Area"
)

const (
	HeaderUserAgent                   = "User-Agent"
	HeaderXForwardedFor               = "X-Forwarded-For"
	HeaderXAccelBuffering             = "X-Accel-Buffering"
	HeaderAuth                        = "HeaderAuth"
	HeaderContentType                 = "Content-Type"
	HeaderTrace                       = "Tracing"
	HeaderTraceID                     = "Tracing-ID"
	HeaderTraceBin                    = "Tracing-Bin"
	HeaderAuthorization               = "Authorization"
	HeaderCookie                      = "Cookie"
	HeaderCookieValueToken            = "token"
	HeaderCookieValueDel              = "del"
	HeaderContentDisposition          = "Content-Disposition"
	HeaderContentEncoding             = "Content-Encoding"
	HeaderReferer                     = "Referer"
	HeaderAccept                      = "Accept"
	HeaderAcceptLanguage              = "Accept-Language"
	HeaderAcceptEncoding              = "Accept-Encoding"
	HeaderCacheControl                = "Cache-Control"
	HeaderSetCookie                   = "Set-Cookie"
	HeaderTrailer                     = "Trailer"
	HeaderTransferEncoding            = "Transfer-Encoding"
	HeaderTransferEncodingChunked     = "chunked"
	HeaderInternal                    = "Internal"
	HeaderTE                          = "TE"
	HeaderLastModified                = "Last-Modified"
	HeaderContentLength               = "Content-Length"
	HeaderAccessControlRequestMethod  = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders = "Access-Control-Request-Headers"
	HeaderOrigin                      = "Origin"
	HeaderConnection                  = "Connection"
	HeaderRange                       = "Range"
	HeaderContentRange                = "Content-Range"
	HeaderAcceptRanges                = "Accept-Ranges"
)

const (
	HeaderGrpcTraceBin = "grpc-trace-bin"
	HeaderGrpcInternal = "grpc-internal"
)

const (
	FormDataFieldTmpl = `form-data; name="%s"`
	FormDataFileTmpl  = `form-data; name="%s"; filename="%s"`
)
