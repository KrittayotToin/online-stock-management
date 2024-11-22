package models

import (
	"time"
)

// Stock represents inventory for products
type Stock struct {
	ID          uint      `json:"id" gorm:"primary_key;auto_increment"`       // Auto-incrementing integer ID
	ProductID   uint      `json:"product_id" gorm:"not null"`                 // Reference to the Product
	Quantity    int       `json:"quantity" gorm:"not null"`                   // Available quantity
	Location    string    `json:"location" gorm:"type:varchar(100);not null"` // Location of stock
	Description string    `json:"description" gorm:"type:text"`               // Additional information about the stock
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`           // Record creation timestamp
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`           // Record last update timestamp
}
