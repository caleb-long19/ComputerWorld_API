package repositories

import (
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Get(uid interface{}) *models.Order {
	order := &models.Order{}
	r.DB.Where("uid = ?", uid).Find(order)
	return order
}

func (r *OrderRepository) GetOrderByUID(order *models.Order, uid string) {
	r.DB.Where("uid = ?", uid).Take(order)
}

func (r *OrderRepository) List(orders *[]models.Order) {
	r.DB.Model(&models.User{}).Find(orders)
}
