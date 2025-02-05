package item

import ic "github.com/fazaalexander/golang-crud-PARI.git/modules/usecase/item"

type ItemHandler struct {
	itemUseCase ic.ItemUseCase
}

func New(itemUseCase ic.ItemUseCase) *ItemHandler {
	return &ItemHandler{
		itemUseCase,
	}
}
