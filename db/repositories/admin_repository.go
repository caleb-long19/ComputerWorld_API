package repositories

import (
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (repo *AdminRepository) Get(id interface{}) *models.Admin {
	admin := &models.Admin{}
	repo.DB.Where("ID = ?", id).Find(admin)
	return admin
}

func (repo *AdminRepository) GetByAdminId(admin *models.Admin, adminId interface{}) {
	repo.DB.Where("id = ?", adminId).Find(admin)
}

func (repo *AdminRepository) List(admins *[]models.Admin) {
	repo.DB.Model(&models.Admin{}).Find(admins)
	return
}
