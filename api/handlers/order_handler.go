package handlers

import (
	s "ComputerWorld_API/api"
	"ComputerWorld_API/api/requests"
	"ComputerWorld_API/api/responses"
	"ComputerWorld_API/api/services"
	"ComputerWorld_API/db"
	"ComputerWorld_API/db/models"
	"ComputerWorld_API/db/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderHandler struct {
	server    *s.Server
	orderRepo *repositories.OrderRepository
	service   *services.OrderService
}

func NewOrderHandler(server *s.Server) *OrderHandler {
	oh := &OrderHandler{server: server}
	oh.orderRepo = repositories.NewOrderRepository(server.Db)
	oh.service = services.NewOrderPrice(server.Db)
	return oh
}

func (h *OrderHandler) Create(c echo.Context) error {
	createOrderRequest := new(requests.CreateOrderRequest)
	if err := c.Bind(&createOrderRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "could not bind order data")
	}
	if err := c.Validate(createOrderRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	order := &models.Order{}
	if err := h.service.Create(createOrderRequest, order); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to store order in the database: %v", err))
	}

	response := responses.NewOrderResponse(order)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *OrderHandler) Update(c echo.Context) error {
	updateOrderRequest := new(requests.UpdateOrderRequest)
	if err := c.Bind(&updateOrderRequest); err != nil {
		return err
	}

	orderUID := c.Param("uid")

	order := h.orderRepo.Get(orderUID)
	if order.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "order does not exist")
	}
	if err := c.Validate(updateOrderRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	order.OrderRef = updateOrderRequest.OrderReference
	order.OrderAmount = updateOrderRequest.OrderAmount
	order.ProductUID = updateOrderRequest.ProductUID

	if err := h.service.Update(order); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong when updating the order in the database")
	}

	return responses.MessageResponse(c, http.StatusOK, "Order successfully updated")
}

func (h *OrderHandler) Get(c echo.Context) error {
	uid := c.Param("uid")

	order := &models.Order{}

	h.orderRepo.GetOrderByUID(order, uid)
	if order.UID == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Order not found")
	}

	response := responses.NewOrderResponse(order)
	return responses.Response(c, http.StatusOK, response)
}

func (h *OrderHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	order := h.orderRepo.Get(uid)

	if order.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "Order not found")
	}
	if err := h.service.Delete(order); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the order from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Order successfully deleted")
}

// Calculations >>
// These are used to automatically calculate the order prices and product stock after creation/updates

func CalculateOrderPrice(order *models.Order) error {
	var product models.Product
	if err := db.Init().First(&product, order.ProductUID).Error; err != nil {
		return err
	}
	order.OrderPrice = float64(order.OrderAmount) * product.Price
	return nil
}

func CalculateProductStock(order *models.Order) error {
	var product models.Product
	if err := db.Init().First(&product, order.ProductUID).Error; err != nil {
		return err
	}

	// Check if there's enough stock to fulfill the order
	//if product.Stock < order.OrderAmount {
	//	return responses.NewHTTPError(http.StatusBadRequest, "insufficient stock for the product")
	//}

	product.Stock -= order.OrderAmount

	// Save the updated product stock in the database
	if err := db.Init().Save(&product).Error; err != nil {
		return err
	}

	return nil
}
