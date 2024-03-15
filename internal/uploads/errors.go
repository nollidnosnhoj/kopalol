package uploads

import "fmt"

type ImageUploadError struct {
	err      string
	filename string
}

func (e *ImageUploadError) Error() string {
	return fmt.Sprintf("%s - %s", e.filename, e.err)
}

func NewImageError(err, filename string) *ImageUploadError {
	return &ImageUploadError{err: err, filename: filename}
}
