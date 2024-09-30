package repositories

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/server/responses"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"regexp"
)

type ManufacturerInterface interface {
	Create(manufacturer *models.Manufacturer, c echo.Context) error
	Get(id interface{}) (*models.Manufacturer, error)
	GetAll() ([]*models.Manufacturer, error)
	Update(manufacturer *models.Manufacturer, c echo.Context) error
	Delete(id interface{}) error
}

type ManufacturerRepository struct {
	DB *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{DB: db}
}

func (repo *ManufacturerRepository) Create(manufacturer *models.Manufacturer, c echo.Context) error {
	// Validate inputs
	err := validateManufacturerInputs(repo.DB, manufacturer)
	if err != nil {
		// Check if it's an HTTPError and use the correct status code and message
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.StatusCode, httpErr.Message)
		}
		return err
	}

	// Proceed with creating the manufacturer if validation passes
	if err := repo.DB.Create(manufacturer).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "could not create manufacturer")
	}

	// Return 201 Created if the manufacturer is successfully created
	return c.JSON(http.StatusCreated, manufacturer)
}

func (repo *ManufacturerRepository) Get(id interface{}) (*models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	if err := repo.DB.Where("manufacturer_id = ?", id).First(&manufacturer).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return &manufacturer, nil
}

func (repo *ManufacturerRepository) GetAll() ([]*models.Manufacturer, error) {
	var manufacturers []*models.Manufacturer
	if err := repo.DB.Find(&manufacturers).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find manufacturers"))
	}
	return manufacturers, nil
}

func (repo *ManufacturerRepository) Update(manufacturer *models.Manufacturer, c echo.Context) error {
	// Validate inputs
	err := validateManufacturerInputs(repo.DB, manufacturer)
	if err != nil {
		// Check if it's an HTTPError and use the correct status code and message
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.StatusCode, httpErr.Message)
		}
		return err
	}

	if err := repo.DB.Save(manufacturer).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "could not update manufacturer")
	}

	return c.JSON(http.StatusCreated, manufacturer)
}

func (repo *ManufacturerRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find manufacturer with id %v", id))
	}
	return repo.DB.Delete(models.Manufacturer{}, "manufacturer_id = ?", id).Error
}

// Validation >>>
// I stored all other validations into one method, so I only need to call it once as is it being used in the create & update methods >>>

func validateManufacturerInputs(db *gorm.DB, manufacturer *models.Manufacturer) error {
	// Check if manufacturer exists
	exists, err := manufacturerExists(db, manufacturer)
	if err != nil {
		return responses.NewHTTPError(http.StatusInternalServerError, "an error occurred while checking manufacturer existence")
	}
	if exists {
		return responses.NewHTTPError(http.StatusConflict, "A manufacturer with this name already exists")
	}

	// Validate manufacturer input
	errVI := isValidManufacturerInput(manufacturer)
	if errVI != nil {
		return errVI // Return the validation error if inputs are invalid
	}

	return err
}

func isValidManufacturerInput(manufacturer *models.Manufacturer) error {
	// Allow only letters
	validNamePattern := `^[a-zA-Z0-9\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, manufacturer.ManufacturerName)
	if !matchedName {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Manufacturer Name is invalid : No Special Characters")
	}

	return nil
}

func manufacturerExists(db *gorm.DB, manufacturer *models.Manufacturer) (bool, error) {
	// Attempt to find the manufacturer in the database
	err := db.Where("manufacturer_name = ?", manufacturer.ManufacturerName).First(&manufacturer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Manufacturer not found, return false
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	// Manufacturer found, return true
	return true, nil
}
