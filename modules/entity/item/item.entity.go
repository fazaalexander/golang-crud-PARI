package item

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	*gorm.Model `json:"-"`
	ID          uint64    `gorm:"primaryKey" json:"Id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Stock       int       `json:"stock" validate:"required"`
	Price       float64   `json:"price" validate:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ItemRequest struct {
	Name        string  `json:"name" form:"name" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required"`
	Stock       int     `json:"stock" form:"stock" validate:"required"`
	Price       float64 `json:"price" form:"price" validate:"required"`
}

type ItemResponse struct {
	ItemID      uint64    `json:"item_id" form:"item_id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       int       `json:"stock" form:"stock"`
	Price       float64   `json:"price" form:"price"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}
