package repositories

import (
	"ComputerWorld_API/db/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

type OrderInterface interface {
	Create(order *models.Order) error
	Get(id interface{}) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(id interface{}) error
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Create(order *models.Order) error {
	err := validateOrderInputs(repo.DB, order)
	if err != nil {
		return err
	}

	errCOP := CalculateOrderPrice(repo.DB, order)
	if errCOP != nil {
		return errCOP
	}

	errCPS := CalculateProductStock(repo.DB, order)
	if errCPS != nil {
		return errCPS
	}

	return repo.DB.Create(order).Error
}

func (repo *OrderRepository) Get(id interface{}) (*models.Order, error) {
	var order models.Order
	if err := repo.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find order with id %v", id))
	}
	return &order, nil
}

func (repo *OrderRepository) GetAll() ([]*models.Order, error) {
	var orders []*models.Order
	if err := repo.DB.Find(&orders).Error; err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find orders %v", orders))
	}
	return orders, nil
}

func (repo *OrderRepository) Update(order *models.Order) error {
	err := validateOrderInputs(repo.DB, order)
	if err != nil {
		return err
	}

	errCOP := CalculateOrderPrice(repo.DB, order)
	if errCOP != nil {
		return errCOP
	}

	errCPS := CalculateProductStock(repo.DB, order)
	if errCPS != nil {
		return errCPS
	}

	return repo.DB.Save(order).Error
}

func (repo *OrderRepository) Delete(id interface{}) error {
	_, err := repo.Get(id)
	if err != nil {
		return err
	}
	return repo.DB.Delete(models.Order{}, "order_id = ?", id).Error
}

// Validation Methods >>>
// I stored all other validations into one method, so I only need to call it once as is it being used in the create & update methods >>>

func validateOrderInputs(db *gorm.DB, order *models.Order) error {
	errVI := isValidOrderInput(order)
	if errVI != nil {
		return errVI
	}

	// Check if order exists
	exists, err := orderExists(db, order)
	if err != nil {
		return errors.New("error: An error occurred while checking order existence")
	}
	if exists {
		return errors.New("error: An order with this name already exists")
	}

	existProduct, errPE := productIDExists(db, order)
	if errPE != nil {
		return errors.New("error: An error occurred while checking product existence")
	}
	if !existProduct {
		return errors.New("error: this product id does not exist")
	}

	return nil
}

func isValidOrderInput(order *models.Order) error {
	// Allow only letters for reference
	validNamePattern := `^[a-zA-Z0-9]+$`
	matchedRef, _ := regexp.MatchString(validNamePattern, order.OrderRef)
	if !matchedRef {
		return errors.New("error: Order reference is invalid")
	}

	// Allow only whole numbers for amount
	validAmountPattern := `^[0-9]+$`
	matchedAmount, _ := regexp.MatchString(validAmountPattern, strconv.Itoa(order.OrderAmount))
	if !matchedAmount {
		return errors.New("error: Order amount is invalid")
	}

	// Allow only whole numbers for product ID
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(order.ProductID))
	if !matchedID {
		return errors.New("error: Product ID is invalid")
	}

	return nil
}

func orderExists(db *gorm.DB, order *models.Order) (bool, error) {
	// Attempt to find the order in the database
	err := db.Where("order_ref = ?", order.OrderRef).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Order not found, return false
			return false, nil
		}
		return false, err
	}
	// Order found, return true
	return true, nil
}

func productIDExists(db *gorm.DB, order *models.Order) (bool, error) {
	product := new(models.Product)

	err := db.Where("product_id = ?", order.ProductID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	// product id found
	return true, nil
}

// Calculations Methods >>
// These are used to automatically calculate the order prices and product stock after creation/updates

func CalculateOrderPrice(db *gorm.DB, order *models.Order) error {
	var product models.Product
	if err := db.First(&product, order.ProductID).Error; err != nil {
		return err
	}
	order.OrderPrice = float64(order.OrderAmount) * product.Price
	return nil
}

func CalculateProductStock(db *gorm.DB, order *models.Order) error {
	var product models.Product
	if err := db.First(&product, order.ProductID).Error; err != nil {
		return err
	}

	// Check if there's enough stock to fulfill the order
	if product.Stock < order.OrderAmount {
		return errors.New("insufficient stock for the product")
	}

	product.Stock -= order.OrderAmount

	// Save the updated product stock in the database
	if err := db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
