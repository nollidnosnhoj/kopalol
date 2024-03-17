package controllers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/components"
	"github.com/nollidnosnhoj/kopalol/internal/queries"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/nollidnosnhoj/kopalol/internal/uploads"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

func UploadFilesHandler(s storage.Storage, q *queries.Queries, l *slog.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		files := form.File["images"]
		uploader := uploads.NewUploader(s, l)
		results := uploader.UploadMultiple(files, ctx)
		for _, result := range results {
			deletionKey, err := utils.GenerateDeletionKey()
			if err != nil {
				result.Error = errors.New("unable to generate deletion key")
				continue
			}
			result.DeletionKey = deletionKey
			_, err = q.InsertFile(ctx, queries.InsertFileParams{
				ID:               result.ID,
				FileName:         result.FileName,
				OriginalFileName: result.OriginalFileName,
				FileSize:         result.FileSize,
				FileType:         result.FileType,
				FileExtension:    result.FileExtension,
				DeletionKey:      deletionKey,
			})
			if err != nil {
				l.Error(err.Error())
				result.Error = errors.New("unable to save file to database")
			}
		}
		return utils.RenderComponent(c, http.StatusOK, components.UploadResults(results))
	}
}
