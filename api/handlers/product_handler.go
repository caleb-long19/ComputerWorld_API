package handlers

import (
	s "ComputerWorld_API/api"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/api/responses"
	"ComputerWorld_API/api/services"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	server      *s.Server
	productRepo *repositories.ProductRepository
	service     *services.ProductService
}

func NewProductHandler(server *s.Server) *ProductHandler {
	ph := &ProductHandler{server: server}
	ph.productRepo = repositories.NewProductRepository(server.Db)
	ph.service = services.NewProductService(server.Db)
	return ph
}

func (h *ProductHandler) Create(c echo.Context) error {
	createProductRequest := new(requests.CreateProductRequest)
	if err := c.Bind(&createProductRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "could not bind product data")
	}
	if err := c.Validate(createProductRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	product := &models.Product{}
	if err := h.service.Create(createProductRequest, product); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to store product in the database: %v", err))
	}

	response := responses.NewProductResponse(product)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *ProductHandler) Update(c echo.Context) error {
	updateProductRequest := new(requests.UpdateProductRequest)
	if err := c.Bind(&updateProductRequest); err != nil {
		return err
	}

	productUID := c.Param("uid")

	product := h.productRepo.Get(productUID)
	if product.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "product does not exist")
	}
	if err := c.Validate(updateProductRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	product.ProductCode = updateProductRequest.ProductCode
	product.ProductName = updateProductRequest.ProductName
	product.ManufacturerUID = updateProductRequest.ManufacturerUID
	product.Stock = updateProductRequest.ProductStock
	product.Price = updateProductRequest.ProductPrice

	if err := h.service.Update(product); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong when updating the product in the database")
	}

	return responses.MessageResponse(c, http.StatusOK, "Product successfully updated")
}

func (h *ProductHandler) Get(c echo.Context) error {
	uid := c.Param("uid")

	product := &models.Product{}

	h.productRepo.GetProductByUID(product, uid)
	if product.UID == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Product not found")
	}

	response := responses.NewProductResponse(product)
	return responses.Response(c, http.StatusOK, response)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	product := h.productRepo.Get(uid)

	if product.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "Product not found")
	}
	if err := h.service.Delete(product); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the product from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Product successfully deleted")
}
