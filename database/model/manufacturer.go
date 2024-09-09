package model

type Manufacturer struct {
	ManufacturerID   int    `gorm:"primaryKey;autoIncrement" json:"manufacturer_id"`
	ManufacturerName string `gorm:"unique" json:"manufacturer_name"`
}
