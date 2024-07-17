package http

const (
	// ContentJavascriptHeaderValue header value for JSONP & Javascript data.
	ContentJavascriptHeaderValue = "text/javascript"
	// ContentHtmlHeaderValue is the  string of text/html response header's content type value.
	ContentHtmlHeaderValue = "text/html"
	ContentCssHeaderValue  = "text/css"
	// ContentTextHeaderValue header value for Text data.
	ContentTextHeaderValue = "text/plain"
	// ContentXmlHeaderValue header value for XML data.
	ContentXmlHeaderValue = "text/xml"
	// ContentMarkdownHeaderValue custom key/content type, the real is the text/html.
	ContentMarkdownHeaderValue = "text/markdown"
	// ContentYamlTextHeaderValue header value for YAML plain text.
	ContentYamlTextHeaderValue = "text/yaml"

	// ContentFormMultipartHeaderValue header value for post multipart form data.
	ContentFormMultipartHeaderValue = "multipart/form-data"

	// ContentBinaryHeaderValue header value for binary data.
	ContentBinaryHeaderValue = "application/octet-stream"
	// ContentWebassemblyHeaderValue header value for web assembly files.
	ContentWebassemblyHeaderValue = "application/wasm"
	// ContentJsonHeaderValue header value for JSON data.
	ContentJsonHeaderValue = "application/json"
	// ContentJsonProblemHeaderValue header value for JSON API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentJsonProblemHeaderValue = "application/problem+json"
	// ContentXmlProblemHeaderValue header value for XML API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentXmlProblemHeaderValue           = "application/problem+xml"
	ContentJavascriptUnreadableHeaderValue = "application/javascript"
	// ContentXmlUnreadableHeaderValue obsolete header value for XML.
	ContentXmlUnreadableHeaderValue = "application/xml"
	// ContentYamlHeaderValue header value for YAML data.
	ContentYamlHeaderValue = "application/x-yaml"
	// ContentProtobufHeaderValue header value for Protobuf messages data.
	ContentProtobufHeaderValue = "application/x-protobuf"
	// ContentMsgPackHeaderValue header value for MsgPack data.
	ContentMsgPackHeaderValue = "application/msgpack"
	// ContentMsgPack2HeaderValue alternative header value for MsgPack data.
	ContentMsgPack2HeaderValue = "application/x-msgpack"
	// ContentFormHeaderValue header value for post form data.
	ContentFormHeaderValue = "application/x-www-form-urlencoded"

	// ContentGrpcHeaderValue Content-Type header value for gRPC.
	ContentGrpcHeaderValue      = "application/grpc"
	ContentGrpcWebHeaderValue   = "application/grpc-web"
	ContentPdfHeaderValue       = "application/pdf"
	ContentJsonUtf8HeaderValue  = "application/json;charset=utf-8"
	ContentFormParamHeaderValue = "application/x-www-form-urlencoded;param=value"

	ContentImagePngHeaderValue  = "image/png"
	ContentImageJpegHeaderValue = "image/jpeg"
	ContentImageGifHeaderValue  = "image/gif"
	ContentImageBmpHeaderValue  = "image/bmp"
	ContentImageWebpHeaderValue = "image/webp"
	ContentImageAvifHeaderValue = "image/avif"
	//ContentImageHeifHeaderValue = "image/heif"
	ContentImageSvgHeaderValue              = "image/svg+xml"
	ContentImageTiffHeaderValue             = "image/tiff"
	ContentImageXIconHeaderValue            = "image/x-icon"
	ContentImageVndMicrosoftIconHeaderValue = "image/vnd.microsoft.icon"

	ContentCharsetUtf8HeaderValue = "charset=UTF-8"
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
