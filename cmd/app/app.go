package app

import (
	"github.com/fazaalexander/golang-crud-PARI.git/cmd/routes"
	"github.com/fazaalexander/golang-crud-PARI.git/common"
	"github.com/fazaalexander/golang-crud-PARI.git/database/mysql"
	itemHandler "github.com/fazaalexander/golang-crud-PARI.git/modules/handler/api/item"
	itemRepo "github.com/fazaalexander/golang-crud-PARI.git/modules/repository/item"
	itemUseCase "github.com/fazaalexander/golang-crud-PARI.git/modules/usecase/item"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	itemRepo := itemRepo.New(mysql.DB)
	itemUseCase := itemUseCase.New(itemRepo)
	itemHandler := itemHandler.New(itemUseCase)

	handler := common.Handler{
		ItemHandler: itemHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
