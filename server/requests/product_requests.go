package requests

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/server/responses"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
)

type ProductRequest struct {
	ProductCode    string  `json:"product_code"`
	ProductName    string  `json:"product_name"`
	ManufacturerID int     `json:"manufacturer_id"`
	ProductStock   int     `json:"product_stock"`
	ProductPrice   float64 `json:"product_price"`
}

// Validation >>>
// Validation contains checking inputs follow a format, checking if duplicates exist, checking if foreign keys exist

func ValidateProductInputs(product *models.Product) error {
	// Validate product inputs
	errVI := isValidProductInput(product)
	if errVI != nil {
		return errVI // Return the validation error if inputs are invalid
	}

	// Check if the product code or name has changed
	changed, err := hasProductChanged(product)
	if err != nil {
		return err
	}
	if changed {
		// If the product code or name has changed, run the duplicate check
		duplicateField, err := productExists(product)
		if err != nil {
			// Return an internal server error if the check fails
			return responses.NewHTTPError(http.StatusInternalServerError, "an error occurred while checking product existence")
		}
		if duplicateField != "" {
			// If there's a duplicate, respond
			if duplicateField == "product_code" {
				return responses.NewHTTPError(http.StatusBadRequest, "A product with this code already exists")
			}
			if duplicateField == "product_name" {
				return responses.NewHTTPError(http.StatusBadRequest, "A product with this name already exists")
			}
		}
	}

	// Check if manufacturer exists
	mfExists, errPE := manufacturerIDExists(product)
	if errPE != nil {
		return errPE // Return the error if manufacturer check fails
	}
	if !mfExists {
		return responses.NewHTTPError(http.StatusBadRequest, "manufacturer does not exist")
	}

	// If all checks pass, return nil (no error)
	return nil
}

func isValidProductInput(product *models.Product) error {
	// Allow only letters for product code
	validCodePattern := `^[a-zA-Z0-9]+$`
	matchedCode, _ := regexp.MatchString(validCodePattern, product.ProductCode)
	if !matchedCode {
		return responses.NewHTTPError(http.StatusBadRequest, "Product code is invalid : No Special Characters")
	}

	// Allow only letters for product name
	validNamePattern := `^[a-zA-Z0-9\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, product.ProductName)
	if !matchedName {
		return responses.NewHTTPError(http.StatusBadRequest, "Product name is invalid : No Special Characters")
	}

	// Allow only whole numbers for manufacturer id
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(product.ManufacturerID))
	if !matchedID {
		return responses.NewHTTPError(http.StatusBadRequest, "Manufacturer ID is invalid : No Special Characters or Letters")
	}

	// Allow only whole numbers for stock
	validStockPattern := `^[0-9]+$`
	matchedStock, _ := regexp.MatchString(validStockPattern, strconv.Itoa(product.Stock))
	if !matchedStock {
		return responses.NewHTTPError(http.StatusBadRequest, "Stock is invalid : No Special Characters or Letters")
	}

	// Allow only numbers for price
	validPricePattern := `^\d+(\.\d{1,2})?$`
	matchedPrice, _ := regexp.MatchString(validPricePattern, strconv.FormatFloat(product.Price, 'f', -1, 64))
	if !matchedPrice {
		return responses.NewHTTPError(http.StatusBadRequest, "Price is invalid : No Special Characters or Letters")
	}

	return nil
}

func hasProductChanged(product *models.Product) (bool, error) {
	// Fetch the existing product from the database
	var existingProduct models.Product
	err := db.DatabaseConnection().First(&existingProduct).Error
	if err != nil {
		// Return error if the product is not found or if there's an issue
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, responses.NewHTTPError(http.StatusBadRequest, "Product not found")
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "an error occurred while fetching the product")
	}

	// Check if the product code or name has changed
	if product.ProductCode != existingProduct.ProductCode || product.ProductName != existingProduct.ProductName {
		return true, nil // Product code or name has changed
	}

	return false, nil // No change in product code or name
}

func productExists(product *models.Product) (string, error) {
	// Check if the product code already exists
	var productWithCode models.Product
	err := db.DatabaseConnection().Where("product_code = ?", product.ProductCode).First(&productWithCode).Error
	if err == nil {
		// Product code exists
		return "product_code", nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", responses.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// Check if the product name already exists
	var productWithName models.Product
	err = db.DatabaseConnection().Where("product_name = ?", product.ProductName).First(&productWithName).Error
	if err == nil {
		// Product name exists
		return "product_name", nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", responses.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// Neither product code nor product name exists
	return "", nil
}

func manufacturerIDExists(product *models.Product) (bool, error) {
	manufacturer := new(models.Manufacturer)
	err := db.DatabaseConnection().Where("manufacturer_id = ?", product.ManufacturerID).First(&manufacturer).Error
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
