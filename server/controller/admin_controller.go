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

type AdminController struct {
	AdminRepository repositories.AdminInterface
}

func (ac *AdminController) Create(c echo.Context) error {
	requestAdmin := new(requests.AdminRequest)

	if err := c.Bind(&requestAdmin); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind admin data"))
	}

	// Validate the request manufacturer data
	validatedAdmin, errV := ValidateAdminRequest(requestAdmin)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	// Call repository method to create the new manufacturer
	err := ac.AdminRepository.Create(validatedAdmin)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed to create admin: %v", err))
	}

	return c.JSON(http.StatusCreated, validatedAdmin)
}

func (ac *AdminController) Get(c echo.Context) error {
	admin, err := ac.AdminRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, admin)
}

func (ac *AdminController) GetAll(c echo.Context) error {
	admins, err := ac.AdminRepository.GetAll()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, admins)
}

func (ac *AdminController) Update(c echo.Context) error {
	existingAdmin, err := ac.AdminRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("admin not found: %v", err))
	}

	var updateAdmin = new(requests.AdminRequest)
	if err := c.Bind(updateAdmin); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind admin data"))
	}
	if updateAdmin == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid admin data"))
	}

	// Validate the request manufacturer data
	validatedExistingAdmin, errV := ValidateAdminRequest(updateAdmin)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	existingAdmin.Email = validatedExistingAdmin.Email
	existingAdmin.Name = validatedExistingAdmin.Name
	existingAdmin.Password = validatedExistingAdmin.Password

	if err := ac.AdminRepository.Update(existingAdmin); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update admin: %v", err))
	}

	return c.JSON(http.StatusCreated, existingAdmin)
}

func (ac *AdminController) Delete(c echo.Context) error {
	err := ac.AdminRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Admin successfully deleted")
}

// ValidateAdminRequest validates the input request for creating or updating a manufacturer.
func ValidateAdminRequest(request *requests.AdminRequest) (*models.Admin, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	admin := new(models.Admin)
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

	admin.Email = request.Email
	admin.Name = request.Name
	admin.Password = request.Password

	err := requests.ValidateAdminInputs(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}
