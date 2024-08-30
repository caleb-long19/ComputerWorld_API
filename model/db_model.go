package model

type Product struct {
	ProductID      int    `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductCode    string `gorm:"unique" json:"product_code"`
	ProductName    string `gorm:"unique" json:"product_name"`
	ManufacturerID int    `json:"manufacturer_id"`
	Manufacturer   Manufacturer
	Stock          int     `json:"stock"`
	Price          float64 `gorm:"not null" json:"price"`
}

type Manufacturer struct {
	ManufacturerID   int    `gorm:"primaryKey;autoIncrement" json:"manufacturer_id"`
	ManufacturerName string `gorm:"unique" json:"manufacturer_name"`
}

type Order struct {
	OrderID      int     `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderRef     string  `gorm:"autoIncrement" json:"order_ref"`
	OrderAmount  int     `json:"order_amount"`
	ProductID    int     `json:"product_id"`
	ProductPrice float64 `json:"price"`
	Product      Product
}

// Ignore this for now
type ProductInformation struct {
	// Variables that will store user input product data
	ProductCode  string
	ProductName  string
	ProductPrice float64
}

/*
type ManufacturerInformation struct {
	ManufacturerID   string
	ManufacturerName string
}
*/
