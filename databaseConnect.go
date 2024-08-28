package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Opens the database connection
func databaseConnection(databaseName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&StoredProduct{})
	db.AutoMigrate(&EmployeeData{})

	return db
}
