package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	CLOUDFLARE_ACCOUNT_ID        string
	CLOUDFLARE_ACCESS_KEY_ID     string
	CLOUDFLARE_ACCESS_KEY_SECRET string
	DATABASE_URL                 string
	UPLOAD_BUCKET_NAME           string
}

func NewConfig() *Config {
	return &Config{
		DATABASE_URL: "postgres://postgres:password@localhost:5432/vgpx?sslmode=disable",
	}
}

func init() {
	viper.AutomaticEnv()
}
