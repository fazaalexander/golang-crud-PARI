package item

import (
	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
	"gorm.io/gorm"
)

type ItemRepo interface {
	GetAllItems(offset, pageSize int) ([]ie.Item, int64, error)
	GetItemById(itemId uint64, item *ie.Item) (*ie.Item, error)
	CreateItem(item *ie.Item) error
	UpdateItem(itemId uint64, item *ie.Item) error
	DeleteItem(itemId uint64, item *ie.Item) error
}

type itemRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ItemRepo {
	return &itemRepo{
		db,
	}
}
