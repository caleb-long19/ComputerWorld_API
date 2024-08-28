package Controller

import (
	"ComputerWorld_API/Console_Application"
	"ComputerWorld_API/Model"
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
func CreateEmployee(c echo.Context) error {
	employeeData := new(Model.EmployeeData)

	if err := c.Bind(employeeData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	newEmployee := &Model.EmployeeData{
		EmployeeName: employeeData.EmployeeName,
		EmployeeRole: employeeData.EmployeeRole,
	}

	if err := Console_Application.DatabaseCN.Create(&newEmployee).Error; err != nil {
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

func GetEmployee(c echo.Context) error {

	id := c.Param("id")

	var employee Model.EmployeeData

	// PACKAGE YOUR FILES TO USE THEM FOR TESTING
	// CREATE MORE TESTS
	// ORGANISE YOUR FILES
	// PREVENT SQL INJECTION USING WHERE
	// UPLOAD PROJECT TO GITHUB

	if res := Console_Application.DatabaseCN.Where("employee_id = ?", id).First(&employee); res.Error != nil {
		return c.String(http.StatusNotFound, id)
	}

	response := map[string]interface{}{
		"employee_data": employee,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateEmployee(c echo.Context) error {

	id := c.Param("id")
	employeeData := new(Model.EmployeeData)

	if err := c.Bind(employeeData); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	existingEmployee := new(Model.EmployeeData)

	if err := Console_Application.DatabaseCN.Where("employee_id = ?", id).First(&existingEmployee).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingEmployee.EmployeeName = employeeData.EmployeeName
	existingEmployee.EmployeeRole = employeeData.EmployeeRole

	if err := Console_Application.DatabaseCN.Save(&existingEmployee).Error; err != nil {
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

func DeleteEmployeeData(c echo.Context) error {
	id := c.Param("id")

	delEmployee := new(Model.EmployeeData)

	err := Console_Application.DatabaseCN.Where("employee_id = ?", id).Delete(&delEmployee).Error
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
