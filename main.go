package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// Database Connection
var databaseCN = databaseConnection("Computer_world.db")

var appType string
var CheckRecordExists bool

func main() {
	// Restful API
	apiServer()

	// Console application
	// cwIntroduction()
}

func cwIntroduction() {

	clearData()

	fmt.Println("Welcome to the Computer World Database!")
	fmt.Println("Choose a page: Product Details, Employee Details, Close Application")

	appType = scanUserInput(appType)
	fmt.Println("You Chose:", appType)

	switch appType {
	case "Product Details":
		fmt.Println("Loading 'Product Details': Please wait...")
		time.Sleep(2 * time.Second)
		ProductInformationApplication()
	case "Employee Details":
		fmt.Println("Loading 'Employee Details': Please wait...")
		time.Sleep(2 * time.Second)
		EmployeeInformationApplication()
	case "Close Application":
		fmt.Println("Closing Application...")
		time.Sleep(2 * time.Second)
		os.Exit(2)
	default:
		fmt.Println("Invalid Input")
	}

	/*
		// Statements that switch between the different app pages
		if appType == "Product Details" {
			fmt.Println("Loading 'Product Details': Please wait...")
			time.Sleep(2 * time.Second)
			ProductInformationApplication()
		} else if appType == "Employee Details" {
			fmt.Println("Loading 'Employee Details': Please wait...")
			time.Sleep(2 * time.Second)
			EmployeeInformationApplication()
		} else if appType == "Close Application" {
			fmt.Println("Closing Application...")
			time.Sleep(2 * time.Second)
			os.Exit(2)
		}
	*/

}

// Error Handling (Prevent duplicates and wrong inputs)
func assertRecordInputError() {
	err := databaseCN.Model(productRecords).
		Select("count(*) > 0").
		Where("id = ? AND `Code` = ? AND `Name` = ?", productRecords.ID, productRecords.Code, productRecords.Name).
		Find(&CheckRecordExists).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record does not exist: %s", err.Error())
		} else {
			log.Fatalf("Database Error Found: %s", err.Error())
		}
	}
}

// Clear string data to reset user inputs
func clearData() {
	appType = ""
	choosePage = ""
	newEmployeeValue = ""
	findEmployee = ""
}
