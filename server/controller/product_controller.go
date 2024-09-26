package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/reponses"
	"ComputerWorld_API/server/requests"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strconv"
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

func (pc *ProductController) Get(c echo.Context) error {
	product, err := pc.ProductRepository.Get(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(c echo.Context) error {
	products, err := pc.ProductRepository.GetAll()
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Update(c echo.Context) error {
	existingProduct, err := pc.ProductRepository.Get(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	var updateProduct = new(requests.ProductRequest)
	if err := c.Bind(updateProduct); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind product")
	}

	if updateProduct == nil {
		return c.JSON(http.StatusBadRequest, updateProduct)
	}

	_, err = pc.validateProductRequest(updateProduct)
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	existingProduct = &models.Product{
		ProductID:      existingProduct.ProductID,
		ProductName:    updateProduct.ProductName,
		ProductCode:    updateProduct.ProductCode,
		ManufacturerID: updateProduct.ManufacturerID,
		Stock:          updateProduct.ProductStock,
		Price:          updateProduct.ProductPrice,
	}

	if err := pc.ProductRepository.Update(existingProduct); err != nil {
		return reponses.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, existingProduct)
}

func (pc *ProductController) Delete(c echo.Context) error {
	err := pc.ProductRepository.Delete(c.Param("id"))
	if err != nil {
		return reponses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Product successfully deleted")
}

// Validation Methods >>>
// Simple validation methods to prevent incorrect values from being requested

func (pc *ProductController) validateProductRequest(request *requests.ProductRequest) (*models.Product, error) {
	if request == nil {
		return nil, errors.New("invalid request body")
	}

	product := new(models.Product)
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
	if request.ProductPrice <= 0.0 {
		return nil, errors.New("error: Invalid product price")
	}
	// Check for invalid characters in product values
	if validName, validCode, validID, validStock, validPrice := isValidProductInput(
		request.ProductName,
		request.ProductCode,
		request.ManufacturerID,
		request.ProductStock,
		request.ProductPrice); !validCode || !validName || !validID || !validStock || !validPrice {
		return nil, errors.New("product input contains invalid characters or format")
	}

	product.ProductName = request.ProductName
	product.ProductCode = request.ProductCode
	product.ManufacturerID = request.ManufacturerID
	product.Stock = request.ProductStock
	product.Price = request.ProductPrice

	return product, nil
}

func isValidProductInput(productName string, productCode string, ManufacturerID int, ProductStock int, ProductPrice float64) (bool, bool, bool, bool, bool) {
	// Allow only letters for product name
	validNamePattern := `^[a-zA-Z0-9]+$`
	matchedName, _ := regexp.MatchString(validNamePattern, productName)

	// Allow only letters for product name
	validCodePattern := `^[a-zA-Z0-9]+$`
	matchedCode, _ := regexp.MatchString(validCodePattern, productCode)

	// Allow only whole numbers for manufacturer id
	validIDPattern := `^[0-9]+$`
	matchedID, _ := regexp.MatchString(validIDPattern, strconv.Itoa(ManufacturerID))

	// Allow only whole numbers for stock
	validStockPattern := `^[0-9]+$`
	matchedStock, _ := regexp.MatchString(validStockPattern, strconv.Itoa(ProductStock))

	// Allow only whole numbers for stock
	validPricePattern := `^[0-9]+$`
	matchedPrice, _ := regexp.MatchString(validPricePattern, strconv.FormatFloat(ProductPrice, 'f', -1, 64))

	return matchedName, matchedCode, matchedID, matchedStock, matchedPrice
}
