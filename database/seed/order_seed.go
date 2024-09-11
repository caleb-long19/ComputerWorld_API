package seed

import (
	"ComputerWorld_API/database/model"
	"log"
)

func (s *Seeding) CreateOrder() {

	orders := []model.Order{
		{
			OrderID:      1,
			OrderAmount:  5,
			OrderRef:     "JOLANDO4",
			ProductID:    1,
			ProductPrice: 500,
		},
		{
			OrderID:      2,
			OrderAmount:  5,
			OrderRef:     "DH4OJ4",
			ProductID:    1,
			ProductPrice: 1750,
		},
		{
			OrderID:      3,
			OrderAmount:  5,
			OrderRef:     "KAUFMAN8",
			ProductID:    1,
			ProductPrice: 1250,
		},
	}

	for _, order := range orders {
		err := s.database.Where("order_id = ?", order.OrderID).FirstOrCreate(&order).Error
		if err != nil {
			log.Printf("Error: Could not create a new order %s: %v", order.OrderRef, err.Error())
		}
	}
}
