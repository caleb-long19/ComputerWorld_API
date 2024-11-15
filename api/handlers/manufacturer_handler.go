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

type ManufacturerHandler struct {
	server           *s.Server
	manufacturerRepo *repositories.ManufacturerRepository
	service          *services.ManufacturerService
}

func NewManufacturerHandler(server *s.Server) *ManufacturerHandler {
	mh := &ManufacturerHandler{server: server}
	mh.manufacturerRepo = repositories.NewManufacturerRepository(server.Db)
	mh.service = services.NewManufacturerService(server.Db)
	return mh
}

func (h *ManufacturerHandler) Create(c echo.Context) error {
	createManufacturerRequest := new(requests.CreateManufacturerRequest)
	if err := c.Bind(&createManufacturerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "could not bind manufacturer data")
	}
	if err := c.Validate(createManufacturerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	manufacturer := &models.Manufacturer{}
	if err := h.service.Create(createManufacturerRequest, manufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed to store manufacturer in the database: %v", err))
	}

	response := responses.NewManufacturerResponse(manufacturer)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *ManufacturerHandler) Update(c echo.Context) error {
	updateManufacturerRequest := new(requests.UpdateManufacturerRequest)
	if err := c.Bind(&updateManufacturerRequest); err != nil {
		return err
	}

	manufacturerUID := c.Param("uid")

	manufacturer := h.manufacturerRepo.Get(manufacturerUID)
	if manufacturer.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "manufacturer does not exist")
	}
	if err := c.Validate(updateManufacturerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	manufacturer.ManufacturerName = updateManufacturerRequest.ManufacturerName

	if err := h.service.Update(manufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong when updating the manufacturer in the database")
	}

	return responses.MessageResponse(c, http.StatusOK, "Manufacturer successfully updated")
}

func (h *ManufacturerHandler) Get(c echo.Context) error {
	uid := c.Param("uid")

	manufacturer := &models.Manufacturer{}

	h.manufacturerRepo.GetManufacturerByUID(manufacturer, uid)
	if manufacturer.UID == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Manufacturer not found")
	}

	response := responses.NewManufacturerResponse(manufacturer)
	return responses.Response(c, http.StatusOK, response)
}

func (h *ManufacturerHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	manufacturer := h.manufacturerRepo.Get(uid)

	if manufacturer.UID == "" {
		return responses.ErrorResponse(c, http.StatusNotFound, "Manufacturer not found")
	}
	if err := h.service.Delete(manufacturer); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Something went wrong deleting the manufacturer from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "Manufacturer successfully deleted")
}
