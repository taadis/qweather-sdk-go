package qweather

import (
	"os"
	"testing"
)

func TestClient_V7HistoricalWeather(t *testing.T) {
	req := NewV7HistoricalWeatherRequest()
	req.Key = os.Getenv("QWEATHER_KEY")
	req.Location = "101010100"
	req.Date = "20200531"
	req.Lang = "zh"
	req.Unit = "m"
	client := NewClient()
	resp, err := client.V7HistoricalWeather(req)
	if err != nil {
		t.Fatalf("client.V7HistoricalWeather error:%+v", err)
	}

	t.Log("code:", resp.Code)
	t.Log("fxLink:", resp.FxLink)
	weatherDaily := resp.WeatherDaily
	t.Log("weatherDaily.date:", weatherDaily.Date)
	t.Log("weatherDaily.sunrise:", weatherDaily.Sunrise)
	t.Log("weatherDaily.sunset:", weatherDaily.Sunset)
	t.Log("weatherDaily.moonrise:", weatherDaily.Moonrise)
	t.Log("weatherDaily.moonset:", weatherDaily.Moonset)
	t.Log("weatherDaily.moonPhase:", weatherDaily.MoonPhase)
	t.Log("weatherDaily.tempMax:", weatherDaily.TempMax)
	t.Log("weatherDaily.tempMin", weatherDaily.TempMin)
	t.Log("weatherDaily.precip", weatherDaily.Precip)
	weatherHourly := resp.WeatherHourly[0]
	t.Log("weatherHourly.time:", weatherHourly.Time)
	t.Log("weatherHourly.temp:", weatherHourly.Temp)
	t.Log("weatherHourly.icon:", weatherHourly.Icon)
	t.Log("weatherHourly.text:", weatherHourly.Text)
	t.Log("weatherHourly.wind360:", weatherHourly.Wind360)
	t.Log("weatherHourly.windDir:", weatherHourly.WindDir)
	t.Log("weatherHourly.windScale:", weatherHourly.WindScale)
	t.Log("weatherHourly.windSpeed:", weatherHourly.WindSpeed)
	t.Log("weatherHourly.humidity:", weatherHourly.Humidity)
	t.Log("weatherHourly.precip:", weatherHourly.Precip)
	t.Log("weatherHourly.pressure:", weatherHourly.Pressure)
	refer := resp.Refer
	t.Log("refer.sources:", refer.Sources)
	t.Log("refer.license:", refer.License)
}
