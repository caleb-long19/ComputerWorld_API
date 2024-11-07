package factories

import (
	m "ComputerWorld_API/db/models"
	"ComputerWorld_API/pkg/rand"
	"github.com/jinzhu/gorm"
	"log"
)

func NewUser(db *gorm.DB, user *m.User) {
	fillUserDefaults(user)

	err := db.Create(user).Error
	if err != nil {
		log.Printf("Error creating user in factory: %v", err)
	}
}

func NewUsers(db *gorm.DB, defaultUser *m.User, total int) []*m.User {
	if total < 1 || total > 100 {
		log.Fatal("Count can only be in the range 1 - 100")
	}
	users := make([]*m.User, total)
	for i := 0; i < total; i++ {
		users[i] = &m.User{
			UID:      "",
			Email:    defaultUser.Email,
			Name:     defaultUser.Name,
			Password: defaultUser.Password,
		}
		NewUser(db, users[i])
	}

	return users
}

func fillUserDefaults(user *m.User) {
	if user.Email == "" {
		user.Email = rand.Email()
	}
	if user.Name == "" {
		user.Name = rand.String()
	}
	if user.Password == "" {
		user.Password = rand.String()
	}
	if user.UID == "" {
		user.UID = rand.String()
	}
}
