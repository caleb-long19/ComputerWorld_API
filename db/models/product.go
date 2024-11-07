package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	UID             string         `gorm:"primaryKey" json:"uid"`
	ProductCode     string         `gorm:"unique" json:"product_code"`
	ProductName     string         `gorm:"unique" json:"product_name"`
	ManufacturerUID string         `json:"manufacturer_uid"`
	Stock           int            `json:"product_stock"`
	Price           float64        `gorm:"not null" json:"product_price"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`

	Manufacturer Manufacturer
}

func (Product) TableName() string {
	return "products"
}

func (t *Product) BeforeCreate(tx *gorm.DB) error {
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
