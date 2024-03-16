package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	S3_ENDPOINT         string
	S3_ACCESS_KEY       string
	S3_SECRET_KEY       string
	S3_IMAGE_URL        string
	S3_FORCE_PATH_STYLE bool
	DATABASE_URL        string
	DATABASE_AUTH_TOKEN string
	UPLOAD_BUCKET_NAME  string
}

func NewConfig() *Config {
	return &Config{
		S3_ENDPOINT:         viper.GetString("S3_ENDPOINT"),
		S3_ACCESS_KEY:       viper.GetString("S3_ACCESS_KEY"),
		S3_SECRET_KEY:       viper.GetString("S3_SECRET_KEY"),
		S3_IMAGE_URL:        viper.GetString("S3_IMAGE_URL"),
		S3_FORCE_PATH_STYLE: viper.GetBool("S3_FORCE_PATH_STYLE"),
		UPLOAD_BUCKET_NAME:  viper.GetString("UPLOAD_BUCKET"),
		DATABASE_AUTH_TOKEN: viper.GetString("DATABASE_AUTH_TOKEN"),
		DATABASE_URL:        viper.GetString("DATABASE_URL"),
	}
}

func init() {
	viper.AutomaticEnv()
}
