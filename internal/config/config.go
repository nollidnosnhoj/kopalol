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
		CLOUDFLARE_ACCOUNT_ID:        viper.GetString("CLOUDFLARE_ACCOUNT_ID"),
		CLOUDFLARE_ACCESS_KEY_ID:     viper.GetString("CLOUDFLARE_ACCESS_KEY_ID"),
		CLOUDFLARE_ACCESS_KEY_SECRET: viper.GetString("CLOUDFLARE_ACCESS_KEY_SECRET"),
		UPLOAD_BUCKET_NAME:           "simplimguploads",
		DATABASE_URL:                 "postgres://postgres:password@localhost:5432/vgpx?sslmode=disable",
	}
}

func init() {
	viper.AutomaticEnv()
}
