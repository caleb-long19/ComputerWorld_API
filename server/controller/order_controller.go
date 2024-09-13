package controller

import (
	"ComputerWorld_API/db/model"
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

//func (oc *OrderController) Read(c echo.Context) error {
//	var order model.Order
//
//	if res := oc.Db.Where("order_id = ?", c.Param("id")).First(&order); res.Error != nil {
//		return c.String(http.StatusNotFound, "Error: Order with ID was not found")
//	}
//
//	return c.JSON(http.StatusOK, order)
//}
//
//func (oc *OrderController) Update(c echo.Context) error {
//	order := new(model.Order)
//	if err := c.Bind(order); err != nil {
//		return c.JSON(http.StatusBadRequest, "Error: Could not bind order")
//	}
//
//	existingOrder := new(model.Order)
//	if err := oc.Db.Where("order_id = ?", c.Param("id")).First(&existingOrder).Error; err != nil {
//		return c.JSON(http.StatusNotFound, existingOrder.OrderID)
//	}
//
//	existingOrder.OrderRef = order.OrderRef
//	existingOrder.OrderAmount = order.OrderAmount
//	existingOrder.ProductID = order.ProductID
//	existingOrder.OrderPrice = order.OrderPrice * float64(order.OrderAmount)
//	if err := oc.Db.Save(&existingOrder).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, "Error: Could not save order")
//	}
//
//	return c.JSON(http.StatusOK, existingOrder)
//}
//
//func (oc *OrderController) Delete(c echo.Context) error {
//	var order model.Order
//
//	result := oc.Db.Where("order_id = ?", c.Param("id")).First(&order)
//	if result.Error != nil {
//		return c.JSON(http.StatusNotFound, order)
//	}
//
//	result = oc.Db.Delete(&order)
//	if result.Error != nil {
//		return c.JSON(http.StatusBadRequest, order)
//	}
//
//	return c.JSON(http.StatusOK, order)
//}

func (oc *OrderController) validateOrderRequest(request *requests.OrderRequest) (*model.Order, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	order := new(model.Order)
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
