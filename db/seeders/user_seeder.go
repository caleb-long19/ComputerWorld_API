package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateUser() {

	users := []models.User{
		{
			UserID:   1,
			Email:    "testuserone@gmail.com",
			Name:     "Jack User",
			Password: "*new",
		},
		{
			UserID:   2,
			Email:    "testusertwo@gmail.com",
			Name:     "Blake User",
			Password: "*new2",
		},
		{
			UserID:   3,
			Email:    "testuserthree@gmail.com",
			Name:     "Jane User",
			Password: "*new3",
		},
	}

	for _, user := range users {
		err := s.database.Where("user_id = ?", user.UserID).FirstOrCreate(&user).Error
		if err != nil {
			log.Printf("Error: Could not create a user %s: %v", user.Email, err.Error())
		}
	}
}
