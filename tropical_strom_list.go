package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7TropicalStormListRequest 台风列表请求
type V7TropicalStormListRequest struct {
	baseRequest
	Basin string `json:"basin"` // 需要查询的台风所在的流域
	Year  string `json:"year"`  // 支持查询本年度和上一年度的台风,例如:year=2020,year=2019
}

func NewV7TropicalStormListRequest() *V7TropicalStormListRequest {
	r := new(V7TropicalStormListRequest)
	return r
}

func (r *V7TropicalStormListRequest) Method() string {
	return http.MethodGet
}

func (r *V7TropicalStormListRequest) Url() string {
	return fmt.Sprintf("%s/v7/tropical/storm-list", Api)
}

func (r *V7TropicalStormListRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7TropicalStormListResponse 台风列表响应
type V7TropicalStormListResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Storm      []struct {
		Id       string `json:"id"`       // 台风ID
		StormId  string `json:"stormId"`  // 台风ID
		Name     string `json:"name"`     // 台风名称
		Basin    string `json:"basin"`    // 台风所处流域
		Year     string `json:"year"`     // 台风所处年份
		IsActive string `json:"isActive"` // 是否为活跃台风,0=停编,2=活跃台风
	} `json:"storm"`
}

func (r *V7TropicalStormListResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7TropicalStormList(req *V7TropicalStormListRequest) (*V7TropicalStormListResponse, error) {
	resp := new(V7TropicalStormListResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
