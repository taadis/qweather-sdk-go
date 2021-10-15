package qweather

import (
	"os"
	"testing"
)

func TestClient_V2LookupCity(t *testing.T) {
	req := NewV2LookupCityRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "西湖区"
	req.Adm = "浙江省"
	client := NewClient()
	resp, err := client.V2LookupCity(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("V2LookupCity resp:%s", resp.String())
}
