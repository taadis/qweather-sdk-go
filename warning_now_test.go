package qweather

import (
	"os"
	"testing"
)

func TestClient_V7WarningNow(t *testing.T) {
	req := NewV7WarningNowRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "116.41,39.92"
	req.Lang = "zh"
	client := NewClient()
	resp, err := client.V7WarningNow(req)
	if err != nil {
		t.Fatalf("client.V7WarningNow error:%+v", err)
	}

	t.Logf("client.V7WarningNow resp:%s", resp.String())
}
