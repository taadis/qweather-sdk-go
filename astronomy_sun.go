package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7AstronomySunRequest 日出日落请求,获取最近60天全球城市日出日落时间
type V7AstronomySunRequest struct {
	baseRequest
	Location string `json:"location"`       // 需要查询地区的LocationID或以英文逗号分隔的经纬度坐标,如:location=101010100或location=116.41,39.92
	Date     string `json:"date"`           // 选择日期
	Lang     string `json:"lang,omitempty"` // 多语言设置,默认中文
}

func NewV7AstronomySunRequest() *V7AstronomySunRequest {
	r := new(V7AstronomySunRequest)
	return r
}

func (r *V7AstronomySunRequest) Method() string {
	return http.MethodGet
}

func (r *V7AstronomySunRequest) Url() string {
	return fmt.Sprintf("%s/v7/astronomy/sun", Api)
}

func (r *V7AstronomySunRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7AstronomySunResponse 日出日落响应
type V7AstronomySunResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Sunrise    string `json:"sunrise"`    // 日出时间
	Sunset     string `json:"sunset"`     // 日落时间
}

func (r *V7AstronomySunResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7AstronomySun(req *V7AstronomySunRequest) (*V7AstronomySunResponse, error) {
	resp := new(V7AstronomySunResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
