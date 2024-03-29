package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets/components"
	"github.com/nollidnosnhoj/kopalol/assets/views"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/queries"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

type FilesController struct {
	storage storage.Storage
	queries *queries.Queries
}

func NewFilesController(container *config.Container) *FilesController {
	return &FilesController{
		storage: container.Storage(),
		queries: container.Database().Queries(),
	}
}

func (f *FilesController) RegisterRoutes(e *echo.Echo) {
	r := e.Group("/files")
	r.GET("/:id/delete/:delete_key", f.showFileDeletionPage())
	r.DELETE("/:id", f.deleteFile())
}

func (f *FilesController) showFileDeletionPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.Param("delete_key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		file, err := f.queries.GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		previewUrl := f.storage.GetImageDir(file.FileName)
		return utils.RenderComponent(c, http.StatusOK, views.ShowFileDeletionConfirmationPage(file, previewUrl))
	}
}

func (f *FilesController) deleteFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.FormValue("key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("deletionKey is required"))
		}
		_, err := f.queries.GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("file not found"))
		}
		err = f.queries.DeleteFile(ctx, id)
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("unable to delete file"))
		}
		return utils.RenderComponent(c, http.StatusOK, components.FileDeletionSuccess())
	}
}
