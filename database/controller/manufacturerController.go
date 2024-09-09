package controller

import (
	"ComputerWorld_API/database/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateManufacturer(c echo.Context) error {
	manufacturerData := new(model.Manufacturer)

	if err := c.Bind(manufacturerData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newManufacturer := &model.Manufacturer{
		ManufacturerName: manufacturerData.ManufacturerName,
	}

	if err := databaseCN.Create(&newManufacturer).Error; err != nil {
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

	var manufacturer model.Manufacturer

	if res := databaseCN.Where("manufacturer_id = ?", id).First(&manufacturer); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	databaseCN.Preload("Product").Preload("Manufacturer").Find(&manufacturer)

	response := map[string]interface{}{
		"Manufacturer_Data": manufacturer,
	}

	return c.JSON(http.StatusOK, response)
}

func PutManufacturer(c echo.Context) error {

	id := c.Param("id")
	manufacturer := new(model.Manufacturer)

	if err := c.Bind(manufacturer); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer := new(model.Manufacturer)

	if err := databaseCN.Where("manufacturer_id = ?", id).First(&existingManufacturer).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingManufacturer.ManufacturerName = manufacturer.ManufacturerName

	if err := databaseCN.Save(&existingManufacturer).Error; err != nil {
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

	deleteManufacturer := new(model.Manufacturer)

	err := databaseCN.Where("manufacturer_id = ?", id).Delete(&deleteManufacturer).Error
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

/*
func GetAll(db *gorm.DB) (model.Product, error) {
	var products model.Product
	err := db.Model(&model.Product{}).Preload("ManufacturerID").Find(&products).Error
	return products, err
}
*/
