package requests

type ManufacturerRequest struct {
	ManufacturerName string `json:"manufacturer_name"`
}

type ProductRequest struct {
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	ManufacturerID int     `json:"manufacturer_id"`
	ProductStock   int     `json:"product_stock"`
	ProductPrice   float64 `json:"product_price"`
}

type OrderRequest struct {
	OrderReference string  `json:"order_reference"`
	OrderAmount    int     `json:"order_amount"`
	ProductID      int     `json:"product_id"`
	ProductPrice   float64 `json:"product_price"`
}
