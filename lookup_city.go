package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V2LookupCityRequest 城市信息查询请求
type V2LookupCityRequest struct {
	baseRequest
	Location string `json:"location,omitempty"` // 需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制）、LocationID或Adcode（仅限中国城市）
	Adm      string `json:"adm,omitempty"`      // 城市的上级行政区划，默认不限定行政区划。
	Range    string `json:"range,omitempty"`    // 搜索范围，可设定只在某个国家范围内进行搜索，国家名称需使用ISO 3166 所定义的国家代码。
	Number   int    `json:"number,omitempty"`   // 返回结果的数量，取值范围1-20，默认返回10个结果
	Lang     string `json:"lang,omitempty"`     // 多语言设置，默认中文，当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
}

func NewV2LookupCityRequest() *V2LookupCityRequest {
	return new(V2LookupCityRequest)
}

func (*V2LookupCityRequest) Method() string {
	return http.MethodGet
}

func (*V2LookupCityRequest) Url() string {
	return fmt.Sprintf("%s/v2/city/lookup", GeoApi)
}

func (r *V2LookupCityRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V2LookupCityResponse 城市信息查询响应
type V2LookupCityResponse struct {
	baseResponse
	Location []struct {
		Name      string `json:"name"`      // 地区/城市名称
		Id        string `json:"id"`        // 地区/城市ID
		Lat       string `json:"lat"`       // 地区/城市纬度
		Lon       string `json:"lon"`       // 地区/城市经度
		Adm2      string `json:"adm2"`      // 地区/城市的上级行政区划名称
		Adm1      string `json:"adm1"`      // 地区/城市所属一级行政区域
		Country   string `json:"country"`   // 地区/城市所属国家名称
		Tz        string `json:"tz"`        // 地区/城市所在时区
		UtcOffset string `json:"utcOffset"` // 地区/城市目前与UTC时间偏移的小时数
		IsDst     string `json:"isDst"`     // 地区/城市是否当前处于夏令时,0表示当前不是夏令时,1表示当前处于夏令时
		Type      string `json:"type"`      // 地区/城市的属性
		Rank      string `json:"rank"`      // 地区评分
		FxLink    string `json:"fxLink"`    // 该地区的天气预报网页链接，便于嵌入你的网站或应用
	} `json:"location"`
}

func (r *V2LookupCityResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V2LookupCity(req *V2LookupCityRequest) (*V2LookupCityResponse, error) {
	resp := new(V2LookupCityResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
