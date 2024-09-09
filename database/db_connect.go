package database

import (
	model2 "ComputerWorld_API/database/model"
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
	db.AutoMigrate(&model2.Manufacturer{})

	db.AutoMigrate(&model2.Product{})

	db.AutoMigrate(&model2.Order{})

	return db
}
