package storage

import (
	"bytes"
	"context"
	"io"
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
