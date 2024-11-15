package services

import (
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type ManufacturerService struct {
	Db *gorm.DB
}

func NewManufacturerService(db *gorm.DB) *ManufacturerService {
	return &ManufacturerService{Db: db}
}

func (s *ManufacturerService) Create(request *requests.CreateManufacturerRequest, manufacturer *models.Manufacturer) error {
	manufacturer.ManufacturerName = request.ManufacturerName

	return s.Db.Create(&manufacturer).Error
}

func (s *ManufacturerService) Update(manufacturer *models.Manufacturer) error {
	return s.Db.Save(manufacturer).Error
}

func (s *ManufacturerService) Delete(manufacturer *models.Manufacturer) error {
	return s.Db.Delete(manufacturer).Error
}
