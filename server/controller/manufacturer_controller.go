package controller

import (
	"ComputerWorld_API/database/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type ManufacturerController struct {
	Db *gorm.DB
}

func NewManufacturerController(db *gorm.DB) *ManufacturerController {
	return &ManufacturerController{Db: db}
}

func (h *ManufacturerController) Create(c echo.Context) error {

	manufacturer := new(model.Manufacturer)
	if err := c.Bind(&manufacturer); err != nil {
		return c.JSON(http.StatusBadRequest, manufacturer)
	}

	newManufacturer := &model.Manufacturer{ManufacturerName: manufacturer.ManufacturerName}
	if err := h.Db.Create(&newManufacturer).Error; err != nil {
		return c.JSON(http.StatusConflict, newManufacturer)
	}

	return c.JSON(http.StatusCreated, newManufacturer)
}

func (h *ManufacturerController) Read(c echo.Context) error {

	var manufacturer model.Manufacturer
	if res := h.Db.Where("manufacturer_id = ?", c.Param("id")).First(&manufacturer); res.Error != nil {
		return c.JSON(http.StatusNotFound, manufacturer)
	}

	return c.JSON(http.StatusOK, manufacturer)
}

func (h *ManufacturerController) Update(c echo.Context) error {

	manufacturer := new(model.Manufacturer)
	if err := c.Bind(manufacturer); err != nil {
		return c.JSON(http.StatusBadRequest, "Error: Could not bind manufacturer data")
	}

	existingManufacturer := new(model.Manufacturer)
	if err := h.Db.Where("manufacturer_id = ?", c.Param("id")).First(&existingManufacturer).Error; err != nil {
		return c.JSON(http.StatusNotFound, existingManufacturer)
	}

	existingManufacturer.ManufacturerName = manufacturer.ManufacturerName
	if err := h.Db.Save(&existingManufacturer).Error; err != nil {
		return c.JSON(http.StatusBadRequest, existingManufacturer)
	}

	return c.JSON(http.StatusOK, existingManufacturer)
}

func (h *ManufacturerController) Delete(c echo.Context) error {

	id := c.Param("id")

	var manufacturer model.Manufacturer
	result := h.Db.Where("manufacturer_id = ?", id).First(&manufacturer)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, manufacturer)
	}

	result = h.Db.Delete(&manufacturer)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, manufacturer)
	}

	return c.JSON(http.StatusOK, manufacturer)
}
