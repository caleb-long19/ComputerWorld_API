package console_application

import "C"
import (
	"ComputerWorld_API/database"
	"ComputerWorld_API/model"
	"fmt"
	"time"
)

var findManufacturer string
var newManufacturer string

func createNewManufacturer() {
	manufacturerInfo := model.Manufacturer{
		ManufacturerName: "",
	}

	clearData()

	fmt.Println("Please Enter The Employee's ProductName: ")
	mName := scanUserInput(manufacturerInfo.ManufacturerName)
	fmt.Println("Product ProductCode: ", mName)

	storeManufacturer(mName)

	//store results of employee
	result := database.DatabaseCN.Create(&manufacturerRecords)

	// Check errors and print results to console
	if result.Error != nil {
		panic(result.Error.Error())
	}

	ManufacturerInformationApplication()
}

func updateManufacturerRecords() {
	clearData()
	fmt.Println("Please Enter The ProductName of the Employee you wish to change: ")
	findManufacturer = scanUserInput(findManufacturer)

	fmt.Println("Please Enter The New ProductName: ")
	newManufacturer = scanUserInput(newManufacturer)
	database.DatabaseCN.Model(&model.Manufacturer{}).Select("manufacturer_name").Where("manufacturer_name = ?", findManufacturer).Updates(map[string]interface{}{"manufacturer_name": newManufacturer})
	fmt.Println("ProductName has been changed!")

	ProductInformationApplication()
}

func deleteManufacturer() {
	selectRecord = ""
	fmt.Println("Deleting Manufacturer Records:")
	selectRecord = scanUserInput(selectRecord)
	fmt.Println("Deleting Manufacturer Data: ", selectRecord)
	database.DatabaseCN.Where("manufacturer_name = ?", selectRecord).Delete(&model.Manufacturer{})
	ManufacturerInformationApplication()
}

func ManufacturerInformationApplication() {
	clearData()
	fmt.Println("Welcome to the Manufacturer Information Page: What would you like to do?")
	fmt.Println("Add Manufacturer, Update Manufacturer, Delete Manufacturer, Exit")
	choosePage = scanUserInput(choosePage)

	fmt.Println("You Chose:", choosePage)

	if choosePage == "Add Manufacturer" {
		fmt.Println("Loading 'Add Manufacturer': Please wait...")
		time.Sleep(1 * time.Second)
		createNewManufacturer()
	} else if choosePage == "Update Manufacturer" {
		fmt.Println("Loading 'Update Manufacturer': Please wait...")
		time.Sleep(1 * time.Second)
		updateManufacturerRecords()
	} else if choosePage == "Delete Manufacturer" {
		fmt.Println("Loading 'Delete Manufacturer': Please wait...")
		time.Sleep(1 * time.Second)
		deleteManufacturer()
	} else if choosePage == "Exit" {
		fmt.Println("Returning to Home Page: Please wait...")
		time.Sleep(1 * time.Second)
		cwIntroduction()
	}
}

func storeManufacturer(name string) {
	manufacturerRecords = model.Manufacturer{ManufacturerName: name}
}
