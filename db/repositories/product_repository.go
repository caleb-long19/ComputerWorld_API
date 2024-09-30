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
	"strconv"
)

type ProductInterface interface {
	Create(product *models.Product, c echo.Context) error
	Get(id interface{}) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(product *models.Product, c echo.Context) error
	Delete(id interface{}) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *models.Product, c echo.Context) error {
	// Validate inputs
	err := validateProductInputs(repo.DB, product)
	if err != nil {
		// Check if it's an HTTPError and use the correct status code and message
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.StatusCode, httpErr.Message)
		}
		return err // Return the error if product check fails
	}

	// Proceed with creating the product if validation passes
	if err := repo.DB.Create(product).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "could not create product")
	}

	// Return 201 Created if the product is successfully created
	return c.JSON(http.StatusCreated, product)
}

func (repo *ProductRepository) Get(id interface{}) (*models.Product, error) {
	var product models.Product
	if err := repo.DB.Where("product_id", id).First(&product, id).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	if err := repo.DB.Find(&products).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find all products"))
	}
	return products, nil
}

func (repo *ProductRepository) Update(product *models.Product, c echo.Context) error {
	// Validate inputs
	err := validateProductInputs(repo.DB, product)
	if err != nil {
		// Check if it's an HTTPError and use the correct status code and message
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return c.JSON(httpErr.StatusCode, httpErr.Message)
		}
		return err // Return the error if product check fails
	}

	// Proceed with creating the product if validation passes
	if err := repo.DB.Save(product).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "could not update product")
	}

	// Return 201 Created if the product is successfully updated
	return c.JSON(http.StatusCreated, product)
}

func (repo *ProductRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find product with id %v", id))
	}
	return repo.DB.Delete(models.Product{}, "product_id = ?", id).Error
}

// Validation Methods >>>
// Validation contains checking inputs follow a format, checking if duplicates exist, checking if foreign keys exist

func validateProductInputs(db *gorm.DB, product *models.Product) error {
	// Validate product inputs
	errVI := isValidProductInput(product)
	if errVI != nil {
		return errVI // Return the validation error if inputs are invalid
	}

	// Check if the product exists
	exists, err := productExists(db, product)
	if err != nil {
		return responses.NewHTTPError(http.StatusBadRequest, "an error occurred while checking product existence")
	}
	if exists {
		return responses.NewHTTPError(http.StatusConflict, "A product with this code or name already exists")
	}

	// Check if manufacturer exists
	mfExists, errPE := manufacturerIDExists(db, product)
	if errPE != nil {
		return errPE // Return the error if manufacturer check fails
	}
	if !mfExists {
		return responses.NewHTTPError(http.StatusNotFound, "manufacturer does not exist")
	}

	// If all checks pass, return nil (no error)
	return nil
}

func isValidProductInput(product *models.Product) error {
	// Allow only letters for product code
	validCodePattern := `^[a-zA-Z0-9]+$`
	matchedCode, _ := regexp.MatchString(validCodePattern, product.ProductCode)
	if !matchedCode {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Product code is invalid : No Special Characters")
	}

	// Allow only letters for product name
	validNamePattern := `^[a-zA-Z0-9\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, product.ProductName)
	if !matchedName {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Product name is invalid : No Special Characters")
	}

	// Allow only whole numbers for manufacturer id
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(product.ManufacturerID))
	if !matchedID {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Manufacturer ID is invalid : No Special Characters or Letters")
	}

	// Allow only whole numbers for stock
	validStockPattern := `^[0-9]+$`
	matchedStock, _ := regexp.MatchString(validStockPattern, strconv.Itoa(product.Stock))
	if !matchedStock {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Stock is invalid : No Special Characters or Letters")
	}

	// Allow only numbers for price
	validPricePattern := `^\d+(\.\d{1,2})?$`
	matchedPrice, _ := regexp.MatchString(validPricePattern, strconv.FormatFloat(product.Price, 'f', -1, 64))
	if !matchedPrice {
		return responses.NewHTTPError(http.StatusNotAcceptable, "Price is invalid : No Special Characters or Letters")
	}

	return nil
}

func productExists(db *gorm.DB, product *models.Product) (bool, error) {
	// Attempt to find the product name or code in the database
	err := db.Where("product_code = ?", product.ProductCode).Or(db.Where("product_name = ?", product.ProductName)).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Product not found - does not exist
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	// Product found - does exist
	return true, nil
}

func manufacturerIDExists(db *gorm.DB, product *models.Product) (bool, error) {
	manufacturer := new(models.Manufacturer)
	err := db.Where("manufacturer_id = ?", product.ManufacturerID).First(&manufacturer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Manufacturer does not exist
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// manufacturer does exist
	return true, nil
}
