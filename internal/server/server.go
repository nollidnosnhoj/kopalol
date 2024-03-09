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
	echo *echo.Echo
}

func NewServer(cache *cache.Cache, uploadStorage storage.Storage) *Server {
	e := echo.New()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	e.GET("/", func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/i/:filename", func(c echo.Context) error {
		filename := c.Param("filename")
		cacheKey := images.GetCacheKey(filename)
		cacheVal, ok := cache.Get(cacheKey)
		if ok {
			log.Printf("Cache hit for %s", cacheKey)
			value := cacheVal.Value.(storage.ImageResult)
			setCacheControlForImage(c)
			return c.Blob(http.StatusOK, value.ContentType, value.Body.Bytes())
		}
		log.Printf("Cache miss for %s", cacheKey)
		result, found, err := uploadStorage.Get(filename, c.Request().Context())
		if !found {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Image not found"})
		}
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		cache.Set(cacheKey, storage.ImageResult{Body: result.Body, ContentType: result.ContentType})
		setCacheControlForImage(c)
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
		err = uploadStorage.Upload(filename, source, ctx)
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		return utils.RenderComponent(c, http.StatusOK, components.ImageUploaded())
	})

	return &Server{
		echo: e,
	}
}

func setCacheControlForImage(c echo.Context) {
	maxAge := 60 * 60 * 24 * 365
	c.Response().Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
}
