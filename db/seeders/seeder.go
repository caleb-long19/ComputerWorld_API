package seeders

import "github.com/jinzhu/gorm"

type Seeding struct {
	DB *gorm.DB
}

func NewSeed(db *gorm.DB) *Seeding {
	return &Seeding{DB: db}
}
