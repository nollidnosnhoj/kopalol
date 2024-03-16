package storage

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/nollidnosnhoj/vgpx/internal/config"
)

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

func NewStorage(typeStorage string, config *config.Config, ctx context.Context) (Storage, error) {
	switch typeStorage {
	case "local":
		storage, err := NewLocalStorage(config.UPLOAD_BUCKET_NAME)
		if err != nil {
			return nil, err
		}
		return storage, nil
	case "s3":
		storage, err := NewS3Storage(ctx, config)
		if err != nil {
			return nil, err
		}
		return storage, nil
	default:
		return nil, errors.New("invalid storage type")
	}

}
