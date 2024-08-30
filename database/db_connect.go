package database

import (
	"ComputerWorld_API/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database Connection
var DatabaseCN = DatabaseConnection("database/computer_world.db")

// DatabaseConnection Opens the database connection
func DatabaseConnection(databaseName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Manufacturer{})

	db.AutoMigrate(&model.Product{})

	db.AutoMigrate(&model.Order{})

	return db
}
