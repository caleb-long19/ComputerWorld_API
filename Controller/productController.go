package Controller

import (
	"ComputerWorld_API/CW_Database"
	"ComputerWorld_API/Model"
	"github.com/labstack/echo/v4"
	"net/http"
)

var databaseCN = CW_Database.DatabaseCN

func CreateProduct(c echo.Context) error {
	productData := new(Model.Product)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newProduct := &Model.Product{
		ProductName: productData.ProductName,
		ProductCode: productData.ProductCode,
		Price:       productData.Price,
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

	var product Model.Product

	if res := databaseCN.Where("product_id = ?", id).First(&product); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"Product_Data": product,
	}

	return c.JSON(http.StatusOK, response)
}

func PutProduct(c echo.Context) error {

	id := c.Param("id")
	product := new(Model.Product)

	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(Model.Product)

	if err := databaseCN.Where("product_id = ?", id).First(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct.ProductName = product.ProductName

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

	deleteProduct := new(Model.Product)

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
