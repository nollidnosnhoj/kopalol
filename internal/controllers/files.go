package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/components"
	"github.com/nollidnosnhoj/kopalol/internal/queries"
	"github.com/nollidnosnhoj/kopalol/internal/storage"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
	"github.com/nollidnosnhoj/kopalol/internal/views"
)

type FilesController struct {
	queries *queries.Queries
	logger  *slog.Logger
	storage storage.Storage
}

func NewFilesController(q *queries.Queries, s storage.Storage, l *slog.Logger) *FilesController {
	return &FilesController{queries: q, storage: s, logger: l}
}

func (h *FilesController) RegisterRoutes(router *echo.Echo) {
	group := router.Group("/files")
	group.GET("/:id/delete/:delete_key", h.showDeletionPage())
	group.DELETE("/:id", h.deleteFileHandler())
}

func (h *FilesController) showDeletionPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.Param("delete_key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		file, err := h.queries.GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		previewUrl := h.storage.GetImageDir(file.FileName)
		return utils.RenderComponent(c, http.StatusOK, views.ShowFileDeletionConfirmationPage(file, previewUrl))
	}
}

func (h *FilesController) deleteFileHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.FormValue("key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("deletionKey is required"))
		}
		_, err := h.queries.GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("file not found"))
		}
		err = h.queries.DeleteFile(ctx, id)
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("unable to delete file"))
		}
		return utils.RenderComponent(c, http.StatusOK, components.FileDeletionSuccess())
	}
}
