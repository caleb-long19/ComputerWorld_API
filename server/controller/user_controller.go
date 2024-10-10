package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/requests"
	"ComputerWorld_API/server/responses"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	UserRepository repositories.UserInterface
}

func (uc *UserController) Create(c echo.Context) error {
	requestUser := new(requests.UserRequest)

	if err := c.Bind(&requestUser); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind user data"))
	}

	// Validate the request manufacturer data
	validatedUser, errV := ValidateUserRequest(requestUser)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	// Call repository method to create the new manufacturer
	err := uc.UserRepository.Create(validatedUser)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed to create user: %v", err))
	}

	return c.JSON(http.StatusCreated, validatedUser)
}

func (uc *UserController) Get(c echo.Context) error {
	user, err := uc.UserRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(c echo.Context) error {
	users, err := uc.UserRepository.GetAll()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) Update(c echo.Context) error {
	existingUser, err := uc.UserRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("user not found: %v", err))
	}

	var updateUser = new(requests.UserRequest)
	if err := c.Bind(updateUser); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind user data"))
	}
	if updateUser == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid user data"))
	}

	// Validate the request manufacturer data
	validatedExistingUser, errV := ValidateUserRequest(updateUser)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	existingUser.Email = validatedExistingUser.Email
	existingUser.Name = validatedExistingUser.Name
	existingUser.Password = validatedExistingUser.Password

	if err := uc.UserRepository.Update(existingUser); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update user: %v", err))
	}

	return c.JSON(http.StatusCreated, existingUser)
}

func (uc *UserController) Delete(c echo.Context) error {
	err := uc.UserRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "User successfully deleted")
}

// ValidateUserRequest validates the input request for creating or updating a manufacturer.
func ValidateUserRequest(request *requests.UserRequest) (*models.User, error) {
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
	if request.Password == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Password is required")
	}

	user.Email = request.Email
	user.Name = request.Name
	user.Password = request.Password

	err := requests.ValidateUserInputs(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
