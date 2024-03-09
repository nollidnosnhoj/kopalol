package utils

import (
	"mime"
	"path/filepath"
)

func GetContentType(filename string) string {
	mType := mime.TypeByExtension(filepath.Ext(filename))
	if mType == "" {
		return "application/octet-stream"
	}
	return mType
}
