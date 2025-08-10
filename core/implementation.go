package core

import (
	"context"
)

type Config struct {
	ConnectionString string `json:"connectionString"`
}

type ConfigServer interface {
	GetConfiguration() (*Config, error)
	WatchConfig(context.Context, *Config, chan string)
}

type Database interface {
	GetUser(context.Context, int) (*User, error)
	WatchConnectionString() chan string
}
