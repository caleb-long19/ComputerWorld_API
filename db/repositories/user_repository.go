package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserInterface interface {
	Create(user *models.User) error
	Get(id interface{}) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id interface{}) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) Get(id interface{}) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find user with id %v", id))
	}
	return &user, nil
}

func (repo *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find users"))
	}
	return users, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find user with id %v", id))
	}
	return repo.DB.Delete(models.User{}, "user_id = ?", id).Error
}
