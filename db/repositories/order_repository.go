package repositories

import (
	"ComputerWorld_API/db/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderInterface interface {
	Create(order *model.Order) error
	Get(id interface{}) (*model.Order, error)
	Update(order *model.Order) error
	Delete(id interface{}) error
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

func (repo *OrderRepository) Get(id interface{}) (*model.Order, error) {
	var order model.Order
	if err := repo.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find order with id %v", id))
	}
	return &order, nil
}

func (repo *OrderRepository) Update(order *model.Order) error {
	return repo.DB.Save(order).Error
}

func (repo *OrderRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return err
	}
	return repo.DB.Delete(model.Order{}, "order_id = ?", id).Error
}
