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

type AdminHandler struct {
	server    *s.Server
	adminRepo *repositories.AdminRepository
	service   *services.AdminService
}

func NewAdminHandler(server *s.Server) *AdminHandler {
	ah := &AdminHandler{server: server}
	ah.adminRepo = repositories.NewAdminRepository(server.Db)
	ah.service = services.NewAdminService(server.Db)
	return ah
}

func (h *AdminHandler) Create(c echo.Context) error {
	createAdminRequest := new(requests.CreateAdminRequest)
	if err := c.Bind(&createAdminRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("could not bind admin data"))
	}
	if err := c.Validate(createAdminRequest); err != nil {
		return responses.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	admin := &models.Admin{}
	err := h.service.Create(createAdminRequest, admin)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Errorf("failed to create admin: %v", err))
	}

	response := responses.NewAdminResponse(admin)
	return responses.Response(c, http.StatusCreated, response)
}

func (h *AdminHandler) Update(c echo.Context) error {
	updateAdminRequest := new(requests.UpdateAdminRequest)
	if err := c.Bind(&updateAdminRequest); err != nil {
		return err
	}

	adminId := c.Param("id")

	admin := h.adminRepo.Get(adminId)
	if admin.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "admin does not exist")
	}
	if err := c.Validate(updateAdminRequest); err != nil {
		return responses.ErrResponse(c, http.StatusBadRequest, fmt.Sprintf("Required fields are empty: %v", err))
	}

	admin.Email = updateAdminRequest.Email
	admin.Name = updateAdminRequest.Name

	if err := h.service.Update(admin); err != nil {
		return responses.ErrResponse(c, http.StatusInternalServerError, "Something went wrong when updating the admin in the database")
	}

	return responses.MessageResponse(c, http.StatusOK, "Admin successfully updated")
}

func (h *AdminHandler) Get(c echo.Context) error {
	id := c.Param("id")

	admin := &models.Admin{}

	h.adminRepo.GetByAdminId(admin, id)

	if admin.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "admin does not exist")
	}

	response := responses.NewAdminResponse(admin)
	return responses.Response(c, http.StatusOK, response)
}

func (h *AdminHandler) Delete(c echo.Context) error {
	uid := c.Param("uid")

	admin := h.adminRepo.Get(uid)

	if admin.UID == "" {
		return responses.ErrResponse(c, http.StatusNotFound, "User not found")
	}

	if err := h.service.Delete(admin); err != nil {
		return responses.ErrResponse(c, http.StatusInternalServerError, "Something went wrong deleting the user from the database.")
	}

	return responses.MessageResponse(c, http.StatusOK, "User successfully deleted")
}
