package handlers

import (
	s "ComputerWorld_API/api"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/api/responses"
	"ComputerWorld_API/api/services"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	server   *s.Server
	userRepo *repositories.UserRepository
	service  *services.UserService
}

func NewUserHandler(server *s.Server) *UserHandler {
	uh := &UserHandler{server: server}
	uh.userRepo = repositories.NewUserRepository(server.Db)
	uh.service = services.NewUserService(server.Db)
	return uh
}

func (h *UserHandler) Create(c echo.Context) error {
	createUserRequest := new(requests.CreateUserRequest)
	if err := c.Bind(&createUserRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "could not bind user data")
	}
	if err := c.Validate(createUserRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	user := &models.User{}
	if err := h.service.Create(createUserRequest, user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to store user in the database: %v", err))
	}

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *UserHandler) Update(c echo.Context) error {
	updateUserRequest := new(requests.UpdateUserRequest)
	if err := c.Bind(&updateUserRequest); err != nil {
		return err
	}

	userID := c.Param("uid")

	user := h.userRepo.Get(userID)
	if user.UID == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User not found")
	}
	if err := c.Validate(updateUserRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	user.Email = updateUserRequest.Email
	user.Name = updateUserRequest.Name

	if err := h.service.Update(user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to update user")
	}

	return responses.MessageResponse(c, http.StatusOK, "User successfully updated")
}

func (h *UserHandler) Get(c echo.Context) error {
	uid := c.Param("uid")

	user := &models.User{}

	h.userRepo.GetUserByUID(user, uid)
	if user.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusOK, response)
}

func (h *UserHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	user := h.userRepo.Get(uid)

	if user.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "User not found")
	}
	if err := h.service.Delete(user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the user from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "User successfully deleted")
}
