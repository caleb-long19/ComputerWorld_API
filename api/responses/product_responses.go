package responses

import (
	"ComputerWorld_API/db/models"
)

type ProductResponse struct {
	ProductCode string  `json:"product_code" example:"XB3JFN"`
	ProductName string  `json:"product_name" example:"Xbox 360"`
	UID         string  `json:"uid" example:"4758ad4a-73ea-4d91-b6bb-eca1fd12f015"`
	Stock       int     `json:"stock" example:"1"`
	Price       float64 `json:"price" example:"250"`
}

func NewProductResponse(product *models.Product) *ProductResponse {
	return &ProductResponse{
		ProductCode: product.ProductCode,
		ProductName: product.ProductName,
		UID:         product.UID,
		Stock:       product.Stock,
		Price:       product.Price,
	}
}

type ProductsResponse struct {
	Data []ProductResponse `json:"data"`
}

func NewProductsResponse(products []models.Product) *ProductsResponse {

	productsData := make([]ProductResponse, 0)
	for i := range products {
		productsData = append(productsData, ProductResponse{
			ProductCode: products[i].ProductCode,
			ProductName: products[i].ProductName,
			UID:         products[i].UID,
			Stock:       products[i].Stock,
			Price:       products[i].Price,
		})
	}

	return &ProductsResponse{
		Data: productsData,
	}
}
