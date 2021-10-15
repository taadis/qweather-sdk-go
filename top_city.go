package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V2TopCityRequest 热门城市查询请求
type V2TopCityRequest struct {
	baseRequest
	Range  string `json:"range,omitempty"`
	Number int    `json:"number,omitempty"`
	Lang   string `json:"lang,omitempty"`
}

func NewV2TopCityRequest() *V2TopCityRequest {
	return new(V2TopCityRequest)
}

func (*V2TopCityRequest) Method() string {
	return http.MethodGet
}

func (*V2TopCityRequest) Url() string {
	return fmt.Sprintf("%s/v2/city/top", GeoApi)
}

func (r *V2TopCityRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V2CityTopResponse 热门城市查询响应
type V2CityTopResponse struct {
	baseResponse
	TopCityList []struct {
		Name      string `json:"name,omitempty"`
		Id        string `json:"id,omitempty"`
		Lat       string `json:"lat,omitempty"`
		Lon       string `json:"lon,omitempty"`
		Adm2      string `json:"adm2,omitempty"`
		Adm1      string `json:"adm1,omitempty"`
		Country   string `json:"country,omitempty"`
		UtcOffset string `json:"utcOffset,omitempty"`
		IsDst     string `json:"isDst,omitempty"`
		Type      string `json:"type,omitempty"`
		Rank      string `json:"rank,omitempty"`
		FxLink    string `json:"fxLink,omitempty"`
	} `json:"topCityList,omitempty"`
}

func (r *V2CityTopResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V2TopCity(req *V2TopCityRequest) (*V2CityTopResponse, error) {
	resp := new(V2CityTopResponse)
	err := c.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
