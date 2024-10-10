package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type AdminInterface interface {
	Create(admin *models.Admin) error
	Get(id interface{}) (*models.Admin, error)
	GetAll() ([]*models.Admin, error)
	Update(ADMIN *models.Admin) error
	Delete(id interface{}) error
}

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (repo *AdminRepository) Create(admin *models.Admin) error {
	return repo.DB.Create(admin).Error
}

func (repo *AdminRepository) Get(id interface{}) (*models.Admin, error) {
	var admin models.Admin
	if err := repo.DB.Where("admin_id = ?", id).First(&admin).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find admin with id %v", id))
	}
	return &admin, nil
}

func (repo *AdminRepository) GetAll() ([]*models.Admin, error) {
	var admins []*models.Admin
	if err := repo.DB.Find(&admins).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find admins"))
	}
	return admins, nil
}

func (repo *AdminRepository) Update(admin *models.Admin) error {
	return repo.DB.Save(admin).Error
}

func (repo *AdminRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find admin with id %v", id))
	}
	return repo.DB.Delete(models.Admin{}, "admin_id = ?", id).Error
}
