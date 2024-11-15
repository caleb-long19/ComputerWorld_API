package requests

type ProductRequest struct {
	ProductCode     string  `json:"product_code"`
	ProductName     string  `json:"product_name"`
	ManufacturerUID string  `json:"manufacturer_uid"`
	ProductStock    int     `json:"product_stock"`
	ProductPrice    float64 `json:"product_price"`
}

type CreateProductRequest struct {
	ProductCode     string  `json:"product_code" validate:"required,max=6" example:"HGN20F"`
	ProductName     string  `json:"product_name" validate:"required,max=50" example:"Xbox 360"`
	ManufacturerUID string  `json:"manufacturer_uid" validate:"required,max=100" example:"HJIGJNSJG"`
	ProductStock    int     `json:"product_stock" validate:"max=500" example:"25"`
	ProductPrice    float64 `json:"product_price" validate:"required" example:"250"`
}

type UpdateProductRequest struct {
	ProductCode     string  `json:"product_code" validate:"required,max=6" example:"HGN20F"`
	ProductName     string  `json:"product_name" validate:"required,max=50" example:"Xbox 360"`
	ManufacturerUID string  `json:"manufacturer_uid" validate:"required,max=100" example:"HJIGJNSJG"`
	ProductStock    int     `json:"product_stock" validate:"max=500" example:"25"`
	ProductPrice    float64 `json:"product_price" validate:"required" example:"250"`
}
