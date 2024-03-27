package uploads

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"mime/multipart"

	"github.com/nollidnosnhoj/kopalol/config"
	"github.com/nollidnosnhoj/kopalol/queries"
	"github.com/nollidnosnhoj/kopalol/storage"
	"github.com/nollidnosnhoj/kopalol/utils"
)

type Uploader struct {
	queries *queries.Queries
	storage storage.Storage
	logger  *slog.Logger
}

func NewUploader(container *config.Container) *Uploader {
	return &Uploader{
		queries: container.Database().Queries(),
		storage: container.Storage(),
		logger:  container.Logger(),
	}
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
	fileInfo, err := validateImage(image)
	if err != nil {
		switch {
		case errors.Is(err, ErrFileTooLarge):
			return &FileUpload{Error: NewFileUploadError("file too large", image.Filename)}
		default:
			return &FileUpload{Error: NewFileUploadError("unable to validate file", image.Filename)}
		}
	}
	id, err := utils.GenerateRandomId(10)
	if err != nil {
		return &FileUpload{Error: NewFileUploadError("unable to generate id", image.Filename)}
	}
	fileName := id + fileInfo.Ext
	source, err := image.Open()
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to open file", image.Filename)}
	}
	defer source.Close()
	fileBytes, err := io.ReadAll(source)
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to read file", image.Filename)}
	}
	md5Hash := utils.EncodeToMd5(fileBytes)
	err = u.storage.Upload(ctx, fileName, fileInfo.Type, source)
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to upload file", image.Filename)}
	}
	deletionKey, err := utils.GenerateDeletionKey()
	if err != nil {
		return &FileUpload{Error: NewFileUploadError("unable to generate deletion key", image.Filename)}
	}

	file, err := u.queries.InsertFile(ctx, queries.InsertFileParams{
		ID:               id,
		FileName:         fileName,
		OriginalFileName: fileInfo.Name,
		FileSize:         fileInfo.Size,
		FileType:         fileInfo.Type,
		FileExtension:    fileInfo.Ext,
		Md5Hash:          md5Hash,
		DeletionKey:      deletionKey,
	})
	if err != nil {
		u.logger.Error(err.Error())
		return &FileUpload{Error: NewFileUploadError("unable to save file to database", image.Filename)}
	}

	url := u.storage.GetImageDir(file.FileName)

	return &FileUpload{
		File:  file,
		Url:   url,
		Error: nil,
	}
}
