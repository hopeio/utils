/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"github.com/hopeio/gox/net/http/client"
)

func GetRequest[RES any](url string) *Request[RES] {
	return (*Request[RES])(client.GetRequest(url))
}

func PostRequest[RES any](url string) *Request[RES] {
	return (*Request[RES])(client.PostRequest(url))
}

func PutRequest[RES any](url string) *Request[RES] {
	return (*Request[RES])(client.PutRequest(url))
}

func DeleteRequest[RES any](url string) *Request[RES] {
	return (*Request[RES])(client.DeleteRequest(url))
}

func Get[RES any](url string, param any) (*RES, error) {
	return GetRequest[RES](url).Do(param)
}

func Post[RES any](url string, param any) (*RES, error) {
	return PostRequest[RES](url).Do(param)
}

func Put[RES any](url string, param any) (*RES, error) {
	return PutRequest[RES](url).Do(param)
}

func Delete[RES any](url string, param any) (*RES, error) {
	return DeleteRequest[RES](url).Do(param)
}
