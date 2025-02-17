/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"net/http"
	"testing"
	"time"
)

type OperatorReq struct {
	Id           uint64    `json:"id" comment:"运营商id" example:"1"`
	Ids          []uint64  `json:"ids" comment:"运营商id" example:"[1,2]"`
	Name         string    `json:"name" comment:"运营商名称" example:"xxx"`
	StartTime    time.Time `json:"startTime" comment:"起始时间"`
	EndTime      time.Time `json:"endTime" comment:"结束时间"`
	PageNo       int       `json:"pageNo" comment:"页数" example:"1"`
	PageNum      int       `json:"pageNum" comment:"每页数量" example:"20"`
	PlatformType uint8     `json:"platformType" comment:"平台类型,必传,传255为所有状态"`
}

type OperatorPublicInfo struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name" comment:"'公司名称"`
	ShortName string `json:"shortName" comment:"运营商名称/公司简称"`
}

type OperatorPublicList struct {
	List  []OperatorPublicInfo `json:"list"`
	Total uint64               `json:"total"`
}

func TestClient(t *testing.T) {
	req := &OperatorReq{
		PageNo:       1,
		PageNum:      10,
		PlatformType: 1,
	}
	res := &OperatorPublicList{}
	err := NewRequest(http.MethodPost, `http://test.xyz/api/list`).
		AddHeader("Auth", "e30=").Do(req, CommonResponse(res))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
