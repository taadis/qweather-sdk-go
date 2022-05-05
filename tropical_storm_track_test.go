package qweather

import (
	"os"
	"testing"
)

func TestClient_V7TropicalStormTrack(t *testing.T) {
	req := NewV7TropicalStormTrackRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.StormId = "NP2018"
	client := NewClient()
	resp, err := client.V7TropicalStormTrack(req)
	if err != nil {
		t.Fatalf("client.V7TropicalStormTrack error:%+v", err)
	}

	t.Logf("client.V7TropicalStormTrack resp:%s", resp.String())
}
