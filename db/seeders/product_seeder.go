package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateProduct() {

	products := []models.Product{
		{
			UID:             "",
			ProductCode:     "XB403",
			ProductName:     "Xbox 360",
			ManufacturerUID: "",
			Stock:           55,
			Price:           100,
		},
		{
			UID:             "",
			ProductCode:     "PS48D",
			ProductName:     "Playstation 5",
			ManufacturerUID: "",
			Stock:           50,
			Price:           350,
		},
		{
			UID:             "",
			ProductCode:     "NS533",
			ProductName:     "Nintendo Switch",
			ManufacturerUID: "",
			Stock:           75,
			Price:           250,
		},
	}

	for _, product := range products {
		err := s.DB.Where("product_id = ?", product.UID).FirstOrCreate(&product).Error
		if err != nil {
			log.Printf("Error: could not create a new product %s: %v", product.ProductName, err.Error())
		}
	}
}
