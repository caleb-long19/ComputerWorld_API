package seeders

import (
	"ComputerWorld_API/db/models"
	"log"
)

func (s *Seeding) CreateManufacturer() {

	manufacturers := []models.Manufacturer{
		{
			UID:              "",
			ManufacturerName: "Microsoft",
		},
		{
			UID:              "",
			ManufacturerName: "Sony",
		},
		{
			UID:              "",
			ManufacturerName: "Nintendo",
		},
	}

	for _, manufacturer := range manufacturers {
		err := s.DB.Where("manufacturer_id = ?", manufacturer.UID).FirstOrCreate(&manufacturer).Error
		if err != nil {
			log.Printf("Error: Could not create a manufacturer %s: %v", manufacturer.ManufacturerName, err.Error())
		}
	}
}
