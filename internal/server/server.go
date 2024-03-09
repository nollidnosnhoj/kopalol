package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/assets"
)

type Server struct {
	router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	e := echo.New()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	return &Server{
		router,
	}
}
