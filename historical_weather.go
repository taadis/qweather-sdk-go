package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type V7HistoricalWeatherRequest struct {
	baseRequest
	Location string `json:"location"`
	Date     string `json:"date"`
	Lang     string `json:"lang,omitempty"`
	Unit     string `json:"unit,omitempty"` // 度量衡单位参数:m=公制单位(默认);i=英制单位
}

func NewV7HistoricalWeatherRequest() *V7HistoricalWeatherRequest {
	r := new(V7HistoricalWeatherRequest)
	return r
}

func (r *V7HistoricalWeatherRequest) Method() string {
	return http.MethodGet
}

func (r *V7HistoricalWeatherRequest) Url() string {
	return fmt.Sprintf("%s/v7/historical/weather", DatasetApi)
}

func (r *V7HistoricalWeatherRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

type V7HistoricalWeatherResponse struct {
	baseResponse
	FxLink       string `json:"fxLink"` // 当前数据的响应式页面,便于嵌入网站或应用
	WeatherDaily struct {
		Date      string `json:"date"`      // 当前日期
		Sunrise   string `json:"sunrise"`   // 当天日出时间
		Sunset    string `json:"sunset"`    // 当天日落时间
		Moonrise  string `json:"moonrise"`  // 当天月升时间
		Moonset   string `json:"moonset"`   // 当天月落时间
		MoonPhase string `json:"moonPhase"` // 当天月相名称
		TempMax   string `json:"tempMax"`   // 当天最高温度
		TempMin   string `json:"tempMin"`   // 当前最低温度
		Precip    string `json:"precip"`    // 当天总降水量,默认单位:毫米
	} `json:"weatherDaily"`
	WeatherHourly []struct {
		Time      string `json:"time"`      // 当天时间
		Temp      string `json:"temp"`      // 当天每小时温度,默认单独:摄氏度
		Icon      string `json:"icon"`      // 当天每小时天气状况图标的代码,图标可通过天气状况和图标下载
		Text      string `json:"text"`      // 当天每小时天气状况的文字描述,包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 当天每小时风向360角度
		WindDir   string `json:"windDir"`   // 当天每小时风向
		WindScale string `json:"windScale"` // 当天每小时风力
		WindSpeed string `json:"windSpeed"` // 当天每小时风速,公里/小时
		Humidity  string `json:"humidity"`  // 当天每小时相对湿度,百分比数值
		Precip    string `json:"precip"`    // 当天每小时累计降水量,默认单位:毫米
		Pressure  string `json:"pressure"`  // 大气压强,默认单位:百帕
	} `json:"weatherHourly"`
}

func (r *V7HistoricalWeatherResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7HistoricalWeather(req *V7HistoricalWeatherRequest) (*V7HistoricalWeatherResponse, error) {
	resp := new(V7HistoricalWeatherResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
