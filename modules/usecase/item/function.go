package item

import (
	"time"

	vld "github.com/fazaalexander/golang-crud-PARI.git/helper/validator"
	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
)

func (ic *itemUseCase) GetAllItems(offset, pageSize int) ([]ie.Item, int64, error) {
	return ic.itemRepo.GetAllItems(offset, pageSize)
}

func (ic *itemUseCase) GetItemById(itemId uint64, item *ie.Item) (*ie.Item, error) {
	return ic.itemRepo.GetItemById(itemId, item)
}

func (ic *itemUseCase) CreateItem(itemRequest *ie.ItemRequest) error {
	if err := vld.Validation(itemRequest); err != nil {
		return err
	}

	item := &ie.Item{
		Name:        itemRequest.Name,
		Description: itemRequest.Description,
		Stock:       itemRequest.Stock,
		Price:       itemRequest.Price,
		CreatedAt:   time.Now(),
	}

	return ic.itemRepo.CreateItem(item)
}

func (ic *itemUseCase) UpdateItem(itemId uint64, itemRequest *ie.ItemRequest) error {
	if err := vld.Validation(itemRequest); err != nil {
		return err
	}

	item := &ie.Item{
		Name:        itemRequest.Name,
		Description: itemRequest.Description,
		Stock:       itemRequest.Stock,
		Price:       itemRequest.Price,
		UpdatedAt:   time.Now(),
	}

	return ic.itemRepo.UpdateItem(itemId, item)
}

func (ic *itemUseCase) DeleteItem(itemId uint64, item *ie.Item) error {
	return ic.itemRepo.DeleteItem(itemId, item)
}
