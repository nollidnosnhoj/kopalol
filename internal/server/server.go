package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/simplimg/assets"
	"github.com/nollidnosnhoj/simplimg/internal/utils"
	"github.com/nollidnosnhoj/simplimg/internal/views"
)

type Server struct {
	Log  *log.Logger
	Echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	logger := log.Default()

	fs := assets.Build()
	e.StaticFS("/dist/", fs)

	e.GET("/", func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	})

	return &Server{
		Log:  logger,
		Echo: e,
	}
}
