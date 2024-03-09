package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/assets"
	"github.com/nollidnosnhoj/vgpx/internal/cache"
	"github.com/nollidnosnhoj/vgpx/internal/controllers"
	"github.com/nollidnosnhoj/vgpx/internal/storage"
	"github.com/nollidnosnhoj/vgpx/internal/utils"
	"github.com/nollidnosnhoj/vgpx/internal/views"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(cache *cache.Cache, uploadStorage storage.Storage) *Server {
	e := echo.New()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	imageController := controllers.NewImageController(e, cache, uploadStorage)
	uploadController := controllers.NewUploadController(e, uploadStorage)

	e.GET("/", func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	})

	e.GET("/i/:filename", imageController.GetImage())
	e.POST("/upload", uploadController.Upload())

	return &Server{
		Router: e,
	}
}
