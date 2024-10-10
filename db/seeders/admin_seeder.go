package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateAdmin() {

	admins := []models.Admin{
		{
			AdminID:  1,
			Email:    "testadminone@gmail.com",
			Name:     "John Admin",
			Password: "*new",
		},
		{
			AdminID:  2,
			Email:    "testadmintwo@gmail.com",
			Name:     "Sarah Admin",
			Password: "*new2",
		},
		{
			AdminID:  3,
			Email:    "testadminthree@gmail.com",
			Name:     "Jake Admin",
			Password: "*new3",
		},
	}

	for _, admin := range admins {
		err := s.database.Where("admin_id = ?", admin.AdminID).FirstOrCreate(&admin).Error
		if err != nil {
			log.Printf("Error: Could not create an admin %s: %v", admin.Email, err.Error())
		}
	}
}
