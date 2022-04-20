package qweather

import (
	"os"
	"testing"
)

func TestClient_V7Air5Days(t *testing.T) {
	req := NewV7Air5DaysRequest()
	req.Key = os.Getenv("QWETHER_KEY")
	req.Location = "116.41,39.92"
	req.Lang = "zh"
	client := NewClient()
	resp, err := client.V7Air5Days(req)
	if err != nil {
		t.Fatalf("client.V7Air5Days error:%+v,", err)
	}

	t.Logf("code:%s", resp.Code)
	t.Logf("updateTime:%s", resp.UpdateTime)
	daily := resp.Daily[0]
	t.Logf("daily.fxDate:%s", daily.FxDate)
	t.Logf("daily.aqi:%s", daily.AQI)
	t.Logf("daily.level:%s", daily.Level)
	t.Logf("daily.category:%s", daily.Category)
	t.Logf("daily.primary:%s", daily.Primary)
	refer := resp.Refer
	t.Logf("refer.sources:%s", refer.Sources)
	t.Logf("refer.license:%s", refer.License)
	t.Logf("client.V7Air5Days resp:%s", resp.String())
}
