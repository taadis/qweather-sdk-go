package qweather

import (
	"os"
	"testing"
)

func TestClient_V7WarningList(t *testing.T) {
	req := NewV7WarningListRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Range = "cn"
	client := NewClient()
	resp, err := client.V7WarningList(req)
	if err != nil {
		t.Fatalf("client.V7WarningList error:%+v", err)
	}

	t.Logf("client.V7WarningList resp:%s", resp.String())
}
