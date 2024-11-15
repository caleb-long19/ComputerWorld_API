package services

import (
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type OrderService struct {
	Db *gorm.DB
}

func NewOrderPrice(db *gorm.DB) *OrderService {
	return &OrderService{Db: db}
}

func (s *OrderService) Create(request *requests.CreateOrderRequest, order *models.Order) error {
	order.OrderRef = request.OrderReference
	order.OrderAmount = request.OrderAmount
	order.ProductUID = request.ProductUID
	order.OrderPrice = request.OrderPrice

	return s.Db.Create(&order).Error
}

func (s *OrderService) Update(order *models.Order) error {
	return s.Db.Save(order).Error
}

func (s *OrderService) Delete(order *models.Order) error {
	return s.Db.Delete(order).Error
}
