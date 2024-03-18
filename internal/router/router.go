package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/controllers"
)

func NewRouter(container *config.Container) *echo.Echo {
	e := echo.New()

	publicDistFs := assets.BuildPublicDistFs()
	e.StaticFS("/dist/", publicDistFs)

	// request logs
	e.Use(RequestLoggingMiddleware(container.Logger()))

	// rate limiter
	e.Use(RateLimitMiddleware())

	// pages
	pagesController := controllers.NewPagesController()
	pagesController.RegisterRoutes(e)

	// uploads
	uploadsController := controllers.NewUploadsController(container)
	uploadsController.RegisterRoutes(e)

	// files
	filesController := controllers.NewFilesController(container)
	filesController.RegisterRoutes(e)

	return e
}
