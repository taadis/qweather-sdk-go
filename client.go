package qweather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"log/slog"
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
	logger     *slog.Logger
	logEnabled bool

	respBodyRaw string
}

func NewClient(options ...Option) *Client {
	c := &Client{
		logger:     slog.Default(),
		logEnabled: false,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *Client) GetResponseBodyRaw() string {
	return c.respBodyRaw
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
	if err != nil {
		return err
	}
	u.RawQuery = ps.Encode()

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err
	}

	if c.logEnabled {
		c.logger.Info("starting request", "method", request.Method(), "url", request.Url())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 使用 `io.TeeReader` 同时读取响应体并将其复制到缓冲区,规避2次解析resp.Body
	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// response json to struct
	err = json.NewDecoder(tee).Decode(response)
	if err != nil {
		return err
	}

	// response raw text
	c.respBodyRaw = buf.String()

	if c.logEnabled {
		c.logger.Info("response", "body raw", c.respBodyRaw)
	}

	return nil
}
