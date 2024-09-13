package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type OrderInterface interface {
	Create(order *models.Order) error
	Get(id interface{}) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(id interface{}) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Create(order *models.Order) error {
	return repo.DB.Create(order).Error
}

func (repo *OrderRepository) Get(id interface{}) (*models.Order, error) {
	var order models.Order
	if err := repo.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find order with id %v", id))
	}
	return &order, nil
}

func (repo *OrderRepository) GetAll() ([]*models.Order, error) {
	var orders []*models.Order
	if err := repo.DB.Find(&orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find orders %v", orders))
	}
	return orders, nil
}

func (repo *OrderRepository) Update(order *models.Order) error {
	return repo.DB.Save(order).Error
}

func (repo *OrderRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return err
	}
	return repo.DB.Delete(models.Order{}, "order_id = ?", id).Error
}
