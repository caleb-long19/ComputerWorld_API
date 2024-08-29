package Controller

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
