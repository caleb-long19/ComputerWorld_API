package controller

import (
	"ComputerWorld_API/database/model"
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

func (h *ProductController) Create(c echo.Context) error {

	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, product)
	}

	newProduct := &model.Product{
		ProductName:    product.ProductName,
		ProductCode:    product.ProductCode,
		ManufacturerID: product.ManufacturerID,
		Stock:          product.Stock,
		Price:          product.Price,
	}
	if err := h.Db.Create(&newProduct).Error; err != nil {
		return c.JSON(http.StatusConflict, newProduct)
	}

	return c.JSON(http.StatusCreated, newProduct)
}

func (h *ProductController) Read(c echo.Context) error {

	var product model.Product
	if res := h.Db.Where("product_id = ?", c.Param("id")).First(&product); res.Error != nil {
		return c.String(http.StatusNotFound, "Error: Product with ID was not found")
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductController) Update(c echo.Context) error {

	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind product")
	}

	existingProduct := new(model.Product)
	if err := h.Db.Where("product_id = ?", c.Param("id")).First(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusNotFound, existingProduct)
	}

	existingProduct.ProductCode = product.ProductCode
	existingProduct.ProductName = product.ProductName
	existingProduct.ManufacturerID = product.ManufacturerID
	existingProduct.Stock = product.Stock
	existingProduct.Price = product.Price
	if err := h.Db.Save(&existingProduct).Error; err != nil {
		return c.JSON(http.StatusBadRequest, existingProduct)
	}

	return c.JSON(http.StatusOK, existingProduct)
}

func (h *ProductController) Delete(c echo.Context) error {

	var product model.Product
	result := h.Db.Where("product_id = ?", c.Param("id")).First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, product)
	}

	result = h.Db.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, product)
	}

	return c.JSON(http.StatusOK, product)
}
