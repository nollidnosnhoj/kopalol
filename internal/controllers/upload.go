package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/internal/components"
	"github.com/nollidnosnhoj/vgpx/internal/images"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
)

type UploadController struct {
	echo    *echo.Echo
	storage storage.Storage
}

func NewUploadController(e *echo.Echo, s storage.Storage) *UploadController {
	return &UploadController{echo: e, storage: s}
}

func (h *UploadController) Upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		image, err := c.FormFile("image")
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		source, err := image.Open()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		defer source.Close()
		id, err := images.GenerateID()
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		filename := images.CreateImageFileName(image.Filename, id)
		err = h.storage.Upload(filename, source, ctx)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		return utils.RenderComponent(c, http.StatusOK, components.ImageUploaded())
	}
}
