package Model

type Product struct {
	ProductID        int          `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductCode      string       `gorm:"unique" json:"code"`
	ProductName      string       `gorm:"unique" json:"name"`
	ManufacturerID   int          `gorm:"foreignKey:ManufacturerID, constraint:OnUpdate:CASCADE,OnDelete:Set Null" json:"manufacturer_id"`
	ManufacturerName string       `gorm:"references:ManufacturerName"`
	Manufacturer     Manufacturer `json:"manufacturer"`
	Stock            int          `json:"stock"`
	Price            float64      `gorm:"not null" json:"price"`
}

type Manufacturer struct {
	// Variables that will store user input product data
	ManufacturerID   int    `gorm:"primaryKey;autoIncrement" json:"manufacturer_id"`
	ManufacturerName string `json:"manufacturer_name"`
}

type Order struct {
	OrderID     int     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderRef    string  `gorm:"unique;autoIncrement" json:"code"`
	ProductID   int     `gorm:"foreignKey:ProductID, constraint:OnUpdate:CASCADE,OnDelete:Set Null" json:"product_id"`
	ProductName string  `gorm:"references:ProductName"`
	Product     Product `json:"product"`
	OrderAmount int     `json:"order_amount"`
	OrderCost   float64 `json:"order_cost"`
}

// Ignore this for now
type ProductInformation struct {
	// Variables that will store user input product data
	ProductCode    string
	ProductName    string
	ManufacturerID int
	Manufacturer   string
	ProductPrice   float64
}

/*
type ManufacturerInformation struct {
	ManufacturerID   string
	ManufacturerName string
}
*/
