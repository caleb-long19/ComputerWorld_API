package factories

import (
	m "ComputerWorld_API/db/models"
	"ComputerWorld_API/pkg/rand"
	"github.com/jinzhu/gorm"
	"log"
)

func NewProduct(db *gorm.DB, product *m.Product) {
	fillProductDefaults(product)

	err := db.Create(product).Error
	if err != nil {
		log.Printf("Error creating product in factory: %v", err)
	}
}

func NewProducts(db *gorm.DB, defaultProduct *m.Product, total int) []*m.Product {
	if total < 1 || total > 100 {
		log.Fatal("Count can only be in the range 1 - 100")
	}
	products := make([]*m.Product, total)
	for i := 0; i < total; i++ {
		products[i] = &m.Product{
			UID: "",
		}
		NewProduct(db, products[i])
	}

	return products
}

func fillProductDefaults(product *m.Product) {
	if product.ProductCode == "" {
		product.ProductCode = rand.String()
	}
	if product.ProductName == "" {
		product.ProductName = rand.String()
	}
	if product.ManufacturerUID == "" {
		product.ManufacturerUID = "MANU1tsypg"
	}
	if product.Stock == 0 {
		product.Stock = rand.Int()
	}
	if product.Price == 0 {
		product.Price = rand.Float()
	}
	if product.UID == "" {
		product.UID = rand.String()
	}
}
