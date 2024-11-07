package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	DB           *gorm.DB
	Manufacturer *ManufacturerRepository
	Product      *ProductRepository
	Order        *OrderRepository
	User         *UserRepository
	Admin        *AdminRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB:           db,
		Manufacturer: NewManufacturerRepository(db),
		Product:      NewProductRepository(db),
		Order:        NewOrderRepository(db),
		User:         NewUserRepository(db),
		Admin:        NewAdminRepository(db),
	}
}
