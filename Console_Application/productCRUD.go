package Console_Application

import (
	"ComputerWorld_API/Model"
	"fmt"
	"strconv"
	"time"
)

func createNewProduct() {
	productData := Model.ProductInformation{
		ProductCode:  "",
		ProductName:  "",
		ProductPrice: 0.0,
	}

	clearData()
	fmt.Println("Please Enter The Product Code: ")
	pCode := scanUserInput(productData.ProductCode)
	fmt.Println("Product Code: ", pCode)

	fmt.Println("Please Enter The Product Name: ")
	pName := scanUserInput(productData.ProductName)
	fmt.Println("Product Name: ", pName)

	fmt.Println("Please Enter The Product Price: ")
	pPrice := scanUserInput(productData.ProductName)
	fmt.Println("Product Price: ", pPrice)
	pPriceF, _ := strconv.ParseFloat(pPrice, 64)

	//Create - Stores selected values into stored_products table
	storeProductDetails(pCode, pName, pPriceF)

	// Check if product record is already inside the  stored_product table
	assertRecordInputError()

	//store results of create data
	result := DatabaseCN.Create(&productRecords)

	// Check errors and print results to console
	if result.Error != nil {
		panic(result.Error.Error())
	}

	ProductInformationApplication()
}

func updateProductRecords() {
	clearData()
	fmt.Println("Please enter the mame of the product you wish to change: ")
	findProduct = scanUserInput(findProduct)

	fmt.Println("Please enter the updated product name: ")
	newProductValue = scanUserInput(newProductValue)
	DatabaseCN.Model(&Model.StoredProduct{}).Select("Name").Where("Name = ?", findProduct).Updates(map[string]interface{}{"Name": newProductValue})
	fmt.Println("Product name has been changed!")

	ProductInformationApplication()
}

func deleteProduct() {
	selectRecord = ""
	fmt.Println("Deleting Product Records:")
	selectRecord = scanUserInput(selectRecord)
	fmt.Println("Deleting Product: ", selectRecord)
	DatabaseCN.Where("Name = ?", selectRecord).Delete(&Model.StoredProduct{})
	ProductInformationApplication()
}

func ProductInformationApplication() {
	clearData()
	fmt.Println("Welcome to the Product Details Page: What would you like to do?")
	fmt.Println("Create Product, Update Product, Delete Product, Exit")
	choosePage = scanUserInput(choosePage)

	fmt.Println("You Chose:", choosePage)

	if choosePage == "Create Product" {
		fmt.Println("Loading 'Create Product': Please wait...")
		time.Sleep(1 * time.Second)
		createNewProduct()
	} else if choosePage == "Update Product" {
		fmt.Println("Loading 'Update Product': Please wait...")
		time.Sleep(1 * time.Second)
		updateProductRecords()
	} else if choosePage == "Delete Product" {
		fmt.Println("Loading 'Delete Product': Please wait...")
		time.Sleep(1 * time.Second)
		deleteProduct()
	} else if choosePage == "Exit" {
		fmt.Println("Returning to Home Page: Please wait...")
		time.Sleep(1 * time.Second)
		cwIntroduction()
	}
}
