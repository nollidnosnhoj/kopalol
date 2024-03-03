package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/assets"
	"github.com/nollidnosnhoj/vgpx/internal/components"
	"github.com/nollidnosnhoj/vgpx/internal/images"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
	"github.com/nollidnosnhoj/vgpx/internal/views"
)

type Server struct {
	Log  *log.Logger
	Echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	logger := log.Default()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	e.Static("/i", "uploads")

	e.GET("/", func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	})

	e.POST("/upload", func(c echo.Context) error {
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
		uploadStorage, err := storage.NewLocalStorage("uploads")
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		err = uploadStorage.Upload(filename, source)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		return utils.RenderComponent(c, http.StatusOK, components.ImageUploaded())
	})

	return &Server{
		Log:  logger,
		Echo: e,
	}
}
