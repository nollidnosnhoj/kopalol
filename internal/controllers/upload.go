package controllers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/internal/components"
	"github.com/nollidnosnhoj/vgpx/internal/queries"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/uploads"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
)

type UploadController struct {
	queries *queries.Queries
	logger  *slog.Logger
	storage storage.Storage
}

func NewUploadController(q *queries.Queries, s storage.Storage, l *slog.Logger) *UploadController {
	return &UploadController{queries: q, storage: s, logger: l}
}

func (h *UploadController) RegisterRoutes(router *echo.Echo) {
	router.POST("/upload", h.uploadHandler())
}

func (h *UploadController) uploadHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		files := form.File["images"]
		uploader := uploads.NewUploader(h.storage, h.logger)
		results := uploader.UploadMultiple(files, ctx)
		for _, result := range results {
			deletionKey, err := utils.GenerateDeletionKey()
			if err != nil {
				result.Error = errors.New("unable to generate deletion key")
				continue
			}
			result.DeletionKey = deletionKey
			_, err = h.queries.InsertFile(ctx, queries.InsertFileParams{
				ID:               result.ID,
				FileName:         result.FileName,
				OriginalFileName: result.OriginalFileName,
				FileSize:         result.FileSize,
				FileType:         result.FileType,
				FileExtension:    result.FileExtension,
				DeletionKey:      deletionKey,
			})
			if err != nil {
				h.logger.Error(err.Error())
				result.Error = errors.New("unable to save file to database")
			}
		}
		return utils.RenderComponent(c, http.StatusOK, components.UploadResults(results))
	}
}
