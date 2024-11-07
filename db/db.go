package db

import (
	"ComputerWorld_API/db/seeders"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import the SQLite dialect
	"os"
)

// Init Opens the db connection
func Init() *gorm.DB {

	dbFile := "computer_world.db"
	if os.Getenv("GO_ENV") == "test" {
		dbFile = "computer_world_test.db"
	}
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		panic("failed to connect to database")
	}

	seeding := seeders.NewSeed(db)
	seeding.CreateManufacturer()
	seeding.CreateProduct()
	seeding.CreateOrder()
	seeding.CreateUser()
	seeding.CreateAdmin()

	return db
}
