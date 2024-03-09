package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/vgpx/assets"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	assetsFs := assets.BuildAssets()
	e.StaticFS("/dist/", assetsFs)

	return e
}
