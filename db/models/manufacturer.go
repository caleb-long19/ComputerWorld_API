package models

type Manufacturer struct {
	ManufacturerID   int    `gorm:"autoIncrement;primaryKey" json:"manufacturer_id"`
	ManufacturerName string `gorm:"unique" json:"manufacturer_name"`
}
