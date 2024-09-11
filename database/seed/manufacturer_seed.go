package seed

import (
	"ComputerWorld_API/database/model"
	"log"
)

func (s *Seeding) CreateManufacturer() {

	manufacturers := []model.Manufacturer{
		{
			ManufacturerID:   1,
			ManufacturerName: "Microsoft",
		},
		{
			ManufacturerID:   2,
			ManufacturerName: "Sony",
		},
		{
			ManufacturerID:   3,
			ManufacturerName: "Nintendo",
		},
	}

	for _, manufacturer := range manufacturers {
		err := s.database.Where("manufacturer_id = ?", manufacturer.ManufacturerID).FirstOrCreate(&manufacturer).Error
		if err != nil {
			log.Printf("Error: Could not create a manufacturer %s: %v", manufacturer.ManufacturerName, err.Error())
		}
	}
}
