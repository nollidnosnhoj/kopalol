package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets/components"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/uploads"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

type UploadsController struct {
	uploader *uploads.Uploader
}

func NewUploadsController(container *config.Container) *UploadsController {
	return &UploadsController{
		uploader: uploads.NewUploader(container),
	}
}

func (u *UploadsController) RegisterRoutes(e *echo.Echo) {
	r := e.Group("/uploads")
	r.POST("/", u.uploadFiles())
}

func (u *UploadsController) uploadFiles() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger()
		ctx := c.Request().Context()
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		files := form.File["images"]
		results := u.uploader.UploadMultiple(files, ctx)
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

func (u *UploadsController) UploadFilesAPIHandler() echo.HandlerFunc {
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
		result := u.uploader.Upload(file, ctx)
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
