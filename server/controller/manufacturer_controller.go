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
	// Bind request body to the ManufacturerRequest struct
	requestManufacturer := new(requests.ManufacturerRequest)

	if err := c.Bind(&requestManufacturer); err != nil {
		// Return bad request if binding fails
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind product data"))
	}

	// Call the validation method
	_, err := mc.validateManufacturerRequest(requestManufacturer)
	if err != nil {
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			// Return the exact status code and message from validation
			return c.JSON(httpErr.StatusCode, echo.Map{
				"error": httpErr.Message,
			})
		}
		// If the error is not a custom HTTPError, return a generic bad request
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("manufacturer validation failed: %v", err))
	}

	// Map the validated request data to the manufacturer model
	manufacturer := &models.Manufacturer{
		ManufacturerName: requestManufacturer.ManufacturerName,
	}

	// Call repository method to create the new manufacturer
	err = mc.ManufacturerRepository.Create(manufacturer, c)
	if err != nil {
		// Return conflict if manufacturer creation fails
		return responses.ErrorResponse(c, http.StatusConflict, fmt.Errorf("failed to create product: %v", err))
	}

	// Return success response with the created manufacturer
	return c.JSON(http.StatusCreated, manufacturer)
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

	// Validate the incoming manufacturer data
	_, err = mc.validateManufacturerRequest(updateManufacturer)
	if err != nil {
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return responses.ErrorResponse(c, httpErr.StatusCode, httpErr)
		}
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("manufacturer validation failed: %v", err))
	}

	existingManufacturer.ManufacturerName = updateManufacturer.ManufacturerName

	if err := mc.ManufacturerRepository.Update(existingManufacturer, c); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update manufacturer: %v", err))
	}

	return responses.SuccessResponse(c, "Manufacturer updated successfully")
}

func (mc *ManufacturerController) Delete(c echo.Context) error {
	err := mc.ManufacturerRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Manufacturer successfully deleted")
}

// Validation Methods >>>
// Simple validation methods to prevent incorrect values from being requested >>>

func (mc *ManufacturerController) validateManufacturerRequest(request *requests.ManufacturerRequest) (*models.Manufacturer, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	manufacturer := new(models.Manufacturer)
	if request.ManufacturerName == "" {
		return nil, responses.NewHTTPError(http.StatusNotAcceptable, "invalid manufacturer name")
	}
	if len(request.ManufacturerName) < 1 || len(request.ManufacturerName) > 25 {
		return nil, responses.NewHTTPError(http.StatusLengthRequired, "manufacturer name must be between 1 and 25 characters")
	}

	manufacturer.ManufacturerName = request.ManufacturerName

	return manufacturer, nil
}
