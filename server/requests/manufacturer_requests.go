package requests

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/server/responses"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"regexp"
)

type ManufacturerRequest struct {
	ManufacturerName string `json:"manufacturer_name"`
}

type Database struct {
	DB *gorm.DB
}

// ValidateManufacturerInputs checks if the manufacturer already exists and if the inputs are valid.
func ValidateManufacturerInputs(manufacturer *models.Manufacturer) error {
	// Check if manufacturer already exists in the database
	if exists, err := manufacturerExists(manufacturer); err != nil {
		return err
	} else if exists {
		return responses.NewHTTPError(http.StatusBadRequest, "A manufacturer with this name already exists")
	}

	// Validate the manufacturer name format
	if err := validateManufacturerFormat(manufacturer.ManufacturerName); err != nil {
		return err
	}

	return nil
}

// validateManufacturerFormat checks if the manufacturer name contains only valid characters.
func validateManufacturerFormat(name string) error {
	validNamePattern := `^[a-zA-Z0-9\s]+$`
	if matched, _ := regexp.MatchString(validNamePattern, name); !matched {
		return responses.NewHTTPError(http.StatusBadRequest, "manufacturer name is invalid: no special characters allowed")
	}
	return nil
}

// manufacturerExists checks if the manufacturer already exists in the database.
func manufacturerExists(manufacturer *models.Manufacturer) (bool, error) {
	err := db.DatabaseConnection().Where("manufacturer_name = ?", manufacturer.ManufacturerName).First(&manufacturer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "error while checking manufacturer existence")
	}
	return true, nil
}
