package services

import (
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type ProductService struct {
	Db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{Db: db}
}

func (s *ProductService) Create(request *requests.ProductRequest, product *models.Product) error {
	product.ProductCode = request.ProductCode
	product.ProductName = request.ProductName
	product.ManufacturerUID = request.ManufacturerUID
	product.Stock = request.ProductStock
	product.Price = request.ProductPrice

	return s.Db.Create(&product).Error
}

func (s *ProductService) Update(product *models.Product) error {
	return s.Db.Save(product).Error
}

func (s *ProductService) Delete(product *models.Product) error {
	return s.Db.Delete(product).Error
}
