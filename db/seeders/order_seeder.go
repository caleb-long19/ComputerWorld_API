package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateOrder() {

	orders := []models.Order{
		{
			UID:         "",
			OrderAmount: 5,
			OrderRef:    "JOLANDO4",
			ProductUID:  "",
			OrderPrice:  500,
		},
		{
			UID:         "",
			OrderAmount: 5,
			OrderRef:    "DH4OJ4",
			ProductUID:  "",
			OrderPrice:  1750,
		},
		{
			UID:         "",
			OrderAmount: 5,
			OrderRef:    "KAUFMAN8",
			ProductUID:  "",
			OrderPrice:  1250,
		},
	}

	for _, order := range orders {
		err := s.DB.Where("order_id = ?", order.UID).FirstOrCreate(&order).Error
		if err != nil {
			log.Printf("Error: Could not create a new order %s: %v", order.OrderRef, err.Error())
		}
	}
}
