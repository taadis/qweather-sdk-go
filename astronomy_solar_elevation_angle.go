package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7AstronomySolarElevationAngleRequest 太阳高度角请求
type V7AstronomySolarElevationAngleRequest struct {
	baseRequest
	Location string `json:"location"` // 需要查询地区的以英文逗号分隔的经纬度坐标,例如:location=116.41,39.92
	Date     string `json:"date"`     // 查询日期,格式为yyyyMMdd,例如date=20170809
	Time     string `json:"time"`     // 查询时间,格式为HHmm,24小时制,例如time=1230
	Tz       string `json:"tz"`       // 查询地区所在的时区,例如tz=0800或tz=0530
	Alt      string `json:"alt"`      // 海拔高度,单位为米,例如alt=43
}

func NewV7AstronomySolarElevationAngleRequest() *V7AstronomySolarElevationAngleRequest {
	r := new(V7AstronomySolarElevationAngleRequest)
	return r
}

func (r *V7AstronomySolarElevationAngleRequest) Method() string {
	return http.MethodGet
}

func (r *V7AstronomySolarElevationAngleRequest) Url() string {
	return fmt.Sprintf("%s/v7/astronomy/solar-elevation-angle", Api)
}

func (r *V7AstronomySolarElevationAngleRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7AstronomySolarElevationAngleResponse 太阳高度角响应
type V7AstronomySolarElevationAngleResponse struct {
	baseResponse
	SolarElevationAngle string `json:"solarElevationAngle"` // 太阳高度角
	SolarAzimuthAngle   string `json:"solarAzimuthAngle"`   // 太阳方位角,正北顺时针方向角度
	SolarHour           string `json:"solarHour"`           // 太阳时,HHmm格式
	HourAngle           string `json:"hourAngle"`           // 时角
}

func (r *V7AstronomySolarElevationAngleResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7AstronomySolarElevationAngle(req *V7AstronomySolarElevationAngleRequest) (*V7AstronomySolarElevationAngleResponse, error) {
	resp := new(V7AstronomySolarElevationAngleResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
