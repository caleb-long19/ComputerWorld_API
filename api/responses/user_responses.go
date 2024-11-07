package responses

import (
	"ComputerWorld_API/db/models"
)

type UserResponse struct {
	Email    string `json:"email" example:"example@gmail.com"`
	Name     string `json:"name" example:"John Example"`
	Password string `json:"password" example:"!ExamplePass123"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}
}

type UsersResponse struct {
	Data []UserResponse `json:"data"`
}

func NewUsersResponse(users []models.User) *UsersResponse {

	usersData := make([]UserResponse, 0)
	for i := range users {
		usersData = append(usersData, UserResponse{
			Email:    users[i].Email,
			Name:     users[i].Name,
			Password: users[i].Password,
		})
	}

	return &UsersResponse{
		Data: usersData,
	}
}
