package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string
}

func NewConfig() *Config {
	return &Config{
		DATABASE_URL: "postgres://postgres:password@localhost:5432/vgpx?sslmode=disable",
	}
}

func init() {
	viper.AutomaticEnv()
}
