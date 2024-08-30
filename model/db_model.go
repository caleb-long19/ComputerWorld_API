package model

type Product struct {
	ProductID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductCode    string `gorm:"unique" json:"product_code"`
	ProductName    string `gorm:"unique" json:"product_name"`
	ManufacturerID int
	Manufacturer   Manufacturer
	Stock          int     `json:"stock"`
	Price          float64 `gorm:"not null" json:"price"`
}

type Manufacturer struct {
	// Variables that will store user input product data
	ID               int    `gorm:"primaryKey;autoIncrement" json:"manufacturer_id"`
	ManufacturerName string `gorm:"unique" json:"manufacturer_name"`
}

type Order struct {
	OrderID     int    `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderRef    string `gorm:"unique;autoIncrement" json:"order_ref"`
	ProductID   int
	Product     Product
	OrderAmount int     `json:"order_amount"`
	OrderCost   float64 `json:"order_cost"`
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
