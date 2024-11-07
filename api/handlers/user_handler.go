package handlers

import (
	s "ComputerWorld_API/api"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/api/responses"
	"ComputerWorld_API/api/services"
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	server  *s.Server
	service *services.UserService
}

func NewUserHandler(server *s.Server) *UserHandler {
	uh := &UserHandler{server: server}
	uh.service = services.NewUserService(server.Db)
	return uh
}

func (h *UserHandler) Create(c echo.Context) error {
	createRequest := new(requests.CreateUserRequest)
	if err := c.Bind(&createRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind user data"))
	}
	_, errV := ValidateUserRequest(createRequest)
	if errV != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, errV)
	}

	user := &models.User{}
	if err := h.service.Create(createRequest, user); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to store user in the database: %v", err))
	}

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *UserHandler) Update(c echo.Context) error {
	updateUserRequest := new(requests.UpdateUserRequest)
	uid := c.Param("uid")

	if err := c.Bind(&updateUserRequest); err != nil {
		return err
	}

	user := models.User{}
	h.server.Repos.User.GetUserByUID(&user, uid)
	if user.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "User not found")
	}

	// Validate the request user data
	//_, errV := ValidateUserRequest(updateUserRequest)
	//if errV != nil {
	//	// Return the validation error directly
	//	return responses.ErrorResponse(c, 0, errV)
	//}

	user.Email = updateUserRequest.Email
	user.Name = updateUserRequest.Name
	if err := h.service.Update(&user); err != nil {
		return responses.ErrResponse(c, http.StatusInternalServerError, "failed to update user")
	}

	return responses.MessageResponse(c, http.StatusOK, "User successfully updated")
}

func (h *UserHandler) Get(c echo.Context) error {
	uid := c.Param("uid")

	user := &models.User{}
	h.server.Repos.User.GetUserByUID(user, uid)
	if user.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "User not found")
	}

	response := responses.NewUserResponse(user)
	return responses.Response(c, http.StatusOK, response)
}

func (h *UserHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	user := models.User{}
	h.server.Repos.User.GetUserByUID(&user, uid)
	if user.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "User not found")
	}

	if err := h.service.Delete(&user); err != nil {
		return responses.ErrResponse(c, http.StatusInternalServerError, "Something went wrong deleting the user from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "User successfully deleted")
}

// ValidateUserRequest validates the input request for creating or updating a manufacturer.
func ValidateUserRequest(request *requests.CreateUserRequest) (*models.User, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	user := new(models.User)
	if request.Email == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "email is required")
	}
	if len(request.Email) < 1 || len(request.Email) > 200 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "email must be between 1 and 200 characters")
	}
	if request.Name == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Name is required")
	}
	if len(request.Name) < 1 || len(request.Email) > 50 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Name must be between 1 and 50 characters")
	}

	user.Email = request.Email
	user.Name = request.Name

	if err := requests.ValidateUserInputs(user); err != nil {
		return nil, err
	}

	return user, nil
}
