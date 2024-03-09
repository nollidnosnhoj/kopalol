package images

import (
	"mime"
	"path/filepath"
)

func CreateImageFileName(filename string, id string) string {
	ext := filepath.Ext(filename)
	return id + ext
}

func GetContentType(filename string) string {
	mType := mime.TypeByExtension(filepath.Ext(filename))
	if mType == "" {
		return "application/octet-stream"
	}
	return mType
}
