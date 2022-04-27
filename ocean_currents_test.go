package qweather

import (
	"os"
	"testing"
)

func TestClient_V7OceanCurrents(t *testing.T) {
	req := NewV7OceanCurrentsRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Location = "P66981"
	req.Date = "20200531"
	client := NewClient()
	resp, err := client.V7OceanCurrents(req)
	if err != nil {
		t.Fatalf("client.V7OceanCurrents error:%+v", err)
	}

	t.Logf("client.V7OceanCurrents resp:%s", resp.String())
}
