package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const Api = "https://api.qweather.com"
const DevApi = "https://devapi.qweather.com"
const GeoApi = "https://geoapi.qweather.com"
const DatasetApi = "https://datasetapi.qweather.com"

type baseRequest struct {
	Key   string `json:"key,omitempty"`
	IsDev bool   `json:"isDev,omitempty"` // 是否为开发环境,默认为线上
}

type baseResponse struct {
	Code  string `json:"code,omitempty"` // 状态码
	Refer struct {
		Sources []string `json:"sources,omitempty"` // 原始数据来源，或数据源说明，可能为空
		License []string `json:"license,omitempty"` // 数据许可或版权声明，可能为空
	} `json:"refer,omitempty"` //
}

type Requester interface {
	Method() string
	Url() string
}

type Responder interface {
}

type Client struct {
}

func NewClient() *Client {
	return new(Client)
}

func (c *Client) Do(request Requester, response Responder) error {
	bs, err := json.Marshal(request)
	if err != nil {
		return err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(bs, &m)
	if err != nil {
		return err
	}

	ps := url.Values{}
	for k, v := range m {
		ps.Set(k, fmt.Sprintf("%s", v))
	}

	method := request.Method()
	u, err := url.Parse(request.Url())
	u.RawQuery = ps.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
