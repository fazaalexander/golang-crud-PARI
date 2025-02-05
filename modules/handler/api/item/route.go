package item

import (
	"github.com/labstack/echo/v4"
)

func (ih *ItemHandler) RegisterRoutes(e *echo.Echo) {
	authorGroup := e.Group("/items")
	authorGroup.GET("", ih.GetAllItems())
	authorGroup.GET("/:id", ih.GetItemById())
	authorGroup.POST("", ih.CreateItem())
	authorGroup.PUT("/:id", ih.UpdateItem())
	authorGroup.DELETE("/:id", ih.DeleteItem())
}
