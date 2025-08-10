package core

import (
	"context"
)

type Config struct {
	ConnectionString string `json:"connectionString"`
}

type ConfigServer interface {
	GetConfiguration() (*Config, error)
	WatchConfig(context.Context, *Config)
}
