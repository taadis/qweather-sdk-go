package qweather

import (
	"os"
	"testing"
)

func TestClient_V7WeatherNow(t *testing.T) {
	req := NewV7WeatherNowRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "120.13026,30.25961"
	req.IsDev = true
	client := NewClient()
	resp, err := client.V7WeatherNow(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("V7WeatherNow resp:%s", resp.String())
}
