package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/assets"
	"github.com/nollidnosnhoj/vgpx/internal/cache"
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

func NewServer(cache *cache.Cache, uploadStorage storage.Storage) *Server {
	e := echo.New()

	logger := log.Default()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	e.GET("/", func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/i/:filename", func(c echo.Context) error {
		filename := c.Param("filename")
		cacheVal, contentType, ok := cache.Get(filename)
		if ok {
			c.Logger().Printf("Cache hit for %s", filename)
			return c.Blob(http.StatusOK, contentType, cacheVal.Bytes())
		}
		c.Logger().Printf("Cache miss for %s", filename)
		result, err := uploadStorage.Get(filename, c.Request().Context())
		if err != nil {
			return err
		}
		cache.Set(filename, result.Body, result.ContentType)
		maxAge := 60 * 60 * 24 * 365
		c.Response().Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
		return c.Blob(http.StatusOK, result.ContentType, result.Body.Bytes())
	})

	e.POST("/upload", func(c echo.Context) error {
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
		uploadStorage, err := storage.NewLocalStorage("uploads")
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		err = uploadStorage.Upload(filename, source, ctx)
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
