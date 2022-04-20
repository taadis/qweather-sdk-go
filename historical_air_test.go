package qweather

import (
	"os"
	"testing"
)

func TestClient_V7HistoricalAir(t *testing.T) {
	req := NewV7HistoricalAirRequest()
	req.Key = os.Getenv("QWEATHER_KEY")
	req.Location = "101010100"
	req.Date = "20200531"
	req.Lang = "zh"
	client := NewClient()
	resp, err := client.V7HistoricalAir(req)
	if err != nil {
		t.Fatalf("client.V7HistoricalAir error:%+v", err)
	}

	airHourly := resp.AirHourly[0]
	refer := resp.Refer
	t.Log("code:", resp.Code)
	t.Log("fxLink:", resp.FxLink)
	t.Log("airHourly.pubTime:", airHourly.PubTime)
	t.Log("airHourly.aqi:", airHourly.AQI)
	t.Log("airHourly.level:", airHourly.Level)
	t.Log("airHourly.category:", airHourly.Category)
	t.Log("airHourly.primary:", airHourly.Primary)
	t.Log("airHourly.pm10:", airHourly.PM10)
	t.Log("airHourly.pm2p5:", airHourly.PM2p5)
	t.Log("airHourly.no2:", airHourly.NO2)
	t.Log("airHourly.so2:", airHourly.SO2)
	t.Log("airHourly.co:", airHourly.CO)
	t.Log("airHourly.o3:", airHourly.O3)
	t.Log("refer.sources:", refer.Sources)
	t.Log("refer.license:", refer.License)
}
