package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7AstronomyMoonRequest 月升月落和月相请求
type V7AstronomyMoonRequest struct {
	baseRequest
	Location string `json:"location"`
	Date     string `json:"date"`
	Lang     string `json:"lang,omitempty"`
}

func NewV7AstronomyMoonRequest() *V7AstronomyMoonRequest {
	r := new(V7AstronomyMoonRequest)
	return r
}

func (r *V7AstronomyMoonRequest) Method() string {
	return http.MethodGet
}

func (r *V7AstronomyMoonRequest) Url() string {
	return fmt.Sprintf("%s/v7/astronomy/moon", Api)
}

func (r *V7AstronomyMoonRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7AstronomyMoonResponse 月升月落和月相响应
type V7AstronomyMoonResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	MoonRise   string `json:"moonRise"`   // 月升时间
	MoonSet    string `json:"moonSet"`    // 月落时间
	MoonPhase  []struct {
		FxTime       string `json:"fxTime"`       // 月相逐小时预报时间
		Value        string `json:"value"`        // 月相数值
		Name         string `json:"name"`         // 月相名称
		Icon         string `json:"icon"`         // 月相图标代码,图标可通过[天气状况和图标]下载
		Illumination string `json:"illumination"` // 月亮照明度,百分比数值
	} `json:"moonPhase"`
}

func (r *V7AstronomyMoonResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7AstronomyMoon(req *V7AstronomyMoonRequest) (*V7AstronomyMoonResponse, error) {
	resp := new(V7AstronomyMoonResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
