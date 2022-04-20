package qweather

import (
	"os"
	"testing"
)

func TestClient_V7AirNow(t *testing.T) {
	req := NewV7AirNowRequest()
	req.Key = os.Getenv("QWETHER_KEY")
	req.Location = ""
	req.Lang = ""
	client := NewClient()
	resp, err := client.V7AirNow(req)
	if err != nil {
		t.Fatalf("client.V7AirNow error:%+v", err)
	}

	t.Logf("client.V7AirNow resp:%s", resp.String())
}
