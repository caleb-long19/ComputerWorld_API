package controller

import (
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/requests"
	"ComputerWorld_API/server/responses"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderController struct {
	OrderRepository repositories.OrderInterface
}

func (oc *OrderController) Create(c echo.Context) error {
	// Bind request body to the OrderRequest struct
	requestOrder := new(requests.OrderRequest)

	if err := c.Bind(&requestOrder); err != nil {
		// Return bad request if binding fails
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind order data"))
	}

	// Validate the request manufacturer data
	validatedRequest, errV := ValidateOrderRequest(requestOrder)
	if errV != nil {
		// Return the validation error directly, with its status code
		return responses.ErrorResponse(c, 0, errV)
	}

	// Call repository method to create the new product
	err := oc.OrderRepository.Create(validatedRequest)
	if err != nil {
		// Return conflict if product creation fails
		return responses.ErrorResponse(c, http.StatusConflict, fmt.Errorf("failed to create order: %v", err))
	}

	// Return success response with the created product
	return c.JSON(http.StatusCreated, validatedRequest)
}

func (oc *OrderController) Get(c echo.Context) error {
	order, err := oc.OrderRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetAll(c echo.Context) error {
	orders, err := oc.OrderRepository.GetAll()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, orders)
}

func (oc *OrderController) Update(c echo.Context) error {
	// Get the existing order by ID
	existingOrder, err := oc.OrderRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("order not found: %v", err))
	}

	// Bind the incoming request to the OrderRequest struct
	var updateOrder = new(requests.OrderRequest)
	if err := c.Bind(updateOrder); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind order data"))
	}

	// Validate the incoming order data
	if updateOrder == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid order data"))
	}

	// Validate the request manufacturer data
	validatedExistingOrder, errV := ValidateOrderRequest(updateOrder)
	if errV != nil {
		// Return the validation error directly
		return responses.ErrorResponse(c, 0, errV)
	}

	existingOrder.OrderRef = validatedExistingOrder.OrderRef
	existingOrder.OrderAmount = validatedExistingOrder.OrderAmount
	existingOrder.ProductID = validatedExistingOrder.ProductID

	// Attempt to update the product in the repository
	if err := oc.OrderRepository.Update(existingOrder); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update order: %v", err))
	}

	// Successfully updated, return the updated order data
	return c.JSON(http.StatusCreated, existingOrder)
}

func (oc *OrderController) Delete(c echo.Context) error {
	err := oc.OrderRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Order successfully deleted")
}

func ValidateOrderRequest(request *requests.OrderRequest) (*models.Order, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	order := new(models.Order)
	if request.OrderReference == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid order reference")
	}
	if len(request.OrderReference) < 3 || len(request.OrderReference) > 12 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Order reference must be between 3 and 12 characters")
	}
	if request.OrderAmount <= 0 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid order amount")
	}
	if request.OrderAmount > 50 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Order amount exceeds maximum limit")
	}
	if request.ProductID <= 0 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid product id")
	}

	order.OrderRef = request.OrderReference
	order.OrderAmount = request.OrderAmount
	order.ProductID = request.ProductID

	err := requests.ValidateOrderInputs(order)
	if err != nil {
		return nil, err
	}

	errCOP := CalculateOrderPrice(order)
	if errCOP != nil {
		return order, errCOP
	}
	errCPS := CalculateProductStock(order)
	if errCPS != nil {
		return order, errCPS
	}

	return order, nil
}

// Calculations >>
// These are used to automatically calculate the order prices and product stock after creation/updates

func CalculateOrderPrice(order *models.Order) error {
	var product models.Product
	if err := db.DatabaseConnection().First(&product, order.ProductID).Error; err != nil {
		return err
	}
	order.OrderPrice = float64(order.OrderAmount) * product.Price
	return nil
}

func CalculateProductStock(order *models.Order) error {
	var product models.Product
	if err := db.DatabaseConnection().First(&product, order.ProductID).Error; err != nil {
		return err
	}

	// Check if there's enough stock to fulfill the order
	if product.Stock < order.OrderAmount {
		return responses.NewHTTPError(http.StatusBadRequest, "insufficient stock for the product")
	}

	product.Stock -= order.OrderAmount

	// Save the updated product stock in the database
	if err := db.DatabaseConnection().Save(&product).Error; err != nil {
		return err
	}

	return nil
}
