package database

import (
	"ComputerWorld_API/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database Connection
var DatabaseCN = DatabaseConnection("computer_world.db")

// DatabaseConnection Opens the database connection
func DatabaseConnection(dbFilePath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	println("Database Name: ", dbFilePath)

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Manufacturer{})

	db.AutoMigrate(&model.Product{})

	db.AutoMigrate(&model.Order{})

	return db
}
