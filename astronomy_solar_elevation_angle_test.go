package qweather

import "testing"

func TestClient_V7AstronomySolarElevationAngle(t *testing.T) {
	req := NewV7AstronomySolarElevationAngleRequest()
	req.Location = "116.41,39,92"
	req.Date = "20170809"
	req.Time = "1230"
	req.Tz = "0800"
	req.Alt = "43"
	client := NewClient()
	resp, err := client.V7AstronomySolarElevationAngle(req)
	if err != nil {
		t.Fatalf("client.V7AstronomySolarElevationAngle error:%+v", err)
	}

	t.Logf("client.V7AstronomySolarElevationAngle resp:%s", resp.String())
}
