package qweather

import (
	"os"
	"testing"
)

func TestClient_V7SolarRadiation(t *testing.T) {
	req := NewV7SolarRadiationRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "116.41,39.92"
	client := NewClient()
	resp, err := client.V7SolarRadiation(req)
	if err != nil {
		t.Fatalf("client.V7SolarRadiation error:%+v", err)
	}

	t.Logf("client.V7SolarRadiation resp:%s", resp.String())
}
