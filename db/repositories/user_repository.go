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

func (r *UserRepository) Get(id interface{}) *models.User {
	user := &models.User{}
	r.DB.Where("ID = ?", id).Find(user)
	return user
}

func (r *UserRepository) GetUserByUID(user *models.User, uid string) {
	r.DB.Where("uid = ?", uid).Take(user)
}

func (r *UserRepository) List(users *[]models.User) {
	r.DB.Model(&models.User{}).Find(users)
	return
}
