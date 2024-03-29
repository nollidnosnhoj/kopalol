package uploads

import (
	"errors"
	"fmt"
)

var ErrFileTooLarge = errors.New("file too large")
var ErrInvalidFileType = errors.New("invalid image type")

type FileUploadError struct {
	err      string
	filename string
}

func (e *FileUploadError) Error() string {
	return fmt.Sprintf("%s - %s", e.filename, e.err)
}

func NewFileUploadError(err, filename string) *FileUploadError {
	return &FileUploadError{err: err, filename: filename}
}
