package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7WeatherNowRequest 实时天气请求
type V7WeatherNowRequest struct {
	baseRequest
	Location string `json:"location,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Gzip     string `json:"gzip,omitempty"`
}

func NewV7WeatherNowRequest() *V7WeatherNowRequest {
	return new(V7WeatherNowRequest)
}

func (r *V7WeatherNowRequest) Method() string {
	return http.MethodGet
}

func (r *V7WeatherNowRequest) Url() string {
	if r.IsDev {
		return fmt.Sprintf("%s/v7/weather/now", DevApi)
	}
	return fmt.Sprintf("%s/v7/weather/now", Api)
}

func (r *V7WeatherNowRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7WeatherNowResponse 实时天气响应
type V7WeatherNowResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime,omitempty"`
	FxLink     string `json:"fxLink,omitempty"`
	Now        struct {
		ObsTime   string `json:"obsTime,omitempty"`
		Temp      string `json:"temp,omitempty"`
		FeelsLike string `json:"feelsLike,omitempty"`
		Icon      string `json:"icon,omitempty"`
		Text      string `json:"text,omitempty"`
		Wind360   string `json:"wind360,omitempty"`
		WindDir   string `json:"windDir,omitempty"`
		WindScale string `json:"windScale,omitempty"`
		WindSpeed string `json:"windSpeed,omitempty"`
		Humidity  string `json:"humidity,omitempty"`
		Precip    string `json:"precip,omitempty"`
		Pressure  string `json:"pressure,omitempty"`
		Vis       string `json:"vis,omitempty"`
		Cloud     string `json:"cloud,omitempty"`
		Dew       string `json:"dew,omitempty"`
	} `json:"now,omitempty"`
}

func (r *V7WeatherNowResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7WeatherNow(req *V7WeatherNowRequest) (*V7WeatherNowResponse, error) {
	resp := new(V7WeatherNowResponse)
	err := c.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
