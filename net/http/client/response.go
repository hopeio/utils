package client

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
)

type ResponseBody httpi.ResAnyData

func CommonResponse(response interface{}) ResponseBodyCheck {
	return &ResponseBody{Details: response}
}

func (res *ResponseBody) CheckError() error {
	if res.Code != 0 {
		return fmt.Errorf("code: %d, message: %s", res.Code, res.Message)
	}
	return nil
}

type ResponseBody2 struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func CommonResponse2(response interface{}) ResponseBodyCheck {
	return &ResponseBody2{Data: response}
}

func (res *ResponseBody2) CheckError() error {
	if res.Code != 0 {
		return fmt.Errorf("status:%d,message:%s", res.Code, res.Msg)
	}
	return nil
}

var (
	ErrNotFound            = fmt.Errorf("not found")
	ErrRangeNotSatisfiable = fmt.Errorf("range not satisfiable")
)

type ResponseBody3 struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func CommonResponse3(response interface{}) ResponseBodyCheck {
	return &ResponseBody3{Data: response}
}

func (res *ResponseBody3) CheckError() error {
	if res.Status != 0 {
		return fmt.Errorf("status:%d,message:%s", res.Status, res.Message)
	}
	return nil
}
