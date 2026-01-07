package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	HTTPPort    string `env:"HTTP_PORT" default:"8080"`
}

func LoadEnv() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL: cfg.DatabaseURL,
		HTTPPort:    cfg.HTTPPort,
	}, nil
}
