package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type V7AirNowRequest struct {
	baseRequest
	Location string `json:"location"`
	Lang     string `json:"lang,omitempty"`
}

func NewV7AirNowRequest() *V7AirNowRequest {
	r := new(V7AirNowRequest)
	return r
}

func (r *V7AirNowRequest) Method() string {
	return http.MethodGet
}

func (r *V7AirNowRequest) Url() string {
	return fmt.Sprintf("%s/v7/air/now", Api)
}

func (r *V7AirNowRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7AirNowResponse 实时空气质量响应
type V7AirNowResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Now        struct {
		PubTime  string `json:"pubTime"`  // 空气质量数据发布时间
		AQI      string `json:"aqi"`      // 空气质量指数
		Level    string `json:"level"`    // 空气质量指数等级
		Category string `json:"category"` // 空间质量指数级别
		Primary  string `json:"primary"`  // 空气质量的主要污染物,空气质量为优时,返回值为NA
		PM10     string `json:"pm10"`     // PM10
		PM2p5    string `json:"pm2p5"`    // PM2.5
		NO2      string `json:"no2"`      // 二氧化氮
		SO2      string `json:"so2"`      // 二氧化硫
		CO       string `json:"co"`       // 一氧化碳
		O3       string `json:"o3"`       // 臭氧
	} `json:"now"`
	Station struct {
		Name     string `json:"name"`     // 监测站名称
		ID       string `json:"id"`       // 监测站ID
		PubTime  string `json:"pubTime"`  // 空气质量数据发布时间
		AQI      string `json:"aqi"`      // 空气质量指数
		Level    string `json:"level"`    // 空气质量指数等级
		Category string `json:"category"` // 空气质量指数级别
		Primary  string `json:"primary"`  // 空气质量的主要污染物,空气质量为优时,返回值为NA
		PM10     string `json:"pm10"`     // PM10
		PM2p5    string `json:"pm2.5"`    // PM2.5
		NO2      string `json:"no2"`      // 二氧化氮
		SO2      string `json:"so2"`      // 二氧化硫
		CO       string `json:"co"`       // 一氧化碳
		O3       string `json:"o3"`       // 臭氧
	} `json:"station"`
}

func (r *V7AirNowResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7AirNow(req *V7AirNowRequest) (*V7AirNowResponse, error) {
	resp := new(V7AirNowResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
