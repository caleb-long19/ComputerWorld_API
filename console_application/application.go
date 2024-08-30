package console_application

import (
	"ComputerWorld_API/database"
	"ComputerWorld_API/model"
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// Database
var productRecords model.Product
var manufacturerRecords model.Manufacturer

// Repeat Application
var choosePage string
var selectRecord string

var appType string
var CheckRecordExists bool

func main() {

	// Console application
	cwIntroduction()
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
		ManufacturerInformationApplication()
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
	err := database.DatabaseCN.Model(productRecords).
		Select("count(*) > 0").
		Where("id = ? AND `ProductCode` = ? AND `ProductName` = ?", productRecords.ProductID, productRecords.ProductCode, productRecords.ProductName).
		Find(&CheckRecordExists).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Fatalf("Record does not exist: %s", err.Error())
		} else {
			log.Fatalf("Database Error Found: %s", err.Error())
		}
	}
}

func scanUserInput(scanInput string) string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanInput = scanner.Text()

	return scanInput
}

// Clear string data to reset user inputs
func clearData() {
	appType = ""
	choosePage = ""
	newManufacturer = ""
	findManufacturer = ""
}
