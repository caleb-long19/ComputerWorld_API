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

	return c.JSON(http.StatusOK, fmt.Sprintf("product_id %v", product.ProductID))
}

func (h *ProductController) PutProduct(c echo.Context) error {

	id := c.Param("id")
	product := new(model.Product)

	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(model.Product)

	if err := h.Db.Where("product_id = ?", id).First(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct.ProductCode = product.ProductCode
	existingProduct.ProductName = product.ProductName
	existingProduct.Stock = product.Stock
	existingProduct.Price = product.Price

	if err := h.Db.Save(&existingProduct).Error; err != nil {
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

func (h *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := h.Db.Where("product_id = ?", id).Delete(&model.Product{}).Error
	if err != nil {

		return c.JSON(http.StatusInternalServerError, "Could not delete product")
	}

	return c.JSON(http.StatusOK, "Product has been deleted")
}
