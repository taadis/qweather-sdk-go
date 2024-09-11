package qweather

import (
	"log/slog"
	"os"
	"testing"
)

const QWEATHER_KEY string = "QWEATHER_KEY"

func TestTopCity(t *testing.T) {
	t.Run("V2TopCity", func(t *testing.T) {
		req := NewV2TopCityRequest()
		req.Key = os.Getenv(QWEATHER_KEY)
		client := NewClient()
		resp, err := client.V2TopCity(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("V2TopCity resp:%s", resp.String())
		// {"code":"401","refer":{}}
	})

	t.Run("V2TopCity-with-enbaled-log", func(t *testing.T) {
		req := NewV2TopCityRequest()
		req.Key = os.Getenv(QWEATHER_KEY)
		client := NewClient(WithLogEnabled(true))
		resp, err := client.V2TopCity(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("V2TopCity resp:%s", resp.String())
		// {"code":"401","refer":{}}
	})

	t.Run("V2TopCity-with-logger", func(t *testing.T) {
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))

		req := NewV2TopCityRequest()
		req.Key = os.Getenv(QWEATHER_KEY)
		client := NewClient(WithLogger(logger), WithLogEnabled(true))
		resp, err := client.V2TopCity(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("V2TopCity resp:%s", resp.String())
		// {"code":"401","refer":{}}
	})
}
