package seed

import (
	"time"

	ie "github.com/fazaalexander/golang-crud-PARI.git/modules/entity/item"
)

func CreateItem() []*ie.Item {
	Items := []*ie.Item{
		{
			ID:          1,
			Name:        "Item 1",
			Description: "Item 1 description",
			Stock:       20,
			Price:       100000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Item 2",
			Description: "Item 2 description",
			Stock:       15,
			Price:       75000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          3,
			Name:        "Item 3",
			Description: "Item 3 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          4,
			Name:        "Item 4",
			Description: "Item 4 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          5,
			Name:        "Item 5",
			Description: "Item 5 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          6,
			Name:        "Item 6",
			Description: "Item 6 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          7,
			Name:        "Item 7",
			Description: "Item 7 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          8,
			Name:        "Item 8",
			Description: "Item 8 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          9,
			Name:        "Item 9",
			Description: "Item 9 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          10,
			Name:        "Item 10",
			Description: "Item 10 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
		{
			ID:          11,
			Name:        "Item 11",
			Description: "Item 11 description",
			Stock:       5,
			Price:       300000,
			CreatedAt:   time.Now(),
		},
	}

	return Items
}
