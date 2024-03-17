package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/assets/components"
	"github.com/nollidnosnhoj/kopalol/assets/views"
	"github.com/nollidnosnhoj/kopalol/internal/config"
	"github.com/nollidnosnhoj/kopalol/internal/queries"
	"github.com/nollidnosnhoj/kopalol/internal/utils"
)

func ShowFileDeletionPageHandler(container *config.Container) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.Param("delete_key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		file, err := container.Queries().GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, views.NotFoundPage())
		}
		previewUrl := container.Storage().GetImageDir(file.FileName)
		return utils.RenderComponent(c, http.StatusOK, views.ShowFileDeletionConfirmationPage(file, previewUrl))
	}
}

func DeleteFileHandler(q *queries.Queries) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		id := c.Param("id")
		deletionKey := c.FormValue("key")
		if deletionKey == "" {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("deletionKey is required"))
		}
		_, err := q.GetFileForDeletion(ctx, queries.GetFileForDeletionParams{
			ID:          id,
			DeletionKey: deletionKey,
		})
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("file not found"))
		}
		err = q.DeleteFile(ctx, id)
		if err != nil {
			return utils.RenderComponent(c, http.StatusOK, components.FileDeletionError("unable to delete file"))
		}
		return utils.RenderComponent(c, http.StatusOK, components.FileDeletionSuccess())
	}
}
