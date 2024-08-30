package controller

import (
	"ComputerWorld_API/database"
	"ComputerWorld_API/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateOrder(c echo.Context) error {
	orderData := new(model.Order)

	if err := c.Bind(orderData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newOrder := &model.Order{
		OrderAmount: orderData.OrderAmount,
	}

	if err := databaseCN.Create(&newOrder).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Order_Data": newOrder,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetOrder(c echo.Context) error {

	id := c.Param("id")

	var order model.Order

	if res := databaseCN.Where("order_id = ?", id).First(&order); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"Order_Data": order,
	}

	return c.JSON(http.StatusOK, response)
}

func PutOrder(c echo.Context) error {

	id := c.Param("id")
	order := new(model.Order)

	if err := c.Bind(order); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingOrder := new(model.Order)

	if err := databaseCN.Where("order_id = ?", id).First(&order).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingOrder.OrderAmount = order.OrderAmount

	if err := database.DatabaseCN.Save(&existingOrder).Error; err != nil {
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

func DeleteOrder(c echo.Context) error {
	id := c.Param("id")

	deleteOrder := new(model.Order)

	err := databaseCN.Where("order_id = ?", id).Delete(&deleteOrder).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Order has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
