package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
	"github.com/nollidnosnhoj/kopalol/internal/views"
)

func ShowHomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	}
}
