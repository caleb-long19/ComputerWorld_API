package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProductInterface interface {
	Create(product *models.Product) error
	Get(id interface{}) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id interface{}) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *models.Product) error {
	return repo.DB.Create(product).Error
}

func (repo *ProductRepository) Get(id interface{}) (*models.Product, error) {
	var product models.Product
	if err := repo.DB.Where("product_id", id).First(&product, id).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	if err := repo.DB.Find(&products).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find all products"))
	}
	return products, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	return repo.DB.Save(product).Error
}

func (repo *ProductRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return repo.DB.Delete(models.Product{}, "product_id = ?", id).Error
}
