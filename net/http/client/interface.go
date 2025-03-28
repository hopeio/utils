/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import "net/http"

type MarshalBody interface {
	MarshalBody(contentType string) ([]byte, error)
}

type UnmarshalBody interface {
	UnmarshalBody(contentType string, body []byte) error
}

type SetRequest interface {
	SetRequest(*http.Request)
}

type FromResponse interface {
	FromResponse(response *http.Response)
}

type ResponseBodyCheck interface {
	CheckError() error
}
