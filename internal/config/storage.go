package config

import (
	"context"

	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/spf13/viper"
)

func NewStorageWithConfig(context context.Context) (*storage.S3Storage, error) {
	settings := &storage.S3StorageSettings{
		Endpoint:       viper.GetString("S3_ENDPOINT"),
		AccessKey:      viper.GetString("S3_ACCESS_KEY"),
		SecretKey:      viper.GetString("S3_SECRET_KEY"),
		ForcePathStyle: viper.GetBool("S3_FORCE_PATH_STYLE"),
		Bucket:         viper.GetString("UPLOAD_BUCKET"),
		ImageUrl:       viper.GetString("S3_IMAGE_URL"),
	}
	return storage.NewS3Storage(context, settings)
}
