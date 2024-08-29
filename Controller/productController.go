package Controller

import (
	"ComputerWorld_API/Console_Application"
	"ComputerWorld_API/Model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateProduct(c echo.Context) error {
	productData := new(Model.Products)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newProduct := &Model.Products{
		Name:  productData.Name,
		Code:  productData.Code,
		Price: productData.Price,
	}

	if err := Console_Application.DatabaseCN.Create(&newProduct).Error; err != nil {
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

	var product Model.Products

	if res := Console_Application.DatabaseCN.Where("product_id = ?", id).First(&product); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"Product_Data": product,
	}

	return c.JSON(http.StatusOK, response)
}

func PutProduct(c echo.Context) error {

	id := c.Param("id")
	product := new(Model.Products)

	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(Model.Products)

	if err := Console_Application.DatabaseCN.Where("product_id = ?", id).First(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct.Name = product.Name

	if err := Console_Application.DatabaseCN.Save(&existingProduct).Error; err != nil {
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

	deleteProduct := new(Model.Products)

	err := Console_Application.DatabaseCN.Where("product_id = ?", id).Delete(&deleteProduct).Error
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
