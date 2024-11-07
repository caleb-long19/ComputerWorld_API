package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateAdmin() {

	admins := []models.Admin{
		{
			UID:      "",
			Email:    "testadminone@gmail.com",
			Name:     "John Admin",
			Password: "*new",
		},
		{
			UID:      "",
			Email:    "testadmintwo@gmail.com",
			Name:     "Sarah Admin",
			Password: "*new2",
		},
		{
			UID:      "",
			Email:    "testadminthree@gmail.com",
			Name:     "Jake Admin",
			Password: "*new3",
		},
	}

	for _, admin := range admins {
		err := s.DB.Where("admin_id = ?", admin.UID).FirstOrCreate(&admin).Error
		if err != nil {
			log.Printf("Error: Could not create an admin %s: %v", admin.Email, err.Error())
		}
	}
}
