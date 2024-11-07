package factories

import (
	m "ComputerWorld_API/db/models"
	"ComputerWorld_API/pkg/rand"
	"github.com/jinzhu/gorm"
	"log"
)

func NewManufacturer(db *gorm.DB, manufacturer *m.Manufacturer) {
	fillManufacturerDefaults(manufacturer)

	err := db.Create(manufacturer).Error
	if err != nil {
		log.Printf("Error creating manufacturer in factory: %v", err)
	}
}

func NewManufacturers(db *gorm.DB, defaultManufacturer *m.Manufacturer, total int) []*m.Manufacturer {
	if total < 1 || total > 100 {
		log.Fatal("Count can only be in the range 1 - 100")
	}
	manufacturers := make([]*m.Manufacturer, total)
	for i := 0; i < total; i++ {
		manufacturers[i] = &m.Manufacturer{
			UID:              "",
			ManufacturerName: defaultManufacturer.ManufacturerName,
		}
		NewManufacturer(db, manufacturers[i])
	}

	return manufacturers
}

func fillManufacturerDefaults(manufacturer *m.Manufacturer) {
	if manufacturer.ManufacturerName == "" {
		manufacturer.ManufacturerName = rand.String()
	}
	if manufacturer.UID == "" {
		manufacturer.UID = rand.String()
	}
}
