package requests

type OrderRequest struct {
	OrderReference string  `json:"order_ref"`
	OrderAmount    int     `json:"order_amount"`
	ProductID      int     `json:"product_id"`
	OrderPrice     float64 `json:"order_price"`
}
