package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

type ProductInterface interface {
	Create(product *models.Product) error
	Get(id interface{}) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id interface{}) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Create(product *models.Product) error {
	err := validateProductInputs(repo.DB, product)
	if err != nil {
		return err
	}

	return repo.DB.Create(product).Error
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

func (repo *ProductRepository) Update(product *models.Product) error {
	err := validateProductInputs(repo.DB, product)
	if err != nil {
		return err
	}

	return repo.DB.Save(product).Error
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
	errVI := isValidProductInput(product)
	if errVI != nil {
		return errVI
	}

	// Check if product exists
	exists, err := productExists(db, product)
	if err != nil {
		return errors.New("error: An error occurred while checking product existence")
	}
	if exists {
		return errors.New("error: A product with this code or name already exists")
	}

	mfExists, errPE := manufacturerIDExists(db, product)
	if errPE != nil {
		return errors.New("error: An error occurred while checking manufacturer existence")
	}
	if !mfExists {
		return errors.New("error: this manufacturer id does not exist")
	}

	return nil
}

func isValidProductInput(product *models.Product) error {
	// Allow only letters for product code
	validCodePattern := `^[a-zA-Z0-9]+$`
	matchedCode, _ := regexp.MatchString(validCodePattern, product.ProductCode)
	if !matchedCode {
		return errors.New("error: Product code is invalid : No Special Characters")
	}

	// Allow only letters for product name
	validNamePattern := `^[a-zA-Z0-9\s]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, product.ProductName)
	if !matchedName {
		return errors.New("error: Product name is invalid : No Special Characters")
	}

	// Allow only whole numbers for manufacturer id
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(product.ManufacturerID))
	if !matchedID {
		return errors.New("error: Manufacturer ID is invalid : No Special Characters or Letters")
	}

	// Allow only whole numbers for stock
	validStockPattern := `^[0-9]+$`
	matchedStock, _ := regexp.MatchString(validStockPattern, strconv.Itoa(product.Stock))
	if !matchedStock {
		return errors.New("error: Stock is invalid : No Special Characters or Letters")
	}

	// Allow only numbers for price
	validPricePattern := `^\d+(\.\d{1,2})?$`
	matchedPrice, _ := regexp.MatchString(validPricePattern, strconv.FormatFloat(product.Price, 'f', -1, 64))
	if !matchedPrice {
		return errors.New("error: Price is invalid : No Special Characters or Letters")
	}

	return nil
}

func productExists(db *gorm.DB, product *models.Product) (bool, error) {
	// Attempt to find the product name or code in the database
	err := db.Where("product_code = ?", product.ProductCode).Or(db.Where("product_name = ?", product.ProductName)).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Product not found, return false
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func manufacturerIDExists(db *gorm.DB, product *models.Product) (bool, error) {
	manufacturer := new(models.Manufacturer)

	err := db.Where("manufacturer_id = ?", product.ManufacturerID).First(&manufacturer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	// manufacturer id found
	return true, nil
}
