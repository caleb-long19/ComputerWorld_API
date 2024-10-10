package db

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/seeders"
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
	db.AutoMigrate(&models.Manufacturer{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Admin{})

	seeding := seeders.NewSeed(db)
	seeding.CreateManufacturer()
	seeding.CreateProduct()
	seeding.CreateOrder()
	seeding.CreateUser()
	seeding.CreateAdmin()

	return db
}
