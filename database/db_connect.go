package database

import (
	model2 "ComputerWorld_API/database/model"
	"ComputerWorld_API/database/seeder"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

// DatabaseConnection Opens the database connection
func DatabaseConnection() *gorm.DB {

	dbFile := "computer_world.db"
	if os.Getenv("GO_ENV") == "test" {
		dbFile = "computer_world_test.db"
	}
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model2.Manufacturer{})
	db.AutoMigrate(&model2.Product{})
	db.AutoMigrate(&model2.Order{})

	seed := seeder.NewSeed(db)
	seed.CreateProduct()
	seed.CreateManufacturer()
	seed.CreateOrder()

	return db
}
