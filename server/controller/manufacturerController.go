package controller

import (
	"ComputerWorld_API/database/model"
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
	log.Println(">>>> MANUFACTURER POST ")

	manufacturerData := new(model.Manufacturer)

	if err := c.Bind(&manufacturerData); err != nil {
		return c.String(http.StatusBadRequest, "Failed to bind manufacturer")
	}

	newManufacturer := &model.Manufacturer{
		ManufacturerName: manufacturerData.ManufacturerName,
	}

	if err := h.Db.Create(&newManufacturer).Error; err != nil {
		return c.String(http.StatusConflict, "Manufacturer already exists")
	}

	return c.JSON(http.StatusCreated, "Manufacturer created successfully")
}

func (h *ManufacturerController) GetManufacturer(c echo.Context) error {
	log.Println(">>>> MANUFACTURER GET ")

	id := c.Param("id")

	var manufacturer model.Manufacturer

	if res := h.Db.Where("manufacturer_id = ?", id).First(&manufacturer); res.Error != nil {
		return c.JSON(http.StatusNotFound, "Could not find manufacturer by that ID")
	}

	response := map[string]interface{}{
		"manufacturer": manufacturer,
	}

	println(c.JSON(http.StatusOK, response))
	return c.JSON(http.StatusOK, fmt.Sprintf("manufacturer_id %v", manufacturer.ManufacturerID))
}

func (h *ManufacturerController) PutManufacturer(c echo.Context) error {

	id := c.Param("id")
	manufacturer := new(model.Manufacturer)

	if err := c.Bind(manufacturer); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind manufacturer data")
	}

	existingManufacturer := new(model.Manufacturer)

	if err := h.Db.Where("manufacturer_id = ?", id).First(&existingManufacturer).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find manufacturer by that ID")
	}

	existingManufacturer.ManufacturerName = manufacturer.ManufacturerName

	if err := h.Db.Save(&existingManufacturer).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not update manufacturer information")
	}

	return c.JSON(http.StatusOK, "Manufacturer updated successfully")
}

func (h *ManufacturerController) DeleteManufacturer(c echo.Context) error {

	var manufacturer model.Manufacturer

	result := h.Db.Where("manufacturer_id = ?", c.Param("id")).First(&manufacturer)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Error: Could not find manufacturer by that ID")
	}

	result = h.Db.Delete(&manufacturer)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not delete manufacturer by that ID")
	}

	return c.JSON(http.StatusOK, "Success: Manufacturer has been deleted")
}
