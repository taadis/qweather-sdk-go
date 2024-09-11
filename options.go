package qweather

import (
	"log/slog"
)

type Option func(*Client)

func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithLogEnabled(logEnabled bool) Option {
	return func(c *Client) {
		c.logEnabled = logEnabled
	}
}
