package uploads

import (
	"context"
	"fmt"
	"log/slog"
	"mime/multipart"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
)

type Uploader struct {
	storage storage.Storage
	logger  *slog.Logger
}

func NewUploader(s storage.Storage, logger *slog.Logger) *Uploader {
	return &Uploader{storage: s, logger: logger}
}

func (u *Uploader) UploadMultiple(images []*multipart.FileHeader, ctx context.Context) ([]ImageUploadResult, error) {
	results := make([]ImageUploadResult, len(images))
	for i, image := range images {
		res, err := u.Upload(image, ctx)
		if err != nil {
			return nil, err
		}
		results[i] = res
	}
	return results, nil
}

func (u *Uploader) Upload(image *multipart.FileHeader, ctx context.Context) (ImageUploadResult, error) {
	if image.Size > 5*1024*1024 {
		return ImageUploadResult{}, fmt.Errorf("image too large: %s", image.Filename)
	}
	source, err := image.Open()
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, fmt.Errorf("unable to open image: %s", image.Filename)
	}
	mimeType, err := mimetype.DetectReader(source)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, fmt.Errorf("unable to detect image type: %s", image.Filename)
	}
	if !mimeType.Is("image/jpeg") && !mimeType.Is("image/png") {
		return ImageUploadResult{}, fmt.Errorf("unsupported image type: %s", mimeType.String())
	}
	defer source.Close()
	id, err := utils.GenerateRandomId(10)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, fmt.Errorf("unable to upload image: %s", image.Filename)
	}
	filename := createImageFileName(image.Filename, id)
	err = u.storage.Upload(filename, source, ctx)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, fmt.Errorf("unable to upload image: %s", image.Filename)
	}
	return ImageUploadResult{Id: id, Url: u.storage.GetImageDir(filename)}, nil
}

func createImageFileName(filename string, id string) string {
	ext := filepath.Ext(filename)
	return id + ext
}
