package requests

type OrderRequest struct {
	OrderReference string  `json:"order_reference"`
	OrderAmount    int     `json:"order_amount"`
	ProductID      int     `json:"product_id"`
	ProductPrice   float64 `json:"product_price"`
}
