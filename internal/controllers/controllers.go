package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
	"github.com/nollidnosnhoj/kopalol/internal/views"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) RegisterRoutes(router *echo.Echo) {
	router.GET("/", h.homeHandler())
}

func (h *HomeController) homeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	}
}
