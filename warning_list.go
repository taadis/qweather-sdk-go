package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7WarningListRequest 天气预警城市列表请求
type V7WarningListRequest struct {
	baseRequest
	Range string `json:"range"` // 选择指定的国家或地区,使用ISO 3166格式化,例如:range=cn,range=fr
}

func NewV7WarningListRequest() *V7WarningListRequest {
	r := new(V7WarningListRequest)
	return r
}

func (r *V7WarningListRequest) Method() string {
	return http.MethodGet
}

func (r *V7WarningListRequest) Url() string {
	return fmt.Sprintf("%s/v7/warning/list", Api)
}

func (r *V7WarningListRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7WarningListResponse 天气预警城市列表响应
type V7WarningListResponse struct {
	baseResponse
	UpdateTime      string `json:"updateTime"` // 当前API的最近更新时间
	WarningLoclList []struct {
		LocationId string `json:"LocationId"` // 当前预警的LocationID
	} `json:"warningLoclList"`
}

func (r *V7WarningListResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7WarningList(req *V7WarningListRequest) (*V7WarningListResponse, error) {
	resp := new(V7WarningListResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
