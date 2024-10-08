package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateProduct() {

	products := []models.Product{
		{
			ProductID:      1,
			ProductCode:    "XB403",
			ProductName:    "Xbox 360",
			ManufacturerID: 1,
			Stock:          55,
			Price:          100,
		},
		{
			ProductID:      2,
			ProductCode:    "PS48D",
			ProductName:    "Playstation 5",
			ManufacturerID: 2,
			Stock:          50,
			Price:          350,
		},
		{
			ProductID:      3,
			ProductCode:    "NS533",
			ProductName:    "Nintendo Switch",
			ManufacturerID: 3,
			Stock:          75,
			Price:          250,
		},
	}

	for _, product := range products {
		err := s.database.Where("product_id = ?", product.ProductID).FirstOrCreate(&product).Error
		if err != nil {
			log.Printf("Error: could not create a new product %s: %v", product.ProductName, err.Error())
		}
	}
}
