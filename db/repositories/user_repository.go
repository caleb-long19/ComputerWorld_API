package repositories

import (
	"ComputerWorld_API/db/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Get(id interface{}) *models.User {
	user := &models.User{}
	repo.DB.Where("ID = ?", id).Find(user)
	return user
}

func (r *UserRepository) GetUserByUID(user *models.User, uid string) {
	r.DB.Where("uid = ?", uid).Take(user)
}

func (repo *UserRepository) List(users *[]models.User) {
	repo.DB.Model(&models.User{}).Find(users)
	return
}
