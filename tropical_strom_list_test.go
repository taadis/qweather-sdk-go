package qweather

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestClient_V7TropicalStormList(t *testing.T) {
	req := NewV7TropicalStormListRequest()
	req.Key = os.Getenv(QWEATHER_KEY)
	req.Basin = "NP"
	req.Year = fmt.Sprint(time.Now().Year() - 1)
	client := NewClient()
	resp, err := client.V7TropicalStormList(req)
	if err != nil {
		t.Fatalf("client.V7TropicalStormList error:%+v", err)
	}

	t.Logf("client.V7TropicalStormList resp:%s", resp.String())
}
