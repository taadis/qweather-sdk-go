package qweather

import (
	"os"
	"testing"
)

func TestClient_V7TropicalStormForecast(t *testing.T) {
	req := NewV7TropicalStormForecastRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.StormId = "NP2018"
	client := NewClient()
	resp, err := client.V7TropicalStormForecast(req)
	if err != nil {
		t.Fatalf("client.V7TropicalStormForecast error:%+v", err)
	}

	t.Logf("client.V7TropicalStormForecast resp:%s", resp.String())
}
