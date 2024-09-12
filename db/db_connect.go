package db

import (
	"ComputerWorld_API/db/model"
	"ComputerWorld_API/db/seed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

// DatabaseConnection Opens the db connection
func DatabaseConnection() *gorm.DB {

	dbFile := "computer_world.db"
	if os.Getenv("GO_ENV") == "test" {
		dbFile = "computer_world_test.db"
	}
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		panic("failed to connect db")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Manufacturer{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Order{})

	seeding := seed.NewSeed(db)
	seeding.CreateManufacturer()
	seeding.CreateProduct()
	seeding.CreateOrder()

	return db
}
