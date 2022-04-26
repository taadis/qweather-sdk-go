package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7OceanTideRequest 潮汐请求
type V7OceanTideRequest struct {
	baseRequest
	Location string `json:"location"`       // 需要查询的潮汐站点,请填写潮汐站点的LocationID
	Date     string `json:"date,omitempty"` // 选择日期,最多可选择未来10天(包含今天)的数据,日期格式为yyyyMMdd,例如date=20200531
}

func NewV7OceanTideRequest() *V7OceanTideRequest {
	r := new(V7OceanTideRequest)
	return r
}

func (r *V7OceanTideRequest) Method() string {
	return http.MethodGet
}

func (r *V7OceanTideRequest) Url() string {
	return fmt.Sprintf("%s/v7/ocean/tide", Api)
}

func (r *V7OceanTideRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7OceanTideResponse 潮汐响应
type V7OceanTideResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新事件
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	TideTable  []struct {
		FxTime string `json:"fxTime"` // 满潮或干潮时间
		Height string `json:"height"` // 海水高度,单位:米
		Type   string `json:"满潮(H)或干潮(L)"`
	} `json:"tideTable"`
	TideHourly []struct {
		FxTime string `json:"fxTime"` // 逐小时预报时间
		Height string `json:"height"` // 海水高度,单位:米
	} `json:"tideHourly"`
}

func (r *V7OceanTideResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7OceanTide(req *V7OceanTideRequest) (*V7OceanTideResponse, error) {
	resp := new(V7OceanTideResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
