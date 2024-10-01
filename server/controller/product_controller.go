package controller

import (
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"ComputerWorld_API/server/requests"
	"ComputerWorld_API/server/responses"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductController struct {
	ProductRepository repositories.ProductInterface
}

func (pc *ProductController) Create(c echo.Context) error {
	// Bind request body to the ProductRequest struct
	requestProduct := new(requests.ProductRequest)

	if err := c.Bind(&requestProduct); err != nil {
		// Return bad request if binding fails
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind product data"))
	}

	// Call the validation method
	_, err := pc.validateProductRequest(requestProduct)
	if err != nil {
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			// Return the exact status code and message from validation
			return c.JSON(httpErr.StatusCode, echo.Map{
				"error": httpErr.Message,
			})
		}
		// If the error is not a custom HTTPError, return a generic bad request
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("product validation failed: %v", err))
	}
	// Map the validated request data to the product model
	product := &models.Product{
		ProductCode:    requestProduct.ProductCode,
		ProductName:    requestProduct.ProductName,
		ManufacturerID: requestProduct.ManufacturerID,
		Stock:          requestProduct.ProductStock,
		Price:          requestProduct.ProductPrice,
	}

	// Call repository method to create the new product
	err = pc.ProductRepository.Create(product, c)
	if err != nil {
		// Return conflict if product creation fails
		return responses.ErrorResponse(c, http.StatusConflict, fmt.Errorf("failed to create product: %v", err))
	}

	// Return success response with the created product
	return c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) Get(c echo.Context) error {
	product, err := pc.ProductRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(c echo.Context) error {
	products, err := pc.ProductRepository.GetAll()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Update(c echo.Context) error {
	// Get the existing product by ID
	existingProduct, err := pc.ProductRepository.Get(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("product not found: %v", err))
	}

	// Bind the incoming request to the ProductRequest struct
	var updateProduct = new(requests.ProductRequest)
	if err := c.Bind(updateProduct); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind product data"))
	}

	// Validate the incoming product data
	if updateProduct == nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid product data"))
	}

	_, err = pc.validateProductRequest(updateProduct)
	if err != nil {
		// Check if the error is of type HTTPError and use the proper status code
		var httpErr *responses.HTTPError
		if errors.As(err, &httpErr) {
			return responses.ErrorResponse(c, httpErr.StatusCode, httpErr)
		}
		// For unexpected validation errors, return a generic bad request
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("product validation failed: %v", err))
	}

	// Update the existing product fields with the new values
	existingProduct.ProductName = updateProduct.ProductName
	existingProduct.ProductCode = updateProduct.ProductCode
	existingProduct.ManufacturerID = updateProduct.ManufacturerID
	existingProduct.Stock = updateProduct.ProductStock
	existingProduct.Price = updateProduct.ProductPrice

	// Attempt to update the product in the repository
	if err := pc.ProductRepository.Update(existingProduct, c); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Errorf("failed to update product: %v", err))
	}

	return c.JSON(http.StatusCreated, existingProduct)
}

func (pc *ProductController) Delete(c echo.Context) error {
	err := pc.ProductRepository.Delete(c.Param("id"))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, err)
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
	if request.ProductCode == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid product code")
	}
	if len(request.ProductCode) < 3 || len(request.ProductCode) > 12 {
		return nil, responses.NewHTTPError(http.StatusLengthRequired, "Product code must be between 3 and 12 characters")
	}
	if request.ProductName == "" {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid product name")
	}
	if len(request.ProductName) < 3 || len(request.ProductName) > 25 {
		return nil, responses.NewHTTPError(http.StatusLengthRequired, "Product name must be between 3 and 25 characters")
	}
	if request.ManufacturerID <= 0 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid manufacturer ID")
	}
	if request.ProductStock < 0 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid stock amount")
	}
	if request.ProductStock > 1000 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Product stock exceeds maximum limit")
	}
	if request.ProductPrice <= 0.0 {
		return nil, responses.NewHTTPError(http.StatusBadRequest, "Invalid product price")
	}

	product.ProductCode = request.ProductCode
	product.ProductName = request.ProductName
	product.ManufacturerID = request.ManufacturerID
	product.Stock = request.ProductStock
	product.Price = request.ProductPrice

	return product, nil
}
