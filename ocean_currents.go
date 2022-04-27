package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7OceanCurrentsRequest 潮流请求
type V7OceanCurrentsRequest struct {
	baseRequest
	Location string `json:"location"` // 需要查询的潮流站点,请填写潮流站点的LocationID
	Date     string `json:"date"`     // 选择日期,最多可选未来10天(包含今天)的数据,日期格式为yyyyMMdd,例如:date=20200531
}

func NewV7OceanCurrentsRequest() *V7OceanCurrentsRequest {
	r := new(V7OceanCurrentsRequest)
	return r
}

func (r *V7OceanCurrentsRequest) Method() string {
	return http.MethodGet
}

func (r *V7OceanCurrentsRequest) Url() string {
	return fmt.Sprintf("%s/v7/ocean/currents", Api)
}

func (r *V7OceanCurrentsRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7OceanCurrentsResponse 潮流响应
type V7OceanCurrentsResponse struct {
	baseResponse
	UpdateTime    string `json:"updateTime"` // 当前API的最近更新时间
	FxLink        string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	CurrentsTable []struct {
		FxTime   string `json:"fxTime"`   // 潮流最大流速时间
		SpeedMax string `json:"speedMax"` // 潮流最大流速,单位:厘米/秒
		Dir360   string `json:"dir360"`   // 潮流360度方向
	} `json:"currentsTable"`
	CurrentsHourly []struct {
		FxTime string `json:"fxTime"` // 逐小时预报时间
		Speed  string `json:"speed"`  // 潮流流速,单位:厘米/秒
		Dir360 string `json:"dir360"` // 潮流360度方向
	} `json:"currentHourly"`
}

func (r *V7OceanCurrentsResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7OceanCurrents(req *V7OceanCurrentsRequest) (*V7OceanCurrentsResponse, error) {
	resp := new(V7OceanCurrentsResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
