package controller

import (
	"ComputerWorld_API/database/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ProductController struct {
	Db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{Db: db}
}

func (h *ProductController) CreateProduct(c echo.Context) error {
	productData := new(model.Product)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	h.Db.Model(&model.Product{}).Association("Product")

	newProduct := &model.Product{
		ProductName:    productData.ProductName,
		ProductCode:    productData.ProductCode,
		ManufacturerID: productData.ManufacturerID,
		Stock:          productData.Stock,
		Price:          productData.Price,
	}

	if err := h.Db.Create(&newProduct).Error; err != nil {
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

func (h *ProductController) GetProduct(c echo.Context) error {

	id := c.Param("id")

	var product model.Product

	if res := h.Db.Where("product_id = ?", id).First(&product); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"product": product,
	}

	println(c.JSON(http.StatusOK, response))
	return c.JSON(http.StatusOK, fmt.Sprintf("product_id %v", product.ProductID))
}

func (h *ProductController) PutProduct(c echo.Context) error {

	id := c.Param("id")

	product := new(model.Product)

	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind product")
	}

	existingProduct := new(model.Product)

	if err := h.Db.Where("product_id = ?", id).First(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find product by ID")
	}

	existingProduct.ProductCode = product.ProductCode
	existingProduct.ProductName = product.ProductName
	existingProduct.ManufacturerID = product.ManufacturerID
	existingProduct.Stock = product.Stock
	existingProduct.Price = product.Price

	if err := h.Db.Save(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind product")
	}

	return c.JSON(http.StatusOK, "Successfully updated product")
}

func (h *ProductController) DeleteProduct(c echo.Context) error {
	var product model.Product

	result := h.Db.Where("product_id = ?", c.Param("id")).First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find product by that ID")
	}

	result = h.Db.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not delete product by that ID")
	}

	return c.JSON(http.StatusOK, "Success: Product has been deleted")
}
