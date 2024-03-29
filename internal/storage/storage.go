package storage

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/nollidnosnhoj/kopalol/internal/config"
)

type StorageType string

var S3StorageType StorageType = "s3"
var LocalStorageType StorageType = "local"

type ImageResult struct {
	Body        bytes.Buffer
	ContentType string
}

type Storage interface {
	GetImageDir(filename string) string
	Get(filename string, context context.Context) (ImageResult, bool, error)
	Upload(context context.Context, filename string, contentType string, source io.Reader) error
	Delete(context context.Context, filename string) error
}

func NewStorage(context context.Context, storageType StorageType, config *config.Configuration) (Storage, error) {
	if err := validateConfiguration(config, storageType); err != nil {
		return nil, err
	}
	switch storageType {
	case S3StorageType:
		return NewS3Storage(context, config)
	case LocalStorageType:
		return NewLocalStorage(config)
	}
	return nil, errors.New("invalid storage type")
}

func validateConfiguration(config *config.Configuration, storageType StorageType) error {
	switch storageType {
	case S3StorageType:
		if config.S3_ENDPOINT == "" {
			return errors.New("S3_ENDPOINT is required")
		}
		if config.S3_ACCESS_KEY == "" {
			return errors.New("S3_ACCESS_KEY is required")
		}
		if config.S3_SECRET_KEY == "" {
			return errors.New("S3_SECRET_KEY is required")
		}
		if config.S3_IMAGE_URL == "" {
			return errors.New("S3_IMAGE_URL is required")
		}
		if config.S3_UPLOAD_BUCKET == "" {
			return errors.New("S3_UPLOAD_BUCKET is required")
		}
	case LocalStorageType:
		if config.S3_UPLOAD_BUCKET == "" {
			return errors.New("S3_UPLOAD_BUCKET is required")
		}
		if config.S3_IMAGE_URL == "" {
			return errors.New("S3_IMAGE_URL is required")
		}
	}
	return nil
}
