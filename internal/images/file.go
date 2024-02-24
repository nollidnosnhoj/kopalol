package images

import "path/filepath"

func CreateImageFileName(filename string, id string) string {
	ext := filepath.Ext(filename)
	return id + ext
}
