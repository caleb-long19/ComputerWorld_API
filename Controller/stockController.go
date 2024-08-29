package Controller

import (
	"ComputerWorld_API/Console_Application"
	"ComputerWorld_API/Model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateStock(c echo.Context) error {
	stockData := new(Model.ProductStock)

	if err := c.Bind(stockData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newStock := &Model.ProductStock{
		Stock: stockData.Stock,
	}

	if err := Console_Application.DatabaseCN.Create(&newStock).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Manufacturer_Dara": newStock,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetStock(c echo.Context) error {

	id := c.Param("id")

	var stock Model.ProductStock

	if res := Console_Application.DatabaseCN.Where("stock_id = ?", id).First(&stock); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"Stock_Data": stock,
	}

	return c.JSON(http.StatusOK, response)
}

func PutStock(c echo.Context) error {

	id := c.Param("id")
	stock := new(Model.ProductStock)

	if err := c.Bind(stock); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingStock := new(Model.ProductStock)

	if err := Console_Application.DatabaseCN.Where("stock_id = ?", id).First(&existingStock).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingStock.Stock = stock.Stock

	if err := Console_Application.DatabaseCN.Save(&existingStock).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Stock_Data": existingStock,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteStock(c echo.Context) error {
	id := c.Param("id")

	deleteStock := new(Model.ProductStock)

	err := Console_Application.DatabaseCN.Where("stock_id = ?", id).Delete(&deleteStock).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Stock has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
