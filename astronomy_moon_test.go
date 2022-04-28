package qweather

import (
	"os"
	"testing"
)

func TestClient_V7AstronomyMoon(t *testing.T) {
	req := NewV7AstronomyMoonRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "101010100"
	req.Date = "2020531"
	req.Lang = "zh"
	client := NewClient()
	resp, err := client.V7AstronomyMoon(req)
	if err != nil {
		t.Fatalf("client.V7AstronomyMoon error:%+v", err)
	}

	t.Logf("client.V7AstronomyMoon resp:%s", resp.String())
}
