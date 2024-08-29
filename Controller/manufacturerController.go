package Controller

import (
	"ComputerWorld_API/Console_Application"
	"ComputerWorld_API/Model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateManufacturer(c echo.Context) error {
	manufacturerData := new(Model.Manufacturer)

	if err := c.Bind(manufacturerData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newManufacturer := &Model.Manufacturer{
		ManufacturerName: manufacturerData.ManufacturerName,
	}

	if err := Console_Application.DatabaseCN.Create(&newManufacturer).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Manufacturer_Dara": newManufacturer,
	}

	return c.JSON(http.StatusCreated, response)
}

func GetManufacturer(c echo.Context) error {

	id := c.Param("id")

	var manufacturer Model.Manufacturer

	if res := Console_Application.DatabaseCN.Where("manufacturer_id = ?", id).First(&manufacturer); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"Manufacturer_Data": manufacturer,
	}

	return c.JSON(http.StatusOK, response)
}

func PutManufacturer(c echo.Context) error {

	id := c.Param("id")
	manufacturer := new(Model.Manufacturer)

	if err := c.Bind(manufacturer); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer := new(Model.Manufacturer)

	if err := Console_Application.DatabaseCN.Where("manufacturer_id = ?", id).First(&existingManufacturer).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer.ManufacturerName = manufacturer.ManufacturerName

	if err := Console_Application.DatabaseCN.Save(&existingManufacturer).Error; err != nil {
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

func DeleteManufacturer(c echo.Context) error {
	id := c.Param("id")

	deleteManufacturer := new(Model.Manufacturer)

	err := Console_Application.DatabaseCN.Where("manufacturer_id = ?", id).Delete(&deleteManufacturer).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Product has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
