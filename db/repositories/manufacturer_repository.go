package repositories

import (
	"ComputerWorld_API/db/model"
	"gorm.io/gorm"
)

type ManufacturerInterface interface {
	Create(manufacturer *model.Manufacturer) error
	// Get(id uint) (*model.Manufacturer, error)
	// Update(manufacturer *model.Manufacturer) error
	// Delete(id uint) error
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

func (repo *ManufacturerRepository) Get(id int) (*model.Manufacturer, error) {
	manufacturer := new(model.Manufacturer)
	repo.DB.First(manufacturer, id)
	return manufacturer, nil
}

func (repo *ManufacturerRepository) Update(manufacturer *model.Manufacturer) error {
	return repo.DB.Save(manufacturer).Error
}

func (repo *ManufacturerRepository) Delete(id int) error {
	_, err := repo.Get(id)
	if err != nil {
		return err
	}
	return repo.DB.Delete(model.Manufacturer{}, "manufacturer_id = ?", id).Error
}
