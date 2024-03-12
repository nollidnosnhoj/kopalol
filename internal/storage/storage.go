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
	Upload(filename string, source io.Reader, context context.Context) error
}
