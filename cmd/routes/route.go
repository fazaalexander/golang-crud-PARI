package routes

import (
	"github.com/fazaalexander/golang-crud-PARI.git/common"
	"github.com/fazaalexander/golang-crud-PARI.git/middleware/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	log.LogMiddleware(e)
	e.Use(middleware.CORS())

	handler.ItemHandler.RegisterRoutes(e)

	return e
}
