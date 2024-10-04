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

type ManufacturerController struct {
	ManufacturerRepository repositories.ManufacturerInterface
}

func (mc *ManufacturerController) Create(c echo.Context) error {
	requestManufacturer := new(requests.ManufacturerRequest)

	if err := c.Bind(&requestManufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind manufacturer data"))
	}

	// Validate the request manufacturer data
	validatedRequest, errV := ValidateManufacturerRequest(requestManufacturer)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	// Call repository method to create the new manufacturer
	err := mc.ManufacturerRepository.Create(validatedRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed to create manufacturer: %v", err))
	}

	return c.JSON(http.StatusCreated, validatedRequest)
}

func (mc *ManufacturerController) Get(c echo.Context) error {
	manufacturer, err := mc.ManufacturerRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, manufacturer)
}

func (mc *ManufacturerController) GetAll(c echo.Context) error {
	manufacturers, err := mc.ManufacturerRepository.GetAll()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, manufacturers)
}

func (mc *ManufacturerController) Update(c echo.Context) error {
	existingManufacturer, err := mc.ManufacturerRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("manufacturer not found: %v", err))
	}

	var updateManufacturer = new(requests.ManufacturerRequest)
	if err := c.Bind(updateManufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind manufacturer data"))
	}
	if updateManufacturer == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid manufacturer data"))
	}

	// Validate the request manufacturer data
	validatedExistingManufacturer, errV := ValidateManufacturerRequest(updateManufacturer)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	existingManufacturer.ManufacturerName = validatedExistingManufacturer.ManufacturerName

	if err := mc.ManufacturerRepository.Update(existingManufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update manufacturer: %v", err))
	}

	return c.JSON(http.StatusCreated, existingManufacturer)
}

func (mc *ManufacturerController) Delete(c echo.Context) error {
	err := mc.ManufacturerRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Manufacturer successfully deleted")
}

// ValidateManufacturerRequest validates the input request for creating or updating a manufacturer.
func ValidateManufacturerRequest(request *requests.ManufacturerRequest) (*models.Manufacturer, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	manufacturer := new(models.Manufacturer)
	if request.ManufacturerName == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "manufacturer name is required")
	}
	if len(request.ManufacturerName) < 1 || len(request.ManufacturerName) > 25 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "manufacturer name must be between 1 and 25 characters")
	}

	manufacturer.ManufacturerName = request.ManufacturerName

	err := requests.ValidateManufacturerInputs(manufacturer)
	if err != nil {
		return nil, err
	}

	return manufacturer, nil
}
