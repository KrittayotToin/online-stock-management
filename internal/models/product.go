package models

import (
	"time"
)

// Product represents a product in the inventory
type Product struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"` // Auto-incrementing integer ID
	Name        string    `json:"name" validate:"required"`
	Price       float64   `json:"price" validate:"gt=0"`
	Description string    `json:"description"`
	SKU         string    `json:"sku" validate:"required"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
