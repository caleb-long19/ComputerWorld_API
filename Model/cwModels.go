package Model

type Products struct {
	ProductID      uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code           string       `gorm:"unique" json:"code"`
	Name           string       `gorm:"unique" json:"name"`
	ManufacturerID int          `gorm:"foreignKey:ManufacturerID,constraint:OnUpdate:CASCADE,OnDelete:Set Null" json:"manufacturer_id"`
	Manufacturer   Manufacturer `json:"manufacturer"`
	Price          float64      `gorm:"not null" json:"price"`
}

type ProductInformation struct {
	// Variables that will store user input product data
	ProductCode    string
	ProductName    string
	ManufacturerID int
	Manufacturer   string
	ProductPrice   float64
}

type ProductStock struct {
	StockID   uint64 `gorm:"primaryKey;autoIncrement" json:"stock_id"`
	ProductID int    `json:"product_id"`
	Product   Products
	Stock     uint64 `json:"stock"`
}

type Manufacturer struct {
	// Variables that will store user input product data
	ManufacturerID   uint   `gorm:"primaryKey;autoIncrement" json:"manufacturer_id"`
	ManufacturerName string `json:"manufacturer_name"`
}

type ManufacturerInformation struct {
	ManufacturerID   string
	ManufacturerName string
}
