package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7Air5DaysRequest 空气质量预报请求
type V7Air5DaysRequest struct {
	baseRequest
	Location string `json:"location"`
	Lang     string `json:"lang"`
}

func NewV7Air5DaysRequest() *V7Air5DaysRequest {
	r := new(V7Air5DaysRequest)
	return r
}

func (r *V7Air5DaysRequest) Method() string {
	return http.MethodGet
}

func (r *V7Air5DaysRequest) Url() string {
	return fmt.Sprintf("%s/v7/air/5d", Api)
}

func (r *V7Air5DaysRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7Air5DaysResponse 空气质量预报响应
type V7Air5DaysResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Daily      []struct {
		FxDate   string `json:"fxDate"`   // 预报日期
		AQI      string `json:"aqi"`      // 空气质量指数
		Level    string `json:"level"`    // 空气质量指数等级
		Category string `json:"category"` // 空气质量指数级别
		Primary  string `json:"primary"`  // 空气质量的主要污染物,空气质量为优时,返回值为NA
	} `json:"daily"`
}

func (r *V7Air5DaysResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7Air5Days(req *V7Air5DaysRequest) (*V7Air5DaysResponse, error) {
	resp := new(V7Air5DaysResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
