package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/internal/components"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/uploads"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
)

type UploadController struct {
	logger  *slog.Logger
	storage storage.Storage
}

func NewUploadController(s storage.Storage, l *slog.Logger) *UploadController {
	return &UploadController{storage: s, logger: l}
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
		results := make([]uploads.ImageUploadResult, len(files))
		for _, image := range files {
			uploader := uploads.NewUploader(h.storage, h.logger)
			res, err := uploader.Upload(image, ctx)
			if err != nil {
				return utils.RenderComponent(c, http.StatusOK, components.Uploader(err))
			}
			results = append(results, res)
		}
		return utils.RenderComponent(c, http.StatusOK, components.ImageUploadSuccess(results))
	}
}
