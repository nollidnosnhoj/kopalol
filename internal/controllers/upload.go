package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/components"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/uploads"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

func UploadFilesHandler(container *config.Container) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		files := form.File["images"]
		uploader := uploads.NewUploader(container)
		results := uploader.UploadMultiple(files, ctx)
		return utils.RenderComponent(c, http.StatusOK, components.UploadResults(results))
	}
}

type UploadFileResponse struct {
	Id          string    `json:"id"`
	ContentType string    `json:"content_type"`
	FileSize    int64     `json:"file_size"`
	DeletionKey string    `json:"deletion_key"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
}

func UploadFilesAPIHandler(container *config.Container) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		files := form.File["file"]
		if len(files) == 0 {
			return c.JSON(http.StatusBadRequest, "no files found")
		}
		file := files[0]
		uploader := uploads.NewUploader(container)
		result := uploader.Upload(file, ctx)
		if result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.Error)
		}
		response := UploadFileResponse{
			Id:          result.ID,
			ContentType: result.FileType,
			FileSize:    result.FileSize,
			DeletionKey: result.DeletionKey,
			Url:         result.Url,
			CreatedAt:   result.CreatedAt,
		}
		return c.JSON(http.StatusOK, response)
	}
}
