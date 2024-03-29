package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Configuration struct {
	DATABASE_URL        string `koanf:"DATABASE_URL"`
	DATABASE_AUTH_TOKEN string `koanf:"DATABASE_AUTH_TOKEN"`
	S3_ENDPOINT         string `koanf:"S3_ENDPOINT"`
	S3_ACCESS_KEY       string `koanf:"S3_ACCESS_KEY"`
	S3_SECRET_KEY       string `koanf:"S3_SECRET_KEY"`
	S3_FORCE_PATH_STYLE bool   `koanf:"S3_FORCE_PATH_STYLE"`
	S3_IMAGE_URL        string `koanf:"S3_IMAGE_URL"`
	S3_UPLOAD_BUCKET    string `koanf:"S3_UPLOAD_BUCKET"`
}

func NewConfiguration() (Configuration, error) {
	var c Configuration
	k := koanf.New(".")
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil); err != nil {
		return c, err
	}
	if err := k.Unmarshal("", &c); err != nil {
		return c, err
	}
	return c, nil
}
