package responses

import (
	"ComputerWorld_API/db/models"
)

type AdminResponse struct {
	Email    string `json:"email" example:"example@gmail.com"`
	Name     string `json:"name" example:"John Example"`
	Password string `json:"password" example:"!ExamplePass123"`
}

func NewAdminResponse(admin *models.Admin) *AdminResponse {
	return &AdminResponse{
		Email:    admin.Email,
		Name:     admin.Name,
		Password: admin.Password,
	}
}

type AdminsResponse struct {
	Data []AdminResponse `json:"data"`
}

func NewAdminsResponse(admins []models.Admin) *AdminsResponse {

	adminsData := make([]AdminResponse, 0)
	for i := range admins {
		adminsData = append(adminsData, AdminResponse{
			Email:    admins[i].Email,
			Name:     admins[i].Name,
			Password: admins[i].Password,
		})
	}

	return &AdminsResponse{
		Data: adminsData,
	}
}
