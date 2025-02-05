package item

import (
	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
	ir "github.com/fazaalexander/golang-crud-PARI.git/modules/repository/item"
)

type ItemUseCase interface {
	GetAllItems(offset, pageSize int) ([]ie.Item, int64, error)
	GetItemById(itemId uint64, item *ie.Item) (*ie.Item, error)
	CreateItem(item *ie.ItemRequest) error
	UpdateItem(itemId uint64, item *ie.ItemRequest) error
	DeleteItem(itemId uint64, item *ie.Item) error
}

type itemUseCase struct {
	itemRepo ir.ItemRepo
}

func New(authorRepo ir.ItemRepo) *itemUseCase {
	return &itemUseCase{
		authorRepo,
	}
}
