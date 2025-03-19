/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gateway

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/net/http/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/textproto"
)

func OutgoingHeaderMatcher(key string) (string, bool) {
	switch key {
	case
		httpi.HeaderSetCookie:
		return key, true
	}
	return "", false
}

var headerMatcher = []string{httpi.HeaderSetCookie}

func HandleForwardResponseServerMetadata(w http.ResponseWriter, md metadata.MD) {
	for _, k := range headerMatcher {
		if vs, ok := md[k]; ok {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
	}
}

func HandleForwardResponseTrailerHeader(w http.ResponseWriter, md metadata.MD) {
	for k := range md {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", grpc.MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func HandleForwardResponseTrailer(w http.ResponseWriter, md metadata.MD) {
	for k, vs := range md {
		tKey := fmt.Sprintf("%s%s", grpc.MetadataTrailerPrefix, k)
		for _, v := range vs {
			w.Header().Add(tKey, v)
		}
	}
}
