package requests

type OrderRequest struct {
	OrderReference string  `json:"order_ref"`
	OrderAmount    int     `json:"order_amount"`
	ProductUID     string  `json:"product_uid"`
	OrderPrice     float64 `json:"order_price"`
}

type CreateOrderRequest struct {
	OrderReference string  `json:"order_ref" validate:"required,max=6" example:"HGN20F"`
	OrderAmount    int     `json:"order_amount" validate:"required,max=100" example:"10"`
	ProductUID     string  `json:"product_uid" validate:"required,max=100" example:"HJIGJNSJG"`
	OrderPrice     float64 `json:"order_price" validate:"required,max=100" example:"10"`
}

type UpdateOrderRequest struct {
	OrderReference string  `json:"order_ref" validate:"required,max=6" example:"HGN20F"`
	OrderAmount    int     `json:"order_amount" validate:"required,max=100" example:"10"`
	ProductUID     string  `json:"product_uid" validate:"required,max=100" example:"HJIGJNSJG"`
	OrderPrice     float64 `json:"order_price" validate:"required,max=100" example:"10"`
}
