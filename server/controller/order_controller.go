package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/requests"
	"ComputerWorld_API/server/responses"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type OrderController struct {
	OrderRepository repositories.OrderInterface
	DB              *gorm.DB
}

func (oc *OrderController) Create(c echo.Context) error {
	// Bind request body to the OrderRequest struct
	requestOrder := new(requests.OrderRequest)

	if err := c.Bind(&requestOrder); err != nil {
		// Return bad request if binding fails
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind order data"))
	}

	// Call the validation method
	_, err := oc.validateOrderRequest(requestOrder)
	if err != nil {
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			// Return the exact status code and message from validation
			return c.JSON(httpErr.StatusCode, echo.Map{
				"error": httpErr.Message,
			})
		}
		// If the error is not a custom HTTPError, return a generic bad request
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("order validation failed: %v", err))
	}

	// Map the validated request data to the product model
	order := &models.Order{
		OrderRef:    requestOrder.OrderReference,
		OrderAmount: requestOrder.OrderAmount,
		ProductID:   requestOrder.ProductID,
	}

	// Call repository method to create the new product
	err = oc.OrderRepository.Create(order, c)
	if err != nil {
		// Return conflict if product creation fails
		return responses.ErrorResponse(c, http.StatusConflict, fmt.Errorf("failed to create order: %v", err))
	}

	// Return success response with the created product
	return c.JSON(http.StatusCreated, order)
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

	_, err = oc.validateOrderRequest(updateOrder)
	if err != nil {
		// Check if the error is of type HTTPError and use the proper status code
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return responses.ErrorResponse(c, httpErr.StatusCode, httpErr)
		}
		// For unexpected validation errors, return a generic bad request
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("order validation failed: %v", err))
	}

	existingOrder = &models.Order{
		OrderID:     existingOrder.OrderID,
		OrderRef:    updateOrder.OrderReference,
		OrderAmount: updateOrder.OrderAmount,
		ProductID:   updateOrder.ProductID,
	}

	// Attempt to update the product in the repository
	if err := oc.OrderRepository.Update(existingOrder, c); err != nil {
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

// Validation Methods >>>
// Simple validation methods to prevent incorrect values from being requested

func (oc *OrderController) validateOrderRequest(request *requests.OrderRequest) (*models.Order, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	order := new(models.Order)
	if request.OrderReference == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid order reference")
	}
	if len(request.OrderReference) < 3 || len(request.OrderReference) > 12 {
		return nil, responses.NewHTTPError(http.StatusLengthRequired, "Order reference must be between 3 and 12 characters")
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

	return order, nil
}
