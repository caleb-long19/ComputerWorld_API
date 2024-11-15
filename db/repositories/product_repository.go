package repositories

import (
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Get(uid interface{}) *models.Product {
	product := &models.Product{}
	r.DB.Where("uid = ?", uid).Find(product)
	return product
}

func (r *ProductRepository) GetProductByUID(product *models.Product, uid string) {
	r.DB.Where("uid = ?", uid).Take(product)
}

func (r *ProductRepository) List(products *[]models.Product) {
	r.DB.Model(&models.User{}).Find(products)
}
