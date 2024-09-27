package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"regexp"
)

type ManufacturerInterface interface {
	Create(manufacturer *models.Manufacturer) error
	Get(id interface{}) (*models.Manufacturer, error)
	GetAll() ([]*models.Manufacturer, error)
	Update(manufacturer *models.Manufacturer) error
	Delete(id interface{}) error
}

type ManufacturerRepository struct {
	DB *gorm.DB
}

func NewManufacturerRepository(db *gorm.DB) *ManufacturerRepository {
	return &ManufacturerRepository{DB: db}
}

func (repo *ManufacturerRepository) Create(manufacturer *models.Manufacturer) error {
	// Validate Manufacturer
	err := validateManufacturerInputs(repo.DB, manufacturer)
	if err != nil {
		return err
	}

	// Create the manufacturer
	return repo.DB.Create(manufacturer).Error
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

func (repo *ManufacturerRepository) Update(manufacturer *models.Manufacturer) error {
	// Validate Manufacturer
	err := validateManufacturerInputs(repo.DB, manufacturer)
	if err != nil {
		return err
	}

	return repo.DB.Save(manufacturer).Error
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
		return errors.New("error: An error occurred while checking manufacturer existence")
	}
	if exists {
		log.Println("CONSOLE: LINE 73: ALREADY EXISTS:", exists)
		return errors.New("error: A manufacturer with this name already exists")
	}

	// Validate input
	if !isValidManufacturerInput(manufacturer) {
		return errors.New("error: Manufacturer name is invalid")
	}

	return err
}

func isValidManufacturerInput(manufacturer *models.Manufacturer) bool {
	// Allow only letters
	validNamePattern := `^[a-zA-Z\s]`
	matched, _ := regexp.MatchString(validNamePattern, manufacturer.ManufacturerName)
	return matched
}

func manufacturerExists(db *gorm.DB, manufacturer *models.Manufacturer) (bool, error) {
	// Attempt to find the manufacturer in the database
	err := db.Where("manufacturer_name = ?", manufacturer.ManufacturerName).First(&manufacturer).Error
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
