package storage

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

func UploadToLocal(filename string, source io.Reader) error {
	uploadDir, err := filepath.Abs("uploads")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return err
	}
	newFileName := path.Join(uploadDir, filename)
	if err != nil {
		return err
	}
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
