package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets"
	"github.com/nollidnosnhoj/kopalol/internal/container"
	"github.com/nollidnosnhoj/kopalol/internal/controllers"
	"github.com/nollidnosnhoj/kopalol/internal/router/middlewares"
)

func NewRouter(container *container.Container) *echo.Echo {
	e := echo.New()

	publicDistFs := assets.BuildPublicDistFs()
	e.StaticFS("/dist/", publicDistFs)

	// request logs
	e.Use(middlewares.RequestLoggingMiddleware(container.Logger))

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
	apiRouter := e.Group("/api", middlewares.ApiKeyAuthMiddleware())
	uploadsController.RegisterAPIRoutes(apiRouter)

	return e
}
