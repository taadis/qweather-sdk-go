package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7WeatherDaysRequest 逐天天气预报请求
type V7WeatherDaysRequest struct {
	baseRequest
	Location string `json:"location,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Duration string `json:"duration,omitempty"`
}

func NewV7WeatherDaysRequest() *V7WeatherDaysRequest {
	return new(V7WeatherDaysRequest)
}

func (r *V7WeatherDaysRequest) Method() string {
	return http.MethodGet
}

func (r *V7WeatherDaysRequest) Url() string {
	if r.IsDev {
		return fmt.Sprintf("%s/v7/weather/%s", DevApi, r.Duration)
	}
	return fmt.Sprintf("%s/v7/weather/%s", Api, r.Duration)
}

func (r *V7WeatherDaysRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7WeatherDaysResponse 逐天天气预报响应
type V7WeatherDaysResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime,omitempty"`
	FxLink     string `json:"fxLink,omitempty"`
	Daily      []struct {
		FxDate         string `json:"fxDate,omitempty"`         // 预报日期
		Sunrise        string `json:"sunrise,omitempty"`        // 日出时间
		Sunset         string `json:"sunset,omitempty"`         // 日落时间
		Moonrise       string `json:"moonrise,omitempty"`       // 月升时间
		Moonset        string `json:"moonset,omitempty"`        // 月落时间
		MoonPhase      string `json:"moonPhase,omitempty"`      // 月相名称
		TempMax        string `json:"tempMax,omitempty"`        // 预报当天最高温度
		TempMin        string `json:"tempMin,omitempty"`        // 预报当天最低温度
		IconDay        string `json:"iconDay,omitempty"`        // 预报白天天气状况的图标代码,图标可通过[天气状况和图标]()下载
		TextDay        string `json:"textDay,omitempty"`        // 预报白天天气状况文字描述,包括阴晴雨雪天气状态的描述
		IconNight      string `json:"iconNight,omitempty"`      // 预报夜间天气状况的图标代码,图标可通过[天气状况和图标]()下载
		TextNight      string `json:"textNight,omitempty"`      // 预报夜间天气状况文字描述,包括阴晴雨雪等天气状态的描述
		Wind360Day     string `json:"wind360Day,omitempty"`     // 预报白天风向360角度
		WindDirDay     string `json:"windDirDay,omitempty"`     // 预报白天风向
		WindScaleDay   string `json:"windScaleDay,omitempty"`   // 预报白天风力等级
		WindSpeedDay   string `json:"windSpeedDay,omitempty"`   // 预报白天风速,单位:公里/小时
		Wind360Night   string `json:"wind360Night,omitempty"`   // 预报夜间风向360角度
		WindDirNight   string `json:"windDirNight,omitempty"`   // 预报夜间当前风向
		WindScaleNight string `json:"windScaleNight,omitempty"` // 预报夜间风力等级
		WindSpeedNight string `json:"windSpeedNight,omitempty"` // 预报夜间风速,单位:公里/小时
		Precip         string `json:"precip,omitempty"`         // 预报当天总降水量,默认单位:毫米
		UvIndex        string `json:"uvIndex,omitempty"`        // 紫外线强度指数
		Humidity       string `json:"humidity,omitempty"`       // 相对湿度,百分比数值
		Pressure       string `json:"pressure,omitempty"`       // 大气压强,默认单位:百帕
		Vis            string `json:"vis,omitempty"`            // 能见度,默认单位:公里
		Cloud          string `json:"cloud,omitempty"`          // 云量,百分比数值
	} `json:"daily,omitempty"`
}

func (r *V7WeatherDaysResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7WeatherDays(req *V7WeatherDaysRequest) (*V7WeatherDaysResponse, error) {
	resp := new(V7WeatherDaysResponse)
	err := c.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
