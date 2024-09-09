package model

type Order struct {
	OrderID      int     `gorm:"primaryKey;autoIncrement" json:"order_id"`
	OrderRef     string  `gorm:"autoIncrement" json:"order_ref"`
	OrderAmount  int     `json:"order_amount"`
	ProductID    int     `json:"product_id"`
	ProductPrice float64 `json:"price"`
	Product      Product
}
