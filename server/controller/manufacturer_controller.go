package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/reponses"
	"ComputerWorld_API/server/requests"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ManufacturerController struct {
	ManufacturerRepository repositories.ManufacturerInterface
}

func (mc *ManufacturerController) Create(c echo.Context) error {
	requestManufacturer := new(requests.ManufacturerRequest)

	if err := c.Bind(&requestManufacturer); err != nil {
		return c.JSON(http.StatusBadRequest, requestManufacturer)
	}
	manufacturer, err := mc.validateManufacturerRequest(requestManufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = mc.ManufacturerRepository.Create(manufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, manufacturer)
}

func (mc *ManufacturerController) Get(c echo.Context) error {
	manufacturer, err := mc.ManufacturerRepository.Get(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, manufacturer)
}

func (mc *ManufacturerController) GetAll(c echo.Context) error {
	manufacturers, err := mc.ManufacturerRepository.GetAll()
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, manufacturers)
}

func (mc *ManufacturerController) Update(c echo.Context) error {
	existingManufacturer, err := mc.ManufacturerRepository.Get(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	updateManufacturer := new(requests.ManufacturerRequest)
	if err := c.Bind(&updateManufacturer); err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if updateManufacturer == nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, errors.New("manufacturer is required"))
	}

	_, err = mc.validateManufacturerRequest(updateManufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	existingManufacturer = &models.Manufacturer{
		ManufacturerID:   existingManufacturer.ManufacturerID,
		ManufacturerName: updateManufacturer.ManufacturerName,
	}

	err = mc.ManufacturerRepository.Update(existingManufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, existingManufacturer)
}

func (mc *ManufacturerController) Delete(c echo.Context) error {
	err := mc.ManufacturerRepository.Delete(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Manufacturer successfully deleted")
}

func (mc *ManufacturerController) validateManufacturerRequest(request *requests.ManufacturerRequest) (*models.Manufacturer, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	manufacturer := new(models.Manufacturer)
	if request.ManufacturerName == "" {
		return nil, errors.New("error: Invalid manufacturer name")
	}

	manufacturer.ManufacturerName = request.ManufacturerName

	return manufacturer, nil
}
