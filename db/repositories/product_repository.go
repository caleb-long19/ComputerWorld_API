package repositories

import (
	"ComputerWorld_API/db/model"
	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(product *model.Product) error
	// Get(id uint) (*model.Manufacturer, error)
	// Update(manufacturer *model.Manufacturer) error
	// Delete(id uint) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *model.Product) error {
	return repo.DB.Create(product).Error
}
