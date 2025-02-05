package item

import (
	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
	"github.com/labstack/echo/v4"
)

func (ir *itemRepo) GetAllItems(offset, pageSize int) ([]ie.Item, int64, error) {
	var items []ie.Item
	var count int64

	if err := ir.db.Model(&ie.Item{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := ir.db.Limit(pageSize).Offset(offset).Order("created_at ASC").Find(&items).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return items, count, nil
}

func (ir *itemRepo) GetItemById(itemId uint64, item *ie.Item) (*ie.Item, error) {
	if err := ir.db.Where("id = ?", itemId).First(&item).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return item, nil
}

func (ir *itemRepo) CreateItem(item *ie.Item) error {
	if err := ir.db.Create(&item).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (ir *itemRepo) UpdateItem(itemId uint64, item *ie.Item) error {
	if err := ir.db.Model(&ie.Item{}).Where("id = ?", itemId).
		Updates(ie.Item{Name: item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			UpdatedAt:   item.UpdatedAt,
		}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (ir *itemRepo) DeleteItem(itemId uint64, item *ie.Item) error {
	if err := ir.db.Where("id = ?", itemId).Delete(&item).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}
