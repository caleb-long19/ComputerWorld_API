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
	"strconv"
)

type OrderController struct {
	OrderRepository repositories.OrderInterface
	DB              *gorm.DB
}

func (oc *OrderController) Create(c echo.Context) error {
	requestOrder := new(requests.OrderRequest)

	if err := c.Bind(&requestOrder); err != nil {
		return c.JSON(http.StatusBadRequest, requestOrder)
	}
	order, err := oc.validateOrderRequest(requestOrder)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// Connect to the database and calculate the prices and stock
	errCPS := CalculateProductStock(db.DatabaseConnection(), order)
	if errCPS != nil {
		return errCPS
	}
	errCOP := CalculateOrderPrice(db.DatabaseConnection(), order)
	if errCOP != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, errCOP)
	}

	err = oc.OrderRepository.Create(order)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, order)
}

func (oc *OrderController) Get(c echo.Context) error {
	order, err := oc.OrderRepository.Get(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetAll(c echo.Context) error {
	orders, err := oc.OrderRepository.GetAll()
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) Update(c echo.Context) error {
	existingOrder, err := oc.OrderRepository.Get(c.Param("id"))

	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	updateOrder := new(requests.OrderRequest)
	if err := c.Bind(&updateOrder); err != nil {
		return c.JSON(http.StatusBadRequest, updateOrder)
	}

	if updateOrder == nil {
		return c.JSON(http.StatusBadRequest, updateOrder)
	}

	_, err = oc.validateOrderRequest(updateOrder)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	existingOrder = &models.Order{
		OrderID:     existingOrder.OrderID,
		OrderRef:    updateOrder.OrderReference,
		OrderAmount: updateOrder.OrderAmount,
		ProductID:   updateOrder.ProductID,
	}

	errCPS := CalculateProductStock(db.DatabaseConnection(), existingOrder)
	if errCPS != nil {
		return errCPS
	}

	errC := CalculateOrderPrice(db.DatabaseConnection(), existingOrder)
	if errC != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, errC)
	}

	if err := oc.OrderRepository.Update(existingOrder); err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, existingOrder)

}

func (oc *OrderController) Delete(c echo.Context) error {
	err := oc.OrderRepository.Delete(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Order successfully deleted")
}

// Validation Methods >>>
// Simple validation methods to prevent incorrect values from being requested

func (oc *OrderController) validateOrderRequest(request *requests.OrderRequest) (*models.Order, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	order := new(models.Order)
	if request.OrderReference == "" {
		return nil, errors.New("error: Invalid order reference")
	}
	if len(request.OrderReference) < 3 || len(request.OrderReference) > 12 {
		return nil, errors.New("error: Order reference must be between 3 and 12 characters")
	}
	if request.OrderAmount <= 0 {
		return nil, errors.New("error: Invalid order amount")
	}
	if request.OrderAmount > 50 {
		return nil, errors.New("error: Order amount exceeds maximum limit")
	}
	if request.ProductID <= 0 {
		return nil, errors.New("error: Invalid product id")
	}
	// Check for invalid characters in order values
	if validRef, validAmount, validID := isValidOrderInput(
		request.OrderReference,
		request.OrderAmount,
		request.ProductID); !validRef || !validAmount || !validID {
		return nil, errors.New("order input contains invalid characters or format")
	}
	// Check if order reference exists
	exists, err := orderExists(request.OrderReference, db.DatabaseConnection(), order)
	if err != nil {
		return nil, errors.New("error: An order with this reference already exists")
	}
	if exists {
		return nil, errors.New("error: An order with this reference already exists")
	}

	order.OrderRef = request.OrderReference
	order.OrderAmount = request.OrderAmount
	order.ProductID = request.ProductID

	return order, nil
}

// TODO: NEED TO MOVE THESE VALIDATIONS AND EXIST CHECKS TO THE REPOSITORY FILE (Order_repository)

func isValidOrderInput(reference string, amount int, productID int) (bool, bool, bool) {
	// Allow only letters for reference
	validNamePattern := `^[a-zA-Z0-9]+$`
	matchedRef, _ := regexp.MatchString(validNamePattern, reference)

	// Allow only whole numbers for amount
	validAmountPattern := `^[0-9]+$`
	matchedAmount, _ := regexp.MatchString(validAmountPattern, strconv.Itoa(amount))

	// Allow only whole numbers for product ID
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(productID))

	return matchedRef, matchedAmount, matchedID
}

func orderExists(orderReference string, db *gorm.DB, order *models.Order) (bool, error) {
	// Attempt to find the order in the database
	err := db.Where("order_ref = ?", orderReference).First(&order).Error
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

// Calculations Methods >>
// These are used to automatically calculate the order prices and product stock after creation/updates
// TODO: Move these calculations into the repository file (order.repository)

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
