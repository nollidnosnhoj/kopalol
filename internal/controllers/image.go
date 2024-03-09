package controllers

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"github.com/nollidnosnhoj/vgpx/internal/cache"
// 	"github.com/nollidnosnhoj/vgpx/internal/storage"
// )

// type ImageController struct {
// 	cache   *cache.Cache
// 	storage storage.Storage
// }

// func NewImageController(c *cache.Cache, s storage.Storage) *ImageController {
// 	return &ImageController{c, s}
// }

// func (h *ImageController) RegisterRoutes(router *echo.Echo) {
// 	i := router.Group("/i")
// 	i.GET("/:filename", h.getImage())
// }

// func (h *ImageController) getImage() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		filename := c.Param("filename")
// 		key := cacheKey(filename)
// 		cacheVal, ok := h.cache.Get(key)
// 		if ok {
// 			res := cacheVal.Value.(storage.ImageResult)
// 			return c.Blob(http.StatusOK, res.ContentType, res.Body.Bytes())
// 		}
// 		result, found, err := h.storage.Get(filename, c.Request().Context())
// 		if !found {
// 			return c.JSON(http.StatusNotFound, map[string]string{"error": "Image not found"})
// 		}
// 		if err != nil {
// 			c.Logger().Error(err)
// 			return err
// 		}
// 		h.cache.Set(key, storage.ImageResult{Body: result.Body, ContentType: result.ContentType})
// 		return c.Blob(http.StatusOK, result.ContentType, result.Body.Bytes())
// 	}
// }

// func cacheKey(filename string) string {
// 	return "image:" + filename
// }
