package responses

import (
	"ComputerWorld_API/db/models"
)

type OrderResponse struct {
	OrderRef    string  `json:"order_ref" example:"XBF834NF"`
	OrderAmount int     `json:"product_name" example:"5"`
	UID         string  `json:"uid" example:"4758ad4a-73ea-4d91-b6bb-eca1fd12f015"`
	OrderPrice  float64 `json:"price" example:"500"`
}

func NewOrderResponse(product *models.Order) *OrderResponse {
	return &OrderResponse{
		OrderRef:    product.OrderRef,
		OrderAmount: product.OrderAmount,
		UID:         product.UID,
		OrderPrice:  product.OrderPrice,
	}
}

type OrdersResponse struct {
	Data []OrderResponse `json:"data"`
}

func NewOrdersResponse(orders []models.Order) *OrdersResponse {

	ordersData := make([]OrderResponse, 0)
	for i := range orders {
		ordersData = append(ordersData, OrderResponse{
			OrderRef:    orders[i].OrderRef,
			OrderAmount: orders[i].OrderAmount,
			UID:         orders[i].UID,
			OrderPrice:  orders[i].OrderPrice,
		})
	}

	return &OrdersResponse{
		Data: ordersData,
	}
}
