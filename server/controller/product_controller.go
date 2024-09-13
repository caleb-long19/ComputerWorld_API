package controller

import (
	"ComputerWorld_API/db/model"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/reponses"
	"ComputerWorld_API/server/requests"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductController struct {
	ProductRepository repositories.ProductInterface
}

func (pc *ProductController) Create(c echo.Context) error {
	requestProduct := new(requests.ProductRequest)

	if err := c.Bind(&requestProduct); err != nil {
		return c.JSON(http.StatusBadRequest, requestProduct)
	}
	product, err := pc.validateProductRequest(requestProduct)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = pc.ProductRepository.Create(product)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, product)
}

//func (h *ProductRepository) Read(c echo.Context) error {
//
//	var product model.Product
//	if res := h.Db.Where("product_id = ?", c.Param("id")).First(&product); res.Error != nil {
//		return c.String(http.StatusNotFound, "Error: Product with ID was not found")
//	}
//
//	return c.JSON(http.StatusOK, product)
//}
//
//func (h *ProductRepository) Update(c echo.Context) error {
//	product := new(model.Product)
//	if err := c.Bind(product); err != nil {
//		return c.JSON(http.StatusBadRequest, "Error: Could not bind product")
//	}
//
//	existingProduct := new(model.Product)
//	if err := h.Db.Where("product_id = ?", c.Param("id")).First(&existingProduct).Error; err != nil {
//		return c.JSON(http.StatusNotFound, existingProduct)
//	}
//
//	existingProduct.ProductCode = product.ProductCode
//	existingProduct.ProductName = product.ProductName
//	existingProduct.ManufacturerID = product.ManufacturerID
//	existingProduct.Stock = product.Stock
//	existingProduct.Price = product.Price
//	if err := h.Db.Save(&existingProduct).Error; err != nil {
//		return c.JSON(http.StatusBadRequest, existingProduct)
//	}
//
//	return c.JSON(http.StatusOK, existingProduct)
//}
//
//func (h *ProductRepository) Delete(c echo.Context) error {
//	var product model.Product
//	result := h.Db.Where("product_id = ?", c.Param("id")).First(&product)
//	if result.Error != nil {
//		return c.JSON(http.StatusNotFound, product)
//	}
//
//	result = h.Db.Delete(&product)
//	if result.Error != nil {
//		return c.JSON(http.StatusBadRequest, product)
//	}
//
//	return c.JSON(http.StatusOK, product)
//}

func (pc *ProductController) validateProductRequest(request *requests.ProductRequest) (*model.Product, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	product := new(model.Product)
	if request.ProductName == "" {
		return nil, errors.New("error: Invalid product reference")
	}
	if request.ProductCode == "" {
		return nil, errors.New("error: Invalid product code")
	}
	if request.ManufacturerID <= 0 {
		return nil, errors.New("error: Invalid manufacturer ID")
	}
	if request.ProductStock < 0 {
		return nil, errors.New("error: Invalid stock amount")
	}
	if request.ProductPrice < 0.0 {
		return nil, errors.New("error: Invalid product price")
	}

	product.ProductName = request.ProductName
	product.ProductCode = request.ProductCode
	product.ManufacturerID = request.ManufacturerID
	product.Stock = request.ProductStock
	product.Price = request.ProductPrice

	return product, nil
}
