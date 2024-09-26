package controller

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/reponses"
	"ComputerWorld_API/server/requests"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"regexp"
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
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
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
	// Check if ManufacturerName is provided and valid
	if updateManufacturer.ManufacturerName == "" {
		return reponses.ErrorResponse(c, http.StatusBadRequest, errors.New("manufacturer name is required"))
	}
	// Check if ManufacturerName isn't about a certain length
	if len(updateManufacturer.ManufacturerName) > 30 {
		return reponses.ErrorResponse(c, http.StatusBadRequest, errors.New("manufacturer name exceeds the maximum length of 30 characters"))
	}
	// Check if the manufacturer name already exists (to prevent duplicates)
	existingByName, _ := mc.ManufacturerRepository.Get(updateManufacturer.ManufacturerName)
	if existingByName != nil && existingByName.ManufacturerID != existingManufacturer.ManufacturerID {
		return reponses.ErrorResponse(c, http.StatusConflict, errors.New("manufacturer name already exists"))
	}
	// Validate the request further
	_, err = mc.validateManufacturerRequest(updateManufacturer)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	existingManufacturer = &models.Manufacturer{
		ManufacturerID:   existingManufacturer.ManufacturerID,
		ManufacturerName: updateManufacturer.ManufacturerName,
	}

	// Perform the update in the repository
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

// Validation Methods >>
// Simple validation methods to prevent incorrect values from being requested>

func (mc *ManufacturerController) validateManufacturerRequest(request *requests.ManufacturerRequest) (*models.Manufacturer, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	manufacturer := new(models.Manufacturer)
	if request.ManufacturerName == "" {
		return nil, errors.New("error: Invalid manufacturer name")
	}
	if len(request.ManufacturerName) < 1 || len(request.ManufacturerName) > 25 {
		return nil, errors.New("error: Manufacturer name must be between 1 and 25 characters")
	}
	// Check for invalid characters in Manufacturer Name
	if !isValidManufacturerName(request.ManufacturerName) {
		return nil, errors.New("manufacturer name contains invalid characters")
	}
	// Check if manufacturer exists
	exists, err := manufacturerExists(request.ManufacturerName, db.DatabaseConnection(), manufacturer)
	if err != nil {
		return nil, errors.New("error: A manufacturer with this name already exists")
	}
	if exists {
		return nil, errors.New("error: An manufacturer with this name already exists")
	}

	manufacturer.ManufacturerName = request.ManufacturerName

	return manufacturer, nil
}

// TODO: NEED TO MOVE THESE VALIDATIONS AND EXIST CHECKS TO THE REPOSITORY FILE (Manufacturer_repository)

func isValidManufacturerName(name string) bool {
	// Allow only letters
	validNamePattern := `^[a-zA-Z\s]`
	matched, _ := regexp.MatchString(validNamePattern, name)
	return matched
}

func manufacturerExists(manufacturerName string, db *gorm.DB, manufacturer *models.Manufacturer) (bool, error) {
	// Attempt to find the manufacturer in the database
	err := db.Where("manufacturer_name = ?", manufacturerName).First(&manufacturer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Manufacturer not found, return false
			return false, nil
		}
		return false, err
	}
	// Manufacturer found, return true
	return true, nil
}
