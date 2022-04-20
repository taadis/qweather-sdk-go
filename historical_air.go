package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7HistoricalAirRequest 历史空气质量请求
type V7HistoricalAirRequest struct {
	baseRequest
	Location string `json:"location"`       // 需要查询的地区,仅支持LocationID
	Date     string `json:"date"`           // 选择日期
	Lang     string `json:"lang,omitempty"` // 多语言设置,默认中文
}

func NewV7HistoricalAirRequest() *V7HistoricalAirRequest {
	r := new(V7HistoricalAirRequest)
	return r
}

func (r *V7HistoricalAirRequest) Method() string {
	return http.MethodGet
}

func (r *V7HistoricalAirRequest) Url() string {
	return fmt.Sprintf("%s/v7/historical/air", DatasetApi)
}

func (r *V7HistoricalAirRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7HistoricalAirResponse 历史空气质量响应
type V7HistoricalAirResponse struct {
	baseResponse
	FxLink    string `json:"fxLink"` // 当前数据的响应式页面,便于嵌入网站或应用
	AirHourly []struct {
		PubTime  string `json:"pubTime"`  // 空气质量数据发布时间
		AQI      string `json:"aqi"`      // 空气质量指数
		Level    string `json:"level"`    // 空气质量指数等级
		Category string `json:"category"` // 空气质量指数级别
		Primary  string `json:"primary"`  // 空气质量的主要污染物,空气质量为优时,返回值为NA
		PM10     string `json:"pm10"`     // PM10
		PM2p5    string `json:"pm2p5"`    // PM2.5
		NO2      string `json:"no2"`      // 二氧化氮
		SO2      string `json:"so2"`      // 二氧化硫
		CO       string `json:"co"`       // 一氧化碳
		O3       string `json:"o3"`       // 臭氧
	} `json:"airHourly"`
}

func (r *V7HistoricalAirResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7HistoricalAir(req *V7HistoricalAirRequest) (*V7HistoricalAirResponse, error) {
	resp := new(V7HistoricalAirResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
