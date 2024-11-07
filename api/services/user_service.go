package services

import (
	"ComputerWorld_API/api/handlers"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{Db: db}
}

func (s *UserService) Create(request *requests.CreateUserRequest, user *models.User) error {
	user.Email = request.Email
	user.Name = request.Name

	return s.Db.Create(&user).Error
}

func (s *UserService) Update(user *models.User) error {
	return s.Db.Save(user).Error
}

func (s *UserService) UpdatePassword(user *models.User, newPassword string) error {
	salt, err := handlers.CreateSalt(10)
	if err != nil {
		return err
	}
	encryptedPassword := handlers.HashPassword(newPassword, salt)

	user.Password = encryptedPassword
	return s.Db.Save(user).Error
}

func (s *UserService) Delete(user *models.User) error {
	return s.Db.Delete(user).Error
}
