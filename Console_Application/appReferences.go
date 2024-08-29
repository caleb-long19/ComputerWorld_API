package Console_Application

import (
	Model2 "ComputerWorld_API/Model"
	"bufio"
	"os"
)

// Repeat Application
var choosePage string
var selectRecord string

var findManufacturer string
var newManufacturer string

var findProduct string
var newProduct string

// Database
var productRecords Model2.Products
var manufacturerRecords Model2.Manufacturer

func scanUserInput(scanInput string) string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanInput = scanner.Text()

	return scanInput
}
