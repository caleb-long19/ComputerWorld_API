package CW_Database

import (
	"ComputerWorld_API/Model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database Connection
var DatabaseCN = DatabaseConnection("CW_Database/Computer_world.db")

// DatabaseConnection Opens the database connection
func DatabaseConnection(databaseName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Model.Manufacturer{})

	db.AutoMigrate(&Model.Product{})

	db.AutoMigrate(&Model.Order{})

	return db
}
