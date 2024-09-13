package repositories

import (
	"ComputerWorld_API/db/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ManufacturerInterface interface {
	Create(manufacturer *model.Manufacturer) error
	Get(id interface{}) (*model.Manufacturer, error)
	Update(manufacturer *model.Manufacturer) error
	Delete(id interface{}) error
}

type ManufacturerRepository struct {
	DB *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{DB: db}
}

func (repo *ManufacturerRepository) Create(manufacturer *model.Manufacturer) error {
	return repo.DB.Create(manufacturer).Error
}

func (repo *ManufacturerRepository) Get(id interface{}) (*model.Manufacturer, error) {
	var manufacturer model.Manufacturer
	if err := repo.DB.Where("manufacturer_id = ?", id).First(&manufacturer).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return &manufacturer, nil
}

func (repo *ManufacturerRepository) Update(manufacturer *model.Manufacturer) error {
	return repo.DB.Save(manufacturer).Error
}

func (repo *ManufacturerRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return repo.DB.Delete(model.Manufacturer{}, "manufacturer_id = ?", id).Error
}
