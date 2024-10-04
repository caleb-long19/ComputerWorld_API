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

type OrderRequest struct {
	OrderReference string  `json:"order_ref"`
	OrderAmount    int     `json:"order_amount"`
	ProductID      int     `json:"product_id"`
	OrderPrice     float64 `json:"order_price"`
}

// Validation Methods
// I stored all other validations into one method, so I only need to call it once as is it being used in the create & update methods >>>

func ValidateOrderInputs(order *models.Order) error {
	errVI := isValidOrderInput(order)
	if errVI != nil {
		return errVI
	}

	// Check if order exists
	exists, err := orderExists(order)
	if err != nil {
		return responses.NewHTTPError(http.StatusBadRequest, "An error occurred while checking order existence")
	}
	if exists {
		return responses.NewHTTPError(http.StatusConflict, "An order with this name already exists")
	}

	existProduct, errPE := productIDExists(order)
	if errPE != nil {
		return responses.NewHTTPError(http.StatusBadRequest, "An error occurred while checking product existence")
	}
	if !existProduct {
		return responses.NewHTTPError(http.StatusNotFound, "this product id does not exist")
	}

	return nil
}

func isValidOrderInput(order *models.Order) error {
	// Allow only letters for reference
	validNamePattern := `^[a-zA-Z0-9]+$`
	matchedRef, _ := regexp.MatchString(validNamePattern, order.OrderRef)
	if !matchedRef {
		return responses.NewHTTPError(http.StatusBadRequest, "Order reference is invalid : No Special Characters")
	}

	// Allow only whole numbers for amount
	validAmountPattern := `^[0-9]+$`
	matchedAmount, _ := regexp.MatchString(validAmountPattern, strconv.Itoa(order.OrderAmount))
	if !matchedAmount {
		return responses.NewHTTPError(http.StatusBadRequest, "Order amount is invalid : No Special Characters or Letters")
	}

	// Allow only whole numbers for product ID
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(order.ProductID))
	if !matchedID {
		return responses.NewHTTPError(http.StatusBadRequest, "Product ID is invalid : No Special Characters or Letters")
	}

	return nil
}

func orderExists(order *models.Order) (bool, error) {
	// Attempt to find the order in the database
	err := db.DatabaseConnection().Where("order_ref = ?", order.OrderRef).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Order not found, return false
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "Internal server error: Try again!")
	}
	// Order found, return true
	return true, nil
}

func productIDExists(order *models.Order) (bool, error) {
	product := new(models.Product)

	err := db.DatabaseConnection().Where("product_id = ?", order.ProductID).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, responses.NewHTTPError(http.StatusInternalServerError, "Internal server error: Try again!")
	}
	// product id found
	return true, nil
}
