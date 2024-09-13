package repositories

import (
	"ComputerWorld_API/db/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(product *model.Product) error
	Get(id interface{}) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id interface{}) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *model.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *ProductRepository) Get(id interface{}) (*model.Product, error) {
	var product model.Product
	if err := repo.DB.Where("product_id", id).First(&product, id).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return &product, nil
}

func (repo *ProductRepository) Update(product *model.Product) error {
	return repo.DB.Save(product).Error
}

func (repo *ProductRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return repo.DB.Delete(model.Product{}, "product_id = ?", id).Error
}
