package qweather

import (
	"os"
	"testing"
)

func TestClient_V7AstronomySun(t *testing.T) {
	req := NewV7AstronomySunRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "101010100"
	req.Date = "20200531"
	req.Lang = "zh"
	client := NewClient()
	resp, err := client.V7AstronomySun(req)
	if err != nil {
		t.Fatalf("client.V7AstronomySun error:%+v", err)
	}

	t.Logf("client.V7AstronomySum resp:%s", resp.String())
}
