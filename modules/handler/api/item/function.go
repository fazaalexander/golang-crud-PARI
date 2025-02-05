package item

import (
	"math"
	"net/http"
	"strconv"

	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
	"github.com/labstack/echo/v4"
)

func (ih *ItemHandler) GetAllItems() echo.HandlerFunc {
	return func(c echo.Context) error {
		var items []ie.Item

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		items, total, err := ih.itemUseCase.GetAllItems(offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Failed to get items due to: " + err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if len(items) == 0 {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Items not found",
				"Status":  http.StatusNotFound,
			})
		}

		var itemResponses []ie.ItemResponse
		for _, item := range items {
			itemResponse := ie.ItemResponse{
				ItemID:      item.ID,
				Name:        item.Name,
				Description: item.Description,
				Stock:       item.Stock,
				Price:       item.Price,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
			}

			itemResponses = append(itemResponses, itemResponse)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Items":     itemResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Message":   "Success get all items",
			"Status":    http.StatusOK,
		})
	}
}

func (ih *ItemHandler) GetItemById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var item *ie.Item
		itemID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		item, err := ih.itemUseCase.GetItemById(itemID, item)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Failed to get item",
				"Error":   err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		itemResponse := ie.ItemResponse{
			ItemID:      item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Item":    itemResponse,
			"Message": "Success get item",
			"Status":  http.StatusOK,
		})
	}
}

func (ih *ItemHandler) CreateItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var item ie.ItemRequest
		if err := c.Bind(&item); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Failed to create item",
				"Error":   err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ih.itemUseCase.CreateItem(&item)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to create item due to: " + err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "New Item Created",
			"Status":  http.StatusOK,
		})
	}
}

func (ih *ItemHandler) UpdateItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var item *ie.Item
		var itemRequest ie.ItemRequest
		itemID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		if err := c.Bind(&itemRequest); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": "Failed to update item",
				"Error":   err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		_, err := ih.itemUseCase.GetItemById(itemID, item)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Failed to update item",
				"Error":   err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		err = ih.itemUseCase.UpdateItem(itemID, &itemRequest)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to update item due to: " + err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "Item Updated Successfully",
			"Status":  http.StatusOK,
		})
	}
}

func (ih *ItemHandler) DeleteItem() echo.HandlerFunc {
	return func(c echo.Context) error {
		var item *ie.Item
		itemID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		item, err := ih.itemUseCase.GetItemById(itemID, item)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Failed to delete item",
				"Error":   err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		err = ih.itemUseCase.DeleteItem(itemID, item)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": "Failed to delete item due to: " + err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Item successfully deleted",
			"Status":  http.StatusOK,
		})
	}
}
