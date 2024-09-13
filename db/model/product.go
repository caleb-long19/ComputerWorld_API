package model

type Product struct {
	ProductID      int    `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductCode    string `gorm:"unique" json:"product_code"`
	ProductName    string `gorm:"unique" json:"product_name"`
	ManufacturerID int    `json:"manufacturer_id"`
	Manufacturer   Manufacturer
	Stock          int     `json:"product_stock"`
	Price          float64 `gorm:"not null" json:"product_price"`
}
