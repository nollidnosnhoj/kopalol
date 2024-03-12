package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	CLOUDFLARE_ACCOUNT_ID        string
	CLOUDFLARE_ACCESS_KEY_ID     string
	CLOUDFLARE_ACCESS_KEY_SECRET string
	CLOUDFLARE_IMAGE_CACHE_URL   string
	DATABASE_URL                 string
	UPLOAD_BUCKET_NAME           string
}

func NewConfig() *Config {
	return &Config{
		CLOUDFLARE_ACCOUNT_ID:        viper.GetString("CLOUDFLARE_ACCOUNT_ID"),
		CLOUDFLARE_ACCESS_KEY_ID:     viper.GetString("CLOUDFLARE_ACCESS_KEY_ID"),
		CLOUDFLARE_ACCESS_KEY_SECRET: viper.GetString("CLOUDFLARE_ACCESS_KEY_SECRET"),
		CLOUDFLARE_IMAGE_CACHE_URL:   viper.GetString("CLOUDFLARE_IMAGE_CACHE_URL"),
		UPLOAD_BUCKET_NAME:           "simplimguploads",
		DATABASE_URL:                 viper.GetString("DATABASE_URL"),
	}
}

func init() {
	viper.AutomaticEnv()
}
