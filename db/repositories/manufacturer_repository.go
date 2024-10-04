package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ManufacturerInterface interface {
	Create(manufacturer *models.Manufacturer) error
	Get(id interface{}) (*models.Manufacturer, error)
	GetAll() ([]*models.Manufacturer, error)
	Update(manufacturer *models.Manufacturer) error
	Delete(id interface{}) error
}

type ManufacturerRepository struct {
	DB *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{DB: db}
}

func (repo *ManufacturerRepository) Create(manufacturer *models.Manufacturer) error {
	return repo.DB.Create(manufacturer).Error
}

func (repo *ManufacturerRepository) Get(id interface{}) (*models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	if err := repo.DB.Where("manufacturer_id = ?", id).First(&manufacturer).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return &manufacturer, nil
}

func (repo *ManufacturerRepository) GetAll() ([]*models.Manufacturer, error) {
	var manufacturers []*models.Manufacturer
	if err := repo.DB.Find(&manufacturers).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find manufacturers"))
	}
	return manufacturers, nil
}

func (repo *ManufacturerRepository) Update(manufacturer *models.Manufacturer) error {
	return repo.DB.Save(manufacturer).Error
}

func (repo *ManufacturerRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return repo.DB.Delete(models.Manufacturer{}, "manufacturer_id = ?", id).Error
}
