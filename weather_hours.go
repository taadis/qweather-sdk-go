package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7WeatherHoursRequest 逐小时天气预报请求
type V7WeatherHoursRequest struct {
	baseRequest
	Location string `json:"location,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Duration string `json:"duration,omitempty"`
}

func NewV7WeatherHoursRequest() *V7WeatherHoursRequest {
	return new(V7WeatherHoursRequest)
}

func (r *V7WeatherHoursRequest) Method() string {
	return http.MethodGet
}

func (r *V7WeatherHoursRequest) Url() string {
	if r.IsDev {
		return fmt.Sprintf("%s/v7/weather/%s", DevApi, r.Duration)
	}
	return fmt.Sprintf("%s/v7/weather/%s", Api, r.Duration)
}

func (r *V7WeatherHoursRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7WeatherHoursResponse 逐小时天气预报响应
type V7WeatherHoursResponse struct {
	baseResponse
	Hourly []struct {
		FxTime    string `json:"fxTime,omitempty"`    // 预报时间
		Temp      string `json:"temp,omitempty"`      // 温度,默认单位:摄氏度
		Icon      string `json:"icon,omitempty"`      // 天气状况和图标的代码,图标可通过[天气状况和图标]()下载
		Text      string `json:"text,omitempty"`      // 天气状况的文字描述,包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360,omitempty"`   // 风向360角度
		WindDir   string `json:"windDir,omitempty"`   // 风向
		WindScale string `json:"windScale,omitempty"` // 风力等级
		WindSpeed string `json:"windSpeed,omitempty"` // 风速,默认单位:公里/小时
		Humidity  string `json:"humidity,omitempty"`  // 相对湿度,百分比数值
		Precip    string `json:"precip,omitempty"`    // 当前小时累计降水量,默认单位:毫米
		Pop       string `json:"pop,omitempty"`       // 逐小时预报降水概率,百分比数值,可能为空
		Pressure  string `json:"pressure,omitempty"`  // 大气压强,默认单位:百帕
		Dew       string `json:"dew,omitempty"`       // 露点温度
	} `json:"hourly,omitempty"`
}

func (r *V7WeatherHoursResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7WeatherHours(req *V7WeatherHoursRequest) (*V7WeatherHoursResponse, error) {
	resp := new(V7WeatherHoursResponse)
	err := c.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
