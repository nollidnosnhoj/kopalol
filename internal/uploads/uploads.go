package uploads

import (
	"context"
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
		res, err := u.upload(image, ctx)
		if err != nil {
			return nil, err
		}
		results[i] = res
	}
	return results, nil
}

func (u *Uploader) upload(image *multipart.FileHeader, ctx context.Context) (ImageUploadResult, error) {
	if image.Size > 5*1024*1024 {
		return ImageUploadResult{}, NewImageError("image too large", image.Filename)
	}
	source, err := image.Open()
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, NewImageError("unable to open image", image.Filename)
	}
	mimeType, err := mimetype.DetectReader(source)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, NewImageError("unable to detect image type", image.Filename)
	}
	if !mimeType.Is("image/jpeg") && !mimeType.Is("image/png") {
		return ImageUploadResult{}, NewImageError("invalid image type", image.Filename)
	}
	source.Seek(0, 0)
	defer source.Close()
	id, err := utils.GenerateRandomId(10)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, NewImageError("unable to generate image id", image.Filename)
	}
	filename := createImageFileName(image.Filename, id)
	err = u.storage.Upload(ctx, filename, mimeType.String(), source)
	if err != nil {
		u.logger.Error(err.Error())
		return ImageUploadResult{}, NewImageError("unable to upload image", image.Filename)
	}
	return ImageUploadResult{Id: id, Url: u.storage.GetImageDir(filename)}, nil
}

func createImageFileName(filename string, id string) string {
	ext := filepath.Ext(filename)
	return id + ext
}
