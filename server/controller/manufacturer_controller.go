package controller

import (
	"ComputerWorld_API/db/model"
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

func (hc *ManufacturerController) Create(c echo.Context) error {
	requestManufacturer := new(requests.ManufacturerRequest)

	if err := c.Bind(&requestManufacturer); err != nil {
		return c.JSON(http.StatusBadRequest, requestManufacturer)
	}
	manufacturer, err := hc.validateManufacturerRequest(requestManufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = hc.ManufacturerRepository.Create(manufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, manufacturer)
}

//func (hc *ManufacturerController) Get(c echo.Context) error {
//	requestManufacturer := new(requests.ManufacturerRequest)
//	if res := hc.ManufacturerRepository.DB.Where("manufacturer_id = ?", c.Param("id")).First(&requestManufacturer); res.Error != nil {
//		return c.JSON(http.StatusNotFound, requestManufacturer)
//	}
//
//	return c.JSON(http.StatusOK, requestManufacturer)
//}

//func (hc *ManufacturerController) Update(c echo.Context) error {
//
//	requestManufacturer := new(requests.ManufacturerRequest)
//	if err := c.Bind(requestManufacturer); err != nil {
//		return c.JSON(http.StatusBadRequest, "Error: Could not bind manufacturer data")
//	}
//
//	existingManufacturer := new(requests.ManufacturerRequest)
//	if err := hc.Db.Where("manufacturer_id = ?", c.Param("id")).First(&existingManufacturer).Error; err != nil {
//		return c.JSON(http.StatusNotFound, existingManufacturer)
//	}
//
//	existingManufacturer.ManufacturerName = requestManufacturer.ManufacturerName
//	if err := hc.Db.Save(&existingManufacturer).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, existingManufacturer)
//	}
//
//	return c.JSON(http.StatusOK, existingManufacturer)
//}

//func (hc *ManufacturerController) Delete(c echo.Context) error {
//	requestManufacturer := new(requests.ManufacturerRequest)
//
//	result := hc.Db.Where("manufacturer_id = ?", c.Param("id")).First(&requestManufacturer)
//	if result.Error != nil {
//		return c.JSON(http.StatusNotFound, requestManufacturer)
//	}
//
//	result = hc.Db.Delete(&requestManufacturer)
//	if result.Error != nil {
//		return c.JSON(http.StatusBadRequest, requestManufacturer)
//	}
//
//	return c.JSON(http.StatusOK, requestManufacturer)
//}

func (hc *ManufacturerController) validateManufacturerRequest(request *requests.ManufacturerRequest) (*model.Manufacturer, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	manufacturer := new(model.Manufacturer)
	if request.ManufacturerName == "" {
		return nil, errors.New("error: Invalid manufacturer name")
	}

	manufacturer.ManufacturerName = request.ManufacturerName

	return manufacturer, nil
}
