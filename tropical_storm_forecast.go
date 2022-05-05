package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7TropicalStormForecastRequest 台风预报请求
type V7TropicalStormForecastRequest struct {
	baseRequest
	StormId string `json:"stormId"` // 需要查询的台风ID
}

func NewV7TropicalStormForecastRequest() *V7TropicalStormForecastRequest {
	r := new(V7TropicalStormForecastRequest)
	return r
}

func (r *V7TropicalStormForecastRequest) Method() string {
	return http.MethodGet
}

func (r *V7TropicalStormForecastRequest) Url() string {
	return fmt.Sprintf("%s/v7/tropical/storm-forecast", Api)
}

func (r *V7TropicalStormForecastRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7TropicalStormForecastResponse 台风预报响应
type V7TropicalStormForecastResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Forecast   []struct {
		FxTime    string `json:"fxTime"`    // 台风预报时间
		Lat       string `json:"lat"`       // 台风所处纬度
		Lon       string `json:"lon"`       // 台风所处经度
		Type      string `json:"type"`      // 台风类型
		Pressure  string `json:"pressure"`  // 台风中心气压
		WindSpeed string `json:"windSpeed"` // 台风附近最大的风速
		MoveSpeed string `json:"moveSpeed"` // 台风移动速度
		MoveDir   string `json:"moveDir"`   // 台风移动方位
		Move360   string `json:"move360"`   // 台风移动方位360度方向
	} `json:"forecast"`
}

func (r *V7TropicalStormForecastResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7TropicalStormForecast(req *V7TropicalStormForecastRequest) (*V7TropicalStormForecastResponse, error) {
	resp := new(V7TropicalStormForecastResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
