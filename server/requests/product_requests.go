package requests

type ProductRequest struct {
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	ManufacturerID int     `json:"manufacturer_id"`
	ProductStock   int     `json:"product_stock"`
	ProductPrice   float64 `json:"product_price"`
}
