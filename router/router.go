package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets"
	"github.com/nollidnosnhoj/kopalol/config"
	"github.com/nollidnosnhoj/kopalol/controllers"
	"github.com/nollidnosnhoj/kopalol/router/middlewares"
)

func NewRouter(container *config.Container) *echo.Echo {
	e := echo.New()

	publicDistFs := assets.BuildPublicDistFs()
	e.StaticFS("/dist/", publicDistFs)

	// request logs
	e.Use(middlewares.RequestLoggingMiddleware(container.Logger()))

	// rate limiter
	e.Use(middlewares.RateLimitMiddleware())

	// pages
	pagesController := controllers.NewPagesController()
	pagesController.RegisterRoutes(e)

	// uploads
	uploadsController := controllers.NewUploadsController(container)
	uploadsController.RegisterRoutes(e)

	// files
	filesController := controllers.NewFilesController(container)
	filesController.RegisterRoutes(e)

	// api
	apiRouter := e.Group("/api")
	uploadsController.RegisterAPIRoutes(apiRouter)

	return e
}
