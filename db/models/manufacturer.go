package models

import (
	"github.com/google/uuid"
	"time"
)
import "gorm.io/gorm"

type Manufacturer struct {
	UID              string         `gorm:"primaryKey" json:"uid"`
	ManufacturerName string         `gorm:"unique" json:"manufacturer_name"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        gorm.DeletedAt `json:"-"`
}

func (Manufacturer) TableName() string {
	return "manufacturers"
}

func (t *Manufacturer) BeforeCreate(tx *gorm.DB) error {
	// If the UID is already set then just return.
	if t.UID != "" {
		return nil
	}
	newUuid, err := uuid.NewV7()
	if err != nil {
		return err
	}
	t.UID = newUuid.String()
	return nil
}
