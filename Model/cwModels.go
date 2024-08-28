package Model

// StoredProduct Model
type StoredProduct struct {
	ID    uint    `gorm:"primaryKey;autoIncrement" json:"ID"`
	Code  string  `gorm:"unique" json:"Code"`
	Name  string  `gorm:"unique" json:"Name"`
	Price float64 `gorm:"not null" json:"Price"`
}

type ProductInformation struct {
	// Variables that will store user input product data
	ProductCode  string
	ProductName  string
	ProductPrice float64
}

// EmployeeData Model
type EmployeeData struct {
	// Variables that will store user input product data
	EmployeeID   uint   `gorm:"primaryKey;autoIncrement" json:"Employee_ID"`
	EmployeeName string `json:"Employee_Name"`
	EmployeeRole string `json:"Employee_Role"`
}

type EmployeeInformation struct {
	EmployeeName string
	EmployeeRole string
}
