package client

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/client"
)

type ResponseBody[RES any] httpi.ResData[RES]

func CommonResponse[RES any]() client.ResponseBodyCheck {
	return &ResponseBody[RES]{}
}

func (res *ResponseBody[RES]) CheckError() error {
	if res.Code != 0 {
		return fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
	}
	return nil
}

func (res *ResponseBody[RES]) GetData() *RES {
	return &res.Data
}

type ResponseBody2[RES any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data RES    `json:"data"`
}

func CommonResponse2[RES any]() client.ResponseBodyCheck {
	return &ResponseBody2[RES]{}
}

func (res *ResponseBody2[RES]) CheckError() error {
	if res.Code != 0 {
		return fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
	}
	return nil
}

func (res *ResponseBody2[RES]) GetData() *RES {
	return &res.Data
}

type ResponseBody3[RES any] struct {
	Status  int    `json:"status"`
	Data    RES    `json:"data"`
	Message string `json:"message"`
}

func CommonResponse3[RES any]() client.ResponseBodyCheck {
	return &ResponseBody3[RES]{}
}

func (res *ResponseBody3[RES]) CheckError() error {
	if res.Status != 0 {
		return fmt.Errorf("status: %d, message: %s", res.Status, res.Message)
	}
	return nil
}
