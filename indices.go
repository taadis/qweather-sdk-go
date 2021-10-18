package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7IndicesRequest 天气生活指数请求
type V7IndicesRequest struct {
	baseRequest
	Location string `json:"location,omitempty"`
	Type     string `json:"type,omitempty"` // 生活指数的类型ID
	Lang     string `json:"lang,omitempty"` // 多语言设置
	Duration string `json:"duration,omitempty"`
}

func NewV7IndicesRequest() *V7IndicesRequest {
	return new(V7IndicesRequest)
}

func (r *V7IndicesRequest) Method() string {
	return http.MethodGet
}

func (r *V7IndicesRequest) Url() string {
	if r.IsDev {
		return fmt.Sprintf("%s/v7/indices/%s", DevApi, r.Duration)
	}
	return fmt.Sprintf("%s/v7/indices/%s", Api, r.Duration)
}

func (r *V7IndicesRequest) Sting() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7IndicesResponse 天气生活指数响应
type V7IndicesResponse struct {
	baseResponse
	Daily []struct {
		Date     string `json:"date,omitempty"`     // 预报日期
		Type     string `json:"type,omitempty"`     // 生活指数类型ID
		Name     string `json:"name,omitempty"`     // 生活指数类型的名称
		Level    string `json:"level,omitempty"`    // 生活指数预报等级
		Category string `json:"category,omitempty"` // 生活指数预报级别名称
		Text     string `json:"text,omitempty"`     // 生活指数预报的详细描述,可能为空
	}
}

func (r *V7IndicesResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7Indices(req *V7IndicesRequest) (*V7IndicesResponse, error) {
	resp := new(V7IndicesResponse)
	err := c.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
