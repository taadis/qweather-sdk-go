package qweather

import (
	"os"
	"testing"
)

func TestClient_V7OceanTide(t *testing.T) {
	req := NewV7OceanTideRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "P2951"
	req.Date = "20200531"
	client := NewClient()
	resp, err := client.V7OceanTide(req)
	if err != nil {
		t.Fatalf("client.V7OceanTide error:%+v", err)
	}

	t.Logf("client.V7OceanTide resp:%s", resp.String())
}
