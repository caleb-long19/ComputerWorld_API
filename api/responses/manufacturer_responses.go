package responses

import (
	"ComputerWorld_API/db/models"
)

type ManufacturerResponse struct {
	ManufacturerName string `json:"manufacturer_name" example:"Microsoft"`
}

func NewManufacturerResponse(manufacturer *models.Manufacturer) *ManufacturerResponse {
	return &ManufacturerResponse{
		ManufacturerName: manufacturer.ManufacturerName,
	}
}

type ManufacturersResponse struct {
	Data []ManufacturerResponse `json:"data"`
}

func NewManufacturersResponse(manufacturers []models.Manufacturer) *ManufacturersResponse {

	manufacturersData := make([]ManufacturerResponse, 0)
	for i := range manufacturers {
		manufacturersData = append(manufacturersData, ManufacturerResponse{
			ManufacturerName: manufacturers[i].ManufacturerName,
		})
	}

	return &ManufacturersResponse{
		Data: manufacturersData,
	}
}
