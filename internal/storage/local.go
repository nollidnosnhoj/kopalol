package storage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type LocalStorage struct {
	Folder string
}

func NewLocalStorage(folder string) (*LocalStorage, error) {
	dir, err := filepath.Abs("uploads")
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &LocalStorage{
		Folder: folder,
	}, nil
}

func (s *LocalStorage) Get(filename string, context context.Context) (ImageResult, error) {
	file, err := os.Open(path.Join(s.Folder, filename))
	if err != nil {
		return ImageResult{}, err
	}
	defer file.Close()
	byteArr, err := io.ReadAll(file)
	if err != nil {
		return ImageResult{}, err
	}
	contentType := http.DetectContentType(byteArr)
	buffer := bytes.NewBuffer(byteArr)
	return ImageResult{Body: *buffer, ContentType: contentType}, nil
}

func (s *LocalStorage) Upload(filename string, source io.Reader, context context.Context) error {
	newFileName := path.Join(s.Folder, filename)
	dest, err := os.Create(newFileName)
	if err != nil {
		return err
	}
	defer dest.Close()
	if _, err := io.Copy(dest, source); err != nil {
		return err
	}
	return nil
}
