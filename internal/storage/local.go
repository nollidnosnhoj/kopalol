package storage

import (
	"io"
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

func (s *LocalStorage) Upload(filename string, source io.Reader) error {
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
