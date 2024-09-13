package repositories

import (
	"ComputerWorld_API/db/model"
	"gorm.io/gorm"
)

type OrderInterface interface {
	Create(order *model.Order) error
	// Get(id uint) (*model.Manufacturer, error)
	// Update(manufacturer *model.Manufacturer) error
	// Delete(id uint) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Create(order *model.Order) error {
	return repo.DB.Create(order).Error
}
