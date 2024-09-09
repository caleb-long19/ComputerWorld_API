package seeding

import "gorm.io/gorm"

type Seeding struct {
	database *gorm.DB
}

func NewSeed(db *gorm.DB) *Seeding {
	return &Seeding{database: db}
}
