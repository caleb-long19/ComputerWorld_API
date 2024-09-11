package controller

import (
	"ComputerWorld_API/database/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type OrderController struct {
	Db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{Db: db}
}

func (h *OrderController) CreateOrder(c echo.Context) error {
	order := new(model.Order)

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var priceValue = model.Product{ProductID: order.ProductID}
	h.Db.Model(priceValue).Where("product_id = ?", order.ProductID).Select("price").Find(&priceValue)

	newOrder := &model.Order{
		OrderRef:     order.OrderRef,
		ProductID:    order.ProductID,
		OrderAmount:  order.OrderAmount,
		ProductPrice: priceValue.Price * float64(order.OrderAmount),
	}

	if err := h.Db.Create(&newOrder).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, fmt.Sprintf("order_id %v", order.OrderID))
}

func (h *OrderController) GetOrder(c echo.Context) error {

	id := c.Param("id")

	var order model.Order

	if res := h.Db.Where("order_id = ?", id).First(&order); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"order": order,
	}

	println(c.JSON(http.StatusOK, response))
	return c.JSON(http.StatusOK, fmt.Sprintf("order_id %v", order.OrderID))
}

func (h *OrderController) PutOrder(c echo.Context) error {

	id := c.Param("id")
	order := new(model.Order)

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not bind order")
	}

	existingOrder := new(model.Order)

	if err := h.Db.Where("order_id = ?", id).First(&order).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find order by that ID")
	}

	existingOrder.OrderAmount = order.OrderAmount

	if err := h.Db.Save(&existingOrder).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not save order")
	}

	return c.JSON(http.StatusOK, "Order updated successfully")
}

func (h *OrderController) DeleteOrder(c echo.Context) error {
	var order model.Order

	result := h.Db.Where("product_id = ?", c.Param("id")).First(&order)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find order by that ID")
	}

	result = h.Db.Delete(&order)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not delete order by that ID")
	}

	return c.JSON(http.StatusOK, "Success: Order has been deleted")
}
