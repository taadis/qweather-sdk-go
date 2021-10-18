package qweather

import (
	"os"
	"testing"
)

func TestClient_V7Indices(t *testing.T) {
	req := NewV7IndicesRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "120.13026,30.25961"
	req.Duration = "1d"
	req.IsDev = true
	req.Type = "3,5"
	client := NewClient()
	resp, err := client.V7Indices(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("V7Indices resp:%s", resp.String())
}
