package services

import (
	"ComputerWorld_API/api/handlers"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type AdminService struct {
	Db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{Db: db}
}

func (s *AdminService) Create(request *requests.AdminRequest, user *models.Admin) error {
	user.Email = request.Email
	user.Name = request.Name

	return s.Db.Create(&user).Error

}

func (s *AdminService) Update(admin *models.Admin) error {
	return s.Db.Save(admin).Error
}

func (s *AdminService) UpdatePassword(admin *models.Admin, newPassword string) error {
	salt, err := handlers.CreateSalt(10)
	if err != nil {
		return err
	}
	encryptedPassword := handlers.HashPassword(newPassword, salt)

	admin.Password = encryptedPassword
	return s.Db.Save(admin).Error
}

func (s *AdminService) Delete(admin *models.Admin) error {
	return s.Db.Delete(admin).Error
}
