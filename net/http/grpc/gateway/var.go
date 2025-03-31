package gateway

import httpi "github.com/hopeio/utils/net/http/consts"

var InComingHeader = []string{"Accept",
	"Accept-Charset",
	"Accept-Language",
	"Accept-Ranges",
	//"Token",
	"Cache-Control",
	"Content-Type",
	//"Cookie",
	"Date",
	"Expect",
	"From",
	"Host",
	"If-Match",
	"If-Modified-Since",
	"If-None-Match",
	"If-Schedule-Tag-Match",
	"If-Unmodified-Since",
	"Max-Forwards",
	"Origin",
	"Pragma",
	"Referer",
	"User-Agent",
	"Via",
	"Warning",
}

var OutgoingHeader = []string{httpi.HeaderSetCookie}
