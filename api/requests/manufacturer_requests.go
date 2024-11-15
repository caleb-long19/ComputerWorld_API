package requests

type ManufacturerRequest struct {
	ManufacturerName string `json:"manufacturer_name"`
}

type CreateManufacturerRequest struct {
	ManufacturerName string `json:"manufacturer_name" validate:"required,max=30" example:"microsoft"`
}

type UpdateManufacturerRequest struct {
	ManufacturerName string `json:"manufacturer_name" validate:"required,max=30" example:"microsoft"`
}
