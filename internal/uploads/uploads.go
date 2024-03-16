package uploads

import (
	"context"
	"errors"
	"log/slog"
	"mime/multipart"

	"github.com/nollidnosnhoj/kopalol/internal/storage"
)

type Uploader struct {
	storage storage.Storage
	logger  *slog.Logger
}

func NewUploader(s storage.Storage, logger *slog.Logger) *Uploader {
	return &Uploader{storage: s, logger: logger}
}

func (u *Uploader) UploadMultiple(images []*multipart.FileHeader, ctx context.Context) []*FileUpload {
	results := []*FileUpload{}
	for _, image := range images {
		res := u.Upload(image, ctx)
		results = append(results, res)
	}
	return results
}

func (u *Uploader) Upload(image *multipart.FileHeader, ctx context.Context) *FileUpload {
	params, err := validateImage(image)
	if err != nil {
		switch {
		case errors.Is(err, ErrFileTooLarge):
			return &FileUpload{Error: NewFileUploadError("file too large", image.Filename)}
		default:
			return &FileUpload{Error: NewFileUploadError("unable to validate file", image.Filename)}
		}
	}
	source, err := image.Open()
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to open file", image.Filename)}
	}
	defer source.Close()
	err = u.storage.Upload(ctx, params.FileName, params.FileType, source)
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to upload file", image.Filename)}
	}
	params.Url = u.storage.GetImageDir(params.FileName)
	return params
}
