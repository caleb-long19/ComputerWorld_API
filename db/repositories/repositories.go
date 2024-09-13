package repositories

import "gorm.io/gorm"

type Repository struct {
	DB           *gorm.DB
	Manufacturer *ManufacturerRepository
	Product      *ProductRepository
	Order        *OrderRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB:           db,
		Manufacturer: NewManufacturerRepository(db),
		Product:      NewProductRepository(db),
		Order:        NewOrderRepository(db),
	}
}
