package qweather

import (
	"os"
	"testing"
)

const QWEATHER_KEY string = "QWEATHER_KEY"

func TestTopCity(t *testing.T) {
	req := NewV2TopCityRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	client := NewClient()
	resp, err := client.V2TopCity(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("V2TopCity resp:%s", resp.String())
}
