package qweather

import (
	"os"
	"testing"
)

func TestClient_V7WeatherHours(t *testing.T) {
	req := NewV7WeatherHoursRequest()
	req.IsDev = true
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "120.13026,30.25961"
	req.Duration = "24h"
	client := NewClient()
	resp, err := client.V7WeatherHours(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("V7WeatherHours resp:%s", resp.String())
}
