package qweather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// V7TropicalStormTrackRequest 台风实况和路径请求
type V7TropicalStormTrackRequest struct {
	baseRequest
	StormId string `json:"stormId"` // 需要查询的台风ID
}

func NewV7TropicalStormTrackRequest() *V7TropicalStormTrackRequest {
	r := new(V7TropicalStormTrackRequest)
	return r
}

func (r *V7TropicalStormTrackRequest) Method() string {
	return http.MethodGet
}

func (r *V7TropicalStormTrackRequest) Url() string {
	return fmt.Sprintf("%s/v7/tropical/storm-track", Api)
}

func (r *V7TropicalStormTrackRequest) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// V7TropicalStormTrackResponse 台风实况和路径响应
type V7TropicalStormTrackResponse struct {
	baseResponse
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面,便于嵌入网站或应用
	IsActive   string `json:"isActive"`   // 是否为活跃台风,0=停编,2=活跃台风
	Now        struct {
		PubTime      string `json:"pubTime"`   // 台风信息发布时间
		Lat          string `json:"lat"`       // 台风所处纬度
		Lon          string `json:"lon"`       // 台风所处经度
		Type         string `json:"type"`      // 台风类型
		Pressure     string `json:"pressure"`  // 台风中心气压
		WindSpeed    string `json:"windSpeed"` // 台风附近最大风速
		MoveSpeed    string `json:"moveSpeed"` // 台风移动速度
		MoveDir      string `json:"moveDir"`   // 台风移动方位
		Move360      string `json:"move360"`   // 台风移动方位360度方向
		WindRadius30 struct {
			NeRadius string `json:"neRadius"` // 台风7级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风7级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风7级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风7级风圈西北半径,可能为空
		} `json:"windRadius30"`
		WindRadius50 struct {
			NeRadius string `json:"neRadius"` // 台风10级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风10级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风10级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风10级风圈西北半径,可能为空
		} `json:"WindRadius50"`
		WindRadius64 struct {
			NeRadius string `json:"neRadius"` // 台风12级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风12级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风12级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风12级风圈西北半径,可能为空
		} `json:"WindRadius50"`
	} `json:"now"`
	Track []struct {
		Time         string `json:"time"`      // 当前台风信息发布时间
		Lat          string `json:"lat"`       // 台风所处纬度
		Lon          string `json:"lon"`       // 台风所处经度
		Type         string `json:"type"`      // 台风类型
		Pressure     string `json:"pressure"`  // 台风中心气压
		WindSpeed    string `json:"windSpeed"` // 台风附近最大风速
		MoveSpeed    string `json:"moveSpeed"` // 台风移动速度
		Move360      string `json:"move360"`   // 台风移动方位360度方向
		WindRadius30 struct {
			NeRadius string `json:"neRadius"` // 台风7级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风7级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风7级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风7级风圈西北半径,可能为空
		} `json:"windRadius30"`
		WindRadius50 struct {
			NeRadius string `json:"neRadius"` // 台风10级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风10级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风10级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风10级风圈西北半径,可能为空
		} `json:"WindRadius50"`
		WindRadius64 struct {
			NeRadius string `json:"neRadius"` // 台风12级风圈东北半径,可能为空
			SeRadius string `json:"seRadius"` // 台风12级风圈东南半径,可能为空
			SwRadius string `json:"swRadius"` // 台风12级风圈西南半径,可能为空
			NwRadius string `json:"nwRadius"` // 台风12级风圈西北半径,可能为空
		} `json:"WindRadius50"`
	} `json:"track"`
}

func (r *V7TropicalStormTrackResponse) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

func (c *Client) V7TropicalStormTrack(req *V7TropicalStormTrackRequest) (*V7TropicalStormTrackResponse, error) {
	resp := new(V7TropicalStormTrackResponse)
	err := c.Do(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
