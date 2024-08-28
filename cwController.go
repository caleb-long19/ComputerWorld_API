package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
// StoredProduct CRUD
func createProduct(c echo.Context) error {
	productData := new(StoredProduct)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newProduct := &StoredProduct{
		Code:  productData.Code,
		Name:  productData.Name,
		Price: productData.Price,
	}

	if err := databaseCN.Create(&newProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Employee_Data": newProduct,
	}

	return c.JSON(http.StatusCreated, response)
}

func getProduct(c echo.Context) error {

	id := c.Param("id")

	var products []*StoredProduct

	if res := databaseCN.Find(&products, id); res.Error != nil {
		return c.String(http.StatusOK, id)
	}

	response := map[string]interface{}{
		"employee_data": products[0],
	}

	return c.JSON(http.StatusOK, response)
}

func putProduct(c echo.Context) error {

	id := c.Param("id")
	productData := new(StoredProduct)

	if err := c.Bind(productData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(StoredProduct)

	if err := databaseCN.First(&existingProduct, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct.Code = productData.Code
	existingProduct.Name = productData.Name
	existingProduct.Price = productData.Price
	if err := databaseCN.Save(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"product_data": existingProduct,
	}

	return c.JSON(http.StatusOK, response)
}

func deleteProductData(c echo.Context) error {
	id := c.Param("id")

	delProduct := new(StoredProduct)

	err := databaseCN.Delete(&delProduct, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Employee has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
*/

// EmployeeData CRUD
func createEmployee(c echo.Context) error {
	employeeData := new(EmployeeData)

	if err := c.Bind(employeeData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newEmployee := &EmployeeData{
		EmployeeName: employeeData.EmployeeName,
		EmployeeRole: employeeData.EmployeeRole,
	}

	if err := databaseCN.Create(&newEmployee).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"Employee_Data": newEmployee,
	}

	return c.JSON(http.StatusCreated, response)
}

func getEmployee(c echo.Context) error {

	id := c.Param("id")

	var employees EmployeeData

	// PACKAGE YOUR FILES TO USE THEM FOR TESTING
	// CREATE MORE TESTS
	// ORGANISE YOUR FILES
	// PREVENT SQL INJECTION USING WHERE
	// UPLOAD PROJECT TO GITHUB

	if res := databaseCN.Debug().Where("Employee_ID = ?", id).First(&employees); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"employee_data": employees,
	}

	return c.JSON(http.StatusOK, response)
}

func updateEmployee(c echo.Context) error {

	id := c.Param("id")
	employeeData := new(EmployeeData)

	if err := c.Bind(employeeData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingEmployee := new(EmployeeData)

	if err := databaseCN.First(&existingEmployee, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingEmployee.EmployeeName = employeeData.EmployeeName
	existingEmployee.EmployeeRole = employeeData.EmployeeRole

	if err := databaseCN.Save(&existingEmployee).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"employee_data": existingEmployee,
	}

	return c.JSON(http.StatusOK, response)
}

func deleteEmployeeData(c echo.Context) error {
	id := c.Param("id")

	delEmployee := new(EmployeeData)

	err := databaseCN.Delete(&delEmployee, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Employee has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}
