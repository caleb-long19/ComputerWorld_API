package controller

import (
	"ComputerWorld_API/database/model"
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

func (h *OrderController) Create(c echo.Context) error {
	order := new(model.Order)

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, order)
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
		return c.JSON(http.StatusConflict, newOrder)
	}

	return c.JSON(http.StatusCreated, newOrder)
}

func (h *OrderController) Read(c echo.Context) error {

	var order model.Order

	if res := h.Db.Where("order_id = ?", c.Param("id")).First(&order); res.Error != nil {
		return c.String(http.StatusNotFound, "Error: Order with ID was not found")
	}

	return c.JSON(http.StatusOK, order)
}

func (h *OrderController) Update(c echo.Context) error {
	order := new(model.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind order")
	}

	existingOrder := new(model.Order)
	if err := h.Db.Where("order_id = ?", c.Param("id")).First(&existingOrder).Error; err != nil {
		return c.JSON(http.StatusNotFound, existingOrder.OrderID)
	}

	existingOrder.OrderRef = order.OrderRef
	existingOrder.OrderAmount = order.OrderAmount
	existingOrder.ProductID = order.ProductID
	existingOrder.ProductPrice = order.ProductPrice * float64(order.OrderAmount)
	if err := h.Db.Save(&existingOrder).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not save order")
	}

	return c.JSON(http.StatusOK, existingOrder)
}

func (h *OrderController) Delete(c echo.Context) error {
	var order model.Order

	result := h.Db.Where("order_id = ?", c.Param("id")).First(&order)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, order)
	}

	result = h.Db.Delete(&order)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, order)
	}

	return c.JSON(http.StatusOK, order)
}
