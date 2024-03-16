package uploads

import (
	"mime/multipart"
	"path/filepath"
	"slices"

	"github.com/gabriel-vasile/mimetype"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
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

func validateImage(file *multipart.FileHeader) (*FileUpload, error) {
	var params = &FileUpload{}

	// validate size
	if file.Size > MAX_UPLOAD_SIZE {
		return nil, ErrFileTooLarge
	}
	params.FileSize = file.Size

	// validate extension
	params.FileExtension = filepath.Ext(file.Filename)
	if params.FileExtension == "" || !slices.Contains(VALID_FILE_EXTENSIONS, params.FileExtension) {
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
	params.FileType = mimeType.String()
	if !slices.Contains(VALID_FILE_TYPES, params.FileType) {
		return nil, ErrInvalidFileType
	}

	// generate id
	id, err := utils.GenerateRandomId(10)
	if err != nil {
		return nil, err
	}
	params.ID = id
	params.FileName = createImageFileName(file.Filename, id)
	params.OriginalFileName = file.Filename

	return params, nil
}

func createImageFileName(filename string, id string) string {
	ext := filepath.Ext(filename)
	return id + ext
}
