package client

import "github.com/hopeio/utils/net/http/client"

type ResponseInterface[T any] interface {
	client.ResponseBodyCheck
	SubData() T
}

// 一个语法糖，一般不用
type SubDataRequest[RES ResponseInterface[T], T any] Request[RES]

func NewSubDataRequest[RES ResponseInterface[T], T any](req *client.Request) *SubDataRequest[RES, T] {
	return (*SubDataRequest[RES, T])(req)
}

func (req *SubDataRequest[RES, T]) Origin() *client.Request {
	return (*client.Request)(req)
}

// Do create a HTTP request
func (r *SubDataRequest[RES, T]) Do(param any) (T, error) {
	var response RES
	err := (*client.Request)(r).Do(param, response)
	if err != nil {
		return response.SubData(), err
	}
	return response.SubData(), err
}

func (req *SubDataRequest[RES, T]) SubData(param any) (T, error) {
	var response RES
	err := (*client.Request)(req).Do(param, &response)
	if err != nil {
		return response.SubData(), err
	}

	return response.SubData(), nil
}
