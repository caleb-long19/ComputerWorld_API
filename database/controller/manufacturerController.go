package controller

import (
	"ComputerWorld_API/database/model"
	rsp "ComputerWorld_API/server/response"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type ManufacturerController struct {
	Db *gorm.DB
}

func NewManufacturerController(db *gorm.DB) *ManufacturerController {
	return &ManufacturerController{Db: db}
}

func (h *ManufacturerController) PostManufacturer(c echo.Context) error {
	log.Println(">>>> HERE ")

	manufacturerData := new(model.Manufacturer)

	if err := c.Bind(&manufacturerData); err != nil {
		return c.String(http.StatusBadRequest, "Failed to bind manufacturer")
	}

	newManufacturer := &model.Manufacturer{
		ManufacturerName: manufacturerData.ManufacturerName,
	}

	if err := h.Db.Create(&newManufacturer).Error; err != nil {
		return c.String(http.StatusOK, "Failed to create manufacturer")
	}

	log.Println(">>>> HERE ")

	return c.JSON(http.StatusCreated, "Manufacturer created successfully")
}

func (h *ManufacturerController) GetManufacturer(c echo.Context) error {
	log.Println(">>>> HERE ")

	id := c.Param("id")

	var manufacturer model.Manufacturer
	if res := h.Db.Where("manufacturer_id = ?", id).First(&manufacturer); res.Error != nil {
		return c.JSON(http.StatusNotFound, "record was not found")
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("manufacturer_id %v", manufacturer.ManufacturerID))
}

func (h *ManufacturerController) PutManufacturer(c echo.Context) error {

	id := c.Param("id")
	manufacturer := new(model.Manufacturer)

	if err := c.Bind(manufacturer); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer := new(model.Manufacturer)

	if err := h.Db.Where("manufacturer_id = ?", id).First(&existingManufacturer).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer.ManufacturerName = manufacturer.ManufacturerName

	if err := h.Db.Save(&existingManufacturer).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Manufacturer_Data": existingManufacturer,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *ManufacturerController) DeleteManufacturer(c echo.Context) error {
	log.Println(">>>> HERE ")

	id := c.Param("id")

	err := h.Db.Where("manufacturer_id = ?", id).Delete(&model.Manufacturer{}).Error
	if err != nil {
		return rsp.ErrorResponse(c, http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, "Manufacturer has been deleted")
}
