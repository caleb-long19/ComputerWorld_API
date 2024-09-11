package controller

import (
	mdl "ComputerWorld_API/database/model"
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
	order := new(mdl.Order)

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var priceValue = mdl.Product{ProductID: order.ProductID}
	h.Db.Model(priceValue).Where("product_id = ?", order.ProductID).Select("price").Find(&priceValue)

	newOrder := &mdl.Order{
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

	var order mdl.Order

	if res := h.Db.Where("order_id = ?", id).First(&order); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("order_id %v", order.OrderID))
}

func (h *OrderController) PutOrder(c echo.Context) error {

	id := c.Param("id")
	order := new(mdl.Order)

	if err := c.Bind(order); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingOrder := new(mdl.Order)

	if err := h.Db.Where("order_id = ?", id).First(&order).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingOrder.OrderAmount = order.OrderAmount

	if err := h.Db.Save(&existingOrder).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Order_Data": existingOrder,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *OrderController) DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	err := h.Db.Where("order_id = ?", id).Delete(&mdl.Order{}).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, "Could not delete order")
	}

	return c.JSON(http.StatusOK, "Order has been deleted")
}
