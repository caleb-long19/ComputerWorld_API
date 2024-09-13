package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/reponses"
	"ComputerWorld_API/server/requests"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderController struct {
	OrderRepository repositories.OrderInterface
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

	existingOrder = &models.Order{
		OrderID:     existingOrder.OrderID,
		OrderRef:    updateOrder.OrderReference,
		OrderAmount: updateOrder.OrderAmount,
		ProductID:   updateOrder.ProductID,
		OrderPrice:  updateOrder.OrderPrice,
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

func (oc *OrderController) validateOrderRequest(request *requests.OrderRequest) (*models.Order, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	order := new(models.Order)
	if request.OrderReference == "" {
		fmt.Printf("Order Reference:", request.OrderReference)
		return nil, errors.New("error: Invalid order reference")
	}
	if request.OrderAmount == 0 {
		return nil, errors.New("error: Invalid order amount")
	}
	if request.ProductID == 0 {
		return nil, errors.New("error: Invalid product id")
	}
	if request.OrderPrice <= 0.0 {
		return nil, errors.New("error: Invalid product price")
	}

	order.OrderRef = request.OrderReference
	order.OrderAmount = request.OrderAmount
	order.ProductID = request.ProductID
	order.OrderPrice = request.OrderPrice

	return order, nil
}
