package repositories

import (
	"github.com/KrittayotToin/online-stock-management/internal/models"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uuid.UUID) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id uuid.UUID) error
	List(page, limit int) ([]models.Product, int64, error)
}
