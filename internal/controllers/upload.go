package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets/templ/components"
	"github.com/nollidnosnhoj/kopalol/internal/container"
	"github.com/nollidnosnhoj/kopalol/internal/models"
	"github.com/nollidnosnhoj/kopalol/internal/uploads"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

type UploadsController struct {
	uploader *uploads.Uploader
}

func NewUploadsController(container *container.Container) *UploadsController {
	return &UploadsController{
		uploader: uploads.NewUploader(container),
	}
}

func (u *UploadsController) RegisterRoutes(e *echo.Echo) {
	e.POST("/upload", u.uploadFiles)
}

func (u *UploadsController) RegisterAPIRoutes(e *echo.Group) {
	e.POST("/upload", u.uploadFilesAPI)
}

func (u *UploadsController) uploadFiles(c echo.Context) error {
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

func (u *UploadsController) uploadFilesAPI(c echo.Context) error {
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
	response := models.UploadFileResponse{
		Id:          result.ID,
		ContentType: result.FileType,
		FileSize:    result.FileSize,
		DeletionKey: result.DeletionKey,
		Md5Hash:     result.Md5Hash,
		Url:         result.Url,
		CreatedAt:   result.CreatedAt,
	}
	return c.JSON(http.StatusOK, response)
}
