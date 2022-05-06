package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7WarningNowRequest 天气灾害预警请求
type V7WarningNowRequest struct {
	baseRequest
	Location string `json:"location"`       // LocationID或经纬度坐标
	Lang     string `json:"lang,omitempty"` // 多语言设置,本数据仅支持中文和英文,zh=中文(默认),en=英文
}

func NewV7WarningNowRequest() *V7WarningNowRequest {
	r := new(V7WarningNowRequest)
	return r
}

func (r *V7WarningNowRequest) Method() string {
	return http.MethodGet
}

func (r *V7WarningNowRequest) Url() string {
	return fmt.Sprintf("%s/v7/warning/now", Api)
}

func (r *V7WarningNowRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7WarningNowResponse 天气灾害预警响应
type V7WarningNowResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	Warning    []struct {
		Id        string `json:"id"`        // 本条预警的唯一标识,可判断本条预警是否已经存在
		Sender    string `json:"sender"`    // 预警发布单位,可能为空
		PubTime   string `json:"pubTime"`   // 预警发布时间
		Title     string `json:"title"`     // 预警信息标题
		StartTime string `json:"startTime"` // 预警开始时间,可能为空
		EndTime   string `json:"endTime"`   // 预警结束时间,可能为空
		Status    string `json:"status"`    // 预警信息的发布状态
		Level     string `json:"level"`     // 预警等级
		Type      string `json:"type"`      // 预警类型ID
		TypeName  string `json:"typeName"`  // 预警类型名称
		Urgency   string `json:"urgency"`   // 预警信息的紧迫程度,可能为空
		Certainty string `json:"certainty"` // 预警信息的确定性,可能为空
		Text      string `json:"text"`      // 预警详细文字描述
		Related   string `json:"related"`   // 与本条预警相关联的预警ID,当预警状态为cancel或update时返回,可能为空
	} `json:"warning"`
}

func (r *V7WarningNowResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7WarningNow(req *V7WarningNowRequest) (*V7WarningNowResponse, error) {
	resp := new(V7WarningNowResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
