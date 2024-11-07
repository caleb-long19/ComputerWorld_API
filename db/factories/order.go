package factories

import (
	m "ComputerWorld_API/db/models"
	"ComputerWorld_API/pkg/rand"
	"github.com/jinzhu/gorm"
	"log"
)

func NewOrder(db *gorm.DB, order *m.Order) {
	fillOrderDefaults(order)

	err := db.Create(order).Error
	if err != nil {
		log.Printf("Error creating order in factory: %v", err)
	}
}

func NewOrders(db *gorm.DB, defaultOrder *m.Order, total int) []*m.Order {
	if total < 1 || total > 100 {
		log.Fatal("Count can only be in the range 1 - 100")
	}
	orders := make([]*m.Order, total)
	for i := 0; i < total; i++ {
		orders[i] = &m.Order{
			UID: "",
		}
		NewOrder(db, orders[i])
	}

	return orders
}

func fillOrderDefaults(order *m.Order) {
	if order.OrderRef == "" {
		order.OrderRef = rand.String()
	}
	if order.OrderAmount == 0 {
		order.OrderAmount = rand.Int()
	}
	if order.ProductUID == "" {
		order.ProductUID = "PRODhjf8f"
	}
	if order.OrderPrice == 0 {
		order.OrderPrice = rand.Float()
	}
	if order.UID == "" {
		order.UID = rand.String()
	}
}
