package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type LocalStorage struct {
	Folder string
	url    string
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
		url:    "http://localhost:8080/uploads",
	}, nil
}

func (s *LocalStorage) GetImageDir(filename string) string {
	return fmt.Sprintf("%s/%s", s.url, filename)
}

func (s *LocalStorage) Get(filename string, context context.Context) (ImageResult, bool, error) {
	file, err := os.Open(path.Join(s.Folder, filename))
	if err != nil {
		if os.IsNotExist(err) {
			return ImageResult{}, false, nil
		}
		return ImageResult{}, false, err
	}
	defer file.Close()
	byteArr, err := io.ReadAll(file)
	if err != nil {
		return ImageResult{}, false, err
	}
	contentType := http.DetectContentType(byteArr)
	buffer := bytes.NewBuffer(byteArr)
	return ImageResult{Body: *buffer, ContentType: contentType}, true, nil
}

func (s *LocalStorage) Upload(context context.Context, filename string, contentType string, source io.Reader) error {
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

func (s *LocalStorage) Delete(context context.Context, filename string) error {
	filePath := path.Join(s.Folder, filename)
	err := os.Remove(filePath)
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			return nil
		}
		return err
	}
	return nil

}
