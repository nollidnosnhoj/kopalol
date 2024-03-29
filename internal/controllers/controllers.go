package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets/views"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

type PagesController struct{}

func NewPagesController() *PagesController {
	return &PagesController{}
}

func (p *PagesController) RegisterRoutes(e *echo.Echo) {
	e.GET("/", p.home())
}

func (p *PagesController) home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return utils.RenderComponent(c, http.StatusOK, views.IndexPage())
	}
}
