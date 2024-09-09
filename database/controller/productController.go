package controller

import (
	"ComputerWorld_API/database"
	"ComputerWorld_API/database/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

var databaseCN = database.DatabaseCN

func CreateProduct(c echo.Context) error {
	productData := new(model.Product)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	databaseCN.Model(&model.Product{}).Association("Product")

	newProduct := &model.Product{
		ProductName:    productData.ProductName,
		ProductCode:    productData.ProductCode,
		ManufacturerID: productData.ManufacturerID,
		Stock:          productData.Stock,
		Price:          productData.Price,
	}

	if err := databaseCN.Create(&newProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Product_Data": newProduct,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetProduct(c echo.Context) error {

	id := c.Param("id")

	var product model.Product

	if res := databaseCN.Where("product_id = ?", id).First(&product); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	databaseCN.Preload("Product").Preload("Manufacturer").Find(&product)

	response := map[string]interface{}{
		"Product_Data": product,
	}

	return c.JSON(http.StatusOK, response)
}

func PutProduct(c echo.Context) error {

	id := c.Param("id")
	product := new(model.Product)

	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(model.Product)

	if err := databaseCN.Where("product_id = ?", id).First(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct.ProductCode = product.ProductCode
	existingProduct.ProductName = product.ProductName
	existingProduct.Stock = product.Stock
	existingProduct.Price = product.Price

	if err := databaseCN.Save(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Product_Data": existingProduct,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	deleteProduct := new(model.Product)

	err := databaseCN.Where("product_id = ?", id).Delete(&deleteProduct).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Product has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
