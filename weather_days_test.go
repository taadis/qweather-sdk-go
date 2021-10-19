package qweather

import (
	"os"
	"testing"
)

func TestClient_V7WeatherDays(t *testing.T) {
	req := NewV7WeatherDaysRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "120.13026,30.25961"
	req.Duration = "7d"
	req.IsDev = true
	client := NewClient()
	resp, err := client.V7WeatherDays(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("V7WeatherDays resp:%s", resp.String())
}
