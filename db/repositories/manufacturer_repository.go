package repositories

import (
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type ManufacturerRepository struct {
	DB *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{DB: db}
}

func (r *ManufacturerRepository) Get(uid interface{}) *models.Manufacturer {
	manufacturer := &models.Manufacturer{}
	r.DB.Where("uid = ?", uid).Find(manufacturer)
	return manufacturer
}

func (r *ManufacturerRepository) GetManufacturerByUID(manufacturer *models.Manufacturer, uid string) {
	r.DB.Where("uid = ?", uid).Take(manufacturer)
}

func (r *ManufacturerRepository) List(manufacturers *[]models.Manufacturer) {
	r.DB.Model(&models.Manufacturer{}).Find(manufacturers)
}
