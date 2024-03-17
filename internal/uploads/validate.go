package uploads

import (
	"mime/multipart"
	"path/filepath"
	"slices"

	"github.com/gabriel-vasile/mimetype"
)

var MAX_UPLOAD_SIZE int64 = 500 * 1024 * 1024
var VALID_FILE_EXTENSIONS = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
}
var VALID_FILE_TYPES = []string{
	"image/jpeg",
	"image/png",
	"image/gif",
}

type FileInfo struct {
	Name string
	Ext  string
	Type string
	Size int64
}

func validateImage(file *multipart.FileHeader) (*FileInfo, error) {
	var fileInfo = &FileInfo{
		Name: file.Filename,
		Size: file.Size,
		Ext:  filepath.Ext(file.Filename),
	}

	// validate size
	if fileInfo.Size > MAX_UPLOAD_SIZE {
		return nil, ErrFileTooLarge
	}

	// validate extension
	if fileInfo.Ext == "" || !slices.Contains(VALID_FILE_EXTENSIONS, fileInfo.Ext) {
		return nil, ErrInvalidFileType
	}

	source, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer source.Close()

	// validate content-type
	mimeType, err := mimetype.DetectReader(source)
	if err != nil {
		return nil, err
	}
	fileInfo.Type = mimeType.String()
	if !slices.Contains(VALID_FILE_TYPES, fileInfo.Type) {
		return nil, ErrInvalidFileType
	}

	return fileInfo, nil
}
