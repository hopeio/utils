package client

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/client"
)

type ResData[RES any] httpi.ResData[RES]

func CommonResponse[RES any]() client.ResponseBodyCheck {
	return &ResData[RES]{}
}

func (res *ResData[RES]) CheckError() error {
	if res.Code != 0 {
		return fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
	}
	return nil
}

func (res *ResData[RES]) GetData() *RES {
	return &res.Data
}
