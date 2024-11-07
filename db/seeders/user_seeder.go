package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateUser() {

	users := []models.User{
		{
			UID:      "",
			Email:    "testuserone@gmail.com",
			Name:     "Jack User",
			Password: "*new",
		},
		{
			UID:      "",
			Email:    "testusertwo@gmail.com",
			Name:     "Blake User",
			Password: "*new2",
		},
		{
			UID:      "",
			Email:    "testuserthree@gmail.com",
			Name:     "Jane User",
			Password: "*new3",
		},
	}

	for _, user := range users {
		err := s.DB.Where("user_id = ?", user.UID).FirstOrCreate(&user).Error
		if err != nil {
			log.Printf("Error: Could not create a user %s: %v", user.Email, err.Error())
		}
	}
}
