package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/controllers"
)

func NewRouter(container *config.Container) *echo.Echo {
	e := echo.New()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	// request logs
	e.Use(RequestLoggingMiddleware(container.Logger()))

	// rate limiter
	e.Use(RateLimitMiddleware())

	e.GET("/", controllers.ShowHomeHandler())
	e.POST("/upload", controllers.UploadFilesHandler(container))

	filesRouter := e.Group("/files")
	filesRouter.GET("/:id/delete/:delete_key", controllers.ShowFileDeletionPageHandler(container))
	filesRouter.DELETE("/:id", controllers.DeleteFileHandler(container.Queries()))

	return e
}
