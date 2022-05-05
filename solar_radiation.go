package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7SolarRadiationRequest 太阳辐射逐小时预报请求
type V7SolarRadiationRequest struct {
	baseRequest
	Location string `json:"location"` // 需要查询地区的经纬度坐标
	Hours    string `json:"hours"`    // 小时数,24h/72h
}

func NewV7SolarRadiationRequest() *V7SolarRadiationRequest {
	r := new(V7SolarRadiationRequest)
	return r
}

func (r *V7SolarRadiationRequest) Method() string {
	return http.MethodGet
}

func (r *V7SolarRadiationRequest) Url() string {
	return fmt.Sprintf("%s/v7/solar-radiation/%s", Api, r.Hours)
}

func (r *V7SolarRadiationRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7SolarRadiationResponse 太阳辐射逐小时预报响应
type V7SolarRadiationResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"`
}

func (r *V7SolarRadiationResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7SolarRadiation(req *V7SolarRadiationRequest) (*V7SolarRadiationResponse, error) {
	resp := new(V7SolarRadiationResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
